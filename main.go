package main

import (
	"context"
	"fmt"
	"os"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func main() {

	modelRunnerBaseUrl := os.Getenv("MODEL_RUNNER_BASE_URL")

	if modelRunnerBaseUrl == "" {
		panic("MODEL_RUNNER_BASE_URL environment variable is not set")
	}
	modelRunnerChatModel := os.Getenv("MODEL_RUNNER_MODEL")
	fmt.Println("Using Model Runner Chat Model:", modelRunnerChatModel)

	if modelRunnerChatModel == "" {
		panic("MODEL_RUNNER_MODEL environment variable is not set")
	}

	bob, err := agents.NewAgent("bob_agent",
		agents.WithDMR(modelRunnerBaseUrl),
		agents.WithParams(openai.ChatCompletionNewParams{
			Model:       modelRunnerChatModel,
			Temperature: openai.Opt(0.5),
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(os.Getenv("AGENT_INSTRUCTION")),
				openai.UserMessage("Tell mee a story about a Werewolf in the forest."),
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	_, err = bob.ChatCompletionStream(context.Background(), func(self *agents.Agent, content string, err error) error {
		fmt.Print(content)
		return nil
	})
	if err != nil {
		panic(err)
	}

}
