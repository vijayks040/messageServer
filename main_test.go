package main

import (
	//"bytes"
	//"fmt"
	//"context"
	//"encoding/json"
	"messageServer/utility"
	//	"net"
	"net/http"

	"net/http/httptest"
	"testing"

	//. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func loadMap() {
	mapLocker.dbMap[1] = utility.Messages{
		Message:     "hi hi",
		Description: "hello",
		User:        "vinay",
	}
	mapLocker.dbMap[2] = utility.Messages{
		Message:     "madam",
		Description: "hello hi",
		User:        "vijay",
	}
}

// func TestStartHTTPServer(t *testing.T) {
// 	Convey("Test function: StartHTTPServer.", t, func() {
// 		Convey("Normal Case1: start http server.", func() {
// 			patches := ApplyFunc(http.ListenAndServe, func(addr string, handler http.Handler) error {
// 				return nil
// 			})
// 			defer patches.Reset()
// 			StartHTTPServer()
// 		})
// 	})
// }

//mocking http server for positive case
// func MockHMServer(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)

// 	messageResponse := []utility.Messages{utility.Messages{
// 		Message:     "hi hi",
// 		Description: "hello",
// 		User:        "vinay",
// 	}, utility.Messages{
// 		Message:     "madam",
// 		Description: "hello",
// 		User:        "vijay",
// 	}}
// 	messageByteResponse, _ := json.Marshal(messageResponse)
// 	w.Write(messageByteResponse)

// }

//test function to check pass scenario
func TestGetMessage(t *testing.T) {
	loadMap()
	assert.NotEmpty(t, mapLocker.dbMap)
}

//test function to check pass scenario
func TestGetOneMessageMap(t *testing.T) {
	mes, _ := GetOneMessageMap(1)
	assert.Equal(t, mes.Message, "hi hi")
}

//func fail test
func TestDeleteOneMessageMap(t *testing.T) {
	err := DeleteOneMessageMap(3)
	assert.Error(t, err)
}

//func pass test
func TestDeleteOneMessageMapPass(t *testing.T) {
	err := DeleteOneMessageMap(1)
	assert.NoError(t, err)
}

func TestGetOneMessageHandlerFail(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/getOneMessage?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOneMessage)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 404)

}
func TestGetOneMessageHandlerpass(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/getOneMessage?id=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOneMessage)
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{
		Message:     "madam",
		Description: "hello hi",
		User:        "vijay",
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
