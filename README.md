# Commit-AI

Um script em Golang que automatiza a criação de mensagens de commit utilizando Inteligência Artificial.

## Descrição

O Commit-AI analisa as alterações feitas em um repositório Git e gera mensagens de commit significativas no formato de Conventional Commits usando Inteligência Artificial. A ferramenta visa melhorar a consistência e qualidade das mensagens de commit, economizando tempo dos desenvolvedores.

## Recursos

- Detecta automaticamente arquivos modificados no repositório
- Gera mensagens de commit no padrão Conventional Commits
- Integra com a API da OpenAI (suporte para outros provedores em desenvolvimento)
- Interface de linha de comando intuitiva
- Configuração personalizável e persistente
- Modo "dry-run" para previsualizar a mensagem antes de fazer o commit

## Instalação

### Pré-requisitos

- Go 1.16 ou superior
- Git instalado e configurado
- Conta e API key em um serviço de IA (OpenAI, por exemplo)

### Compilação e instalação

```bash
# Clone o repositório
git clone https://github.com/user/commit-ai.git

# Navegue até o diretório
cd commit-ai

# Compile e instale
go install
```

## Uso

### Configuração inicial

Na primeira execução, configure a ferramenta:

```bash
commit-ai --configure
```

### Uso básico

```bash
# No diretório do seu repositório Git
commit-ai
```

### Opções

- `--configure`: Modo de configuração
- `--dry-run`: Gera a mensagem sem fazer o commit
- `--repo=PATH`: Especifica um caminho para o repositório

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE). 