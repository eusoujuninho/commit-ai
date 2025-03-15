#!/bin/bash

# Cores para saída no terminal
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Função para exibir textos com cores
echo_color() {
    echo -e "${1}${2}${NC}"
}

# Exibir cabeçalho
show_header() {
    clear
    echo_color $BLUE "
    ╔═══════════════════════════════════════════════╗
    ║               COMMIT-AI WIZARD                ║
    ║          Assistente de Configuração           ║
    ╚═══════════════════════════════════════════════╝
    "
    echo_color $CYAN "Este assistente irá guiá-lo pela configuração do Commit-AI"
    echo_color $CYAN "para que você possa começar a usar rapidamente."
    echo
}

# Verificar se o commit-ai está instalado
check_installation() {
    if ! command -v commit-ai &> /dev/null; then
        echo_color $RED "Commit-AI não foi encontrado no seu PATH."
        echo_color $YELLOW "Por favor, instale-o primeiro usando o script install.sh"
        exit 1
    fi
    
    echo_color $GREEN "✓ Commit-AI está instalado e pronto para ser configurado."
    echo
}

# Introdução e explicação das opções de configuração
explain_configuration() {
    echo_color $MAGENTA "Como funciona o Commit-AI:"
    echo_color $NC "O Commit-AI analisa as alterações em seu repositório Git e"
    echo_color $NC "usa inteligência artificial para gerar mensagens de commit"
    echo_color $NC "significativas seguindo o padrão Conventional Commits."
    echo
    echo_color $MAGENTA "Provedores de IA disponíveis:"
    echo_color $CYAN "- gemini (Google Gemini) - Padrão, API key fornecida"
    echo_color $CYAN "- openai (OpenAI GPT) - Requer sua própria API key"
    echo_color $CYAN "- claude (Anthropic Claude) - Requer sua própria API key"
    echo_color $CYAN "- deepseek (DeepSeek) - Requer sua própria API key"
    echo_color $CYAN "- openrouter (OpenRouter) - Requer sua própria API key"
    echo_color $CYAN "- grok (xAI Grok) - Requer sua própria API key"
    echo_color $CYAN "- ollama (Modelos locais) - Não requer API key, precisa ter Ollama instalado"
    echo
    echo_color $MAGENTA "Idiomas suportados:"
    echo_color $CYAN "- pt-br (Português Brasil) - Padrão"
    echo_color $CYAN "- en (Inglês)"
    echo_color $CYAN "- es (Espanhol)"
    echo_color $CYAN "- fr (Francês)"
    echo_color $CYAN "- de (Alemão)"
    echo
    
    echo_color $YELLOW "Pressione ENTER para continuar..."
    read -r
}

# Selecionar provedor de IA
select_provider() {
    show_header
    echo_color $MAGENTA "1. Selecione seu provedor de IA preferido:"
    echo
    echo_color $CYAN "1) gemini  - Google Gemini (padrão, API key fornecida)"
    echo_color $CYAN "2) openai  - OpenAI GPT (requer sua própria API key)"
    echo_color $CYAN "3) claude  - Anthropic Claude (requer sua própria API key)"
    echo_color $CYAN "4) deepseek - DeepSeek (requer sua própria API key)"
    echo_color $CYAN "5) openrouter - OpenRouter (requer sua própria API key)"
    echo_color $CYAN "6) grok    - xAI Grok (requer sua própria API key)"
    echo_color $CYAN "7) ollama  - Modelos locais (não requer API key)"
    echo
    echo_color $YELLOW "Digite o número do provedor [1-7] (ou pressione ENTER para padrão): "
    read -r provider_choice
    
    case $provider_choice in
        1|"") PROVIDER="gemini" ;;
        2) PROVIDER="openai" ;;
        3) PROVIDER="claude" ;;
        4) PROVIDER="deepseek" ;;
        5) PROVIDER="openrouter" ;;
        6) PROVIDER="grok" ;;
        7) PROVIDER="ollama" ;;
        *)
            echo_color $RED "Opção inválida. Usando provedor padrão (gemini)."
            PROVIDER="gemini"
            ;;
    esac
    
    echo_color $GREEN "✓ Provedor selecionado: $PROVIDER"
    echo
    
    # Se for ollama, configurar URL e modelo
    if [ "$PROVIDER" == "ollama" ]; then
        echo_color $YELLOW "URL do servidor Ollama [http://localhost:11434]: "
        read -r ollama_url
        if [ -z "$ollama_url" ]; then
            OLLAMA_URL="http://localhost:11434"
        else
            OLLAMA_URL="$ollama_url"
        fi
        
        echo_color $YELLOW "Modelo Ollama a ser usado [llama3]: "
        read -r ollama_model
        if [ -z "$ollama_model" ]; then
            OLLAMA_MODEL="llama3"
        else
            OLLAMA_MODEL="$ollama_model"
        fi
    # Se não for gemini nem ollama, solicitar API key
    elif [ "$PROVIDER" != "gemini" ]; then
        echo_color $YELLOW "Digite sua API key para $PROVIDER: "
        read -r api_key
        if [ -z "$api_key" ]; then
            echo_color $RED "API key é obrigatória para $PROVIDER. Usando provedor padrão (gemini)."
            PROVIDER="gemini"
        else
            API_KEY="$api_key"
        fi
    fi
    
    echo_color $YELLOW "Pressione ENTER para continuar..."
    read -r
}

# Selecionar idioma
select_language() {
    show_header
    echo_color $MAGENTA "2. Selecione o idioma para as mensagens de commit:"
    echo
    echo_color $CYAN "1) pt-br - Português Brasil (padrão)"
    echo_color $CYAN "2) en    - Inglês"
    echo_color $CYAN "3) es    - Espanhol"
    echo_color $CYAN "4) fr    - Francês"
    echo_color $CYAN "5) de    - Alemão"
    echo
    echo_color $YELLOW "Digite o número do idioma [1-5] (ou pressione ENTER para padrão): "
    read -r language_choice
    
    case $language_choice in
        1|"") LANGUAGE="pt-br" ;;
        2) LANGUAGE="en" ;;
        3) LANGUAGE="es" ;;
        4) LANGUAGE="fr" ;;
        5) LANGUAGE="de" ;;
        *)
            echo_color $RED "Opção inválida. Usando idioma padrão (pt-br)."
            LANGUAGE="pt-br"
            ;;
    esac
    
    echo_color $GREEN "✓ Idioma selecionado: $LANGUAGE"
    echo
    
    echo_color $YELLOW "Pressione ENTER para continuar..."
    read -r
}

# Configurar commit automático
configure_auto_commit() {
    show_header
    echo_color $MAGENTA "3. Você deseja habilitar commit automático?"
    echo
    echo_color $CYAN "Com o commit automático habilitado, o Commit-AI"
    echo_color $CYAN "irá fazer commits automaticamente sem pedir confirmação."
    echo_color $CYAN "Isso é útil para fluxos de trabalho automatizados."
    echo
    echo_color $YELLOW "Habilitar commit automático? (S/n): "
    read -r auto_commit_choice
    
    if [[ "$auto_commit_choice" =~ ^[Nn]$ ]]; then
        AUTO_COMMIT="false"
    else
        AUTO_COMMIT="true"
    fi
    
    echo_color $GREEN "✓ Commit automático: $AUTO_COMMIT"
    echo
    
    echo_color $YELLOW "Pressione ENTER para continuar..."
    read -r
}

# Configurar caminho padrão do repositório
configure_repo_path() {
    show_header
    echo_color $MAGENTA "4. Configurar caminho padrão do repositório"
    echo
    echo_color $CYAN "Você pode definir um caminho padrão para o repositório Git"
    echo_color $CYAN "que o Commit-AI usará se não for especificado."
    echo_color $CYAN "Deixe em branco para usar o diretório atual como padrão."
    echo
    echo_color $YELLOW "Caminho padrão do repositório: "
    read -r repo_path
    
    REPO_PATH="$repo_path"
    
    if [ -z "$REPO_PATH" ]; then
        echo_color $GREEN "✓ Usando diretório atual como padrão"
    else
        echo_color $GREEN "✓ Caminho padrão definido: $REPO_PATH"
    fi
    echo
    
    echo_color $YELLOW "Pressione ENTER para continuar..."
    read -r
}

# Aplicar configurações
apply_configuration() {
    show_header
    echo_color $MAGENTA "Resumo da configuração:"
    echo
    echo_color $CYAN "Provedor de IA: $PROVIDER"
    if [ "$PROVIDER" == "ollama" ]; then
        echo_color $CYAN "URL do Ollama: $OLLAMA_URL"
        echo_color $CYAN "Modelo Ollama: $OLLAMA_MODEL"
    elif [ "$PROVIDER" != "gemini" ]; then
        echo_color $CYAN "API Key: ****$(echo $API_KEY | tail -c 5)"
    fi
    echo_color $CYAN "Idioma: $LANGUAGE"
    echo_color $CYAN "Commit automático: $AUTO_COMMIT"
    if [ -n "$REPO_PATH" ]; then
        echo_color $CYAN "Caminho padrão: $REPO_PATH"
    else
        echo_color $CYAN "Caminho padrão: <diretório atual>"
    fi
    echo
    echo_color $YELLOW "Confirmar e aplicar estas configurações? (S/n): "
    read -r confirm
    
    if [[ "$confirm" =~ ^[Nn]$ ]]; then
        echo_color $RED "Configuração cancelada pelo usuário."
        exit 0
    fi
    
    echo_color $BLUE "Aplicando configurações..."
    
    # Construir comando de configuração baseado nas escolhas
    CONFIG_CMD="commit-ai --configure"
    
    # Usar expect para automatizar a entrada
    EXPECT_SCRIPT="$(
        cat <<EOT
#!/usr/bin/expect -f
set timeout -1
spawn $CONFIG_CMD

# Provedor de IA
expect "Provedor de IA"
send "$PROVIDER\r"

# API Key (dependendo do provedor)
if {"$PROVIDER" == "openai"} {
    expect "Chave da API OpenAI"
    send "$API_KEY\r"
} elseif {"$PROVIDER" == "claude"} {
    expect "Chave da API Claude"
    send "$API_KEY\r"
} elseif {"$PROVIDER" == "deepseek"} {
    expect "Chave da API DeepSeek"
    send "$API_KEY\r"
} elseif {"$PROVIDER" == "openrouter"} {
    expect "Chave da API OpenRouter"
    send "$API_KEY\r"
} elseif {"$PROVIDER" == "grok"} {
    expect "Chave da API Grok"
    send "$API_KEY\r"
} elseif {"$PROVIDER" == "ollama"} {
    expect "URL do servidor Ollama"
    send "$OLLAMA_URL\r"
    expect "Modelo Ollama"
    send "$OLLAMA_MODEL\r"
}

# Idioma
expect "Idioma para mensagens"
send "$LANGUAGE\r"

# Caminho do repositório
expect "Caminho padrão do repositório"
send "$REPO_PATH\r"

# Commit automático
expect "Commit automático"
send "$AUTO_COMMIT\r"

expect eof
EOT
    )"
    
    # Verificar se expect está instalado
    if ! command -v expect &> /dev/null; then
        echo_color $RED "O programa 'expect' não está instalado, mas é necessário para configuração automática."
        echo_color $YELLOW "Por favor, instale-o com: sudo apt-get install expect (Debian/Ubuntu)"
        echo_color $YELLOW "Ou execute manualmente: commit-ai --configure"
        exit 1
    fi
    
    # Salvar o script expect temporariamente
    TEMP_EXPECT=$(mktemp)
    echo "$EXPECT_SCRIPT" > "$TEMP_EXPECT"
    chmod +x "$TEMP_EXPECT"
    
    # Executar o script expect
    "$TEMP_EXPECT"
    
    # Limpar arquivo temporário
    rm "$TEMP_EXPECT"
    
    echo_color $GREEN "✓ Configuração aplicada com sucesso!"
}

# Mostrar dicas e finalizar
show_tips_and_finish() {
    show_header
    echo_color $MAGENTA "🎉 Configuração concluída! 🎉"
    echo
    echo_color $CYAN "Dicas rápidas para usar o Commit-AI:"
    echo
    echo_color $GREEN "Commit básico:"
    echo_color $NC "  commit-ai"
    echo
    echo_color $GREEN "Visualizar mensagem sem commitar:"
    echo_color $NC "  commit-ai --dry-run"
    echo
    echo_color $GREEN "Usar outro provedor temporariamente:"
    echo_color $NC "  commit-ai --provider=openai"
    echo
    echo_color $GREEN "Usar outro idioma temporariamente:"
    echo_color $NC "  commit-ai --language=en"
    echo
    echo_color $GREEN "Modo contínuo (watcher):"
    echo_color $NC "  commit-ai --watch"
    echo
    echo_color $GREEN "Watcher com configuração personalizada:"
    echo_color $NC "  commit-ai --watch --interval=5m --min-changes=3"
    echo
    echo_color $BLUE "Obrigado por usar o Commit-AI! Aproveite a ferramenta! 💻✨"
    echo
}

# Função principal
main() {
    show_header
    check_installation
    explain_configuration
    select_provider
    select_language
    configure_auto_commit
    configure_repo_path
    apply_configuration
    show_tips_and_finish
}

# Executar o script
main 