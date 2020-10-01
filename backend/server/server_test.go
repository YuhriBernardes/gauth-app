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

func TestAuthenticateRoute(t *testing.T) {

	s := server.Server{Router: server.Router{}}

	s.Init()

	w := httptest.NewRecorder()

	requestBody := server.AuthenticateRequest{UserName: "Yurhi", Password: "Bernards"}

	req, _ := http.NewRequest("POST", "/api/auth", jsonBody(requestBody))

	s.Server.ServeHTTP(w, req)

	var body server.AuthenticateResponse

	responseBody(w, &body)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, len(body.Token), sha512.Size*2)
}
