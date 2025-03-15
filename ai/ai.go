// Package ai fornece funcionalidades para interação com APIs de IA
package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

// OllamaProvider implementa a interface Provider para o Ollama (modelos locais)
type OllamaProvider struct {
	ServerURL string
	Model     string
}

// NewOllamaProvider cria uma nova instância de OllamaProvider
func NewOllamaProvider(serverURL string, model string) *OllamaProvider {
	return &OllamaProvider{
		ServerURL: serverURL,
		Model:     model,
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

// Estrutura para requisição ao Ollama
type ollamaRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"`
}

// Estrutura para resposta do Ollama
type ollamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

// getLanguagePrompt retorna o prompt adequado para o idioma solicitado
func getLanguagePrompt(changes string, language string) string {
	switch language {
	case "en":
		return fmt.Sprintf(`Analyze the following code changes in detail and generate a comprehensive, meaningful commit message following the Conventional Commits format:

%s

You must:
1. Follow the Conventional Commits format (type: description)
2. Be specific and precise about what was changed and why
3. Look at the actual code changes in the diff, not just the file names
4. Focus on the technical details and functional impacts of the changes
5. Limit to a maximum of 100 characters (for the first line)
6. Start with one of these types based on the nature of the changes:
   - feat: A new feature 
   - fix: A bug fix
   - docs: Documentation only changes
   - style: Changes that do not affect the meaning of the code (white-space, formatting, etc)
   - refactor: A code change that neither fixes a bug nor adds a feature
   - perf: A code change that improves performance
   - test: Adding missing tests or correcting existing tests
   - build: Changes that affect the build system or external dependencies
   - ci: Changes to CI configuration files and scripts
   - chore: Other changes that don't modify src or test files

Only respond with the commit message without any additional explanation or commentary.`, changes)
	case "es":
		return fmt.Sprintf(`Analiza en detalle los siguientes cambios de código y genera un mensaje de commit completo y significativo siguiendo el formato de Conventional Commits:

%s

Debes:
1. Seguir el formato de Conventional Commits (tipo: descripción)
2. Ser específico y preciso sobre qué se cambió y por qué
3. Analizar los cambios reales en el código del diff, no solo los nombres de los archivos
4. Enfocarte en los detalles técnicos y el impacto funcional de los cambios
5. Limitar a un máximo de 100 caracteres (para la primera línea)
6. Comenzar con uno de estos tipos basado en la naturaleza de los cambios:
   - feat: Una nueva característica
   - fix: Corrección de un error
   - docs: Cambios solo en documentación
   - style: Cambios que no afectan el significado del código (espacios en blanco, formato, etc)
   - refactor: Un cambio de código que no corrige un error ni agrega una característica
   - perf: Un cambio de código que mejora el rendimiento
   - test: Agregar pruebas faltantes o corregir pruebas existentes
   - build: Cambios que afectan el sistema de compilación o dependencias externas
   - ci: Cambios en los archivos de configuración y scripts de CI
   - chore: Otros cambios que no modifican los archivos src o test

Responde solamente con el mensaje de commit sin ninguna explicación o comentario adicional.`, changes)
	case "fr":
		return fmt.Sprintf(`Analysez en détail les modifications de code suivantes et générez un message de commit complet et significatif suivant le format Conventional Commits:

%s

Vous devez:
1. Suivre le format Conventional Commits (type: description)
2. Être spécifique et précis sur ce qui a été modifié et pourquoi
3. Examiner les changements réels dans le code du diff, pas seulement les noms de fichiers
4. Vous concentrer sur les détails techniques et les impacts fonctionnels des modifications
5. Limiter à un maximum de 100 caractères (pour la première ligne)
6. Commencer par l'un de ces types en fonction de la nature des changements:
   - feat: Une nouvelle fonctionnalité
   - fix: Correction d'un bug
   - docs: Modifications de la documentation uniquement
   - style: Changements qui n'affectent pas la signification du code (espace, formatage, etc)
   - refactor: Une modification du code qui ne corrige pas un bug et n'ajoute pas de fonctionnalité
   - perf: Une modification du code qui améliore les performances
   - test: Ajout de tests manquants ou correction de tests existants
   - build: Modifications affectant le système de build ou les dépendances externes
   - ci: Modifications des fichiers de configuration CI et des scripts
   - chore: Autres changements qui ne modifient pas les fichiers src ou test

Répondez uniquement avec le message de commit sans aucune explication ou commentaire supplémentaire.`, changes)
	case "de":
		return fmt.Sprintf(`Analysieren Sie die folgenden Codeänderungen im Detail und generieren Sie eine umfassende, aussagekräftige Commit-Nachricht im Conventional Commits-Format:

%s

Sie müssen:
1. Das Conventional Commits-Format befolgen (Typ: Beschreibung)
2. Spezifisch und präzise angeben, was geändert wurde und warum
3. Die tatsächlichen Codeänderungen im Diff betrachten, nicht nur die Dateinamen
4. Sich auf die technischen Details und funktionalen Auswirkungen der Änderungen konzentrieren
5. Auf maximal 100 Zeichen beschränken (für die erste Zeile)
6. Mit einem dieser Typen beginnen, basierend auf der Art der Änderungen:
   - feat: Eine neue Funktion
   - fix: Eine Fehlerbehebung
   - docs: Änderungen nur an der Dokumentation
   - style: Änderungen, die die Bedeutung des Codes nicht beeinflussen (Leerzeichen, Formatierung, usw.)
   - refactor: Eine Codeänderung, die weder einen Fehler behebt noch eine Funktion hinzufügt
   - perf: Eine Codeänderung, die die Leistung verbessert
   - test: Hinzufügen fehlender Tests oder Korrigieren vorhandener Tests
   - build: Änderungen, die das Build-System oder externe Abhängigkeiten betreffen
   - ci: Änderungen an CI-Konfigurationsdateien und Skripten
   - chore: Andere Änderungen, die keine src- oder test-Dateien modifizieren

Antworten Sie nur mit der Commit-Nachricht ohne zusätzliche Erklärungen oder Kommentare.`, changes)
	default: // Padrão é português pt-br
		return fmt.Sprintf(`Analise detalhadamente as seguintes mudanças no código e gere uma mensagem de commit abrangente e significativa seguindo o formato de Conventional Commits:

%s

Você deve:
1. Seguir o formato de Conventional Commits (tipo: descrição)
2. Ser específico e preciso sobre o que foi alterado e por quê
3. Analisar as mudanças reais no código do diff, não apenas os nomes dos arquivos
4. Focar nos detalhes técnicos e impactos funcionais das alterações
5. Limitar a um máximo de 100 caracteres (para a primeira linha)
6. Começar com um destes tipos com base na natureza das alterações:
   - feat: Uma nova funcionalidade
   - fix: Correção de um bug
   - docs: Alterações apenas na documentação
   - style: Alterações que não afetam o significado do código (espaço em branco, formatação, etc)
   - refactor: Uma alteração de código que não corrige um bug nem adiciona uma funcionalidade
   - perf: Uma alteração de código que melhora o desempenho
   - test: Adição de testes ausentes ou correção de testes existentes
   - build: Alterações que afetam o sistema de compilação ou dependências externas
   - ci: Alterações nos arquivos de configuração de CI e scripts
   - chore: Outras alterações que não modificam arquivos src ou test

Responda apenas com a mensagem de commit sem nenhuma explicação ou comentário adicional.`, changes)
	}
}

// getSystemPrompt retorna a mensagem do sistema no idioma apropriado
func getSystemPrompt(language string) string {
	switch language {
	case "en":
		return "You are an expert Git commit message generator with deep understanding of software development principles. Your task is to analyze code changes and create precise, informative commit messages that accurately describe what was changed and why. Focus on technical details and functional impacts, not just superficial changes."
	case "es":
		return "Eres un experto generador de mensajes de commit de Git con profundo conocimiento de los principios de desarrollo de software. Tu tarea es analizar los cambios de código y crear mensajes de commit precisos e informativos que describan con exactitud qué se cambió y por qué. Enfócate en los detalles técnicos y los impactos funcionales, no solo en los cambios superficiales."
	case "fr":
		return "Vous êtes un expert en génération de messages de commit Git avec une compréhension approfondie des principes de développement logiciel. Votre tâche consiste à analyser les modifications de code et à créer des messages de commit précis et informatifs qui décrivent avec précision ce qui a été modifié et pourquoi. Concentrez-vous sur les détails techniques et les impacts fonctionnels, pas seulement sur les changements superficiels."
	case "de":
		return "Sie sind ein Experte für Git-Commit-Nachrichten mit tiefem Verständnis für Software-Entwicklungsprinzipien. Ihre Aufgabe ist es, Codeänderungen zu analysieren und präzise, informative Commit-Nachrichten zu erstellen, die genau beschreiben, was geändert wurde und warum. Konzentrieren Sie sich auf technische Details und funktionale Auswirkungen, nicht nur auf oberflächliche Änderungen."
	default: // Padrão é português pt-br
		return "Você é um especialista em geração de mensagens de commit Git com profundo entendimento dos princípios de desenvolvimento de software. Sua tarefa é analisar mudanças de código e criar mensagens de commit precisas e informativas que descrevam com precisão o que foi alterado e por quê. Concentre-se nos detalhes técnicos e impactos funcionais, não apenas em mudanças superficiais."
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

// GenerateCommitMessage gera uma mensagem de commit com base nas mudanças usando Ollama
func (p *OllamaProvider) GenerateCommitMessage(changes string, language string) (string, error) {
	// Se a URL do servidor não estiver configurada, retornar um exemplo de mensagem
	if p.ServerURL == "" {
		return "feat: implementação inicial (URL do servidor Ollama não configurada)", nil
	}

	// Se o modelo não estiver especificado, usar um padrão
	model := p.Model
	if model == "" {
		model = "llama3" // Modelo padrão
	}

	prompt := getLanguagePrompt(changes, language)

	// Ollama usa uma API compatível com OpenAI para chat
	reqBody := ollamaRequest{
		Model: model,
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
		Stream: false, // Não usar streaming para simplificar
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Montar a URL completa para o endpoint de chat
	url := fmt.Sprintf("%s/api/chat", p.ServerURL)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second, // Modelos locais podem ser mais lentos
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao conectar ao servidor Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("erro na API do Ollama (status %d): %s", resp.StatusCode, body)
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta do Ollama: %w", err)
	}

	// Verificar se a resposta contém uma mensagem
	if ollamaResp.Message.Content == "" {
		return "", fmt.Errorf("nenhuma resposta gerada pelo Ollama")
	}

	return ollamaResp.Message.Content, nil
}
