package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	openai "github.com/sashabaranov/go-openai"
// )

// func main() {
// 	client := openai.NewClient("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")
// 	resp, err := client.CreateChatCompletion(
// 		context.Background(),
// 		openai.ChatCompletionRequest{
// 			Model: openai.GPT3Dot5Turbo,
// 			Messages: []openai.ChatCompletionMessage{
// 				{
// 					Role:    openai.ChatMessageRoleUser,
// 					Content: "请你解释一下docker的原理",
// 				},
// 			},
// 		},
// 	)

// 	if err != nil {
// 		fmt.Printf("Error: %v", err)
// 	}
// 	log.Println("This is the response:")
// 	fmt.Println(resp)
// }
