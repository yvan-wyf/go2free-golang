package openai

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type OpenAIRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      int      `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             int      `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// openai`s api for Q&A, only support EN
func ChatAPI(question string) ChatResponse {
	host := "https://api.openai.com/v1/completions"

	params := OpenAIRequest{
		Model:            "text-davinci-003",
		Prompt:           "Q:" + question + "\n" + " A:",
		Temperature:      0,
		MaxTokens:        100,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
		Stop:             []string{"\n"},
	}
	jsons, _ := json.Marshal(params)
	println("request:", string(jsons))
	bodyStr := strings.NewReader(string(jsons))

	// 发送 HTTP GET 请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", host, bodyStr)
	if err != nil {
		// handle error
		return ChatResponse{}
	}

	// 设置HTTP请求的header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return ChatResponse{}
	}

	defer resp.Body.Close()

	// 读取请求返回的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := ChatResponse{}

	json.Unmarshal(body, &response)

	json, _ := json.Marshal(response)
	println(string(json))

	return response
}
