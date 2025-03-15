# Commit-AI

A Golang tool that automates Git commit messages using Artificial Intelligence.

## Description

Commit-AI analyzes changes made in a Git repository and generates meaningful commit messages in Conventional Commits format using AI. The tool aims to improve the consistency and quality of commit messages, saving developers time and effort.

## Features

- Automatically detects modified files in repositories
- Generates commit messages following Conventional Commits standard
- Supports multiple AI providers (OpenAI, Gemini, Claude, DeepSeek, OpenRouter, Grok, Ollama)
- Offers multilingual support (Portuguese, English, Spanish, French, German)
- Provides an intuitive command-line interface
- Includes persistent and customizable configuration
- Offers "dry-run" mode to preview messages before committing

## Installation

### Prerequisites

- Git installed and configured
- For building from source: Go 1.16 or later
- API key for at least one of the supported AI providers (not needed for Ollama)

### Pre-compiled Binaries

For convenience, we provide pre-compiled binaries for major operating systems:

- **Linux**: `bin/commit-ai-linux-amd64`
- **macOS**: `bin/commit-ai-darwin-amd64`
- **Windows**: `bin/commit-ai-windows-amd64.exe`

You don't need to have Go installed to use the pre-compiled binaries. Simply download the appropriate binary for your operating system, make it executable (Linux/macOS), and move it to a directory in your PATH.

### Building from Source

```bash
# Clone the repository
git clone https://github.com/user/commit-ai.git

# Navigate to the directory
cd commit-ai

# Compile and install
go install
```

## Usage

### Initial Configuration

On first run, configure the tool:

```bash
commit-ai --configure
```

### Basic Usage

```bash
# In your Git repository directory
commit-ai
```

### Options

- `--configure`: Configuration mode
- `--dry-run`: Generates the message without committing
- `--repo=PATH`: Specifies a path to the repository
- `--provider=NAME`: Overrides the AI provider for this run
- `--language=CODE`: Overrides the language for this run
- `--version`: Displays version information

## Contributing

We welcome contributions to Commit-AI! Here's how you can help:

### Reporting Issues

If you find a bug or have a suggestion for improvement:

1. Check if the issue already exists in the [Issues](https://github.com/user/commit-ai/issues) section
2. If not, create a new issue with a descriptive title and detailed information:
   - Steps to reproduce the problem
   - Expected behavior
   - Actual behavior
   - Environment details (OS, Go version, etc.)

### Submitting Pull Requests

1. Fork the repository
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
   or
   ```bash
   git checkout -b fix/issue-description
   ```
3. Make your changes following our code style
4. Add tests for your changes when applicable
5. Run the existing tests to ensure nothing breaks:
   ```bash
   go test ./...
   ```
6. Commit your changes using Commit-AI itself:
   ```bash
   commit-ai
   ```
7. Push your branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
8. Open a pull request against the `master` branch
9. In your PR description, explain the changes and reference any related issues

### Development Guidelines

- Follow Go best practices and conventions
- Maintain backward compatibility when possible
- Document new features or changes in behavior
- Write tests for new functionality
- Keep performance in mind, especially with large repositories

### Adding New AI Providers

To add support for a new AI provider:

1. Implement the `Provider` interface in the `ai/ai.go` file
2. Add appropriate configuration options in `config/config.go`
3. Update the provider selection in `main.go`
4. Add documentation for the new provider

### Code of Conduct

- Be respectful and inclusive in your interactions
- Focus on constructive feedback
- Help maintain a welcoming environment for all contributors

## License

This project is licensed under the [MIT License](LICENSE). 