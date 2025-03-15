// Package ai fornece funcionalidades para interação com APIs de IA
package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Provider define a interface para provedores de IA
type Provider interface {
	GenerateCommitMessage(changes string) (string, error)
}

// OpenAIProvider implementa a interface Provider para a API da OpenAI
type OpenAIProvider struct {
	APIKey string
}

// NewOpenAIProvider cria uma nova instância de OpenAIProvider
func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	return &OpenAIProvider{
		APIKey: apiKey,
	}
}

// GeminiProvider implementa a interface Provider para a API do Google Gemini
type GeminiProvider struct {
	APIKey string
}

// NewGeminiProvider cria uma nova instância de GeminiProvider
func NewGeminiProvider(apiKey string) *GeminiProvider {
	return &GeminiProvider{
		APIKey: apiKey,
	}
}

// Estruturas para a requisição à API da OpenAI e Gemini (compatível)
type openAIRequest struct {
	Model     string    `json:"model"`
	Messages  []message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando OpenAI
func (p *OpenAIProvider) GenerateCommitMessage(changes string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key não configurada)", nil
	}

	prompt := fmt.Sprintf(`Analise as seguintes mudanças no código e gere uma mensagem de commit concisa e significativa no formato convencional:

%s

A mensagem deve:
1. Seguir o formato de Conventional Commits (tipo: descrição)
2. Ser clara e descritiva
3. Ter no máximo 72 caracteres
4. Começar com um dos seguintes tipos: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Apenas responda com a mensagem de commit, sem explicações adicionais.`, changes)

	reqBody := openAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []message{
			{
				Role:    "system",
				Content: "Você é um assistente especializado em gerar mensagens de commit claras e significativas no formato de Conventional Commits.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 100,
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na API da OpenAI (status %d): %s", resp.StatusCode, body)
	}

	var openAIResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return "", err
	}

	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return openAIResp.Choices[0].Message.Content, nil
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando Gemini
func (p *GeminiProvider) GenerateCommitMessage(changes string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do Gemini não configurada)", nil
	}

	prompt := fmt.Sprintf(`Analise as seguintes mudanças no código e gere uma mensagem de commit concisa e significativa no formato convencional:

%s

A mensagem deve:
1. Seguir o formato de Conventional Commits (tipo: descrição)
2. Ser clara e descritiva
3. Ter no máximo 72 caracteres
4. Começar com um dos seguintes tipos: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Apenas responda com a mensagem de commit, sem explicações adicionais.`, changes)

	reqBody := openAIRequest{
		Model: "gemini-2.0-flash", // Usando o modelo do Gemini
		Messages: []message{
			{
				Role:    "system",
				Content: "Você é um assistente especializado em gerar mensagens de commit claras e significativas no formato de Conventional Commits.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 100,
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Usando o endpoint do Gemini com compatibilidade OpenAI
	req, err := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na API do Gemini (status %d): %s", resp.StatusCode, body)
	}

	var geminiResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return "", err
	}

	if len(geminiResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return geminiResp.Choices[0].Message.Content, nil
}
