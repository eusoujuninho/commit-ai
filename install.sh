#!/bin/bash

# Cores para saída no terminal
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Função para exibir textos com cores
echo_color() {
    echo -e "${1}${2}${NC}"
}

# Função para exibir o banner
show_banner() {
    echo_color $BLUE "
    ╔═══════════════════════════════════════════════╗
    ║                 COMMIT-AI                     ║
    ║       Instalação e Configuração Global        ║
    ╚═══════════════════════════════════════════════╝
    "
}

# Verificar se é root
check_root() {
    if [ "$EUID" -eq 0 ]; then
        echo_color $RED "Este script não deve ser executado como root (sudo)!"
        echo_color $YELLOW "Por favor, execute-o como um usuário normal."
        exit 1
    fi
}

# Detectar sistema operacional
detect_os() {
    echo_color $BLUE "Detectando sistema operacional..."
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        OS="linux"
        BINARY="commit-ai-linux-amd64"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        OS="darwin"
        BINARY="commit-ai-darwin-amd64"
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
        OS="windows"
        BINARY="commit-ai-windows-amd64.exe"
    else
        echo_color $RED "Sistema operacional não suportado: $OSTYPE"
        exit 1
    fi
    
    echo_color $GREEN "Sistema operacional detectado: $OS"
}

# Verificar dependências
check_dependencies() {
    echo_color $BLUE "Verificando dependências..."
    
    if ! command -v git &> /dev/null; then
        echo_color $RED "Git não encontrado! Por favor, instale o Git antes de continuar."
        exit 1
    fi
    
    echo_color $GREEN "Todas as dependências estão instaladas!"
}

# Criar diretório de instalação
create_install_dir() {
    echo_color $BLUE "Criando diretório de instalação..."
    
    if [ "$OS" == "windows" ]; then
        INSTALL_DIR="$HOME/bin"
    else
        INSTALL_DIR="$HOME/.local/bin"
    fi
    
    mkdir -p "$INSTALL_DIR"
    
    echo_color $GREEN "Diretório de instalação criado: $INSTALL_DIR"
}

# Verificar se o diretório está no PATH
check_path() {
    echo_color $BLUE "Verificando PATH..."
    
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo_color $YELLOW "O diretório $INSTALL_DIR não está no seu PATH."
        echo_color $YELLOW "Adicionando ao seu arquivo de perfil..."
        
        if [ -f "$HOME/.zshrc" ]; then
            echo "export PATH=\$PATH:$INSTALL_DIR" >> "$HOME/.zshrc"
            SHELL_FILE="$HOME/.zshrc"
        elif [ -f "$HOME/.bashrc" ]; then
            echo "export PATH=\$PATH:$INSTALL_DIR" >> "$HOME/.bashrc"
            SHELL_FILE="$HOME/.bashrc"
        elif [ -f "$HOME/.bash_profile" ]; then
            echo "export PATH=\$PATH:$INSTALL_DIR" >> "$HOME/.bash_profile"
            SHELL_FILE="$HOME/.bash_profile"
        else
            echo "export PATH=\$PATH:$INSTALL_DIR" >> "$HOME/.profile"
            SHELL_FILE="$HOME/.profile"
        fi
        
        echo_color $GREEN "Diretório adicionado ao PATH em $SHELL_FILE"
        PATH_UPDATED=1
    else
        echo_color $GREEN "Diretório já está no PATH, prosseguindo..."
        PATH_UPDATED=0
    fi
}

# Instalar o binário
install_binary() {
    echo_color $BLUE "Instalando Commit-AI..."
    
    # Verificar se estamos no diretório do repositório
    if [ -f "bin/$BINARY" ]; then
        # Estamos no repositório, copiar o binário diretamente
        SRC_PATH="bin/$BINARY"
    else
        echo_color $RED "Binário não encontrado no diretório local."
        echo_color $RED "Este script deve ser executado a partir do diretório do repositório commit-ai."
        exit 1
    fi
    
    # Copiar o binário para o diretório de instalação
    if [ "$OS" == "windows" ]; then
        DEST_PATH="$INSTALL_DIR/commit-ai.exe"
    else
        DEST_PATH="$INSTALL_DIR/commit-ai"
    fi
    
    cp "$SRC_PATH" "$DEST_PATH"
    
    # Tornar o binário executável (não necessário no Windows)
    if [ "$OS" != "windows" ]; then
        chmod +x "$DEST_PATH"
    fi
    
    echo_color $GREEN "Commit-AI instalado com sucesso em: $DEST_PATH"
}

# Configuração inicial assistida
run_configuration() {
    echo_color $BLUE "Vamos configurar o Commit-AI para você!"
    echo_color $YELLOW "Você gostaria de configurar o Commit-AI agora? [S/n]: "
    read -r configure_now
    
    if [[ "$configure_now" =~ ^[Nn]$ ]]; then
        echo_color $YELLOW "Você pode configurar manualmente mais tarde usando o comando: commit-ai --configure"
        return
    fi
    
    if [ $PATH_UPDATED -eq 1 ]; then
        echo_color $YELLOW "Já que seu PATH foi atualizado, precisamos usar o caminho completo para o binário."
        "$DEST_PATH" --configure
    else
        commit-ai --configure
    fi
    
    echo_color $GREEN "Configuração inicial concluída com sucesso!"
}

# Mostrar instruções finais
show_final_instructions() {
    echo_color $BLUE "╔═════════════════════════════════════════════════════════════════╗"
    echo_color $BLUE "║                   INSTALAÇÃO CONCLUÍDA!                        ║"
    echo_color $BLUE "╚═════════════════════════════════════════════════════════════════╝"
    echo
    echo_color $GREEN "O Commit-AI foi instalado com sucesso em seu sistema!"
    echo
    
    if [ $PATH_UPDATED -eq 1 ]; then
        echo_color $YELLOW "Para começar a usar o Commit-AI, você precisa:"
        echo_color $YELLOW "  1. Reiniciar seu terminal, OU"
        echo_color $YELLOW "  2. Executar o comando: source $SHELL_FILE"
        echo
    fi
    
    echo_color $GREEN "Uso básico:"
    echo_color $NC "  commit-ai               # Gerar mensagem de commit no diretório atual"
    echo_color $NC "  commit-ai --dry-run     # Apenas mostrar a mensagem sem fazer commit"
    echo_color $NC "  commit-ai --watch       # Modo de monitoramento contínuo"
    echo
    echo_color $YELLOW "Para mais informações, visite: https://github.com/user/commit-ai"
    echo
    echo_color $BLUE "Obrigado por instalar o Commit-AI! 💻✨"
}

# Execução principal
main() {
    show_banner
    check_root
    detect_os
    check_dependencies
    create_install_dir
    check_path
    install_binary
    run_configuration
    show_final_instructions
}

# Iniciar o script
main 