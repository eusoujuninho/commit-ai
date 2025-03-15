// Package watcher implementa funcionalidades para monitoramento automático de repositórios Git
package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/user/commit-ai/ai"
	"github.com/user/commit-ai/config"
	"github.com/user/commit-ai/git"
)

// DebounceEvent representa um evento de alteração após debouncing
type DebounceEvent struct {
	RepoPath string
	Time     time.Time
}

// CommitOptions contém as opções para commits automáticos
type CommitOptions struct {
	Interval       time.Duration // Intervalo mínimo entre commits
	MinChanges     int           // Número mínimo de alterações para acionar um commit
	Provider       string        // Provedor de IA para gerar mensagens
	Language       string        // Idioma para mensagens de commit
	IgnorePatterns []string      // Padrões de arquivos para ignorar
	Silent         bool          // Modo silencioso (sem output no console)
	CommitMsg      string        // Mensagem de commit personalizada (opcional)
	AutoStage      bool          // Adicionar arquivos automaticamente
	DoCommit       bool          // Realizar o commit (ou apenas mostrar mensagem)
}

// DefaultCommitOptions retorna as opções padrão para commits automáticos
func DefaultCommitOptions() CommitOptions {
	return CommitOptions{
		Interval:       30 * time.Second,
		MinChanges:     1,
		Provider:       "",
		Language:       "",
		IgnorePatterns: []string{".git/", "*.tmp", "*.log", "*~"},
		Silent:         false,
		AutoStage:      true,
		DoCommit:       true,
	}
}

// Debounce agrupa múltiplos eventos em um único evento após um período de inatividade
func Debounce(events <-chan DebounceEvent, interval time.Duration) <-chan DebounceEvent {
	debounced := make(chan DebounceEvent)

	go func() {
		var latestEvent DebounceEvent
		var timer *time.Timer
		timerActive := false

		for {
			select {
			case event := <-events:
				latestEvent = event
				if !timerActive {
					timer = time.AfterFunc(interval, func() {
						debounced <- latestEvent
						timerActive = false
					})
					timerActive = true
				} else {
					timer.Reset(interval)
				}
			}
		}
	}()

	return debounced
}

// ShouldIgnoreFile verifica se um arquivo deve ser ignorado
func ShouldIgnoreFile(path string, ignorePatterns []string) bool {
	// Sempre ignorar o diretório .git
	if strings.Contains(path, "/.git/") || strings.HasPrefix(path, ".git/") {
		return true
	}

	// Verificar padrões personalizados
	for _, pattern := range ignorePatterns {
		if strings.HasSuffix(pattern, "/") {
			// É um diretório
			dirPattern := strings.TrimSuffix(pattern, "/")
			if strings.Contains(path, "/"+dirPattern+"/") || strings.HasPrefix(path, dirPattern+"/") {
				return true
			}
		} else if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			// É um padrão de arquivo
			return true
		}
	}

	return false
}

// StartWatcher inicia o monitoramento de um repositório
func StartWatcher(repoPath string, cfg *config.Config, options CommitOptions) error {
	// Use os valores do config quando não especificados nas opções
	if options.Provider == "" {
		options.Provider = cfg.AIProvider
	}

	if options.Language == "" {
		options.Language = cfg.Language
	}

	// Criar o watcher para monitorar alterações no sistema de arquivos
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("erro ao criar watcher: %w", err)
	}
	defer fsWatcher.Close()

	// Canal para eventos de alteração após debouncing
	rawEvents := make(chan DebounceEvent)
	debouncedEvents := Debounce(rawEvents, options.Interval)

	// Adicionar todos os diretórios ao watcher, exceto os ignorados
	log.Printf("Iniciando monitoramento do repositório: %s", repoPath)
	filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Continuar mesmo com erro
		}

		if info.IsDir() && !ShouldIgnoreFile(path, options.IgnorePatterns) {
			if err := fsWatcher.Add(path); err != nil {
				log.Printf("Erro ao adicionar diretório %s ao watcher: %v", path, err)
			} else if !options.Silent {
				log.Printf("Monitorando diretório: %s", path)
			}
		}
		return nil
	})

	// Criar canal para sinal de término
	done := make(chan bool)

	// Processar eventos de alteração
	go func() {
		for {
			select {
			case event, ok := <-fsWatcher.Events:
				if !ok {
					return
				}

				// Ignorar arquivos que devem ser ignorados
				if ShouldIgnoreFile(event.Name, options.IgnorePatterns) {
					continue
				}

				// Verificar se é uma alteração relevante
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
					if !options.Silent {
						log.Printf("Alteração detectada: %s (%s)", event.Name, event.Op)
					}
					rawEvents <- DebounceEvent{RepoPath: repoPath, Time: time.Now()}
				}

			case err, ok := <-fsWatcher.Errors:
				if !ok {
					return
				}
				log.Printf("Erro no watcher: %v", err)

			case <-done:
				return
			}
		}
	}()

	// Processar eventos debounced e fazer commits quando apropriado
	go func() {
		for event := range debouncedEvents {
			// Abrir o repositório
			repo, err := git.OpenRepository(event.RepoPath)
			if err != nil {
				log.Printf("Erro ao abrir repositório: %v", err)
				continue
			}

			// Verificar alterações
			changedFiles, err := repo.GetChangedFiles()
			if err != nil {
				log.Printf("Erro ao obter alterações: %v", err)
				continue
			}

			// Filtrar arquivos que devem ser ignorados
			var relevantChanges []string
			for _, file := range changedFiles {
				if !ShouldIgnoreFile(file, options.IgnorePatterns) {
					relevantChanges = append(relevantChanges, file)
				}
			}

			// Verificar se há alterações suficientes para fazer um commit
			if len(relevantChanges) >= options.MinChanges {
				if !options.Silent {
					log.Printf("Detectadas %d alterações relevantes. Preparando commit automático...", len(relevantChanges))
				}

				// Auto-stage se configurado
				if options.AutoStage {
					for _, file := range relevantChanges {
						if !options.Silent {
							log.Printf("Adicionando arquivo: %s", file)
						}
					}
					// Adicionar todas as alterações ao stage
					err = repo.AddFilesToStage(relevantChanges)
					if err != nil {
						log.Printf("Erro ao adicionar arquivos ao stage: %v", err)
						continue
					}
				}

				// Gerar mensagem de commit com IA ou usar mensagem personalizada
				var commitMsg string
				if options.CommitMsg != "" {
					commitMsg = options.CommitMsg
				} else {
					// Obter alterações para análise
					changes, err := repo.GetChanges()
					if err != nil {
						log.Printf("Erro ao obter mudanças: %v", err)
						continue
					}

					// Criar provedor de IA
					var provider ai.Provider
					switch options.Provider {
					case "openai":
						provider = ai.NewOpenAIProvider(cfg.OpenAIKey)
					case "gemini":
						provider = ai.NewGeminiProvider(cfg.GeminiKey)
					case "claude":
						provider = ai.NewClaudeProvider(cfg.ClaudeKey)
					case "deepseek":
						provider = ai.NewDeepSeekProvider(cfg.DeepSeekKey)
					case "openrouter":
						provider = ai.NewOpenRouterProvider(cfg.OpenRouterKey)
					case "grok":
						provider = ai.NewGrokProvider(cfg.GrokKey)
					case "ollama":
						provider = ai.NewOllamaProvider(cfg.OllamaURL, cfg.OllamaModel)
					default:
						provider = ai.NewGeminiProvider(cfg.GeminiKey)
					}

					// Gerar mensagem de commit
					if !options.Silent {
						log.Printf("Gerando mensagem de commit com IA (%s)...", options.Provider)
					}

					commitMsg, err = provider.GenerateCommitMessage(changes, options.Language)
					if err != nil {
						log.Printf("Erro ao gerar mensagem de commit: %v", err)
						continue
					}
				}

				if !options.Silent {
					log.Printf("Mensagem de commit gerada: %s", commitMsg)
				}

				// Realizar o commit se configurado
				if options.DoCommit {
					if !options.Silent {
						log.Printf("Realizando commit...")
					}

					err = repo.Commit(commitMsg)
					if err != nil {
						log.Printf("Erro ao fazer commit: %v", err)
						continue
					}

					if !options.Silent {
						log.Printf("Commit realizado com sucesso!")
					}
				} else {
					if !options.Silent {
						log.Printf("Modo dry-run: commit não realizado")
					}
				}
			} else if !options.Silent {
				log.Printf("Apenas %d alterações detectadas. Mínimo para commit: %d", len(relevantChanges), options.MinChanges)
			}
		}
	}()

	fmt.Printf("Watcher iniciado. Pressione Ctrl+C para sair.\n")

	// Manter o watcher em execução
	select {}
}
