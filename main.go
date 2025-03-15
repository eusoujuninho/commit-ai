package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/user/commit-ai/ai"
	"github.com/user/commit-ai/config"
	"github.com/user/commit-ai/git"
)

// Versão do aplicativo
const version = "0.1.0"

// Chave da API Gemini fornecida
const defaultGeminiAPIKey = "AIzaSyBqFvUFUHK6Cj6pk7JVpjpWmOaJJIiq6rQ"

func main() {
	// Definir flags da linha de comando
	configureFlag := flag.Bool("configure", false, "Configurar a aplicação")
	dryRunFlag := flag.Bool("dry-run", false, "Apenas mostrar a mensagem de commit sem fazer o commit")
	repoPathFlag := flag.String("repo", "", "Caminho para o repositório Git")
	versionFlag := flag.Bool("version", false, "Mostrar a versão do aplicativo")
	providerFlag := flag.String("provider", "", "Provedor de IA a ser usado (openai, gemini, claude, deepseek, openrouter, grok)")
	languageFlag := flag.String("language", "", "Idioma para a mensagem de commit (pt-br, en, es, fr, de)")

	// Analisar flags
	flag.Parse()

	// Mostrar versão e sair
	if *versionFlag {
		fmt.Printf("Commit-AI versão %s\n", version)
		os.Exit(0)
	}

	// Carregar configuração
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar configuração: %v\n", err)
		os.Exit(1)
	}

	// Configurar o provedor a partir da flag, se fornecida
	if *providerFlag != "" {
		cfg.AIProvider = *providerFlag
	}

	// Configurar idioma a partir da flag, se fornecida
	if *languageFlag != "" {
		cfg.Language = *languageFlag
	}

	// Verificar se a chave do Gemini está definida e usar o padrão se necessário
	if cfg.AIProvider == "gemini" && cfg.GeminiKey == "" {
		cfg.GeminiKey = defaultGeminiAPIKey
		if err := config.SaveConfig(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao salvar configuração: %v\n", err)
		}
	}

	// Modo de configuração
	if *configureFlag {
		configureApp(cfg)
		os.Exit(0)
	}

	// Determinar o caminho do repositório
	repoPath := *repoPathFlag
	if repoPath == "" {
		repoPath = cfg.RepoPath
		if repoPath == "" {
			var err error
			repoPath, err = os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao obter diretório atual: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Exibir mensagem de dry-run se necessário
	if *dryRunFlag {
		fmt.Println("Modo dry-run ativado - nenhum commit será realizado")
	}

	// Abrir repositório Git
	fmt.Printf("Analisando repositório em: %s\n", repoPath)
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir repositório Git: %v\n", err)
		os.Exit(1)
	}

	// Obter arquivos alterados
	fmt.Println("Detectando alterações...")
	changedFiles, err := repo.GetChangedFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao obter arquivos alterados: %v\n", err)
		os.Exit(1)
	}

	if len(changedFiles) == 0 {
		fmt.Println("Nenhuma alteração detectada para commit.")
		os.Exit(0)
	}

	fmt.Println("Arquivos alterados:")
	for _, file := range changedFiles {
		fmt.Printf("  - %s\n", file)
	}

	// Obter mudanças para análise
	changes, err := repo.GetChanges()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao obter mudanças: %v\n", err)
		os.Exit(1)
	}

	// Criar provedor de IA
	var provider ai.Provider
	switch cfg.AIProvider {
	case "openai":
		fmt.Println("Usando provedor OpenAI...")
		provider = ai.NewOpenAIProvider(cfg.OpenAIKey)
	case "gemini":
		fmt.Println("Usando provedor Gemini...")
		provider = ai.NewGeminiProvider(cfg.GeminiKey)
	case "claude":
		fmt.Println("Usando provedor Claude (Anthropic)...")
		provider = ai.NewClaudeProvider(cfg.ClaudeKey)
	case "deepseek":
		fmt.Println("Usando provedor DeepSeek...")
		provider = ai.NewDeepSeekProvider(cfg.DeepSeekKey)
	case "openrouter":
		fmt.Println("Usando provedor OpenRouter...")
		provider = ai.NewOpenRouterProvider(cfg.OpenRouterKey)
	case "grok":
		fmt.Println("Usando provedor Grok (xAI)...")
		provider = ai.NewGrokProvider(cfg.GrokKey)
	default:
		fmt.Fprintf(os.Stderr, "Provedor de IA desconhecido: %s, usando Gemini como fallback\n", cfg.AIProvider)
		provider = ai.NewGeminiProvider(cfg.GeminiKey)
	}

	// Mostrar idioma sendo usado
	var idiomaTexto string
	switch cfg.Language {
	case "en":
		idiomaTexto = "inglês"
	case "es":
		idiomaTexto = "espanhol"
	case "fr":
		idiomaTexto = "francês"
	case "de":
		idiomaTexto = "alemão"
	default:
		idiomaTexto = "português"
	}
	fmt.Printf("Idioma para mensagens: %s\n", idiomaTexto)

	// Gerar mensagem de commit
	fmt.Println("Gerando mensagem de commit com IA...")
	commitMsg, err := provider.GenerateCommitMessage(changes, cfg.Language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao gerar mensagem de commit: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Mensagem de commit gerada: %s\n", commitMsg)

	// No modo dry-run, apenas exibir a mensagem e sair
	if *dryRunFlag {
		os.Exit(0)
	}

	// Confirmar commit
	if !cfg.AutoCommit {
		// Adaptar a mensagem de confirmação e opções conforme o idioma
		var confirmPrompt string
		var acceptedResponses []string

		switch cfg.Language {
		case "en":
			confirmPrompt = "Make commit with this message? [Y/n]: "
			acceptedResponses = []string{"", "y", "yes"}
		case "es":
			confirmPrompt = "¿Hacer commit con este mensaje? [S/n]: "
			acceptedResponses = []string{"", "s", "si", "sí"}
		case "fr":
			confirmPrompt = "Effectuer un commit avec ce message? [O/n]: "
			acceptedResponses = []string{"", "o", "oui"}
		case "de":
			confirmPrompt = "Commit mit dieser Nachricht durchführen? [J/n]: "
			acceptedResponses = []string{"", "j", "ja"}
		default: // português
			confirmPrompt = "Fazer commit com esta mensagem? [S/n]: "
			acceptedResponses = []string{"", "s", "sim"}
		}

		fmt.Print(confirmPrompt)
		reader := bufio.NewReader(os.Stdin)
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		// Verificar se a resposta é aceita (incluindo vazio como valor padrão)
		isAccepted := false
		for _, response := range acceptedResponses {
			if confirm == response {
				isAccepted = true
				break
			}
		}

		if !isAccepted {
			// Adaptar a mensagem de cancelamento ao idioma
			switch cfg.Language {
			case "en":
				fmt.Println("Operation canceled by the user.")
			case "es":
				fmt.Println("Operación cancelada por el usuario.")
			case "fr":
				fmt.Println("Opération annulée par l'utilisateur.")
			case "de":
				fmt.Println("Vorgang vom Benutzer abgebrochen.")
			default: // português
				fmt.Println("Operação cancelada pelo usuário.")
			}
			os.Exit(0)
		}
	}

	// Realizar commit
	fmt.Println("Realizando commit...")
	if err := repo.Commit(commitMsg); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer commit: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Commit realizado com sucesso!")
}

// configureApp configura a aplicação de forma interativa
func configureApp(cfg *config.Config) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Configuração do Commit-AI ===")

	// Configurar provedor de IA
	fmt.Printf("Provedor de IA (openai/gemini/claude/deepseek/openrouter/grok) [%s]: ", cfg.AIProvider)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "openai" || input == "gemini" || input == "claude" || input == "deepseek" || input == "openrouter" || input == "grok" {
		cfg.AIProvider = input
	}

	// Configurar chave da API dependendo do provedor
	switch cfg.AIProvider {
	case "openai":
		fmt.Printf("Chave da API OpenAI [%s]: ", maskAPIKey(cfg.OpenAIKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.OpenAIKey = input
		}
	case "gemini":
		currentKey := cfg.GeminiKey
		if currentKey == "" {
			currentKey = defaultGeminiAPIKey
		}
		fmt.Printf("Chave da API Gemini [%s]: ", maskAPIKey(currentKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.GeminiKey = input
		} else if cfg.GeminiKey == "" {
			// Se a chave ainda não estiver configurada, usar a padrão
			cfg.GeminiKey = defaultGeminiAPIKey
		}
	case "claude":
		fmt.Printf("Chave da API Claude (Anthropic) [%s]: ", maskAPIKey(cfg.ClaudeKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.ClaudeKey = input
		}
	case "deepseek":
		fmt.Printf("Chave da API DeepSeek [%s]: ", maskAPIKey(cfg.DeepSeekKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.DeepSeekKey = input
		}
	case "openrouter":
		fmt.Printf("Chave da API OpenRouter [%s]: ", maskAPIKey(cfg.OpenRouterKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.OpenRouterKey = input
		}
	case "grok":
		fmt.Printf("Chave da API Grok (xAI) [%s]: ", maskAPIKey(cfg.GrokKey))
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			cfg.GrokKey = input
		}
	}

	// Configurar idioma
	fmt.Printf("Idioma para mensagens (pt-br/en/es/fr/de) [%s]: ", cfg.Language)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "pt-br" || input == "en" || input == "es" || input == "fr" || input == "de" {
		cfg.Language = input
	}

	// Configurar caminho padrão do repositório
	fmt.Printf("Caminho padrão do repositório [%s]: ", cfg.RepoPath)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		cfg.RepoPath = input
	}

	// Configurar commit automático
	fmt.Printf("Commit automático (true/false) [%t]: ", cfg.AutoCommit)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "true" {
		cfg.AutoCommit = true
	} else if input == "false" {
		cfg.AutoCommit = false
	}

	// Salvar configuração
	if err := config.SaveConfig(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao salvar configuração: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Configuração salva com sucesso!")
}

// maskAPIKey mascara a chave da API para exibição
func maskAPIKey(key string) string {
	if key == "" {
		return ""
	}
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "..." + key[len(key)-4:]
}
