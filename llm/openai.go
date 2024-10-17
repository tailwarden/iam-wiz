package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/tailwarden/iam-wiz/policy"
)

func readPromptFromFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %v", err)
	}
	return string(data), nil
}

// GenerateIAMPolicy uses OpenAI to generate an IAM policy based on a prompt
func GenerateIAMPolicy(prompt string) (policy.IAMPolicy, error) {
	systemPrompt, err := readPromptFromFile("config/prompt.txt") // Make sure the file path is correct
	if err != nil {
		return policy.IAMPolicy{}, err
	}

	// Initialize OpenAI client with your API key
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Create a request with the provided prompt
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4oMini, // Choose the model you want to use
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})

	if err != nil {
		return policy.IAMPolicy{}, fmt.Errorf("failed to generate policy: %v", err)
	}

	// Assume the response is a JSON-structured IAM policy
	var iamPolicy policy.IAMPolicy
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &iamPolicy)
	if err != nil {
		return policy.IAMPolicy{}, fmt.Errorf("failed to unmarshal policy: %v", err)
	}

	return iamPolicy, nil
}
