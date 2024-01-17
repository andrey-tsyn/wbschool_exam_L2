package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

type errorResponse struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
}

func sendErrorResponse(w http.ResponseWriter, e errorResponse) {
	w.WriteHeader(e.StatusCode)
	setDefaultHeaders(w)

	response := struct {
		Error errorResponse `json:"error"`
	}{e}

	bytesResp, err := json.Marshal(response)
	if err != nil {
		log.Errorf("error occured while sending response: %+v, error: %s", response, err)
	}

	_, err = w.Write(bytesResp)
	if err != nil {
		log.Error(err)
	}
}

type successResponse struct {
	ObjectToSerialize any    `json:"obj"`
	Msg               string `json:"msg"`
}

func sendSuccessResponse(w http.ResponseWriter, s successResponse) {
	w.WriteHeader(http.StatusOK)
	setDefaultHeaders(w)

	response := struct {
		Result successResponse `json:"result"`
	}{s}

	bytesResp, err := json.Marshal(response)
	if err != nil {
		log.Errorf("error occured while sending s: %+v, error: %s", response, err)
	}

	_, err = w.Write(bytesResp)
	if err != nil {
		log.Error(err)
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	for k, v := range defaultHeaders {
		w.Header().Set(k, v)
	}
}
