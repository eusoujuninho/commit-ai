# Commit-AI Documentation

Commit-AI is a Go-based utility that automates the generation of Git commit messages using Artificial Intelligence. It analyzes code changes in a Git repository and generates meaningful, well-formatted commit messages following the Conventional Commits standard.

**Version:** 0.1.0

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Pre-compiled Binaries](#pre-compiled-binaries)
  - [Building from Source](#building-from-source)
  - [Installation Options](#installation-options)
- [Configuration](#configuration)
  - [First-time Setup](#first-time-setup)
  - [Configuration File](#configuration-file)
  - [AI Providers](#ai-providers)
  - [Language Support](#language-support)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Command-line Options](#command-line-options)
  - [Exit Codes](#exit-codes)
- [Use Cases](#use-cases)
  - [Individual Developers](#individual-developers)
  - [Team Consistency](#team-consistency)
  - [CI/CD Integration](#cicd-integration)
- [AI Providers](#ai-providers-1)
  - [Google Gemini](#google-gemini)
  - [OpenAI](#openai)
  - [Claude (Anthropic)](#claude-anthropic)
  - [DeepSeek](#deepseek)
  - [OpenRouter](#openrouter)
  - [Grok (xAI)](#grok-xai)
  - [Ollama (Local Models)](#ollama-local-models)
  - [Adding New Providers](#adding-new-providers)
- [Multilingual Support](#multilingual-support)
  - [Available Languages](#available-languages)
  - [Changing Language](#changing-language)
- [Examples](#examples)
  - [Basic Commit](#basic-commit)
  - [Dry Run Mode](#dry-run-mode)
  - [Custom Repository Path](#custom-repository-path)
  - [Auto-commit Mode](#auto-commit-mode)
  - [Using Different Languages](#using-different-languages)
  - [Using Ollama Local Models](#using-ollama-local-models)
- [Architecture](#architecture)
  - [Component Overview](#component-overview)
  - [Data Flow](#data-flow)
- [Development](#development)
  - [Project Structure](#project-structure)
  - [Extending Commit-AI](#extending-commit-ai)
- [Troubleshooting](#troubleshooting)
  - [Common Issues](#common-issues)
  - [Logs and Debugging](#logs-and-debugging)
- [FAQ](#faq)
- [License](#license)
- [Contributing](#contributing)
- [Modos de Operação](#modos-de-operação)
  - [Commit Padrão](#commit-padrão)
  - [Modo Dry-Run](#modo-dry-run)
  - [Modo Watcher](#modo-watcher)
- [Opções de Linha de Comando](#opções-de-linha-de-comando)

## Introduction

Commit-AI simplifies the process of writing meaningful Git commit messages by leveraging AI to analyze code changes and generate appropriate commit messages following the Conventional Commits standard. This helps maintain a consistent commit history while saving developers time and effort.

## Features

- Automatically detects modified files in Git repositories
- Generates commit messages in the Conventional Commits format
- Supports multiple AI providers (Google Gemini, OpenAI, Claude, DeepSeek, OpenRouter, Grok)
- Supports local AI models through Ollama integration
- Provides multilingual support (Portuguese, English, Spanish, French, German)
- Provides an intuitive command-line interface
- Offers customizable settings through a configuration file
- Includes a "dry-run" mode to preview messages before committing
- Handles both staged and unstaged changes
- Truncates large diffs to avoid token limits with AI providers

## Installation

### Prerequisites

Before installing Commit-AI, ensure you have the following:

- Go 1.16 or later
- Git installed and configured
- API key for at least one of the supported AI providers:
  - Google Gemini (default)
  - OpenAI
  - Claude (Anthropic)
  - DeepSeek
  - OpenRouter
  - Grok (xAI)
- (Optional) Ollama installed locally for using local AI models without API keys

### Pre-compiled Binaries

For convenience, Commit-AI provides pre-compiled binaries for major operating systems. These binaries are available in the `bin/` directory of the project repository:

- **Linux**: `bin/commit-ai-linux-amd64`
- **macOS**: `bin/commit-ai-darwin-amd64`
- **Windows**: `bin/commit-ai-windows-amd64.exe`

To use a pre-compiled binary:

1. Download the appropriate binary for your operating system
2. Verify the integrity of the binary using the checksums in `bin/checksums.txt`:
   ```bash
   sha256sum -c checksums.txt
   ```
3. Make it executable (Linux/macOS):
   ```bash
   chmod +x commit-ai-linux-amd64
   ```
4. Move it to a location in your PATH:

   **Linux/macOS**:
   ```bash
   sudo mv commit-ai-linux-amd64 /usr/local/bin/commit-ai
   ```
   or without sudo:
   ```bash
   mv commit-ai-linux-amd64 ~/bin/commit-ai
   ```

   **Windows**:
   - Move `commit-ai-windows-amd64.exe` to a directory in your PATH
   - Optionally rename it to `commit-ai.exe` for convenience

Using pre-compiled binaries eliminates the need to have Go installed on your system and simplifies the installation process.

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/user/commit-ai.git
   ```

2. Navigate to the project directory:
   ```bash
   cd commit-ai
   ```

3. Build the project:
   ```bash
   go build
   ```

4. (Optional) Install the binary to your Go bin directory:
   ```bash
   go install
   ```
   This will compile and install the binary to your Go bin directory (usually `~/go/bin/`).

### Installation Options

#### Global Installation

You can install Commit-AI globally on your system using one of these methods:

##### Using Go Install (Recommended)

If you have Go installed, the easiest way to install Commit-AI globally is:

```bash
# From anywhere
go install github.com/user/commit-ai@latest

# Or from the project directory
cd path/to/commit-ai
go install
```

This will:
- Compile the application
- Place the binary in your `$GOPATH/bin` directory (usually `~/go/bin/`)
- Make it available as `commit-ai` command if your `$GOPATH/bin` is in your PATH

To verify the installation:
```bash
which commit-ai     # Should show something like /home/juninho/go/bin/commit-ai
commit-ai --version # Should display: "Commit-AI versão 0.1.0"
```

> **Note**: Make sure your `$GOPATH/bin` directory is in your PATH. You can check your GOPATH with `go env GOPATH`. You can add it to your PATH by adding this line to your `.bashrc` or `.zshrc` file:
> ```bash
> export PATH="$PATH:$(go env GOPATH)/bin"
> ```

##### Manual Installation

After building, you can move the binary to a location in your PATH:

```bash
# Linux/macOS (requires admin privileges)
sudo mv commit-ai /usr/local/bin/

# Alternative (no sudo required)
mv commit-ai ~/bin/  # Ensure ~/bin is in your PATH
```

#### Docker Installation

For containerized usage, you can build and run a Docker image:

```bash
# Build the Docker image
docker build -t commit-ai .

# Run Commit-AI in a Docker container
# Mount your Git repository to /repo inside the container
docker run -it -v $(pwd):/repo commit-ai --repo=/repo
```

## Configuration

### First-time Setup

When you first run Commit-AI, it will create a default configuration file. You can also set up the configuration explicitly:

```bash
commit-ai --configure
```

This will prompt you for:
- AI provider selection (Gemini, OpenAI, Claude, DeepSeek, OpenRouter, Grok, or Ollama)
- API key for the selected provider (or server URL and model for Ollama)
- Preferred language for commit messages
- Default repository path (optional)
- Auto-commit preference (true/false)

### Configuration File

The configuration is stored in `~/.commit-ai/config.json` with the following structure:

```json
{
  "ai_provider": "gemini",
  "openai_key": "",
  "gemini_key": "your-gemini-key",
  "claude_key": "",
  "deepseek_key": "",
  "openrouter_key": "",
  "grok_key": "",
  "ollama_url": "http://localhost:11434",
  "ollama_model": "llama3",
  "language": "pt-br",
  "repo_path": "",
  "auto_commit": false,
  "commit_style": "conventional"
}
```

### AI Providers

Commit-AI supports the following AI providers:

- **Gemini**: Google's Gemini AI model (default)
- **OpenAI**: OpenAI's GPT models
- **Claude**: Anthropic's Claude models
- **DeepSeek**: DeepSeek AI models
- **OpenRouter**: Unified API access to multiple AI models
- **Grok**: xAI's Grok model
- **Ollama**: Local AI models running on your own machine

### Language Support

Commit-AI now supports generating commit messages in multiple languages:

- Portuguese (pt-br) - Default
- English (en)
- Spanish (es)
- French (fr)
- German (de)

## Usage

### Basic Usage

To generate a commit message and perform a commit:

```bash
# Navigate to your Git repository
cd your-repository

# Run Commit-AI
commit-ai
```

The tool will:
1. Detect changes in the repository
2. Generate a commit message using the configured AI provider in the selected language
3. Display the generated message
4. Ask for confirmation (unless auto-commit is enabled)
5. Perform the commit if confirmed

### Command-line Options

Commit-AI supports the following command-line options:

| Option | Description |
|--------|-------------|
| `--configure` | Run the configuration wizard |
| `--dry-run` | Generate a commit message without performing a commit |
| `--repo=PATH` | Specify a custom repository path |
| `--provider=NAME` | Override the AI provider (gemini/openai/claude/deepseek/openrouter/grok/ollama) for this run |
| `--language=CODE` | Override the language (pt-br/en/es/fr/de) for this run |
| `--version` | Display the version information |
| `--watch` | Ativa o modo watcher (monitoramento contínuo) |
| `--interval=DURATION` | Define o intervalo entre verificações no modo watcher (padrão: 30s) |
| `--min-changes=N` | Número mínimo de alterações para acionar um commit (padrão: 1) |
| `--ignore=PATTERNS` | Lista de padrões a ignorar, separados por vírgula |
| `--silent` | Modo silencioso para watcher (menos output) |

### Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | General error (configuration, repository access, etc.) |

## Use Cases

### Individual Developers

- **Quick Commits**: Generate meaningful commit messages in seconds
- **Consistency**: Maintain a consistent commit style across all your projects
- **Learning**: Learn best practices for commit messages by example
- **Multilingual**: Generate commit messages in your preferred language

### Team Consistency

- **Standardization**: Ensure all team members follow the same commit message format
- **Readability**: Make the repository history more readable and navigable
- **Automation**: Simplify the code review process with clear commit messages
- **Internationalization**: Support for multinational teams with different language preferences

### CI/CD Integration

- **Automated Commits**: Use Commit-AI in CI/CD pipelines for automated tasks
- **Dependency Updates**: Generate appropriate commit messages for dependency updates
- **Code Formatting**: Create meaningful commits for automated code formatting

## AI Providers

### Google Gemini

Commit-AI uses Google's Gemini model by default. To use Gemini:

1. Obtain a Gemini API key from [Google AI Studio](https://ai.google.dev/)
2. Configure Commit-AI to use Gemini:
   ```bash
   commit-ai --configure
   # Select "gemini" as provider and enter your API key
   ```

### OpenAI

To use OpenAI as your AI provider:

1. Obtain an API key from [OpenAI](https://platform.openai.com/)
2. Configure Commit-AI to use OpenAI:
   ```bash
   commit-ai --configure
   # Select "openai" as provider and enter your API key
   ```

### Claude (Anthropic)

To use Claude as your AI provider:

1. Obtain an API key from [Anthropic](https://console.anthropic.com/)
2. Configure Commit-AI to use Claude:
   ```bash
   commit-ai --configure
   # Select "claude" as provider and enter your API key
   ```

Claude is known for its strong natural language understanding and reasoning capabilities, which can produce particularly insightful commit messages.

### DeepSeek

To use DeepSeek as your AI provider:

1. Obtain an API key from [DeepSeek](https://platform.deepseek.com/)
2. Configure Commit-AI to use DeepSeek:
   ```bash
   commit-ai --configure
   # Select "deepseek" as provider and enter your API key
   ```

DeepSeek provides specialized code-aware models that are particularly good at understanding code changes.

### OpenRouter

To use OpenRouter as your AI provider:

1. Obtain an API key from [OpenRouter](https://openrouter.ai/)
2. Configure Commit-AI to use OpenRouter:
   ```bash
   commit-ai --configure
   # Select "openrouter" as provider and enter your API key
   ```

OpenRouter provides unified access to multiple AI models from different providers, allowing you to use various models without managing multiple API keys.

### Grok (xAI)

To use Grok as your AI provider:

1. Obtain an API key from [xAI/Grok](https://grok.x.ai/)
2. Configure Commit-AI to use Grok:
   ```bash
   commit-ai --configure
   # Select "grok" as provider and enter your API key
   ```

Grok is the AI model developed by xAI, offering strong capabilities for code understanding and generating concise, meaningful commit messages.

### Ollama (Local Models)

To use Ollama as your AI provider:

1. Install Ollama on your local machine from [Ollama.ai](https://ollama.ai/)
2. Pull and run a model (e.g., llama3, mistral, codellama):
   ```bash
   ollama pull llama3
   ```
3. Configure Commit-AI to use Ollama:
   ```bash
   commit-ai --configure
   # Select "ollama" as provider and configure URL and model
   ```

Using Ollama provides several advantages:
- **Privacy**: Your code changes never leave your machine
- **No API Key Required**: Run models without needing external API keys
- **Offline Usage**: Generate commit messages without internet access
- **Cost-Free**: No usage charges or token limits
- **Model Selection**: Choose from various models based on your needs

Commit-AI is configured to use Ollama on `http://localhost:11434` by default, which is Ollama's standard address. The default model is `llama3`, but you can configure any model you've pulled to Ollama.

### Adding New Providers

Commit-AI is designed to be extensible. New AI providers can be added by implementing the `Provider` interface in the `ai` package:

```go
// Provider defines the interface for AI providers
type Provider interface {
    GenerateCommitMessage(changes string, language string) (string, error)
}
```

## Multilingual Support

Commit-AI now supports generating commit messages in multiple languages.

### Available Languages

The currently supported languages are:

- Portuguese (pt-br) - Default
- English (en)
- Spanish (es)
- French (fr)
- German (de)

### Changing Language

You can change the language in several ways:

1. **Configuration Wizard**:
   ```bash
   commit-ai --configure
   # Select your preferred language when prompted
   ```

2. **Command-line Flag**:
   ```bash
   commit-ai --language=en  # For English
   commit-ai --language=es  # For Spanish
   commit-ai --language=fr  # For French
   commit-ai --language=de  # For German
   commit-ai --language=pt-br  # For Portuguese (default)
   ```

3. **Configuration File**:
   Edit the `language` field in `~/.commit-ai/config.json`

## Examples

### Basic Commit

```bash
# Navigate to a Git repository with changes
cd your-repository

# Run Commit-AI
commit-ai
```

Example output:
```
Analisando repositório em: /path/to/your-repository
Detectando alterações...
Arquivos alterados:
  - src/main.go
  - README.md
Usando provedor Gemini...
Idioma para mensagens: português
Gerando mensagem de commit com IA...
Mensagem de commit gerada: docs: atualiza README e adiciona novos parâmetros em main.go
Fazer commit com esta mensagem? (s/n): s
Realizando commit...
Commit realizado com sucesso!
```

### Dry Run Mode

```bash
# Preview commit message without committing
commit-ai --dry-run
```

### Custom Repository Path

```bash
# Generate a commit message for a specific repository
commit-ai --repo=/path/to/repository
```

### Auto-commit Mode

When configured with `auto_commit: true`, Commit-AI will automatically commit changes without asking for confirmation.

### Using Different Languages

```bash
# Generate commit message in English
commit-ai --language=en

# Generate commit message in Spanish
commit-ai --language=es

# Generate commit message with a specific provider in French
commit-ai --provider=openai --language=fr
```

### Using Ollama Local Models

```bash
# Generate commit message using Ollama local model with default settings
commit-ai --provider=ollama

# Generate commit message using Ollama with a specific model
# (assuming you've already pulled the model with 'ollama pull codellama')
commit-ai --provider=ollama
# Then configure the model during the commit process

# Use Ollama with a specific language
commit-ai --provider=ollama --language=en
```

To set up Ollama as your default provider:

```bash
commit-ai --configure
# Select "ollama" as the provider
# Specify the URL (default: http://localhost:11434)
# Choose your preferred model (e.g., llama3, mistral, codellama)
```

When using Ollama, your code changes remain on your local machine, providing enhanced privacy.

## Architecture

### Component Overview

Commit-AI consists of several key components:

1. **Main Application**: Handles CLI parsing and orchestration
2. **Git Integration**: Detects and analyzes changes in repositories
3. **AI Providers**: Interfaces with various AI services
4. **Config Management**: Manages user preferences and API keys
5. **Language Processing**: Handles multilingual support

### Data Flow

The data flow in Commit-AI is as follows:

1. Load user configuration
2. Open and analyze Git repository
3. Detect changed files and generate diff
4. Send diff to AI provider
5. Receive generated commit message
6. Display message to user and request confirmation (if not in auto-commit mode)
7. Perform Git commit with the generated message

## Development

### Project Structure

```
commit-ai/
├── main.go            # Main application entry point
├── git/               # Git interaction package
│   └── git.go
├── ai/                # AI providers package
│   └── ai.go
├── config/            # Configuration management
│   └── config.go
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── README.md          # Project documentation
└── DOCUMENTATION.md   # Detailed documentation
```

### Extending Commit-AI

Commit-AI is designed to be extensible in several ways:

1. **Adding AI Providers**: Implement the `Provider` interface in the `ai` package
2. **Enhancing Git Integration**: Extend the `Repository` struct in the `git` package
3. **Adding Configuration Options**: Modify the `Config` struct in the `config` package

## Troubleshooting

### Common Issues

#### Authentication Errors

If you encounter AI provider authentication errors:

1. Verify your API key is correct
2. Ensure your API key has the necessary permissions
3. Check that your internet connection is working
4. Confirm the API provider is not experiencing downtime

#### Git Repository Issues

If Commit-AI cannot detect your Git repository:

1. Ensure you're in a valid Git repository directory
2. Check that you have proper permissions to the repository
3. Verify that git is installed and accessible in your PATH
4. For custom repository paths, ensure the path is correct and accessible

### Logs and Debugging

Commit-AI provides error messages to help diagnose issues. For more detailed information, review the error output when running commands.

## FAQ

**Q: Does Commit-AI work with private repositories?**  
A: Yes, Commit-AI works with any Git repository you have access to, including private ones.

**Q: Does Commit-AI send my code to external services?**  
A: Commit-AI sends the diff of your changes to the configured AI provider for analysis. This includes the code changes but not your entire codebase. Large diffs are truncated to 4000 characters.

**Q: Can I use Commit-AI without an internet connection?**  
A: No, Commit-AI requires an internet connection to communicate with AI providers.

**Q: How secure is my API key?**  
A: Your API key is stored locally in the configuration file (~/.commit-ai/config.json). Take standard precautions to protect this file.

**Q: Can I customize the commit message format?**  
A: Currently, Commit-AI generates messages in the Conventional Commits format. Customization options may be added in future versions.

## License

Commit-AI is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions to Commit-AI are welcome! Please feel free to submit pull requests, report bugs, or suggest features.

## Modos de Operação

### Commit Padrão

### Modo Dry-Run

### Modo Watcher

O Commit-AI oferece um modo de monitoramento contínuo (watcher) que observa as alterações no repositório em tempo real e realiza commits automáticos quando detectadas mudanças significativas.

Para ativar o modo watcher:

```bash
commit-ai --watch
```

Opções adicionais para o modo watcher:

- `--interval=DURATION`: Define o intervalo mínimo entre verificações de alterações (exemplo: 30s, 1m, 5m). Padrão: 30s.
- `--min-changes=N`: Define o número mínimo de arquivos alterados para acionar um commit. Padrão: 1.
- `--ignore=PATTERNS`: Lista de padrões de arquivos a serem ignorados, separados por vírgula (exemplo: "*.log,tmp/,*.bak"). Padrão: .git/,*.tmp,*.log,*~
- `--silent`: Reduz a quantidade de mensagens de log exibidas durante o monitoramento.

O modo watcher pode ser combinado com outras opções como `--provider`, `--language` e `--dry-run`:

```bash
commit-ai --watch --provider=claude --interval=2m --min-changes=3 --language=en
```

No exemplo acima, o Commit-AI verificará alterações a cada 2 minutos e realizará commits automáticos quando houver pelo menos 3 arquivos alterados, utilizando o provedor Claude para gerar mensagens de commit em inglês.

## Opções de Linha de Comando

| Flag                 | Descrição                                                                             |
|----------------------|---------------------------------------------------------------------------------------|
| --watch              | Ativa o modo watcher (monitoramento contínuo)                                         |
| --interval=DURATION  | Define o intervalo entre verificações no modo watcher (padrão: 30s)                   |
| --min-changes=N      | Número mínimo de alterações para acionar um commit (padrão: 1)                        |
| --ignore=PATTERNS    | Lista de padrões a ignorar, separados por vírgula                                     |
| --silent             | Modo silencioso para watcher (menos output)                                           | 