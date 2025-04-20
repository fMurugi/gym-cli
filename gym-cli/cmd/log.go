package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var duration,date,note string

var logCmd = &cobra.Command{
	Use : "log [type] [distance]",
	Short: "Log a  new workout",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		workout :=map[string]string{
			"type": args[0],
			"distance": args[1],
			"duration": duration,
			"date": date,
			"note": note,
		}

		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(workout).
			Post("http://localhost:8080/workouts")

		if err != nil{
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("workout logged:", resp.String())
	},
}

func init(){
	logCmd.Flags().StringVar(&duration, "duration", "", "Workout duration")
	logCmd.Flags().StringVar(&date, "date", "", "Workout date (yyyy-mm-dd)")
	logCmd.Flags().StringVar(&note, "note", "", "Optional note")
	rootCmd.AddCommand(logCmd)
}