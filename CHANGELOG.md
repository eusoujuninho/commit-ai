# Changelog

All notable changes to the Commit-AI project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2024-03-16

### Added
- Watcher mode for continuous monitoring and automatic commits
- User-friendly installation script (`install.sh`)
- Interactive setup wizard for easy configuration (`setup-wizard.sh`)
- Support for ignoring specific files/patterns in watcher mode
- Improved documentation with detailed examples for all new features
- Better handling of repository path resolution

### Changed
- Updated repository references to point to the correct GitHub account
- Improved error handling in file monitoring
- Enhanced configuration process with clearer instructions

## [0.1.0] - 2024-03-15

### Added
- Initial release of Commit-AI
- Support for multiple AI providers:
  - Google Gemini (default)
  - OpenAI
  - Claude (Anthropic)
  - DeepSeek
  - OpenRouter
  - Grok (xAI)
  - Ollama (local models)
- Multilingual support:
  - Portuguese (pt-br) - Default
  - English (en)
  - Spanish (es)
  - French (fr)
  - German (de)
- Command-line interface with multiple options:
  - Configuration mode (`--configure`)
  - Dry-run mode (`--dry-run`)
  - Repository path specification (`--repo=PATH`)
  - Provider selection (`--provider=NAME`)
  - Language selection (`--language=CODE`)
  - Version information (`--version`)
- Pre-compiled binaries for major operating systems:
  - Linux (AMD64)
  - macOS (AMD64)
  - Windows (AMD64)
- Comprehensive documentation:
  - Installation guides
  - Usage examples
  - Contribution guidelines
- Automatic detection of Git repository changes
- Support for both staged and unstaged changes
- Implementation of the Conventional Commits format
- Persistent configuration through JSON file
- Customizable settings 