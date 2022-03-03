package cmd

import (
	"encoding/json"
	"fmt"
	"go-web-template/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var devMode bool
var AppCfg config.AppConfig

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "Start this gin server.",
	Long:  `todo`,
	Run: func(cmd *cobra.Command, args []string) {
		Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is .conf)")
	rootCmd.PersistentFlags().StringP("port", "p", "", "listen port (default is :3333)")
	rootCmd.PersistentFlags().BoolVarP(&devMode, "dev", "d", false, "development mode")
	viper.BindPFlag("app.port", rootCmd.PersistentFlags().Lookup("port"))

	// rootCmd.AddCommand(otherCmd)
}

func initConfig() {
	rootDir, err := os.Getwd()
	cobra.CheckErr(err)

	viper.AddConfigPath(rootDir)
	viper.SetConfigType("yaml")

	var cfgFileName string = "config"
	if cfgFile != "" {
		cfgFileName = cfgFile
	}
	if devMode {
		cfgFileName += ".dev"
	}

	viper.SetConfigName(cfgFileName)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

		} else {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

	}
	AppCfg.Bind(viper.GetViper())
	cfg, _ := json.Marshal(AppCfg)
	fmt.Println(string(cfg))
}
