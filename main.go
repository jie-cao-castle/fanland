package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var userLicense string

var appCmd = &cobra.Command{
	Use:   "fanland",
	Short: "Start fanland server",
	Long:  "Start fanland server",
	Run: func(cmd *cobra.Command, args []string) {
		app := &App{}
		app.Start(cmd, args)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	appCmd.PersistentFlags().StringVar(&cfgFile, "server", "", "config file (default is $HOME/config/server.yaml)")
	appCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	appCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	appCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", appCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", appCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config/server")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	if err := appCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
