package app

import (
	"dddapib/internal/domain/service"
	"dddapib/internal/infrastructure/persistence"
	"dddapib/internal/infrastructure/transport/http"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

var (
	flagCfgFile string
	// add more program flag here

	rootCmd = &cobra.Command{
		// main entry of this application
		Run: func(cmd *cobra.Command, args []string) {
			initLogger()
			svc := service.NewService(persistence.NewPersistence())
			http.NewServer(svc).ListenAndServe()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&flagCfgFile, "config", "", "toml config file (default is ./config.toml)")
}

func initLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func initConfig() {
	if flagCfgFile != "" {
		viper.SetConfigFile(flagCfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("fail to read config file, reaosn: %s", err))
	}
}

func Run() error {
	return rootCmd.Execute()
}
