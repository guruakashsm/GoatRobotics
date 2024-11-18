package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/guruakashsm/GoatRobotics/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
)

func ChooseLogLevel() string {
	loglvl := viper.GetString("Logging.level")
	if loglvl == "" {
		log.Logger.Info().Msgf("Log Level not chosen, defaulting to: %v", constants.Info)
		return constants.Info
	}
	log.Logger.Info().Msgf("Chosen Log Level: %v", loglvl)
	return loglvl
}

func ChoosePort() int {
	portNo := viper.GetInt("Port")
	if portNo == 0 {
		log.Logger.Info().Msg("Port not mentioned in config file. Defaulting to the constant value.")
		return constants.Port
	}
	log.Logger.Info().Msgf("Chosen Port: %v", portNo)
	return portNo
}

func ChooseHostName() string {
	hostname := viper.GetString("Host")
	if hostname == "" {
		log.Logger.Info().Msg("Host not mentioned in config file. Defaulting to the constant value.")
		return constants.Host
	}
	log.Logger.Info().Msgf("Chosen Host: %v", hostname)
	return hostname
}

func ChooseBaseURL() string {
	baseURL := viper.GetString("BaseURL")
	if baseURL == "" {
		log.Logger.Info().Msg("BaseURL not mentioned in config file. Defaulting to the constant value.")
		return constants.BaseURL
	}
	log.Logger.Info().Msgf("Chosen BaseURL: %v", baseURL)
	return baseURL
}

func GenerateRequestID() string {
	now := time.Now()
	datePart := now.Format(time.DateOnly)
	timePart := now.Format(time.Kitchen)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length := rand.Intn(6) + 5
	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		sb.WriteByte(charset[randomIndex])
	}
	randomID := sb.String()

	requestID := fmt.Sprintf("%s%s%s%s", viper.GetString("Appcode"), datePart, timePart, randomID)
	log.Logger.Info().Msgf("Generated Request ID: %v", requestID)
	return requestID
}

func GetServerVersion() string {
	version := viper.GetString("Version")
	if version == "" {
		log.Logger.Info().Msg("Version not mentioned in config file. Defaulting to the constant value.")
		return constants.Version
	}
	log.Logger.Info().Msgf("Server Version: %v", version)
	return version
}

func GetMaxMessage() int {
	maxMessage := viper.GetInt("MaxMessage")
	if maxMessage == 0 {
		log.Logger.Info().Msg("MaxMessage not mentioned in config file. Defaulting to the constant value.")
		return constants.MaxMessage
	}
	log.Logger.Info().Msgf("MaxMessage: %v", maxMessage)
	return maxMessage
}

func GetAuditFilePath() string {
	auditFilePath := viper.GetString("AuditFilePath")
	if auditFilePath == "" {
		log.Logger.Info().Msg("AuditFilePath not mentioned in config file. Defaulting to the constant value.")
		return constants.AuditFilePath
	}
	log.Logger.Info().Msgf("AuditFilePath: %v", auditFilePath)
	return auditFilePath
}
