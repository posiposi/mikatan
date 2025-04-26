package openai

type OpenAICommunicator interface {
	SendBooks(prompt []*Prompt) (*ChatResponse, error)
}

type Role string

const (
	System    Role = "system"
	User      Role = "user"
	Assistant Role = "assistant"
)

type Prompt struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type openAIRequest struct {
	Model    string    `json:"model"`
	Messages []*Prompt `json:"messages"`
}

type openAIResponse struct {
	ID      string   `json:"id"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

type choice struct {
	Index        int    `json:"index"`
	Message      Prompt `json:"message"`
	FinishReason string `json:"finish_reason"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponse struct {
	Content string `json:"content"`
	Usage   usage  `json:"usage"`
}
