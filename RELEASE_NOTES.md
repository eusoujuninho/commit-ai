# Release Notes - Commit-AI v0.1.0

We're excited to announce the first stable release of Commit-AI, a Golang tool that automates Git commit messages using Artificial Intelligence.

## Overview

Commit-AI analyzes changes made in a Git repository and generates meaningful commit messages in Conventional Commits format using AI. The tool aims to improve the consistency and quality of commit messages, saving developers time and effort.

## Key Features

### Multiple AI Providers

Commit-AI supports a wide range of AI providers, giving you the flexibility to choose the one that best fits your needs:

- **Google Gemini** (default): Fast and effective for most commit message generation
- **OpenAI**: Leverages GPT models for high-quality commit messages
- **Claude (Anthropic)**: Known for strong natural language understanding
- **DeepSeek**: Specialized in code-aware models
- **OpenRouter**: Unified access to various AI models
- **Grok (xAI)**: New offering with strong code understanding
- **Ollama**: Local models for privacy and offline usage

### Multilingual Support

Generate commit messages in multiple languages:

- Portuguese (pt-br) - Default
- English (en)
- Spanish (es)
- French (fr)
- German (de)

### User-Friendly Interface

- Intuitive command-line interface
- Customizable configuration
- Dry-run mode to preview messages
- Support for repository path specification

### Pre-compiled Binaries

Ready-to-use binaries for all major platforms:

- Linux (AMD64)
- macOS (AMD64)
- Windows (AMD64)

## Installation

Check the README.md and DOCUMENTATION.md files for detailed installation instructions.

## Getting Started

After installation, you can easily configure Commit-AI:

```bash
commit-ai --configure
```

And start using it in your Git repository:

```bash
commit-ai
```

## Feedback and Contributions

We welcome your feedback and contributions! Please check our contribution guidelines in the README.md file. 