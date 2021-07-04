package cmd

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/server"
	"github.com/kwong21/graphql-go-cardkeeper/service"
	"github.com/spf13/cobra"
)

var (
	configFile string
	config     models.Config
	rootCmd    = &cobra.Command{
		Use:   "cardkeeper-server",
		Short: "GraphQL server for Cardkeeper",
		Long: `This is the GraphQL server for maintaining players and teams to track cards for. 
		The Cardkeeper-Server is an API service for UIs and other tools to communicate with`,
		Run: func(cmd *cobra.Command, args []string) {
			service := service.New(config)
			server.Init(service)
		},
	}
)

// Execute the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConifg)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file")
	rootCmd.MarkPersistentFlagRequired("config")
}

func initConifg() {
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatalf("Unable to read configuration file: %s", err)
	}
}
