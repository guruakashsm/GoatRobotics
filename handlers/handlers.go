package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/GURUAKASHSM/ChatApp/models"
	"github.com/GURUAKASHSM/ChatApp/service"
	"github.com/GURUAKASHSM/ChatApp/utils"
	"github.com/rs/zerolog/log"
)

func Handle() {
	log.Logger.Info().Msg("Initialize routes ...")
	baseUrl := utils.ChooseBaseURL()

	chatRoom := service.NewChatRoom()
	go chatRoom.Run()

	http.Handle((baseUrl + "/join"), Middleware(http.HandlerFunc(chatRoom.JoinClient)))
	http.Handle((baseUrl + "/leave"), Middleware(http.HandlerFunc(chatRoom.LeaveClient)))
	http.Handle((baseUrl + "/send"), Middleware(http.HandlerFunc(chatRoom.SendMessage)))
	http.Handle((baseUrl + "/messages"), Middleware(http.HandlerFunc(chatRoom.GetMessages)))
	http.Handle("/ping", Middleware(http.HandlerFunc(service.Ping)))
	http.Handle("/version", Middleware(http.HandlerFunc(service.ServerVersion)))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		htmlPath := filepath.Join("UI", "index.html")
		http.ServeFile(w, r, htmlPath)
	})

	log.Logger.Info().Msg("Routes Initialized Successfully")

}

// Middleware logs the details of incoming requests
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		start := time.Now()
		requestBody, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		customResponseWriter := &models.CustomResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
			Body:           bytes.NewBuffer([]byte{}),
		}

		log.Logger.Info().Msgf("[AUDIT] Request started: %s %s", r.Method, r.URL.Path)
		log.Logger.Info().Msgf("[AUDIT] Request headers: %v", r.Header)
		log.Logger.Info().Msgf("[AUDIT] Request body: %s", string(requestBody))

		next.ServeHTTP(customResponseWriter, r)

		duration := time.Since(start)

		go utils.Audit(r, string(requestBody), customResponseWriter, start)
		log.Logger.Info().Msgf("[AUDIT] Request completed: %s %s", r.Method, r.URL.Path)
		log.Logger.Info().Msgf("[AUDIT] Response status: %d", customResponseWriter.StatusCode)
		log.Logger.Info().Msgf("[AUDIT] Response body: %s", customResponseWriter.Body.String())
		log.Logger.Info().Msgf("[AUDIT] Request duration: %v", duration)

	})
}
