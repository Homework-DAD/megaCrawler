package megaCrawler

import (
	"encoding/json"
	"net/http"
)

type errorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func contain[T comparable](slice []T, check T) bool {
	for _, a := range slice {
		if a == check {
			return true
		}
	}
	return false
}

func mapToKeySlice[T any](m map[string]T) (slice []string) {
	for s, _ := range m {
		slice = append(slice, s)
	}
	return
}

func successResponse(msg string) (b []byte, err error) {
	errorJson := errorResp{
		StatusCode: 200,
		Message:    msg,
	}
	return json.Marshal(errorJson)
}

func errorResponse(w http.ResponseWriter, statusCode int, msg string) (err error) {
	errorJson := errorResp{
		StatusCode: statusCode,
		Message:    msg,
	}
	b, err := json.Marshal(errorJson)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	err = L.Error(msg)
	if err != nil {
		return err
	}
	return nil
}
