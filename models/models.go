package models

import (
	"bytes"
	"net/http"
	"time"
)

// Message to Store and Send Data
type Message struct {
	UserID  string    `json:"userId"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Read    bool      `json:"read"`
}

// Audit represents the structure of an audit
type Audit struct {
	RequestMethod    string        `json:"requestMethod"`
	RequestURL       string        `json:"requestURL"`
	RequestBody      string        `json:"requestBody"`
	RequestHeaders   string        `json:"requestHeaders"`
	QueryParameters  string        `json:"queryParameters"`
	RequestTime      time.Time     `json:"requestTime"`
	UserID           string        `json:"userID,omitempty"`
	ResponseBody     string        `json:"responseBody"`
	ResponseHeaders  string        `json:"responseHeaders"`
	StatusCode       int           `json:"statusCode"`
	ResponseTime     time.Time     `json:"responseTime"`
	ResponseDuration time.Duration `json:"responseDuration"`
	RequestID        string        `json:"requestID,omitempty"`
	Version          string        `json:"version,omitempty"`
	RequestSize      int64         `json:"requestSize,omitempty"`
	ResponseSize     int64         `json:"responseSize,omitempty"`
}

type Clients struct {
	Messages []Message
	Ch       chan Message
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func (rw *CustomResponseWriter) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *CustomResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type PingResponse struct{
	Message string `json:"message"`
	ResponseTime time.Time `json:"ReponseTime"`
}

type ServerVersion struct{
	Version string `json:"Version"`
	ResponseTime time.Time `json:"ReponseTime"`
}

type JoinClientResponse struct{
	ID string `json:"userId"`
	Message string `json:"message"`
	ResponseTime time.Time `json:"ReponseTime"`
}

type LeaveClientResponse struct{
	ID string `json:"userId"`
	Message string `json:"message"`
	ResponseTime time.Time `json:"ReponseTime"`
}

type SendMessageResponse struct{
	ID string `json:"userId"`
	Message string `json:"message"`
	ResponseTime time.Time `json:"ReponseTime"`
}

type GetMessagesResponse struct{
	Messages []Message `json:"messages"`
	Message string `json:"message"`
	ResponseTime time.Time `json:"ReponseTime"`
	ID string `json:"userId"`
}