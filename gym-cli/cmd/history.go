/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List past workouts",

	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		var workouts[]map[string]interface{}

		_, err := client.R().
			SetResult(&workouts).
			Get("http://localhost:8080/workouts")
		
			if err != nil{
				fmt.Println("Error:", err)
				return
			}
		fmt.Println("Workout History:")

		for _, w := range workouts {
			fmt.Printf("- %s: %s, %s, %s, Note: %s (ID: %s)\n",
				w["date"], w["type"], w["distance"], w["duration"], w["note"], w["id"])
		}
		
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

}
