package main

import (
	//"bytes"
	//"fmt"
	//"context"
	//"encoding/json"
	"messageServer/utility"
	//	"net"
	//"net/http"

	//"net/http/httptest"
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
	assert.NotEmpty(t, mapLocker.dbMap[1])
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
