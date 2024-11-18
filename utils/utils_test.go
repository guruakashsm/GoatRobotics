package utils_test

import (
	"testing"
	"time"

	"github.com/guruakashsm/GoatRobotics/constants" 
	"github.com/guruakashsm/GoatRobotics/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestChooseLogLevel(t *testing.T) {
	viper.Set("Logging.level", "")
	logLevel := utils.ChooseLogLevel()
	assert.Equal(t, constants.Info, logLevel)

	viper.Set("Logging.level", "debug")
	logLevel = utils.ChooseLogLevel()
	assert.Equal(t, "debug", logLevel)
}

func TestChoosePort(t *testing.T) {
	viper.Set("Port", 0)
	port := utils.ChoosePort()
	assert.Equal(t, constants.Port, port)

	viper.Set("Port", 8080)
	port = utils.ChoosePort()
	assert.Equal(t, 8080, port)
}

func TestChooseHostName(t *testing.T) {
	viper.Set("Host", "")
	host := utils.ChooseHostName()
	assert.Equal(t, constants.Host, host)

	viper.Set("Host", "example.com")
	host = utils.ChooseHostName()
	assert.Equal(t, "example.com", host)
}

func TestChooseBaseURL(t *testing.T) {
	viper.Set("BaseURL", "")
	baseURL := utils.ChooseBaseURL()
	assert.Equal(t, constants.BaseURL, baseURL)

	viper.Set("BaseURL", "http://localhost")
	baseURL = utils.ChooseBaseURL()
	assert.Equal(t, "http://localhost", baseURL)
}

func TestGenerateRequestID(t *testing.T) {
	viper.Set("Appcode", "APP")
	now := time.Now()
	datePart := now.Format("2006-01-02") // Format for `time.DateOnly`
	timePart := now.Format("3:04PM")     // Format for `time.Kitchen`

	requestID := utils.GenerateRequestID()
	assert.Contains(t, requestID, "APP")
	assert.Contains(t, requestID, datePart)
	assert.Contains(t, requestID, timePart)
	assert.True(t, len(requestID) > 20, "Request ID should be sufficiently long")
}

func TestGetServerVersion(t *testing.T) {
	viper.Set("Version", "")
	version := utils.GetServerVersion()
	assert.Equal(t, constants.Version, version)

	viper.Set("Version", "v1.0.0")
	version = utils.GetServerVersion()
	assert.Equal(t, "v1.0.0", version)
}

func TestGetMaxMessage(t *testing.T) {
	viper.Set("MaxMessage", 0)
	maxMessage := utils.GetMaxMessage()
	assert.Equal(t, constants.MaxMessage, maxMessage)

	viper.Set("MaxMessage", 100)
	maxMessage = utils.GetMaxMessage()
	assert.Equal(t, 100, maxMessage)
}

func TestGetAuditFilePath(t *testing.T) {
	viper.Set("AuditFilePath", "")
	auditFilePath := utils.GetAuditFilePath()
	assert.Equal(t, constants.AuditFilePath, auditFilePath)

	viper.Set("AuditFilePath", "/var/log/audit.log")
	auditFilePath = utils.GetAuditFilePath()
	assert.Equal(t, "/var/log/audit.log", auditFilePath)
}
