#!/bin/bash

# Cores para saída no terminal
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Definir idioma padrão (pt-br, en, es)
LANGUAGE="pt-br"

# Função para detectar idioma do sistema
detect_system_language() {
    local system_lang=$(echo $LANG | cut -d'_' -f1)
    
    case "$system_lang" in
        pt|pt-BR|pt-br|pt_BR|pt_br)
            LANGUAGE="pt-br"
            ;;
        es|es-ES|es-es|es_ES|es_es)
            LANGUAGE="es"
            ;;
        *)
            LANGUAGE="en"
            ;;
    esac
}

# Detectar idioma do sistema
detect_system_language

# Função para exibir textos com cores
echo_color() {
    echo -e "${1}${2}${NC}"
}

# Strings em português
load_pt_strings() {
    BANNER="
    ╔═══════════════════════════════════════════════╗
    ║                 COMMIT-AI                     ║
    ║       Instalação e Configuração Global        ║
    ╚═══════════════════════════════════════════════╝
    "
    
    ERR_ROOT="Este script não deve ser executado como root (sudo)!"
    ERR_ROOT_HINT="Por favor, execute-o como um usuário normal."
    
    DETECTING_OS="Detectando sistema operacional..."
    OS_DETECTED="Sistema operacional detectado:"
    OS_NOT_SUPPORTED="Sistema operacional não suportado:"
    
    CHECKING_DEPS="Verificando dependências..."
    GIT_NOT_FOUND="Git não encontrado! Por favor, instale o Git antes de continuar."
    DEPS_OK="Todas as dependências estão instaladas!"
    
    CREATING_DIR="Criando diretório de instalação..."
    DIR_CREATED="Diretório de instalação criado:"
    
    CHECKING_PATH="Verificando PATH..."
    DIR_NOT_IN_PATH="O diretório $INSTALL_DIR não está no seu PATH."
    ADDING_TO_PROFILE="Adicionando ao seu arquivo de perfil..."
    DIR_ADDED_TO_PATH="Diretório adicionado ao PATH em"
    DIR_IN_PATH="Diretório já está no PATH, prosseguindo..."
    
    INSTALLING="Instalando Commit-AI..."
    BINARY_NOT_FOUND="Binário não encontrado no diretório local."
    RUN_FROM_REPO="Este script deve ser executado a partir do diretório do repositório commit-ai."
    INSTALL_COMPLETE="Commit-AI instalado com sucesso em:"
    
    CONFIG_TITLE="Vamos configurar o Commit-AI para você!"
    CONFIG_PROMPT="Você gostaria de configurar o Commit-AI agora? [S/n]:"
    CONFIG_LATER="Você pode configurar manualmente mais tarde usando o comando: commit-ai --configure"
    PATH_UPDATED="Já que seu PATH foi atualizado, precisamos usar o caminho completo para o binário."
    CONFIG_SUCCESS="Configuração inicial concluída com sucesso!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║                   INSTALAÇÃO CONCLUÍDA!                        ║
╚═════════════════════════════════════════════════════════════════╝"
    
    SUCCESS_MSG="O Commit-AI foi instalado com sucesso em seu sistema!"
    
    RESTART_TERMINAL="Para começar a usar o Commit-AI, você precisa:"
    RESTART_TERMINAL_1="  1. Reiniciar seu terminal, OU"
    RESTART_TERMINAL_2="  2. Executar o comando: source $SHELL_FILE"
    
    BASIC_USAGE="Uso básico:"
    USAGE_1="  commit-ai               # Gerar mensagem de commit no diretório atual"
    USAGE_2="  commit-ai --dry-run     # Apenas mostrar a mensagem sem fazer commit"
    USAGE_3="  commit-ai --watch       # Modo de monitoramento contínuo"
    
    MORE_INFO="Para mais informações, visite: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="Obrigado por instalar o Commit-AI! 💻✨"
    
    # Prompts para seleção de idioma
    LANGUAGE_SELECTION="Selecione seu idioma / Select your language / Seleccione su idioma:"
    LANG_PT="Português"
    LANG_EN="Inglês"
    LANG_ES="Espanhol"
    LANG_PROMPT="Digite o número do idioma desejado [1-3]:"
}

# Strings em inglês
load_en_strings() {
    BANNER="
    ╔═══════════════════════════════════════════════╗
    ║                 COMMIT-AI                     ║
    ║        Global Installation and Setup          ║
    ╚═══════════════════════════════════════════════╝
    "
    
    ERR_ROOT="This script should not be run as root (sudo)!"
    ERR_ROOT_HINT="Please run it as a normal user."
    
    DETECTING_OS="Detecting operating system..."
    OS_DETECTED="Operating system detected:"
    OS_NOT_SUPPORTED="Operating system not supported:"
    
    CHECKING_DEPS="Checking dependencies..."
    GIT_NOT_FOUND="Git not found! Please install Git before continuing."
    DEPS_OK="All dependencies are installed!"
    
    CREATING_DIR="Creating installation directory..."
    DIR_CREATED="Installation directory created:"
    
    CHECKING_PATH="Checking PATH..."
    DIR_NOT_IN_PATH="Directory $INSTALL_DIR is not in your PATH."
    ADDING_TO_PROFILE="Adding to your profile file..."
    DIR_ADDED_TO_PATH="Directory added to PATH in"
    DIR_IN_PATH="Directory already in PATH, proceeding..."
    
    INSTALLING="Installing Commit-AI..."
    BINARY_NOT_FOUND="Binary not found in local directory."
    RUN_FROM_REPO="This script must be run from the commit-ai repository directory."
    INSTALL_COMPLETE="Commit-AI successfully installed at:"
    
    CONFIG_TITLE="Let's configure Commit-AI for you!"
    CONFIG_PROMPT="Would you like to configure Commit-AI now? [Y/n]:"
    CONFIG_LATER="You can manually configure later using the command: commit-ai --configure"
    PATH_UPDATED="Since your PATH was updated, we need to use the full path to the binary."
    CONFIG_SUCCESS="Initial configuration completed successfully!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║                 INSTALLATION COMPLETED!                      ║
╚═════════════════════════════════════════════════════════════════╝"
    
    SUCCESS_MSG="Commit-AI has been successfully installed on your system!"
    
    RESTART_TERMINAL="To start using Commit-AI, you need to:"
    RESTART_TERMINAL_1="  1. Restart your terminal, OR"
    RESTART_TERMINAL_2="  2. Run the command: source"
    
    BASIC_USAGE="Basic usage:"
    USAGE_1="  commit-ai               # Generate commit message in current directory"
    USAGE_2="  commit-ai --dry-run     # Just show the message without committing"
    USAGE_3="  commit-ai --watch       # Continuous monitoring mode"
    
    MORE_INFO="For more information, visit: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="Thank you for installing Commit-AI! 💻✨"
    
    # Language selection prompts
    LANGUAGE_SELECTION="Selecione seu idioma / Select your language / Seleccione su idioma:"
    LANG_PT="Portuguese"
    LANG_EN="English"
    LANG_ES="Spanish"
    LANG_PROMPT="Enter the number of your preferred language [1-3]:"
}

# Strings em espanhol
load_es_strings() {
    BANNER="
    ╔═══════════════════════════════════════════════╗
    ║                 COMMIT-AI                     ║
    ║       Instalación y Configuración Global      ║
    ╚═══════════════════════════════════════════════╝
    "
    
    ERR_ROOT="¡Este script no debe ejecutarse como root (sudo)!"
    ERR_ROOT_HINT="Por favor, ejecútelo como usuario normal."
    
    DETECTING_OS="Detectando sistema operativo..."
    OS_DETECTED="Sistema operativo detectado:"
    OS_NOT_SUPPORTED="Sistema operativo no soportado:"
    
    CHECKING_DEPS="Verificando dependencias..."
    GIT_NOT_FOUND="¡Git no encontrado! Por favor, instale Git antes de continuar."
    DEPS_OK="¡Todas las dependencias están instaladas!"
    
    CREATING_DIR="Creando directorio de instalación..."
    DIR_CREATED="Directorio de instalación creado:"
    
    CHECKING_PATH="Verificando PATH..."
    DIR_NOT_IN_PATH="El directorio $INSTALL_DIR no está en su PATH."
    ADDING_TO_PROFILE="Añadiendo a su archivo de perfil..."
    DIR_ADDED_TO_PATH="Directorio añadido al PATH en"
    DIR_IN_PATH="Directorio ya está en PATH, continuando..."
    
    INSTALLING="Instalando Commit-AI..."
    BINARY_NOT_FOUND="Binario no encontrado en el directorio local."
    RUN_FROM_REPO="Este script debe ejecutarse desde el directorio del repositorio commit-ai."
    INSTALL_COMPLETE="Commit-AI instalado exitosamente en:"
    
    CONFIG_TITLE="¡Vamos a configurar Commit-AI para usted!"
    CONFIG_PROMPT="¿Le gustaría configurar Commit-AI ahora? [S/n]:"
    CONFIG_LATER="Puede configurar manualmente más tarde usando el comando: commit-ai --configure"
    PATH_UPDATED="Como su PATH fue actualizado, necesitamos usar la ruta completa al binario."
    CONFIG_SUCCESS="¡Configuración inicial completada con éxito!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║                 ¡INSTALACIÓN COMPLETADA!                    ║
╚═════════════════════════════════════════════════════════════════╝"
    
    SUCCESS_MSG="¡Commit-AI ha sido instalado exitosamente en su sistema!"
    
    RESTART_TERMINAL="Para comenzar a usar Commit-AI, necesita:"
    RESTART_TERMINAL_1="  1. Reiniciar su terminal, O"
    RESTART_TERMINAL_2="  2. Ejecutar el comando: source"
    
    BASIC_USAGE="Uso básico:"
    USAGE_1="  commit-ai               # Generar mensaje de commit en el directorio actual"
    USAGE_2="  commit-ai --dry-run     # Solo mostrar el mensaje sin hacer commit"
    USAGE_3="  commit-ai --watch       # Modo de monitoreo continuo"
    
    MORE_INFO="Para más información, visite: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="¡Gracias por instalar Commit-AI! 💻✨"
    
    # Prompts de selección de idioma
    LANGUAGE_SELECTION="Selecione seu idioma / Select your language / Seleccione su idioma:"
    LANG_PT="Portugués"
    LANG_EN="Inglés"
    LANG_ES="Español"
    LANG_PROMPT="Introduzca el número del idioma deseado [1-3]:"
}

# Função para selecionar idioma manualmente
select_language() {
    echo
    echo_color $BLUE "Selecione seu idioma / Select your language / Seleccione su idioma:"
    echo
    echo_color $CYAN "1) Português"
    echo_color $CYAN "2) English"
    echo_color $CYAN "3) Español"
    echo
    echo_color $YELLOW "Digite o número do idioma desejado / Enter the language number / Introduzca el número del idioma [1-3]: "
    read -r lang_choice
    
    case $lang_choice in
        1) LANGUAGE="pt-br" ;;
        2) LANGUAGE="en" ;;
        3) LANGUAGE="es" ;;
        *) echo "Idioma inválido, usando Português / Invalid language, using English / Idioma inválido, usando Español" 
           LANGUAGE="pt-br" ;;
    esac
}

# Solicitar seleção de idioma
select_language

# Carregar strings do idioma selecionado
case $LANGUAGE in
    "pt-br") load_pt_strings ;;
    "en") load_en_strings ;;
    "es") load_es_strings ;;
    *) load_en_strings ;;
esac

# Função para exibir o banner
show_banner() {
    echo_color $BLUE "$BANNER"
}

# Verificar se é root
check_root() {
    if [ "$EUID" -eq 0 ]; then
        echo_color $RED "$ERR_ROOT"
        echo_color $YELLOW "$ERR_ROOT_HINT"
        exit 1
    fi
}

# Detectar sistema operacional
detect_os() {
    echo_color $BLUE "$DETECTING_OS"
    
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
        echo_color $RED "$OS_NOT_SUPPORTED $OSTYPE"
        exit 1
    fi
    
    echo_color $GREEN "$OS_DETECTED $OS"
}

# Verificar dependências
check_dependencies() {
    echo_color $BLUE "$CHECKING_DEPS"
    
    if ! command -v git &> /dev/null; then
        echo_color $RED "$GIT_NOT_FOUND"
        exit 1
    fi
    
    echo_color $GREEN "$DEPS_OK"
}

# Criar diretório de instalação
create_install_dir() {
    echo_color $BLUE "$CREATING_DIR"
    
    if [ "$OS" == "windows" ]; then
        INSTALL_DIR="$HOME/bin"
    else
        INSTALL_DIR="$HOME/.local/bin"
    fi
    
    mkdir -p "$INSTALL_DIR"
    
    echo_color $GREEN "$DIR_CREATED $INSTALL_DIR"
}

# Verificar se o diretório está no PATH
check_path() {
    echo_color $BLUE "$CHECKING_PATH"
    
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo_color $YELLOW "$DIR_NOT_IN_PATH"
        echo_color $YELLOW "$ADDING_TO_PROFILE"
        
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
        
        echo_color $GREEN "$DIR_ADDED_TO_PATH $SHELL_FILE"
        PATH_UPDATED=1
    else
        echo_color $GREEN "$DIR_IN_PATH"
        PATH_UPDATED=0
    fi
}

# Instalar o binário
install_binary() {
    echo_color $BLUE "$INSTALLING"
    
    # Verificar se estamos no diretório do repositório
    if [ -f "bin/$BINARY" ]; then
        # Estamos no repositório, copiar o binário diretamente
        SRC_PATH="bin/$BINARY"
    else
        echo_color $RED "$BINARY_NOT_FOUND"
        echo_color $RED "$RUN_FROM_REPO"
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
    
    echo_color $GREEN "$INSTALL_COMPLETE $DEST_PATH"
}

# Configuração inicial assistida
run_configuration() {
    echo_color $BLUE "$CONFIG_TITLE"
    echo_color $YELLOW "$CONFIG_PROMPT"
    read -r configure_now
    
    if [[ "$LANGUAGE" == "pt-br" && "$configure_now" =~ ^[Nn]$ ]] || 
       [[ "$LANGUAGE" == "en" && "$configure_now" =~ ^[Nn]$ ]] || 
       [[ "$LANGUAGE" == "es" && "$configure_now" =~ ^[Nn]$ ]]; then
        echo_color $YELLOW "$CONFIG_LATER"
        return
    fi
    
    if [ $PATH_UPDATED -eq 1 ]; then
        echo_color $YELLOW "$PATH_UPDATED"
        "$DEST_PATH" --configure
    else
        commit-ai --configure
    fi
    
    echo_color $GREEN "$CONFIG_SUCCESS"
}

# Mostrar instruções finais
show_final_instructions() {
    echo_color $BLUE "$FINAL_HEADER"
    echo
    echo_color $GREEN "$SUCCESS_MSG"
    echo
    
    if [ $PATH_UPDATED -eq 1 ]; then
        echo_color $YELLOW "$RESTART_TERMINAL"
        echo_color $YELLOW "$RESTART_TERMINAL_1"
        echo_color $YELLOW "$RESTART_TERMINAL_2 $SHELL_FILE"
        echo
    fi
    
    echo_color $GREEN "$BASIC_USAGE"
    echo_color $NC "$USAGE_1"
    echo_color $NC "$USAGE_2"
    echo_color $NC "$USAGE_3"
    echo
    echo_color $YELLOW "$MORE_INFO"
    echo
    echo_color $BLUE "$THANK_YOU"
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