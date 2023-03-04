package chatgpt

import "net/http"

type ClientInterface interface {
	Chat(message string) (string, error)
}

type Client struct {
	host string
	cli  *http.Client
	cfg  *config
}

const host = "https://api.openai.com"

type Model string

const (
	GPT35Turbo     Model = "gpt-3.5-turbo"
	GPT35Turbo0301 Model = "gpt-3.5-turbo-0301"
)

type role string

const (
	user      role = "user"
	assistant role = "assistant"
)

type config struct {
	apiKey string
}

type chatGPTRequest struct {
	Model    Model            `json:"model"`
	Messages []messageContent `json:"messages"`
}

type messageContent struct {
	Role    role   `json:"role"`
	Content string `json:"content"`
}

type chatGPTResponse struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Usage   usageResponse    `json:"usage"`
	Choices []choiceResponse `json:"choices"`
	Error   errorResponse    `json:"error"`
}

type usageResponse struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type choiceResponse struct {
	Message      messageContent `json:"message"`
	FinishReason string         `json:"finish_reason"`
	Index        int32          `json:"index"`
}

type errorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
