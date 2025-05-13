/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	build "github.com/vassilidzuba/yacicctl/cmd/build"
	config "github.com/vassilidzuba/yacicctl/cmd/config"
	project "github.com/vassilidzuba/yacicctl/cmd/project"
	step "github.com/vassilidzuba/yacicctl/cmd/step"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yacicctl",
	Short: "client application for service yacic",
	Long:  `client application for service yacic.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("No home directory")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(home, ".yacicctl"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config not found")
		} else {
			log.Println("config found but an error occurred")
		}
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(project.Cmd)
	rootCmd.AddCommand(build.Cmd)
	rootCmd.AddCommand(step.Cmd)
	rootCmd.AddCommand(config.Cmd)

	rootCmd.PersistentFlags().StringP("username", "u", "", "username, facultative if you have a ~/.yacicctl/config.json file")
	rootCmd.PersistentFlags().StringP("password", "p", "", "password, facultative if you have a ~/.yacicctl/config.json file")

	err := viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	if err != nil {
		log.Fatal(err)
	}

	err = viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	if err != nil {
		log.Fatal(err)
	}
}
