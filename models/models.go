package models

import (
	"bytes"
	"net/http"
	"time"
)

// Message to Store and Send Data
type Message struct {
	UserID  string    `json:"userId,omitempty"`
	Message string    `json:"message,omitempty"`
	Time    time.Time `json:"time,omitempty"`
}

// Audit represents the structure of an audit
type Audit struct {
	RequestMethod    string        `json:"requestMethod,omitempty"`
	RequestURL       string        `json:"requestURL,omitempty"`
	RequestBody      string        `json:"requestBody,omitempty"`
	RequestHeaders   string        `json:"requestHeaders,omitempty"`
	QueryParameters  string        `json:"queryParameters,omitempty"`
	RequestTime      time.Time     `json:"requestTime,omitempty"`
	UserID           string        `json:"userID,omitempty"`
	ResponseBody     string        `json:"responseBody,omitempty"`
	ResponseHeaders  string        `json:"responseHeaders,omitempty"`
	StatusCode       int           `json:"statusCode,omitempty"`
	ResponseTime     time.Time     `json:"responseTime,omitempty"`
	ResponseDuration time.Duration `json:"responseDuration,omitempty"`
	RequestID        string        `json:"requestID,omitempty"`
	Version          string        `json:"version,omitempty"`
	RequestSize      int64         `json:"requestSize,omitempty"`
	ResponseSize     int64         `json:"responseSize,omitempty"`
}

type Clients struct {
	Messages []Message `json:"messages,omitempty"`
	Ch       chan Message `json:"ch,omitempty"`
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int          `json:"statusCode,omitempty"`
	Body       *bytes.Buffer `json:"body,omitempty"`
}

func (rw *CustomResponseWriter) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *CustomResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type PingResponse struct {
	Message     string    `json:"message,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
}

type ServerVersion struct {
	Version     string    `json:"Version,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
}

type JoinClientResponse struct {
	ID           string    `json:"userId,omitempty"`
	Message      string    `json:"message,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
}

type LeaveClientResponse struct {
	ID           string    `json:"userId,omitempty"`
	Message      string    `json:"message,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
}

type SendMessageResponse struct {
	ID           string    `json:"userId,omitempty"`
	Message      string    `json:"message,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
}

type GetMessagesResponse struct {
	Messages    []*Message `json:"messages,omitempty"`
	Message     string     `json:"message,omitempty"`
	ResponseTime time.Time `json:"ReponseTime,omitempty"`
	ID           string    `json:"userId,omitempty"`
}
