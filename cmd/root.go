package main

import (
	"os"

	"github.com/senadi/limiter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var configPath string

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", configPath, "TCP address to listen on")
}

var rootCmd = &cobra.Command{
	Use:   limiter.Name,
	Short: limiter.Description,
	Long:  limiter.LongDescription,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Configure some basic logging
		// TODO: log.go - extract proper logging configuraion
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)

		// Read the config file
		if configPath == "" {
			log.Error("Missing config file!")
			os.Exit(1)
		}
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Error("Error reading config file", err)
			os.Exit(1)
		}
		// Confirm which config file is used
		log.WithField("cfg", viper.ConfigFileUsed()).Info("Config file loaded")

	},
}
