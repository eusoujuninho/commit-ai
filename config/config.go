// Package config gerencia as configurações do aplicativo
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config armazena a configuração da aplicação
type Config struct {
	AIProvider    string `json:"ai_provider"`
	OpenAIKey     string `json:"openai_key"`
	GeminiKey     string `json:"gemini_key"`
	ClaudeKey     string `json:"claude_key"`     // Nova chave para API do Claude
	DeepSeekKey   string `json:"deepseek_key"`   // Nova chave para API do DeepSeek
	OpenRouterKey string `json:"openrouter_key"` // Nova chave para API do OpenRouter
	GrokKey       string `json:"grok_key"`       // Nova chave para API do Grok
	OllamaURL     string `json:"ollama_url"`     // URL do servidor Ollama
	OllamaModel   string `json:"ollama_model"`   // Modelo Ollama a ser usado
	RepoPath      string `json:"repo_path"`
	AutoCommit    bool   `json:"auto_commit"`
	CommitStyle   string `json:"commit_style"`
	Language      string `json:"language"` // Idioma para as mensagens de commit
}

// DefaultConfig retorna uma configuração padrão
func DefaultConfig() *Config {
	return &Config{
		AIProvider:    "gemini",
		OpenAIKey:     "",
		GeminiKey:     "",
		ClaudeKey:     "",                       // Inicializado vazio
		DeepSeekKey:   "",                       // Inicializado vazio
		OpenRouterKey: "",                       // Inicializado vazio
		GrokKey:       "",                       // Inicializado vazio
		OllamaURL:     "http://localhost:11434", // URL padrão do Ollama
		OllamaModel:   "llama3",                 // Modelo padrão do Ollama
		RepoPath:      "",
		AutoCommit:    false,
		CommitStyle:   "conventional",
		Language:      "pt-br", // Português Brasil como idioma padrão
	}
}

// LoadConfig carrega a configuração do arquivo
func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// Verificar se o arquivo existe
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Criar configuração padrão
		config := DefaultConfig()

		// Salvar configuração padrão
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return nil, err
		}

		file, err := os.Create(configPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(config); err != nil {
			return nil, err
		}

		return config, nil
	}

	// Carregar configuração existente
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig salva a configuração no arquivo
func SaveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// Criar diretório se não existir
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

// getConfigPath retorna o caminho para o arquivo de configuração
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".commit-ai", "config.json"), nil
}
