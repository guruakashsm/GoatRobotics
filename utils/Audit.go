package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/guruakashsm/GoatRobotics/models"
	"github.com/rs/zerolog/log"
)

// Mutex to synchronize writes to the audit file
var fileMutex sync.Mutex

// Audit writes audit data to a JSON file
func Audit(r *http.Request, requestBody string, response *models.CustomResponseWriter, requestTime time.Time) {
	log.Logger.Info().Msg("Auditing the Request & Response")

	audit := models.Audit{
		RequestMethod:    r.Method,
		RequestURL:       r.RequestURI,
		RequestBody:      requestBody,
		RequestHeaders:   fmt.Sprintf("%v", r.Header),
		QueryParameters:  fmt.Sprintf("%v", r.URL.Query()),
		RequestTime:      requestTime.UTC(),
		UserID:           r.URL.Query().Get("id"),
		ResponseBody:     response.Body.String(),
		ResponseHeaders:  fmt.Sprintf("%v", response.Header()),
		StatusCode:       response.StatusCode,
		ResponseTime:     time.Now().UTC(),
		ResponseDuration: time.Since(requestTime),
		RequestID:        GenerateRequestID(),
		Version:          GetServerVersion(),
		RequestSize:      r.ContentLength,
		ResponseSize:     int64(response.Body.Len()),
	}

	auditData, err := json.MarshalIndent(audit, "", "  ")
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to marshal audit data")
		return
	}

	file, err := os.OpenFile(GetAuditFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to open audit file")
		return
	}
	defer file.Close()

	
	_, err = file.WriteString(string(auditData) + "\n")
	if err != nil {
		log.Logger.Error().Err(err).Msg("Failed to write audit data to file")
		return
	}

	log.Logger.Info().Msg("Audit data successfully written to audit.log file.")
}
