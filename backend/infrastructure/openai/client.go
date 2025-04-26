package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenAIClient struct {
	apiKey     string
	model      string
	url        string
	httpClient *http.Client
}

func NewOpenAIClient(apiKey string, model string, httpClient *http.Client) OpenAICommunicator {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &OpenAIClient{
		apiKey:     apiKey,
		model:      model,
		url:        "https://api.openai.com/v1/chat/completions",
		httpClient: httpClient,
	}
}

func (c *OpenAIClient) SendBooks(prompts []*Prompt) (*ChatResponse, error) {
	reqBody := &openAIRequest{
		Model:    c.model,
		Messages: prompts,
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %v: %s", resp.StatusCode, body)
	}

	var openAIResp openAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, err
	}
	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("API response does not contain any choices")
	}

	chatResp := &ChatResponse{
		Content: openAIResp.Choices[0].Message.Content,
		Usage:   openAIResp.Usage,
	}

	return chatResp, nil
}
