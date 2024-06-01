/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	godaniel "github.com/49pctber/godaniel/internal"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the godaniel server",
	Long: `Start a server on localhost.
By default, the server runs on localhost:8052, but you may specify a port using --port.`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", godaniel.Handler)

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		godaniel.DefaultName = name

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}

		ps := fmt.Sprintf(":%d", port) // port string
		fmt.Printf("Serving on http://localhost%s\n", ps)

		panic(http.ListenAndServe(ps, nil))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().IntP("port", "p", 8052, "the port on which to run the server")
}
