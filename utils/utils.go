package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/GURUAKASHSM/ChatApp/constants"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
)

func ChooseLogLevel() string {
	loglvl := viper.GetString("Logging.level")
	if loglvl == "" {
		log.Logger.Info().Msgf("Log Level not Choosen So Default is Choosen : %v", constants.Info)
		return constants.Info
	}
	log.Logger.Info().Msgf("Choosen Log Level : %v", loglvl)
	return loglvl
}

func ChoosePort() int {
	portNo := viper.GetInt("Port")
	if portNo == 0 {
		log.Logger.Info().Msg("Port not Mentioned in config file .")
		return constants.Port
	}

	log.Logger.Info().Msgf("Port Choosen : %v", portNo)
	return portNo
}

func ChooseHostName() string {
	hostname := viper.GetString("Host")
	if hostname == "" {
		log.Logger.Info().Msg("Host not Mentioned in config file .")
		return constants.Host
	}

	log.Logger.Info().Msgf("Host Choosen : %v", hostname)
	return hostname
}

func ChooseBaseURL() string {
	baseURL := viper.GetString("BaseURL")
	if baseURL == "" {
		log.Logger.Info().Msg("BaseURL not Mentioned in config file .")
		return constants.BaseURL
	}
	log.Logger.Info().Msgf("BaseURL Choosen : %v", baseURL)
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

	return fmt.Sprintf("%s%s%s%s", viper.GetString("Appcode"), datePart, timePart, randomID)
}

func GetServerVersion()string{
	version := viper.GetString("Version")
	if version == "" {
		return constants.Version
	}
	return version
}


func GetMaxMessage()int{
	maxMessage := viper.GetInt("MaxMessage")
	if maxMessage == 0{
		return constants.MaxMessage
	}
	return maxMessage
}