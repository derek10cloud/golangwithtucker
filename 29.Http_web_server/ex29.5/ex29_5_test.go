package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // 1./경로테스트

	mux := MakeWebHandler() //2.
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) // 3. 코드 확인
	data, _ := io.ReadAll(res.Body)       //4. 데이터를 읽어서 확인
	assert.Equal("Hello World", string(data))
}

func TesBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // /bar 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
