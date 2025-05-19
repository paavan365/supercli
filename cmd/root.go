package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "supercli",
	Short: "Automates Infra login and K8s log fetching",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loginCmd) // <- Don't forget this line
}
