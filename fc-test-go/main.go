package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	requestPath := req.URL.Path
	requestURI := req.RequestURI
	requestClientIP := req.RemoteAddr

	result := fmt.Sprintf("Path: %s\n Uri: %s\n IP: %s\n", requestPath, requestURI, requestClientIP)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(result))
	return nil
}
