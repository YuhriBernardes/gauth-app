package server_test

import (
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YuhriBernardes/gauth-app/server"
)

func jsonBody(entity interface{}) io.Reader {
	jsonBytes, err := json.Marshal(entity)

	if err != nil {
		panic(fmt.Sprintf("Failed to convert body to json %v", err))
	}

	contentBuffer := bytes.NewBuffer(jsonBytes)

	return contentBuffer
}

func responseBody(w *httptest.ResponseRecorder, i interface{}) {

	result := w.Result()

	defer result.Body.Close()

	bodyBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic("Failed to decode body")
	}

	json.Unmarshal(bodyBytes, i)
}

func TestAuthenticateRouteAuthorized(t *testing.T) {
	s := server.Server{Router: server.Router{}}

	s.Init()

	w := httptest.NewRecorder()

	requestBody := server.AuthenticateRequest{UserName: "edoraoff", Password: "14jkl;"}

	req, _ := http.NewRequest("POST", "/api/auth", jsonBody(requestBody))

	s.Server.ServeHTTP(w, req)

	var body server.AuthenticateResponse

	responseBody(w, &body)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(body.Token), sha512.Size*2)
}

func TestAuthenticateRouteUnauthorized(t *testing.T) {
	s := server.Server{Router: server.Router{}}

	s.Init()

	w := httptest.NewRecorder()

	requestBody := server.AuthenticateRequest{UserName: "irar", Password: "14jkl;"}

	req, _ := http.NewRequest("POST", "/api/auth", jsonBody(requestBody))

	s.Server.ServeHTTP(w, req)

	responseBody := w.Result().Body

	defer responseBody.Close()

	bodyBytes, _ := ioutil.ReadAll(responseBody)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Empty(t, bodyBytes)
}

func TestAuthenticateRouteBadRequest(t *testing.T) {

	tests := []struct {
		name     string
		request  server.AuthenticateRequest
		reqError server.RequestError
	}{
		{
			name:     "Missing password field",
			request:  server.AuthenticateRequest{UserName: "irar"},
			reqError: server.RequestError{Message: "Field password is required"},
		},
		{
			name:     "Missing userName field",
			request:  server.AuthenticateRequest{Password: "irar"},
			reqError: server.RequestError{Message: "Field userName is required"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			s := server.Server{Router: server.Router{}}

			s.Init()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("POST", "/api/auth", jsonBody(tt.request))

			s.Server.ServeHTTP(w, req)

			resBody := server.RequestError{}

			responseBody(w, &resBody)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, tt.reqError, resBody)
		})
	}
}
