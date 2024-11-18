package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	gError "github.com/guruakashsm/GoatRobotics/errors"
	"github.com/guruakashsm/GoatRobotics/models"
	"github.com/guruakashsm/GoatRobotics/utils"
	"github.com/rs/zerolog/log"
)

// ChatRoom represents the central chat room
type ChatRoom struct {
	Clients map[string]struct {
		Messages []models.Message // All messages received by the client (read and unread)
		Ch       chan models.Message
	}
	Mu         sync.RWMutex
	Broadcast  chan models.Message
	Register   chan string
	Unregister chan string
}

// NewChatRoom creates a new chat room
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		Clients: make(map[string]struct {
			Messages []models.Message
			Ch       chan models.Message
		}),
		Broadcast:  make(chan models.Message),
		Register:   make(chan string),
		Unregister: make(chan string),
	}
}

// Run starts the chat room's main loop
func (c *ChatRoom) Run() {
	for {
		select {
		case id := <-c.Register:
			c.Mu.Lock()
			c.Clients[id] = models.Clients{
				Messages: []models.Message{},
				Ch:       make(chan models.Message, utils.GetMaxMessage()),
			}
			c.Mu.Unlock()

		case id := <-c.Unregister:
			c.Mu.Lock()
			if _, exists := c.Clients[id]; exists {
				close(c.Clients[id].Ch)
				delete(c.Clients, id)
			}
			c.Mu.Unlock()

		case message := <-c.Broadcast:
			c.Mu.RLock()
			for id, client := range c.Clients {
				client.Messages = append(client.Messages, models.Message{
					UserID:  message.UserID,
					Message: message.Message,
					Time:    message.Time,
					Read:    false,
				})
				client.Ch <- message
				c.Clients[id] = client
			}
			c.Mu.RUnlock()
		}
	}
}

// Ping to Check the Server Status
func Ping(w http.ResponseWriter, r *http.Request) {
	response := &models.PingResponse{
		Message:      "Pinged Successfully",
		ResponseTime: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// To Cheack Server Version Currently Running
func ServerVersion(w http.ResponseWriter, r *http.Request) {
	var serverVersionRes = models.ServerVersion{
		Version:      utils.GetServerVersion(),
		ResponseTime: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serverVersionRes)
}

// JoinClient adds a client to the chat room
func (c *ChatRoom) JoinClient(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("********** JOIN CLIENT **********")

	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	log.Logger.Info().Msgf("User Joined to Chat Romm with ID : %v", id)

	c.Register <- id

	response := models.JoinClientResponse{
		ID:           id,
		Message:      "Joined Chat Successfully",
		ResponseTime: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LeaveClient removes a client from the chat room
func (c *ChatRoom) LeaveClient(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("********** LEAVE CLIENT **********")

	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	c.Unregister <- id
	response := models.LeaveClientResponse{
		ID:           id,
		Message:      "Left Chat Successfully",
		ResponseTime: time.Now(),
	}

	log.Logger.Info().Msgf("User Left Chat Romm with ID : %v", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SendMessage broadcasts a message from a client
func (c *ChatRoom) SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("********** SEND MESSAGE **********")

	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	message := strings.TrimSpace(r.URL.Query().Get("message"))
	if message == "" {
		log.Logger.Err(gError.MESSAGE_REQUIRED.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.MESSAGE_REQUIRED.Error())
		return
	}

	log.Logger.Info().Msgf("User with ID : %v Sends Message : %v to the Chat Room", id, message)

	c.Mu.RLock()
	_, exists := c.Clients[id]
	c.Mu.RUnlock()
	if !exists {
		log.Logger.Err(gError.USER_NOT_FOUND.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.USER_NOT_FOUND.Error())
		return
	}

	c.Broadcast <- models.Message{UserID: id, Message: message, Time: time.Now()}
	response := models.SendMessageResponse{
		ID:           id,
		Message:      message,
		ResponseTime: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetMessages retrieves all messages for a client (both read and unread) and marks them as read.
func (c *ChatRoom) GetMessages(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	c.Mu.Lock()
	client, exists := c.Clients[id]
	c.Mu.Unlock()
	if !exists {
		log.Logger.Err(gError.USER_NOT_FOUND.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.USER_NOT_FOUND.Error())
		return
	}

	log.Logger.Info().Msgf("User with ID : %v aked to Get all Messages", id)
	response := models.GetMessagesResponse{
		Messages:     client.Messages,
		ResponseTime: time.Now(),
		ID:           id,
	}

	if len(client.Messages) == 0 {
		response.Message = "No messages"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	c.Mu.Lock()
	for i := range client.Messages {
		client.Messages[i].Read = true
	}
	c.Clients[id] = client
	c.Mu.Unlock()
}
