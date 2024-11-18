package utils

import (
	"fmt"
	"net/http"
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

	log.Logger.Println(audit)

	log.Logger.Info().Msg("Audit data successfully written to file.")
}
