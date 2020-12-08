package graph_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"firebase.google.com/go/auth"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gitlab.slade360emr.com/go/base"
	"gitlab.slade360emr.com/go/profile/graph"
	"gitlab.slade360emr.com/go/profile/graph/profile"
)

const (
	testHTTPClientTimeout = 180
)

var allowedOrigins = []string{
	"https://healthcloud.co.ke",
	"https://bewell.healthcloud.co.ke",
	"http://localhost:5000",
	"https://api-gateway-test.healthcloud.co.ke",
	"https://api-gateway-prod.healthcloud.co.ke",
	"https://profile-testing-uyajqt434q-ew.a.run.app",
	"https://profile-prod-uyajqt434q-ew.a.run.app",
}

// these are set up once in TestMain and used by all the acceptance tests in
// this package
var srv *http.Server
var baseURL string
var serverErr error

func TestMain(m *testing.M) {
	// setup
	ctx := context.Background()
	srv, baseURL, serverErr = base.StartTestServer(ctx, graph.PrepareServer, allowedOrigins) // set the globals
	if serverErr != nil {
		log.Printf("unable to start test server: %s", serverErr)
	}

	// run the tests
	log.Printf("about to run tests")
	code := m.Run()
	log.Printf("finished running tests")

	// cleanup here
	defer func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("test server shutdown error: %s", err)
		}
	}()
	os.Exit(code)
}

func mapToJSONReader(m map[string]interface{}) (io.Reader, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal map to JSON: %w", err)
	}

	buf := bytes.NewBuffer(bs)
	return buf, nil
}

func TestGraphQLPractitionerSignUp(t *testing.T) {
	ctx, _ := base.GetAuthenticatedContextAndToken(t)
	if ctx == nil {
		t.Errorf("nil context")
		return
	}

	graphQLURL := fmt.Sprintf("%s/%s", baseURL, "graphql")
	headers, err := base.GetGraphQLHeaders(ctx)
	if err != nil {
		t.Errorf("nil context")
		return
	}
	gql := map[string]interface{}{}
	gql["query"] = `
	mutation practitionerSignUp($signupInput: PractitionerSignupInput!){
		practitionerSignUp(input:$signupInput)
	  }
	`
	gql["variables"] = map[string]interface{}{
		"signupInput": map[string]interface{}{
			"license":   "fake license",
			"cadre":     profile.PractitionerCadreDoctor,
			"specialty": base.PractitionerSpecialtyAnaesthesia,
			"emails":    []string{"mike.farad@healthcloud.co.ke"},
		},
	}

	validQueryReader, err := mapToJSONReader(gql)
	if err != nil {
		t.Errorf("unable to get GQL JSON io Reader: %s", err)
		return
	}
	client := http.Client{
		Timeout: time.Second * testHTTPClientTimeout,
	}

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid query",
			args: args{
				body: validQueryReader,
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				http.MethodPost,
				graphQLURL,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("unable to compose request: %s", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range headers {
				r.Header.Add(k, v)
			}
			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("request error: %s", err)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("nil response")
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read request body: %s", err)
				return
			}
			assert.NotNil(t, data)
			if data == nil {
				t.Errorf("nil response data")
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}

func TestGraphQLApprovePractitionerSignUp(t *testing.T) {
	ctx, _ := base.GetAuthenticatedContextAndToken(t)
	if ctx == nil {
		t.Errorf("nil context")
		return
	}

	graphQLURL := fmt.Sprintf("%s/%s", baseURL, "graphql")
	headers, err := base.GetGraphQLHeaders(ctx)
	if err != nil {
		t.Errorf("nil context")
		return
	}
	gql := map[string]interface{}{}
	gql["query"] = `
	mutation approvePracticionerSignUp{
		approvePractitionerSignup(practitionerID: "a7942fb4-61b4-4cf2-ab39-a2904d3090c3")
	  }
	`

	validQueryReader, err := mapToJSONReader(gql)
	if err != nil {
		t.Errorf("unable to get GQL JSON io Reader: %s", err)
		return
	}
	client := http.Client{
		Timeout: time.Second * testHTTPClientTimeout,
	}

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid query",
			args: args{
				body: validQueryReader,
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				http.MethodPost,
				graphQLURL,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("unable to compose request: %s", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range headers {
				r.Header.Add(k, v)
			}
			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("request error: %s", err)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("nil response")
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read request body: %s", err)
				return
			}
			assert.NotNil(t, data)
			if data == nil {
				t.Errorf("nil response data")
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}

func TestGetProfileAttributesHandler(t *testing.T) {
	client := http.DefaultClient
	attribute := "emails"

	uids := profile.UserUIDs{
		UIDs: []string{"some-uids"},
	}
	bs, err := json.Marshal(uids)
	if err != nil {
		t.Errorf("unable to marshal test item to JSON: %s", err)
	}
	payload := bytes.NewBuffer(bs)

	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "successful get confirmed email addresses",
			args: args{
				url: fmt.Sprintf(
					"%s/internal/contactdetails/%s/",
					baseURL,
					attribute,
				),
				httpMethod: http.MethodPost,
				body:       payload,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "failed get confirmed email addresses",
			args: args{
				url: fmt.Sprintf(
					"%s/internal/contactdetails/%s/",
					baseURL,
					attribute,
				),
				httpMethod: http.MethodPost,
				body:       nil,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)

			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range base.GetDefaultHeaders(t, baseURL, "profile") {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			if tt.wantErr {
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("expected status %d, got %d and response %s", tt.wantStatus, resp.StatusCode, string(data))
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}
		})
	}
}

func TestGraphQLRequestPinReset(t *testing.T) {
	ctx := base.GetPhoneNumberAuthenticatedContext(t)
	if ctx == nil {
		t.Errorf("nil context")
		return
	}

	graphQLURL := fmt.Sprintf("%s/%s", baseURL, "graphql")
	headers, err := base.GetGraphQLHeaders(ctx)
	if err != nil {
		t.Errorf("nil context")
		return
	}
	gql := map[string]interface{}{}
	gql["query"] = `
	query requestPinReset{
		requestPinReset(msisdn: "+254711223344")
	}
	`

	validQueryReader, err := mapToJSONReader(gql)
	if err != nil {
		t.Errorf("unable to get GQL JSON io Reader: %s", err)
		return
	}
	client := http.Client{
		Timeout: time.Second * testHTTPClientTimeout,
	}

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid query",
			args: args{
				body: validQueryReader,
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				http.MethodPost,
				graphQLURL,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("unable to compose request: %s", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range headers {
				r.Header.Add(k, v)
			}
			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("request error: %s", err)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("nil response")
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read request body: %s", err)
				return
			}
			assert.NotNil(t, data)
			if data == nil {
				t.Errorf("nil response data")
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}

}

func TestGraphQLUpdatePin(t *testing.T) {
	ctx := base.GetPhoneNumberAuthenticatedContext(t)
	if ctx == nil {
		t.Errorf("nil context")
		return
	}

	graphQLURL := fmt.Sprintf("%s/%s", baseURL, "graphql")
	headers, err := base.GetGraphQLHeaders(ctx)
	if err != nil {
		t.Errorf("nil context")
		return
	}
	gql := map[string]interface{}{}
	gql["query"] = `
	mutation updateUserPIN{
		updateUserPIN(msisdn: "+254711223344", pin: "1234", otp: "654789")
	}
	`

	validQueryReader, err := mapToJSONReader(gql)
	if err != nil {
		t.Errorf("unable to get GQL JSON io Reader: %s", err)
		return
	}
	client := http.Client{
		Timeout: time.Second * testHTTPClientTimeout,
	}

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid query",
			args: args{
				body: validQueryReader,
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				http.MethodPost,
				graphQLURL,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("unable to compose request: %s", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range headers {
				r.Header.Add(k, v)
			}
			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("request error: %s", err)
				return
			}

			if resp == nil && !tt.wantErr {
				t.Errorf("nil response")
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read request body: %s", err)
				return
			}
			assert.NotNil(t, data)
			if data == nil {
				t.Errorf("nil response data")
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}

}

func TestUpdatePinHandler(t *testing.T) {
	client := http.DefaultClient
	pinRecovery := profile.PinRecovery{
		MSISDN:    base.TestUserPhoneNumber,
		PINNumber: "4565",
		OTP:       strconv.Itoa(rand.Int()),
	}
	bs, err := json.Marshal(pinRecovery)
	if err != nil {
		t.Errorf("unable to marshal test item to JSON: %s", err)
	}
	payload := bytes.NewBuffer(bs)

	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "successful update pin",
			args: args{
				url:        fmt.Sprintf("%s/update_pin", baseURL),
				httpMethod: http.MethodPost,
				body:       payload,
			},
			wantStatus: http.StatusBadRequest, // Not a verified otp code
			wantErr:    false,
		},
		{
			name: "failed generate and send otp",
			args: args{
				url:        fmt.Sprintf("%s/update_pin", baseURL),
				httpMethod: http.MethodPost,
				body:       nil,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)

			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range base.GetDefaultHeaders(t, baseURL, "profile") {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			if tt.wantErr {
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("expected status %d, got %d and response %s", tt.wantStatus, resp.StatusCode, string(data))
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}
		})
	}
}

func TestRequestPinResetHandler(t *testing.T) {
	client := http.DefaultClient
	srv := profile.NewService()
	assert.NotNil(t, srv, "service is nil")

	ctx, _ := base.GetAuthenticatedContextAndToken(t)
	if ctx == nil {
		t.Errorf("nil context")
		return
	}
	set, err := srv.SetUserPIN(ctx, base.TestUserPhoneNumber, "1234")
	if !set {
		t.Errorf("setting a pin for test user failed. It returned false")
	}
	if err != nil {
		t.Errorf("setting a pin for test user failed: %v", err)
	}
	pinRecovery := profile.PinRecovery{
		MSISDN: base.TestUserPhoneNumber,
	}
	bs, err := json.Marshal(pinRecovery)
	if err != nil {
		t.Errorf("unable to marshal test item to JSON: %s", err)
	}
	payload := bytes.NewBuffer(bs)

	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid pin reset request",
			args: args{
				url:        fmt.Sprintf("%s/request_pin_reset", baseURL),
				httpMethod: http.MethodPost,
				body:       payload,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "failed generate and send otp",
			args: args{
				url:        fmt.Sprintf("%s/request_pin_reset", baseURL),
				httpMethod: http.MethodPost,
				body:       nil,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)

			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range base.GetDefaultHeaders(t, baseURL, "profile") {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			if tt.wantErr {
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("expected status %d, got %d and response %s", tt.wantStatus, resp.StatusCode, string(data))
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}
		})
	}
}

func TestRetrieveUserProfileFirebaseDocSnapshotHandler(t *testing.T) {

	ctx := base.GetAuthenticatedContext(t)
	assert.NotNil(t, ctx)
	auth := ctx.Value(base.AuthTokenContextKey).(*auth.Token)
	assert.NotNil(t, auth)
	profileUid := &profile.BusinessPartnerUID{
		UID: auth.UID,
	}
	assert.NotNil(t, profileUid)
	srv := profile.NewService()
	assert.NotNil(t, srv)
	handler := graph.RetrieveUserProfileHandler(ctx, srv)

	assert.NotNil(t, handler)

	uidJson, err := json.Marshal(profileUid)
	assert.NotNil(t, uidJson)
	assert.Nil(t, err)

	validRequest := httptest.NewRequest(http.MethodPost, "/", nil)
	validRequest.Body = ioutil.NopCloser(bytes.NewReader(uidJson))

	type args struct {
		rw http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid case",
			args: args{
				rw: httptest.NewRecorder(),
				r:  validRequest,
			},
			want: http.StatusOK,
		},

		{
			name: "invalid case",
			args: args{
				rw: httptest.NewRecorder(),
				r:  httptest.NewRequest(http.MethodPost, "/", ioutil.NopCloser(bytes.NewReader([]byte{}))),
			},
			want: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.rw, tt.args.r)

			response, ok := tt.args.rw.(*httptest.ResponseRecorder)
			assert.True(t, ok)
			assert.NotNil(t, response)

			assert.Equal(t, tt.want, response.Code)
		})
	}
}

func TestSaveMemberCoverToFirestoreHandler(t *testing.T) {
	ctx := base.GetAuthenticatedContext(t)
	if ctx == nil {
		t.Error("nil context")
		return
	}

	aut := ctx.Value(base.AuthTokenContextKey).(*auth.Token)
	if aut == nil {
		t.Errorf("nil auth token")
		return
	}

	srv := profile.NewService()
	if srv == nil {
		t.Errorf("nil profile service")
		return
	}

	handler := graph.SaveMemberCoverToFirestoreHandler(ctx, srv)

	type Payload struct {
		PayerName      string `json:"payerName"`
		MemberName     string `json:"memberName"`
		MemberNumber   string `json:"memberNumber"`
		PayerSladeCode int    `json:"payerSladeCode"`
		UID            string `json:"uid"`
	}

	type args struct {
		payload Payload
		rw      http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid case",
			args: args{
				payload: Payload{
					PayerName:      "UAP",
					MemberName:     "Jakaya",
					MemberNumber:   "133",
					PayerSladeCode: 457,
					UID:            aut.UID,
				},
				rw: httptest.NewRecorder(),
			},
			want: http.StatusOK,
		},

		{
			name: "invalid case",
			args: args{
				payload: Payload{
					MemberName:     "Jak",
					MemberNumber:   "132",
					PayerName:      "APA",
					PayerSladeCode: 111,
					UID:            "",
				},
				rw: httptest.NewRecorder(),
			},
			want: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadJson, err := json.Marshal(tt.args.payload)
			if err != nil {
				t.Errorf("can't marshal payload to JSON")
				return
			}
			if payloadJson == nil {
				t.Errorf("nil JSON payload")
				return
			}

			request := httptest.NewRequest(http.MethodPost, "/", nil)
			request.Body = ioutil.NopCloser(bytes.NewReader(payloadJson))
			handler(tt.args.rw, request)

			response, ok := tt.args.rw.(*httptest.ResponseRecorder)
			if response == nil {
				t.Errorf("nil response")
				return
			}
			if !ok {
				t.Errorf(
					"expected response to be a *httptest.ResponseRecorder")
				return
			}

			if response.Code != tt.want {
				t.Errorf(
					"expected status code %d, got %d", tt.want, response.Code)

				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					t.Errorf("can't read response body")
					return
				}

				log.Printf("raw response data: \n%s\n", string(data))

				return
			}
		})
	}
}

func TestIsUnderAgeHandler(t *testing.T) {
	ctx := base.GetAuthenticatedContext(t)
	if ctx == nil {
		t.Error("nil context")
		return
	}

	aut := ctx.Value(base.AuthTokenContextKey).(*auth.Token)
	if aut == nil {
		t.Errorf("nil auth token")
		return
	}

	srv := profile.NewService()
	if srv == nil {
		t.Errorf("nil profile service")
		return
	}

	fc := &base.FirebaseClient{}
	fa, err := fc.InitFirebase()
	if err != nil {
		t.Errorf("can't initialize Firebase app: %s", err)
		return
	}
	firestore, err := fa.Firestore(context.Background())
	if err != nil {
		t.Errorf(
			"can't initialize Firestore client: %s", err)
		return
	}

	profile, err := srv.UserProfile(ctx)
	if profile == nil {
		t.Errorf("expected a profile")
		return
	}
	if err != nil {
		t.Errorf("did not expect an error: %v", err)
		return
	}

	date := &base.Date{
		Year:  1997,
		Month: 12,
		Day:   13,
	}
	profile.DateOfBirth = date
	dsnap, err := srv.RetrieveUserProfileFirebaseDocSnapshot(ctx)
	if err != nil {
		t.Errorf("can't retrieve user profile firebase snapshot: %v", err)
		return
	}

	err = base.UpdateRecordOnFirestore(
		firestore, srv.GetUserProfileCollectionName(),
		dsnap.Ref.ID, profile,
	)
	if err != nil {
		t.Errorf("can't update user profile on Firebase: %v", err)
		return
	}

	handler := graph.IsUnderAgeHandler(ctx, srv)

	type UserContext struct {
		UID string `json:"uid"`
	}

	type args struct {
		userContext UserContext
	}
	tests := []struct {
		name string
		args args
		want int
		rw   http.ResponseWriter
	}{
		{
			name: "valid case - valid UID",
			args: args{
				UserContext{
					UID: aut.UID,
				},
			},
			rw:   httptest.NewRecorder(),
			want: http.StatusOK,
		},

		{
			name: "invalid case - blank UID",
			args: args{
				UserContext{
					UID: "",
				},
			},
			rw:   httptest.NewRecorder(),
			want: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadJson, err := json.Marshal(tt.args.userContext)
			if err != nil {
				t.Errorf("can't marshal payload to JSON")
				return
			}
			if payloadJson == nil {
				t.Errorf("nil JSON payload")
				return
			}

			request := httptest.NewRequest(http.MethodPost, "/", nil)
			request.Body = ioutil.NopCloser(bytes.NewReader(payloadJson))

			handler(tt.rw, request)

			response, ok := tt.rw.(*httptest.ResponseRecorder)
			if response == nil {
				t.Errorf("nil response")
				return
			}
			if !ok {
				t.Errorf(
					"expected response to be a *httptest.ResponseRecorder")
				return
			}

			if response.Code != tt.want {
				t.Errorf(
					"expected status code %d, got %d", tt.want, response.Code)

				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					t.Errorf("can't read response body")
					return
				}

				log.Printf("raw response data: \n%s\n", string(data))

				return
			}
		})
	}
}

func TestUserProfileHandler(t *testing.T) {
	ctx := base.GetAuthenticatedContext(t)
	assert.NotNil(t, ctx, "context should not be nil")

	authToken := ctx.Value(base.AuthTokenContextKey).(*auth.Token)
	assert.NotNil(t, authToken, "authToken should not be nil")

	srv := profile.NewService()
	assert.NotNil(t, srv, "service is nil")

	handler := graph.UserProfileHandler(ctx, srv)

	type UserContext struct {
		Token *auth.Token `json:"token"`
	}

	type args struct {
		userContext UserContext
	}
	tests := []struct {
		name string
		args args
		rw   http.ResponseWriter
		want int
	}{
		{
			name: "valid case",
			args: args{
				userContext: UserContext{
					Token: authToken,
				},
			},
			rw:   httptest.NewRecorder(),
			want: http.StatusOK,
		},
		{
			name: "invalid case",
			args: args{
				userContext: UserContext{
					Token: nil,
				},
			},
			want: http.StatusBadRequest,
			rw:   httptest.NewRecorder(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadJson, err := json.Marshal(tt.args.userContext)
			assert.Nil(t, err, "failed to marshal payload")
			assert.NotNil(t, payloadJson, "payload is nil")

			request := httptest.NewRequest(http.MethodPost, "/", nil)
			request.Body = ioutil.NopCloser(bytes.NewReader(payloadJson))

			handler(tt.rw, request)

			response, ok := tt.rw.(*httptest.ResponseRecorder)

			assert.True(t, ok)
			assert.NotNil(t, response, "response should not be nil")

			assert.Equal(t, tt.want, response.Code)
		})
	}
}

func TestSendRetryOTPHandler(t *testing.T) {
	client := http.DefaultClient
	sendOTPRetry := profile.SendRetryOTP{
		Msisdn:    base.TestUserPhoneNumber,
		RetryStep: 1,
	}
	bs, err := json.Marshal(sendOTPRetry)
	if err != nil {
		t.Errorf("unable to marshal test item to JSON: %s", err)
	}
	payload := bytes.NewBuffer(bs)

	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid generate and send retry OTPs request",
			args: args{
				url:        fmt.Sprintf("%s/send_retry_otp", baseURL),
				httpMethod: http.MethodPost,
				body:       payload,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "invalid generate and send retry OTPs request",
			args: args{
				url:        fmt.Sprintf("%s/send_retry_otp", baseURL),
				httpMethod: http.MethodPost,
				body:       nil,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)

			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range base.GetDefaultHeaders(t, baseURL, "profile") {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			if tt.wantErr {
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("expected status %d, got %d and response %s", tt.wantStatus, resp.StatusCode, string(data))
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}
		})
	}
}

func TestRetrieveUserProfileHandler(t *testing.T) {
	client := http.DefaultClient

	_, authToken := base.GetAuthenticatedContextAndToken(t)
	if authToken == nil {
		t.Errorf("nil auth token")
		return
	}

	bpUID := &profile.BusinessPartnerUID{
		UID: authToken.UID,
	}
	bs, err := json.Marshal(bpUID)
	if err != nil {
		t.Errorf("unable to marshal BP UID payload to JSON: %s", err)
		return
	}
	payload := bytes.NewBuffer(bs)

	type args struct {
		url        string
		httpMethod string
		body       io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "valid user profile retrieve request - valid UID",
			args: args{
				url:        fmt.Sprintf("%s/internal/retrieve_user_profile", baseURL),
				httpMethod: http.MethodPost,
				body:       payload,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "invalid user profile retrieve request - nil body",
			args: args{
				url:        fmt.Sprintf("%s/internal/retrieve_user_profile", baseURL),
				httpMethod: http.MethodPost,
				body:       nil,
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)

			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range base.GetDefaultHeaders(t, baseURL, "profile") {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			if tt.wantErr {
				return
			}

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("expected status %d, got %d and response %s", tt.wantStatus, resp.StatusCode, string(data))
				return
			}

			if !tt.wantErr && resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}
		})
	}
}
