package service_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/guruakashsm/GoatRobotics/models"
	"github.com/guruakashsm/GoatRobotics/service"
	"github.com/guruakashsm/GoatRobotics/utils"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.Ping)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.PingResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Pinged Successfully", response.Message)
}

func TestServerVersion(t *testing.T) {
	version := utils.GetServerVersion()

	req, err := http.NewRequest("GET", "/version", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ServerVersion)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.ServerVersion
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, version, response.Version)
}

func TestJoinClient(t *testing.T) {
	chatRoom := service.NewChatRoom()
	go chatRoom.Run()

	req, err := http.NewRequest("GET", "/join?id=user123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatRoom.JoinClient)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.JoinClientResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "user123", response.ID)
	assert.Equal(t, "Joined Chat Successfully", response.Message)
}

func TestJoinClientMissingID(t *testing.T) {
	chatRoom := service.NewChatRoom()
	go chatRoom.Run()

	req, err := http.NewRequest("GET", "/join", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatRoom.JoinClient)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "ID_REQUIRED")
}

func TestLeaveClient(t *testing.T) {
	chatRoom := service.NewChatRoom()
	go chatRoom.Run()

	chatRoom.Register <- "user123"

	req, err := http.NewRequest("GET", "/leave?id=user123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatRoom.LeaveClient)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.LeaveClientResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "user123", response.ID)
	assert.Equal(t, "Left Chat Successfully", response.Message)
}

func TestSendMessage(t *testing.T) {
	chatRoom := service.NewChatRoom()
	chatRoom.Clients["user123"] = true

	req, err := http.NewRequest("GET", "/send?id=user123&message=Hello", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatRoom.SendMessage)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.SendMessageResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "user123", response.ID)
	assert.Equal(t, "Message Sent Successfully", response.Message)
}

func TestGetMessages(t *testing.T) {
	chatRoom := service.NewChatRoom()
	chatRoom.Clients["user123"] = true
	chatRoom.Messages = append(chatRoom.Messages, &models.Message{
		UserID:  "user123",
		Message: "Hello",
		Time:    time.Now(),
	})

	req, err := http.NewRequest("GET", "/messages?id=user123", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatRoom.GetMessages)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.GetMessagesResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "user123", response.ID)
	assert.NotEmpty(t, response.Messages)
}
