package main

import (
	"bufio"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const COPY_TO_ENV_ARGUMENTS_EXPECTED = 2

func APIKey() *cobra.Command {
	var command = &cobra.Command{
		Use:   "api-key",
		Short: "API key commands",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return errIncorrectUsage
		},
	}

	command.AddCommand(makeAPIKeyCommand)
	command.AddCommand(copyToEnvCommand)

	return command
}

var copyToEnvCommand = &cobra.Command{
	Use:   "copy-to-env",
	Short: "Copy api keys to env",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != COPY_TO_ENV_ARGUMENTS_EXPECTED {
			return fmt.Errorf("expected %d arguments but got %d instead", COPY_TO_ENV_ARGUMENTS_EXPECTED, len(args))
		}

		secretsFile, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		var secretsMap map[string]map[string]string
		err = json.Unmarshal(secretsFile, &secretsMap)
		if err != nil {
			return err
		}

		secretsJSON, err := json.Marshal(secretsMap)
		if err != nil {
			return err
		}

		return writeToEnvFile(args[1], "APP_API_KEYS", string(secretsJSON))

	},
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

func writeToEnvFile(envFilePath string, key string, value string) error {
	envFile, err := os.Open(envFilePath)
	if err != nil {
		return err
	}
	defer envFile.Close()

	envFileScanner := bufio.NewScanner(envFile)
	var envFileContent []string
	var envTargetLine *int
	var lineNumber int
	for envFileScanner.Scan() {
		text := envFileScanner.Text()
		if strings.HasPrefix(text, fmt.Sprintf("%s=", key)) && envTargetLine == nil {
			envTargetLine = &lineNumber
		}
		envFileContent = append(envFileContent, text)
		if envTargetLine == nil {
			lineNumber += 1
		}
	}
	err = envFileScanner.Err()
	if err != nil {
		return err
	}

	envTargetLineContent := fmt.Sprintf("%s='%s'", key, value)
	if envTargetLine != nil {
		envFileContent[*envTargetLine] = envTargetLineContent
	} else {
		envFileContent = append(envFileContent, envTargetLineContent)
	}

	return os.WriteFile(".env", []byte(strings.Join(envFileContent, "\n")), 0777)
}
