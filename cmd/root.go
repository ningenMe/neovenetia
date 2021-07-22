package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ningenme/neovenetia/pkg/version"
)

var rootCmd = &cobra.Command{
	Use:   "neovenetia",
	Short: "Neovenetia is a validator of package config",
	Long: `Neovenetia is a cli application that manages package config, directory structures`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			version.Main()
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("version", "v", false, "An alias for the `version` subcommand.")
}

func initConfig() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
