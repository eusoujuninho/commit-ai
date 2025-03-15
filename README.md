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
- Features continuous monitoring with automatic commits (Watcher mode)
- User-friendly installation and configuration wizards

## Installation

### Prerequisites

- Git installed and configured
- For building from source: Go 1.16 or later
- API key for at least one of the supported AI providers (not needed for Ollama)

### Easy Installation (Recommended)

For a quick and user-friendly installation:

```bash
# Clone the repository
git clone https://github.com/eusoujuninho/commit-ai.git

# Navigate to the directory
cd commit-ai

# Run the installation script
./install.sh

# After installation, run the setup wizard for guided configuration
./setup-wizard.sh
```

The installation script will automatically:
- Detect your operating system
- Install the appropriate binary globally
- Add it to your PATH
- Guide you through initial configuration

### Pre-compiled Binaries

For convenience, we provide pre-compiled binaries for major operating systems:

- **Linux**: `bin/commit-ai-linux-amd64`
- **macOS**: `bin/commit-ai-darwin-amd64`
- **Windows**: `bin/commit-ai-windows-amd64.exe`

You don't need to have Go installed to use the pre-compiled binaries. Simply download the appropriate binary for your operating system, make it executable (Linux/macOS), and move it to a directory in your PATH.

### Building from Source

```bash
# Clone the repository
git clone https://github.com/eusoujuninho/commit-ai.git

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

For a more user-friendly guided setup:

```bash
./setup-wizard.sh
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
- `--watch`: Activates watcher mode for continuous monitoring
- `--interval=DURATION`: Sets the time between checks in watcher mode (default: 30s)
- `--min-changes=N`: Minimum number of changes to trigger a commit (default: 1)
- `--ignore=PATTERNS`: List of patterns to ignore, separated by commas
- `--silent`: Silent mode for watcher (less output)

### Watcher Mode

Commit-AI can continuously monitor your repository and automatically make commits when changes are detected:

```bash
# Basic watcher mode with default settings
commit-ai --watch

# Custom configuration with longer interval and more changes required
commit-ai --watch --interval=5m --min-changes=3

# Watcher mode with specific provider and language
commit-ai --watch --provider=claude --language=en

# Ignore specific file patterns
commit-ai --watch --ignore=*.log,tmp/,node_modules/

# Silent mode with less console output
commit-ai --watch --silent
```

The watcher mode is perfect for:
- Continuous development sessions where you want automatic versioning
- Ensuring regular commits without manual intervention
- Creating detailed commit history automatically
- Projects where you prefer frequent, smaller commits

Watcher mode can be combined with `--dry-run` to test the behavior without making actual commits.

### Installation Wizard

The installation wizard (`install.sh`) provides a seamless experience for installing Commit-AI by:

- Automatically detecting your operating system
- Installing the appropriate binary for your platform
- Setting up the correct permissions
- Adding Commit-AI to your PATH
- Guiding you through initial configuration

### Setup Wizard

The setup wizard (`setup-wizard.sh`) offers a friendly, interactive configuration experience:

- Visual, step-by-step configuration process
- Clear explanations of all available options
- Simple selection of AI providers and languages
- Easy API key configuration for various providers
- Guided setup of repository paths and automatic commit options
- Helpful tips for using Commit-AI after setup

## Contributing

We welcome contributions to Commit-AI! Here's how you can help:

### Reporting Issues

If you find a bug or have a suggestion for improvement:

1. Check if the issue already exists in the [Issues](https://github.com/eusoujuninho/commit-ai/issues) section
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