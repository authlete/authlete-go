package api_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/authlete/authlete-go/api"
	"github.com/authlete/authlete-go/conf"
	"github.com/authlete/authlete-go/dto"
	"github.com/stretchr/testify/assert"
)

func createAuthleteConfiguration(url string) conf.AuthleteConfiguration {
	cnf := conf.AuthleteSimpleConfiguration{}
	cnf.SetBaseUrl(url)
	cnf.SetServiceApiKey(`SERVICE_API_KEY`)
	cnf.SetServiceApiSecret(`SERVICE_API_SECRET`)

	return &cnf
}

func httpHeaderCheckHelper(t *testing.T, key string, value []string) {
	if key == "User-Agent" {
		return
	}
	assert := assert.New(t)
	wantHeader := map[string][]string{
		"Accept":          {"application/json"},
		"Content-Type":    {"application/json;charset=UTF-8"},
		"Authorization":   {"Basic U0VSVklDRV9BUElfS0VZOlNFUlZJQ0VfQVBJX1NFQ1JFVA=="},
		"Accept-Encoding": {"gzip"},
	}
	assert.Equal(wantHeader[key], value)
}

func TestAuthorization_Request(t *testing.T) {
	assert := assert.New(t)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request check
		for k, v := range r.Header {
			httpHeaderCheckHelper(t, k, v)
		}
		assert.Equal("/api/auth/authorization", r.URL.Path)
		assert.Equal(url.Values{}, r.URL.Query())
		assert.Equal("POST", r.Method)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			assert.NoError(err)
		}
		assert.Equal("{\"parameters\":\"client_id=26478243745571\\u0026code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM\\u0026code_challenge_method=S256\\u0026redirect_uri=https%3A%2F%2Fmy-client.example.com%2Fcb1\\u0026response_type=code\\u0026scope=timeline.read+history.read\"}", string(body))
		// dummy response
		fmt.Fprintln(w, `{"dummy": "dummy"}`)
	}))
	defer svr.Close()
	api, err := api.New(createAuthleteConfiguration(svr.URL))
	assert.NoError(err)
	parameters := dto.AuthorizationRequestParameters{
		ResponseType:        "code",
		ClientId:            "26478243745571",
		RedirectUri:         "https://my-client.example.com/cb1",
		Scope:               "timeline.read history.read",
		CodeChallenge:       "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM",
		CodeChallengeMethod: "S256",
	}
	resp, aErr := api.Authorization(&dto.AuthorizationRequest{Parameters: parameters})
	assert.Nil(aErr)
	assert.NotNil(resp)
}
