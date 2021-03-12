package mock

import (
	"context"
	"net/http"

	"cloud.google.com/go/pubsub"
	"gitlab.slade360emr.com/go/base"
	"gitlab.slade360emr.com/go/profile/pkg/onboarding/application/extension"
)

// FakeBaseExtensionImpl is a `base` library fake  .
type FakeBaseExtensionImpl struct {
	GetLoggedInUserUIDFn   func(ctx context.Context) (string, error)
	NormalizeMSISDNFn      func(msisdn string) (*string, error)
	FetchDefaultCurrencyFn func(c base.Client) (*base.FinancialYearAndCurrency, error)
	FetchUserProfileFn     func(authClient base.Client) (*base.EDIUserProfile, error)
	LoginClientFn          func(username string, password string) (base.Client, error)
	LoadDepsFromYAMLFn     func() (*base.DepsConfig, error)
	SetupISCclientFn       func(config base.DepsConfig, serviceName string) (*base.InterServiceClient, error)
	GetEnvVarFn            func(envName string) (string, error)
	NewServerClientFn      func(
		clientID string,
		clientSecret string,
		apiTokenURL string,
		apiHost string,
		apiScheme string,
		grantType string,
		username string,
		password string,
		extraHeaders map[string]string,
	) (*base.ServerClient, error)
	EnsureTopicsExistFn func(
		ctx context.Context,
		pubsubClient *pubsub.Client,
		topicIDs []string,
	) error
	GetRunningEnvironmentFn     func() string
	NamespacePubsubIdentifierFn func(
		serviceName string,
		topicID string,
		environment string,
		version string,
	) string
	PublishToPubsubFn func(
		ctx context.Context,
		pubsubClient *pubsub.Client,
		topicID string,
		environment string,
		serviceName string,
		version string,
		payload []byte,
	) error
	GoogleCloudProjectIDEnvVarNameFn func() (string, error)
	EnsureSubscriptionsExistFn       func(
		ctx context.Context,
		pubsubClient *pubsub.Client,
		topicSubscriptionMap map[string]string,
		callbackURL string,
	) error
	SubscriptionIDsFn                 func(topicIDs []string) map[string]string
	PubSubHandlerPathFn               func() string
	VerifyPubSubJWTAndDecodePayloadFn func(
		w http.ResponseWriter,
		r *http.Request,
	) (*base.PubSubPayload, error)
	GetPubSubTopicFn    func(m *base.PubSubPayload) (string, error)
	ErrorMapFn          func(err error) map[string]string
	WriteJSONResponseFn func(
		w http.ResponseWriter,
		source interface{},
		status int,
	)
}

// GetLoggedInUserUID ...
func (b *FakeBaseExtensionImpl) GetLoggedInUserUID(ctx context.Context) (string, error) {
	return b.GetLoggedInUserUIDFn(ctx)
}

// NormalizeMSISDN ...
func (b *FakeBaseExtensionImpl) NormalizeMSISDN(msisdn string) (*string, error) {
	return b.NormalizeMSISDNFn(msisdn)
}

// FetchDefaultCurrency ...
func (b *FakeBaseExtensionImpl) FetchDefaultCurrency(c base.Client,
) (*base.FinancialYearAndCurrency, error) {
	return b.FetchDefaultCurrencyFn(c)
}

// FetchUserProfile ...
func (b *FakeBaseExtensionImpl) FetchUserProfile(authClient base.Client) (*base.EDIUserProfile, error) {
	return b.FetchUserProfileFn(authClient)
}

// LoginClient returns a logged in client with the supplied username and password
func (b *FakeBaseExtensionImpl) LoginClient(username, password string) (base.Client, error) {
	return b.LoginClientFn(username, password)
}

// LoadDepsFromYAML ...
func (b *FakeBaseExtensionImpl) LoadDepsFromYAML() (*base.DepsConfig, error) {
	return b.LoadDepsFromYAMLFn()
}

// SetupISCclient ...
func (b *FakeBaseExtensionImpl) SetupISCclient(config base.DepsConfig, serviceName string) (*base.InterServiceClient, error) {
	return b.SetupISCclientFn(config, serviceName)
}

// GetEnvVar ...
func (b *FakeBaseExtensionImpl) GetEnvVar(envName string) (string, error) {
	return b.GetEnvVarFn(envName)
}

// NewServerClient ...
func (b *FakeBaseExtensionImpl) NewServerClient(
	clientID string,
	clientSecret string,
	apiTokenURL string,
	apiHost string,
	apiScheme string,
	grantType string,
	username string,
	password string,
	extraHeaders map[string]string,
) (*base.ServerClient, error) {
	return b.NewServerClientFn(clientID, clientSecret, apiTokenURL, apiHost, apiScheme, grantType, username, password, extraHeaders)
}

// EnsureTopicsExist ...
func (b *FakeBaseExtensionImpl) EnsureTopicsExist(
	ctx context.Context,
	pubsubClient *pubsub.Client,
	topicIDs []string,
) error {
	return b.EnsureTopicsExistFn(ctx, pubsubClient, topicIDs)
}

// GetRunningEnvironment ..
func (b *FakeBaseExtensionImpl) GetRunningEnvironment() string {
	return b.GetRunningEnvironmentFn()
}

// NamespacePubsubIdentifier ..
func (b *FakeBaseExtensionImpl) NamespacePubsubIdentifier(
	serviceName string,
	topicID string,
	environment string,
	version string,
) string {
	return b.NamespacePubsubIdentifierFn(
		serviceName,
		topicID,
		environment,
		version,
	)
}

// PublishToPubsub ..
func (b *FakeBaseExtensionImpl) PublishToPubsub(
	ctx context.Context,
	pubsubClient *pubsub.Client,
	topicID string,
	environment string,
	serviceName string,
	version string,
	payload []byte,
) error {
	return b.PublishToPubsubFn(
		ctx,
		pubsubClient,
		topicID,
		environment,
		serviceName,
		version,
		payload,
	)
}

// GoogleCloudProjectIDEnvVarName ..
func (b *FakeBaseExtensionImpl) GoogleCloudProjectIDEnvVarName() (string, error) {
	return b.GoogleCloudProjectIDEnvVarNameFn()
}

// EnsureSubscriptionsExist ...
func (b *FakeBaseExtensionImpl) EnsureSubscriptionsExist(
	ctx context.Context,
	pubsubClient *pubsub.Client,
	topicSubscriptionMap map[string]string,
	callbackURL string,
) error {
	return b.EnsureSubscriptionsExistFn(
		ctx,
		pubsubClient,
		topicSubscriptionMap,
		callbackURL,
	)
}

// SubscriptionIDs ..
func (b *FakeBaseExtensionImpl) SubscriptionIDs(topicIDs []string) map[string]string {
	return b.SubscriptionIDsFn(topicIDs)
}

// PubSubHandlerPath ..
func (b *FakeBaseExtensionImpl) PubSubHandlerPath() string {
	return b.PubSubHandlerPathFn()
}

// VerifyPubSubJWTAndDecodePayload ..
func (b *FakeBaseExtensionImpl) VerifyPubSubJWTAndDecodePayload(
	w http.ResponseWriter,
	r *http.Request,
) (*base.PubSubPayload, error) {
	return b.VerifyPubSubJWTAndDecodePayloadFn(w, r)
}

// GetPubSubTopic ..
func (b *FakeBaseExtensionImpl) GetPubSubTopic(m *base.PubSubPayload) (string, error) {
	return b.GetPubSubTopicFn(m)
}

// ErrorMap ..
func (b *FakeBaseExtensionImpl) ErrorMap(err error) map[string]string {
	return b.ErrorMapFn(err)
}

// WriteJSONResponse ..
func (b *FakeBaseExtensionImpl) WriteJSONResponse(
	w http.ResponseWriter,
	source interface{},
	status int,
) {
	b.WriteJSONResponseFn(w, source, status)
}

// PINExtensionImpl is a `PIN` fake  .
type PINExtensionImpl struct {
	EncryptPINFn func(rawPwd string, options *extension.Options) (string, string)
	ComparePINFn func(rawPwd string, salt string, encodedPwd string, options *extension.Options) bool
}

// EncryptPIN ...
func (p *PINExtensionImpl) EncryptPIN(rawPwd string, options *extension.Options) (string, string) {
	return p.EncryptPINFn(rawPwd, options)
}

// ComparePIN ...
func (p *PINExtensionImpl) ComparePIN(rawPwd string, salt string, encodedPwd string, options *extension.Options) bool {
	return p.ComparePINFn(rawPwd, salt, encodedPwd, options)
}

// ISCClientExtension is an ISC fake
type ISCClientExtension struct {
	MakeRequestFn func(method string, path string, body interface{}) (*http.Response, error)
}

// MakeRequest ...
func (i *ISCClientExtension) MakeRequest(method string, path string, body interface{}) (*http.Response, error) {
	return i.MakeRequestFn(method, path, body)
}