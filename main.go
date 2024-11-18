package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/guruakashsm/GoatRobotics/constants"
	"github.com/guruakashsm/GoatRobotics/handlers"
	"github.com/guruakashsm/GoatRobotics/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	figure.NewFigure("GOAT ROBOTICS", "", true).Print()
	ReadConfig()
	LoadLogConfiguration()
}

func main() {
	handlers.Handle()
	ListenAndServe(LoadServerConfiguration())
}

// Read config.json file for Configurations
func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	log.Logger.Info().Msg("Reading config.json ...")
	if err := viper.ReadInConfig(); err != nil {
		log.Logger.Err(err).Msg("Error reading config file")
	}
}

// Loading Log Congiurations
func LoadLogConfiguration() {
	log.Logger.Info().Msg("Loading Log Configuration ...")

	logLevel := constants.LogMap[utils.ChooseLogLevel()]
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var multiLog zerolog.LevelWriter
	if viper.GetBool("Logging.storeLogs") {
		logFile := viper.GetString("Logging.file")
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to open log file")
		}
		log.Logger.Info().Msgf("Logs will also stores in Path : %v", logFile)
		multiLog = zerolog.MultiLevelWriter(os.Stderr, file)
	}

	log.Logger = zerolog.New(multiLog).Level(logLevel)
	log.Logger.Info().Msg("Log Configuration loaded Successfully ....")
}

// Loading Server Configuration
func LoadServerConfiguration() string {
	log.Logger.Info().Msg("Loading Server Configuation ..")
	serverDetails := fmt.Sprintf("%v:%v", utils.ChooseHostName(), utils.ChoosePort())
	log.Logger.Info().Msgf("Loaaded Server Configuation Successfully : %v", serverDetails)
	return serverDetails
}

// Listen & Serve and GraceFul ShutDown
func ListenAndServe(configuration string) {

	server := &http.Server{
		Addr:    configuration,
		Handler: http.DefaultServeMux,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msgf("Chat server running on %v", configuration)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	<-signalChan
	log.Info().Msg("Received shutdown signal, initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server Shutdown failed")
	} else {
		log.Info().Msg("Server gracefully shut down")
	}
}
