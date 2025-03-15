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
	GenerateCommitMessage(changes string, language string) (string, error)
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

// ClaudeProvider implementa a interface Provider para a API da Anthropic (Claude)
type ClaudeProvider struct {
	APIKey string
}

// NewClaudeProvider cria uma nova instância de ClaudeProvider
func NewClaudeProvider(apiKey string) *ClaudeProvider {
	return &ClaudeProvider{
		APIKey: apiKey,
	}
}

// DeepSeekProvider implementa a interface Provider para a API DeepSeek
type DeepSeekProvider struct {
	APIKey string
}

// NewDeepSeekProvider cria uma nova instância de DeepSeekProvider
func NewDeepSeekProvider(apiKey string) *DeepSeekProvider {
	return &DeepSeekProvider{
		APIKey: apiKey,
	}
}

// OpenRouterProvider implementa a interface Provider para a API OpenRouter
type OpenRouterProvider struct {
	APIKey string
}

// NewOpenRouterProvider cria uma nova instância de OpenRouterProvider
func NewOpenRouterProvider(apiKey string) *OpenRouterProvider {
	return &OpenRouterProvider{
		APIKey: apiKey,
	}
}

// GrokProvider implementa a interface Provider para a API Grok da xAI
type GrokProvider struct {
	APIKey string
}

// NewGrokProvider cria uma nova instância de GrokProvider
func NewGrokProvider(apiKey string) *GrokProvider {
	return &GrokProvider{
		APIKey: apiKey,
	}
}

// Estruturas para a requisição à API da OpenAI e Gemini (compatível)
type openAIRequest struct {
	Model     string    `json:"model"`
	Messages  []message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// Estrutura para a requisição à API da Anthropic (Claude)
type claudeRequest struct {
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

type claudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

// getLanguagePrompt retorna o prompt adequado para o idioma solicitado
func getLanguagePrompt(changes string, language string) string {
	switch language {
	case "en":
		return fmt.Sprintf(`Analyze the following code changes and generate a concise and meaningful commit message in the conventional format:

%s

The message should:
1. Follow the Conventional Commits format (type: description)
2. Be clear and descriptive
3. Have a maximum of 72 characters
4. Start with one of the following types: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Only respond with the commit message, without additional explanations.`, changes)
	case "es":
		return fmt.Sprintf(`Analiza los siguientes cambios en el código y genera un mensaje de commit conciso y significativo en formato convencional:

%s

El mensaje debe:
1. Seguir el formato de Conventional Commits (tipo: descripción)
2. Ser claro y descriptivo
3. Tener un máximo de 72 caracteres
4. Comenzar con uno de los siguientes tipos: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Solo responde con el mensaje de commit, sin explicaciones adicionales.`, changes)
	case "fr":
		return fmt.Sprintf(`Analysez les modifications de code suivantes et générez un message de commit concis et significatif au format conventionnel:

%s

Le message doit:
1. Suivre le format Conventional Commits (type: description)
2. Être clair et descriptif
3. Avoir un maximum de 72 caractères
4. Commencer par l'un des types suivants: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Répondez uniquement avec le message de commit, sans explications supplémentaires.`, changes)
	case "de":
		return fmt.Sprintf(`Analysieren Sie die folgenden Codeänderungen und generieren Sie eine prägnante und aussagekräftige Commit-Nachricht im konventionellen Format:

%s

Die Nachricht sollte:
1. Dem Format von Conventional Commits folgen (Typ: Beschreibung)
2. Klar und beschreibend sein
3. Maximal 72 Zeichen haben
4. Mit einem der folgenden Typen beginnen: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Antworten Sie nur mit der Commit-Nachricht, ohne zusätzliche Erklärungen.`, changes)
	default: // Padrão é português pt-br
		return fmt.Sprintf(`Analise as seguintes mudanças no código e gere uma mensagem de commit concisa e significativa no formato convencional:

%s

A mensagem deve:
1. Seguir o formato de Conventional Commits (tipo: descrição)
2. Ser clara e descritiva
3. Ter no máximo 72 caracteres
4. Começar com um dos seguintes tipos: feat, fix, docs, style, refactor, perf, test, build, ci, chore

Apenas responda com a mensagem de commit, sem explicações adicionais.`, changes)
	}
}

// getSystemPrompt retorna a mensagem do sistema no idioma apropriado
func getSystemPrompt(language string) string {
	switch language {
	case "en":
		return "You are an assistant specialized in generating clear and meaningful commit messages in the Conventional Commits format."
	case "es":
		return "Eres un asistente especializado en generar mensajes de commit claros y significativos en el formato de Conventional Commits."
	case "fr":
		return "Vous êtes un assistant spécialisé dans la génération de messages de commit clairs et significatifs au format Conventional Commits."
	case "de":
		return "Sie sind ein Assistent, der sich auf die Erstellung klarer und aussagekräftiger Commit-Nachrichten im Conventional Commits-Format spezialisiert hat."
	default: // Padrão é português pt-br
		return "Você é um assistente especializado em gerar mensagens de commit claras e significativas no formato de Conventional Commits."
	}
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando OpenAI
func (p *OpenAIProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	reqBody := openAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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
func (p *GeminiProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do Gemini não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	reqBody := openAIRequest{
		Model: "gemini-2.0-flash", // Usando o modelo do Gemini
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando Claude
func (p *ClaudeProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do Claude não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	reqBody := claudeRequest{
		Model: "claude-3-haiku-20240307",
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na API do Claude (status %d): %s", resp.StatusCode, body)
	}

	var claudeResp claudeResponse
	if err := json.NewDecoder(resp.Body).Decode(&claudeResp); err != nil {
		return "", err
	}

	if len(claudeResp.Content) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return claudeResp.Content[0].Text, nil
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando DeepSeek
func (p *DeepSeekProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do DeepSeek não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	// DeepSeek usa o formato compatível com OpenAI
	reqBody := openAIRequest{
		Model: "deepseek-coder",
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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

	req, err := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(reqJSON))
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
		return "", fmt.Errorf("erro na API do DeepSeek (status %d): %s", resp.StatusCode, body)
	}

	var deepSeekResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&deepSeekResp); err != nil {
		return "", err
	}

	if len(deepSeekResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return deepSeekResp.Choices[0].Message.Content, nil
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando OpenRouter
func (p *OpenRouterProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do OpenRouter não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	// OpenRouter usa uma API compatível com OpenAI mas permite acesso a vários modelos
	reqBody := openAIRequest{
		Model: "openai/gpt-4-turbo", // Pode ser substituído por outros modelos disponíveis no OpenRouter
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("HTTP-Referer", "https://github.com/user/commit-ai") // Requerido pelo OpenRouter
	req.Header.Set("X-Title", "Commit-AI")                              // Recomendado pelo OpenRouter

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na API do OpenRouter (status %d): %s", resp.StatusCode, body)
	}

	var openRouterResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openRouterResp); err != nil {
		return "", err
	}

	if len(openRouterResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return openRouterResp.Choices[0].Message.Content, nil
}

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando Grok
func (p *GrokProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a chave de API não estiver configurada, retornar um exemplo de mensagem
	if p.APIKey == "" {
		return "feat: implementação inicial (API key do Grok não configurada)", nil
	}

	prompt := getLanguagePrompt(changes, language)

	// Grok usa formato compatível com OpenAI
	reqBody := openAIRequest{
		Model: "grok-1",
		Messages: []message{
			{
				Role:    "system",
				Content: getSystemPrompt(language),
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

	req, err := http.NewRequest("POST", "https://api.grok.ai/v1/chat/completions", bytes.NewBuffer(reqJSON))
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
		return "", fmt.Errorf("erro na API do Grok (status %d): %s", resp.StatusCode, body)
	}

	var grokResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&grokResp); err != nil {
		return "", err
	}

	if len(grokResp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta gerada pela IA")
	}

	return grokResp.Choices[0].Message.Content, nil
}
