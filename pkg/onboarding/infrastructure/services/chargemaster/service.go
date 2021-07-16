package chargemaster

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/savannahghi/firebasetools"
	"github.com/savannahghi/serverutils"
	log "github.com/sirupsen/logrus"
	"gitlab.slade360emr.com/go/apiclient"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/dto"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/domain"
)

const (
	// ChargeMasterHostEnvVarName is the name of an environment variable that
	//points at the API root e.g "https://base.chargemaster.slade360emr.com/v1"
	ChargeMasterHostEnvVarName = "CHARGE_MASTER_API_HOST"

	// ChargeMasterAPISchemeEnvVarName points at an environment variable that
	// indicates whether the API is "http" or "https". It is used when our code
	// needs to construct custom API paths from scratch.
	ChargeMasterAPISchemeEnvVarName = "CHARGE_MASTER_API_SCHEME"

	// ChargeMasterTokenURLEnvVarName is an environment variable that contains
	// the path to the OAuth 2 token URL for the charge master base. This URL
	// could be the same as that used by other Slade 360 products e.g EDI.
	// It could also be different.
	ChargeMasterTokenURLEnvVarName = "CHARGE_MASTER_TOKEN_URL"

	// ChargeMasterClientIDEnvVarName is the name of an environment variable that holds
	// the OAuth2 client ID for a charge master API application.
	ChargeMasterClientIDEnvVarName = "CHARGE_MASTER_CLIENT_ID"

	// ChargeMasterClientSecretEnvVarName is the name of an environment variable that holds
	// the OAuth2 client secret for a charge master API application.
	ChargeMasterClientSecretEnvVarName = "CHARGE_MASTER_CLIENT_SECRET"

	// ChargeMasterUsernameEnvVarName is the name of an environment variable that holds the
	// username of a charge master API user.
	ChargeMasterUsernameEnvVarName = "CHARGE_MASTER_USERNAME"

	// ChargeMasterPasswordEnvVarName is the name of an environment variable that holds the
	// password of the charge master API user referred to by `ChargeMasterUsernameEnvVarName`.
	ChargeMasterPasswordEnvVarName = "CHARGE_MASTER_PASSWORD"

	// ChargeMasterGrantTypeEnvVarName should be "password" i.e the only type of OAuth 2
	// "application" that will work for this client is a confidential one that supports
	// password authentication.
	ChargeMasterGrantTypeEnvVarName = "CHARGE_MASTER_GRANT_TYPE"

	// ChargeMasterBusinessPartnerPath endpoint for business partners on charge master
	ChargeMasterBusinessPartnerPath = "/v1/business_partners/"
)

// ServiceChargeMaster represents logic required to communicate with chargemaster
type ServiceChargeMaster interface {
	FetchChargeMasterClient() *apiclient.ServerClient
	FindProvider(ctx context.Context, pagination *firebasetools.PaginationInput, filter []*dto.BusinessPartnerFilterInput,
		sort []*dto.BusinessPartnerSortInput) (*dto.BusinessPartnerConnection, error)
	FindBranch(ctx context.Context, pagination *firebasetools.PaginationInput, filter []*dto.BranchFilterInput,
		sort []*dto.BranchSortInput) (*dto.BranchConnection, error)
	FetchProviderByID(ctx context.Context, id string) (*domain.BusinessPartner, error)
}

// ServiceChargeMasterImpl ..
type ServiceChargeMasterImpl struct {
	ChargeMasterClient *apiclient.ServerClient
}

// NewChargeMasterUseCasesImpl ...
func NewChargeMasterUseCasesImpl() ServiceChargeMaster {

	clientID := serverutils.MustGetEnvVar(ChargeMasterClientIDEnvVarName)
	clientSecret := serverutils.MustGetEnvVar(ChargeMasterClientSecretEnvVarName)
	apiTokenURL := serverutils.MustGetEnvVar(ChargeMasterTokenURLEnvVarName)
	apiHost := serverutils.MustGetEnvVar(ChargeMasterHostEnvVarName)
	apiScheme := serverutils.MustGetEnvVar(ChargeMasterAPISchemeEnvVarName)
	grantType := serverutils.MustGetEnvVar(ChargeMasterGrantTypeEnvVarName)
	username := serverutils.MustGetEnvVar(ChargeMasterUsernameEnvVarName)
	password := serverutils.MustGetEnvVar(ChargeMasterPasswordEnvVarName)
	extraHeaders := make(map[string]string)
	client, err := apiclient.NewServerClient(
		clientID, clientSecret, apiTokenURL, apiHost, apiScheme, grantType, username, password, extraHeaders)
	if err != nil {
		log.Panicf("unable to initialize Chargemaster client for profile service: %s", err)
		os.Exit(1)
	}

	return &ServiceChargeMasterImpl{ChargeMasterClient: client}
}

// FetchChargeMasterClient ...
func (chr ServiceChargeMasterImpl) FetchChargeMasterClient() *apiclient.ServerClient {
	return chr.ChargeMasterClient
}

// FindProvider search for a provider in chargemaster using their name
//
// Example https://base.chargemaster.slade360emr.com/v1/business_partners/?bp_type=PROVIDER&search={name}
func (chr ServiceChargeMasterImpl) FindProvider(ctx context.Context, pagination *firebasetools.PaginationInput,
	filter []*dto.BusinessPartnerFilterInput, sort []*dto.BusinessPartnerSortInput) (*dto.BusinessPartnerConnection, error) {

	paginationParams, err := firebasetools.GetAPIPaginationParams(pagination)
	if err != nil {
		return nil, err
	}

	defaultParams := url.Values{}
	defaultParams.Add("fields", "id,name,slade_code,parent")
	defaultParams.Add("is_active", "True")
	defaultParams.Add("bp_type", "PROVIDER")

	queryParams := []url.Values{defaultParams, paginationParams}
	for _, fp := range filter {
		queryParams = append(queryParams, fp.ToURLValues())
	}
	for _, fp := range sort {
		queryParams = append(queryParams, fp.ToURLValues())
	}

	mergedParams := apiclient.MergeURLValues(queryParams...)
	queryFragment := mergedParams.Encode()

	type apiResp struct {
		apiclient.SladeAPIListRespBase

		Results []*domain.BusinessPartner `json:"results,omitempty"`
	}

	r := apiResp{}
	err = apiclient.ReadRequestToTarget(chr.FetchChargeMasterClient(), "GET", ChargeMasterBusinessPartnerPath, queryFragment, nil, &r)
	if err != nil {
		return nil, err
	}

	startOffset := firebasetools.CreateAndEncodeCursor(r.StartIndex)
	endOffset := firebasetools.CreateAndEncodeCursor(r.EndIndex)
	hasNextPage := r.Next != ""
	hasPreviousPage := r.Previous != ""

	edges := []*dto.BusinessPartnerEdge{}
	for pos, org := range r.Results {
		edge := &dto.BusinessPartnerEdge{
			Node: &domain.BusinessPartner{
				ID:        org.ID,
				Name:      org.Name,
				SladeCode: org.SladeCode,
				Parent:    org.Parent,
			},
			Cursor: firebasetools.CreateAndEncodeCursor(pos + 1),
		}
		edges = append(edges, edge)
	}
	pageInfo := &firebasetools.PageInfo{
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
		StartCursor:     startOffset,
		EndCursor:       endOffset,
	}
	connection := &dto.BusinessPartnerConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	return connection, nil
}

// FetchProviderByID returns details of a specific provider given the ID
func (chr ServiceChargeMasterImpl) FetchProviderByID(ctx context.Context, id string) (*domain.BusinessPartner, error) {

	partner := &domain.BusinessPartner{}
	BusinessPartnerPath := fmt.Sprintf("%v%v/", ChargeMasterBusinessPartnerPath, id)

	err := apiclient.ReadRequestToTarget(chr.FetchChargeMasterClient(), "GET", BusinessPartnerPath, "", nil, partner)
	if err != nil {
		return nil, err
	}

	// check if partner is empty
	if (*partner == domain.BusinessPartner{}) {
		return nil, fmt.Errorf("business partner not found. invalid id: %v", id)
	}

	return partner, nil
}

// FindBranch lists all locations known to Slade 360 Charge Master
// Example URL: https://base.chargemaster.slade360emr.com/v1/business_partners/?format=json&page_size=100&parent=6ba48d97-93d2-4815-a447-f51240cbcab8&fields=id,name,slade_code
func (chr ServiceChargeMasterImpl) FindBranch(ctx context.Context, pagination *firebasetools.PaginationInput,
	filter []*dto.BranchFilterInput, sort []*dto.BranchSortInput) (*dto.BranchConnection, error) {

	paginationParams, err := firebasetools.GetAPIPaginationParams(pagination)
	if err != nil {
		return nil, err
	}
	defaultParams := url.Values{}
	defaultParams.Add("fields", "id,name,slade_code")
	defaultParams.Add("is_active", "True")
	defaultParams.Add("is_branch", "True")

	queryParams := []url.Values{defaultParams, paginationParams}
	for _, fp := range filter {
		queryParams = append(queryParams, fp.ToURLValues())
	}
	for _, fp := range sort {
		queryParams = append(queryParams, fp.ToURLValues())
	}
	mergedParams := apiclient.MergeURLValues(queryParams...)
	queryFragment := mergedParams.Encode()

	type apiResp struct {
		apiclient.SladeAPIListRespBase

		Results []*domain.BusinessPartner `json:"results,omitempty"`
	}

	r := apiResp{}
	err = apiclient.ReadRequestToTarget(chr.FetchChargeMasterClient(), "GET", ChargeMasterBusinessPartnerPath, queryFragment, nil, &r)
	if err != nil {
		return nil, err
	}
	startOffset := firebasetools.CreateAndEncodeCursor(r.StartIndex)
	endOffset := firebasetools.CreateAndEncodeCursor(r.EndIndex)
	hasNextPage := r.Next != ""
	hasPreviousPage := r.Previous != ""

	edges := []*dto.BranchEdge{}
	for pos, branch := range r.Results {
		orgSladeCode, err := parentOrgSladeCodeFromBranch(branch)
		if err != nil {
			return nil, err
		}

		edge := &dto.BranchEdge{
			Node: &domain.Branch{
				ID:                    branch.ID,
				Name:                  branch.Name,
				BranchSladeCode:       branch.SladeCode,
				OrganizationSladeCode: orgSladeCode,
			},
			Cursor: firebasetools.CreateAndEncodeCursor(pos + 1),
		}
		edges = append(edges, edge)
	}
	pageInfo := &firebasetools.PageInfo{
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
		StartCursor:     startOffset,
		EndCursor:       endOffset,
	}
	connection := &dto.BranchConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	return connection, nil
}

func parentOrgSladeCodeFromBranch(branch *domain.BusinessPartner) (string, error) {
	if !strings.HasPrefix(branch.SladeCode, "BRA-") {
		return "", fmt.Errorf("%s is not a valid branch Slade Code; expected a BRA- prefix", branch.SladeCode)
	}
	trunc := strings.TrimPrefix(branch.SladeCode, "BRA-")
	split := strings.Split(trunc, "-")
	if len(split) != 3 {
		return "", fmt.Errorf("expected the branch Slade Code to split into 3 parts on -; got %s", split)
	}
	orgParts := split[0:2]
	orgSladeCode := strings.Join(orgParts, "-")
	return orgSladeCode, nil
}
