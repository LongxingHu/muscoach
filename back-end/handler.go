package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type RequestPayload struct {
	Lyrics string `json:"lyrics"`
}

func (app *Config) HandleLyrics(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Go")
	var requestPayload RequestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Printf("This is the request payload: %v", requestPayload)

	lyrics := ParseLyric(requestPayload.Lyrics)
	cleanedLyrics := strings.Join(lyrics, "\n")

	log.Println("This is the request prompt:")
	response := openAi(cleanedLyrics)
	fmt.Println("response:" + response)
	var payload jsonResponse
	payload.Message = response
	app.writeJSON(w, http.StatusOK, payload)
}

func ParseLyric(lyric string) (sentences []string) {
	//对一首歌原文进行解析，分出时间戳和歌词
	lyricSliceInOneSong := strings.Split(lyric, "\n")
	//fmt.Printf("%v", lyricSliceInOneSong)
	for _, sentenceInOneSong := range lyricSliceInOneSong {
		re := regexp.MustCompile(`^\[(\d{2}:\d{2}\.\d{2,3})\](.*)$`)
		match := re.FindSubmatch([]byte(sentenceInOneSong))
		if len(match) == 3 && len(match[2]) > 0 {
			// fmt.Println(string(match[1]))
			sentences = append(sentences, string(match[2]))
			// fmt.Println(string(match[2]))
		}
	}
	return sentences
}

func openAi(cleanedLyrics string) string {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "请逐句深度解析下面的歌词(中文歌词不用翻译)：\n" + cleanedLyrics,
				},
			},
		},
	)

	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	log.Println("This is the response:")
	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}
