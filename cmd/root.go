/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	godaniel "github.com/49pctber/godaniel/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godaniel",
	Short: "Prints some ✨good vibes✨ to the console",
	Long: `Prints some ✨good vibes✨ to the console.
You can customize the name using the --name flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}
		godaniel.DefaultName = name
		td := godaniel.GetTemplateData(name)
		godaniel.PrintAffirmations(td)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("name", "n", "Daniel", "the name of the person to affirm")
}
