package main

import (
	"encoding/hex"
	"fmt"

	cryptoRand "crypto/rand"

	"github.com/spf13/cobra"
)

func APIKey() *cobra.Command {
	var command = &cobra.Command{
		Use:   "api-key",
		Short: "API key commands",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return errIncorrectUsage
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	command.AddCommand(makeAPIKeyCommand)

	return command
}

var makeAPIKeyCommand = &cobra.Command{
	Use:   "make",
	Short: "Make a API key",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKeyBytes := make([]byte, 32)
		_, err := cryptoRand.Read(apiKeyBytes)
		if err != nil {
			return err
		}

		apiKey := hex.EncodeToString(apiKeyBytes)
		fmt.Printf("API key: %s\n", apiKey)
		return nil
	},
}
