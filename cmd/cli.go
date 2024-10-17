package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	openai "github.com/tailwarden/iam-wiz/llm"
)

// cliCmd is the main command for the CLI
var CliCmd = &cobra.Command{
	Use:   "generate-iam",
	Short: "Generate IAM policies using LLM based on user prompts",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			// Prompt the user for input
			fmt.Print("Enter your prompt to generate IAM policy (or type 'exit' to quit): ")
			prompt, _ := reader.ReadString('\n')
			prompt = strings.TrimSpace(prompt)

			// Exit if the user types 'exit'
			if strings.ToLower(prompt) == "exit" {
				fmt.Println("Exiting...")
				break
			}

			// Generate the IAM policy
			iamPolicy, err := openai.GenerateIAMPolicy(prompt)
			if err != nil {
				log.Fatalf("Error generating IAM policy: %v", err)
			}

			// Pretty print the generated policy
			policyJSON, err := iamPolicy.PrettyPrint()
			if err != nil {
				fmt.Println("Error formatting policy:", err)
				return
			}

			// Print the formatted JSON string
			fmt.Println(policyJSON)
		}
	},
}
