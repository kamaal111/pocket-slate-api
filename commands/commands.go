package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var pocketSlateAPICommand = &cobra.Command{
		Use:   "pocket-slate-api",
		Short: "Pocket Slate API CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	pocketSlateAPICommand.AddCommand(APIKey())
	err := pocketSlateAPICommand.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
