/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
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
		http.HandleFunc("GET /", godaniel.GoDanielHandler)

		http.HandleFunc("GET /{name}", func(w http.ResponseWriter, r *http.Request) {
			var td godaniel.TemplateData
			rname := godaniel.RemoveNonLetters(r.PathValue("name"))

			if len(rname) != 0 {
				// render template for name
				td = godaniel.GetTemplateData(rname)
				tmpl, err := template.ParseFS(godaniel.Templates, "static/base.html", "static/godaniel.html")
				if err != nil {
					panic(err)
				}
				tmpl.Execute(w, td)
			} else {
				// get name
				tmpl, err := template.ParseFS(godaniel.Templates, "static/base.html", "static/getname.html")
				if err != nil {
					panic(err)
				}
				tmpl.Execute(w, nil)
			}
		})

		http.HandleFunc("GET /json/{name}", func(w http.ResponseWriter, r *http.Request) {
			var td godaniel.TemplateData
			rname := godaniel.RemoveNonLetters(r.PathValue("name"))

			if len(rname) != 0 {
				td = godaniel.GetTemplateData(rname)
			} else {
				td = godaniel.GetTemplateData(godaniel.DefaultName)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(td)
		})

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
