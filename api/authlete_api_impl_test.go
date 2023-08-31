package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestAuthorization(t *testing.T) {
	assert := assert.New(t)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// [todo] request check
		fmt.Fprintln(w, `{"resultCode": "RESULT_CODE"}`)
	}))
	defer svr.Close()
	api, err := api.New(createAuthleteConfiguration(svr.URL))
	assert.NoError(err)
	resp, aErr := api.Authorization(&dto.AuthorizationRequest{Parameters: `parameters`})
	assert.Nil(aErr)
	assert.NotNil(resp)
}
