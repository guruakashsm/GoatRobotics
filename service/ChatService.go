package service

import (
	"context"
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
    Broadcast  chan models.Message
    Register   chan string
    Unregister chan string
    Mu         sync.RWMutex
    Clients    map[string]bool 
    Messages   []*models.Message
    msgMu      sync.Mutex 
}

// NewChatRoom creates a new chat room
func NewChatRoom() *ChatRoom {
    return &ChatRoom{
        Broadcast:  make(chan models.Message, 100), 
        Register:   make(chan string),
        Unregister: make(chan string),
        Clients:    make(map[string]bool),
        Messages:   make([]*models.Message, 0), 
    }
}

// Run starts the chat room's main loop
func (c *ChatRoom) Run() {
    for {
        select {
        case id := <-c.Register:
            log.Logger.Info().Msgf("Registering client with ID: %v", id) // Log when a client is registered
            c.Mu.Lock()
            c.Clients[id] = true // Register the client
            c.Mu.Unlock()

        case id := <-c.Unregister:
            log.Logger.Info().Msgf("Unregistering client with ID: %v", id) // Log when a client is unregistered
            c.Mu.Lock()
            if exists := c.Clients[id]; exists {
                delete(c.Clients, id) // Unregister the client
            }
            c.Mu.Unlock()

        case message := <-c.Broadcast:
            log.Logger.Info().Msgf("Broadcasting message: %v", message) // Log the broadcasted message
            c.msgMu.Lock()
            c.Messages = append(c.Messages, &message)
            c.msgMu.Unlock()
        }
    }
}

// Ping to Check the Server Status
func Ping(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("Received ping request") // Log when the ping request is received
	response := &models.PingResponse{
		Message:      "Pinged Successfully",
		ResponseTime: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// To Check Server Version Currently Running
func ServerVersion(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("Received request for server version") // Log when the server version request is received
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
		log.Logger.Err(gError.ID_REQUIRED.Error()) // Log error if ID is missing
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	log.Logger.Info().Msgf("User Joined to Chat Room with ID: %v", id) // Log the user's join event

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
		log.Logger.Err(gError.ID_REQUIRED.Error()) // Log error if ID is missing
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gError.ID_REQUIRED.Error())
		return
	}

	log.Logger.Info().Msgf("User Left Chat Room with ID: %v", id) // Log the user's leave event

	c.Unregister <- id
	response := models.LeaveClientResponse{
		ID:           id,
		Message:      "Left Chat Successfully",
		ResponseTime: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SendMessage broadcasts a message from a client
func (c *ChatRoom) SendMessage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error()) // Log error if ID is missing
		http.Error(w, gError.ID_REQUIRED.Code, http.StatusBadRequest)
		return
	}

	message := strings.TrimSpace(r.URL.Query().Get("message"))
	if message == "" {
		log.Logger.Err(gError.MESSAGE_REQUIRED.Error()) // Log error if message is missing
		http.Error(w, gError.MESSAGE_REQUIRED.Code, http.StatusBadRequest)
		return
	}

	log.Logger.Info().Msgf("User with ID: %v sent message: %v", id, message) // Log the sent message

	// Send message to Broadcast channel
	c.Broadcast <- models.Message{UserID: id, Message: message, Time: time.Now()}

	response := models.SendMessageResponse{
		ID:           id,
		Message:      "Message Sent Successfully",
		ResponseTime: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetMessages retrieves all messages for a client.
func (c *ChatRoom) GetMessages(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.URL.Query().Get("id"))
	if id == "" {
		log.Logger.Err(gError.ID_REQUIRED.Error()) // Log error if ID is missing
		http.Error(w, gError.ID_REQUIRED.Code, http.StatusBadRequest)
		return
	}

	c.Mu.RLock()
	_, exists := c.Clients[id]
	c.Mu.RUnlock()
	if !exists {
		log.Logger.Err(gError.ID_REQUIRED.Error()).Msgf("User not found: %v", id) // Log error if the user is not found
		http.Error(w, gError.USER_NOT_FOUND.Code, http.StatusNotFound)
		return
	}

	log.Logger.Info().Msgf("Retrieving messages for client ID: %v", id) // Log message retrieval request

	c.Mu.RLock()
	messages := append([]*models.Message{}, c.Messages...) // Create a copy of messages to avoid race conditions
	c.Mu.RUnlock()

	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	responseChan := make(chan models.GetMessagesResponse)
	go func() {
		response := models.GetMessagesResponse{
			ID:           id,
			Messages:     messages,
			ResponseTime: time.Now(),
		}
		if len(messages) == 0 {
			response.Message = "No new messages"
		}
		responseChan <- response
	}()

	select {
	case response := <-responseChan:
		log.Logger.Info().Msgf("Sending messages response for client ID: %v", id) // Log message response sent
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	case <-ctx.Done():
		log.Logger.Err(gError.NO_MESSAGE_FOUND.Error()).Msgf("Request timed out while retrieving messages for client ID: %v", id) // Log timeout error
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	}
}
