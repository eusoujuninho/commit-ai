#!/bin/bash

# Cores para saída no terminal
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
BOLD='\033[1m'
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
    ║          Assistente de Configuração           ║
    ╚═══════════════════════════════════════════════╝
    "
    
    WELCOME="Bem-vindo ao assistente de configuração do Commit-AI!"
    WIZARD_DESC="Este assistente irá ajudá-lo a configurar o Commit-AI para o seu ambiente."
    
    STEP_PROVIDER="PASSO 1: Escolha o provedor de IA"
    PROVIDER_DESC="O Commit-AI suporta vários provedores de IA para gerar mensagens de commit."
    PROVIDER_OPTIONS="Por favor, escolha um provedor de IA:"
    
    PROVIDER_OPENAI="OpenAI (requer chave API)"
    PROVIDER_GEMINI="Google Gemini (requer chave API, padrão)"
    PROVIDER_CLAUDE="Claude/Anthropic (requer chave API)"
    PROVIDER_DEEPSEEK="DeepSeek (requer chave API)"
    PROVIDER_OPENROUTER="OpenRouter (requer chave API)"
    PROVIDER_GROK="Grok/xAI (requer chave API)"
    PROVIDER_OLLAMA="Ollama (local, sem chave API)"
    
    PROVIDER_PROMPT="Digite o número do provedor desejado [1-7] (padrão: 2):"
    
    STEP_API_KEY="PASSO 2: Configure a chave API"
    API_KEY_DESC="Você selecionou um provedor que requer uma chave API."
    API_KEY_INSTRUCTIONS="Digite sua chave API para"
    API_KEY_SKIP="Pressione ENTER para pular (você poderá configurar depois)"
    
    OLLAMA_SELECTED="Você selecionou Ollama como seu provedor. Nenhuma chave API é necessária!"
    
    STEP_LANGUAGE="PASSO 3: Escolha o idioma preferido"
    LANGUAGE_DESC="Escolha o idioma para as mensagens de commit geradas:"
    
    LANG_OPTION_PT="Português (pt-br)"
    LANG_OPTION_EN="Inglês (en)"
    LANG_OPTION_ES="Espanhol (es)"
    LANG_OPTION_FR="Francês (fr)"
    LANG_OPTION_DE="Alemão (de)"
    
    LANGUAGE_PROMPT="Digite o número do idioma desejado [1-5] (padrão: 1):"
    
    STEP_REPO="PASSO 4: Defina o caminho padrão do repositório"
    REPO_DESC="Você pode definir um caminho padrão para o repositório Git:"
    REPO_CURRENT="Seu diretório atual é:"
    REPO_PROMPT="Digite o caminho padrão do repositório ou deixe em branco para usar o diretório atual:"
    
    STEP_AUTO_COMMIT="PASSO 5: Configurar commit automático"
    AUTO_COMMIT_DESC="O Commit-AI pode fazer commits automaticamente após gerar mensagens."
    AUTO_COMMIT_PROMPT="Ativar commit automático? [s/N]:"
    
    SAVING_CONFIG="Salvando configuração..."
    CONFIG_SAVED="Configuração salva com sucesso!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║              CONFIGURAÇÃO CONCLUÍDA!                         ║
╚═════════════════════════════════════════════════════════════════╝"
    
    NEXT_STEPS="Próximos passos:"
    BASIC_USAGE="Uso básico:"
    USAGE_BASIC="  commit-ai               # Gerar mensagem de commit no diretório atual"
    USAGE_DRY_RUN="  commit-ai --dry-run     # Apenas mostrar a mensagem sem fazer commit"
    USAGE_WATCH="  commit-ai --watch       # Modo de monitoramento contínuo"
    
    MORE_INFO="Para mais informações, visite: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="Obrigado por configurar o Commit-AI! 💻✨"
    
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
    ║              Setup Wizard                     ║
    ╚═══════════════════════════════════════════════╝
    "
    
    WELCOME="Welcome to the Commit-AI setup wizard!"
    WIZARD_DESC="This wizard will help you configure Commit-AI for your environment."
    
    STEP_PROVIDER="STEP 1: Choose AI Provider"
    PROVIDER_DESC="Commit-AI supports several AI providers to generate commit messages."
    PROVIDER_OPTIONS="Please choose an AI provider:"
    
    PROVIDER_OPENAI="OpenAI (requires API key)"
    PROVIDER_GEMINI="Google Gemini (requires API key, default)"
    PROVIDER_CLAUDE="Claude/Anthropic (requires API key)"
    PROVIDER_DEEPSEEK="DeepSeek (requires API key)"
    PROVIDER_OPENROUTER="OpenRouter (requires API key)"
    PROVIDER_GROK="Grok/xAI (requires API key)"
    PROVIDER_OLLAMA="Ollama (local, no API key required)"
    
    PROVIDER_PROMPT="Enter the provider number [1-7] (default: 2):"
    
    STEP_API_KEY="STEP 2: Configure API Key"
    API_KEY_DESC="You selected a provider that requires an API key."
    API_KEY_INSTRUCTIONS="Enter your API key for"
    API_KEY_SKIP="Press ENTER to skip (you can configure it later)"
    
    OLLAMA_SELECTED="You selected Ollama as your provider. No API key is required!"
    
    STEP_LANGUAGE="STEP 3: Choose Preferred Language"
    LANGUAGE_DESC="Choose the language for generated commit messages:"
    
    LANG_OPTION_PT="Portuguese (pt-br)"
    LANG_OPTION_EN="English (en)"
    LANG_OPTION_ES="Spanish (es)"
    LANG_OPTION_FR="French (fr)"
    LANG_OPTION_DE="German (de)"
    
    LANGUAGE_PROMPT="Enter the language number [1-5] (default: 1):"
    
    STEP_REPO="STEP 4: Set Default Repository Path"
    REPO_DESC="You can set a default path for the Git repository:"
    REPO_CURRENT="Your current directory is:"
    REPO_PROMPT="Enter the default repository path or leave blank to use current directory:"
    
    STEP_AUTO_COMMIT="STEP 5: Configure Auto Commit"
    AUTO_COMMIT_DESC="Commit-AI can automatically commit after generating messages."
    AUTO_COMMIT_PROMPT="Enable auto commit? [y/N]:"
    
    SAVING_CONFIG="Saving configuration..."
    CONFIG_SAVED="Configuration saved successfully!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║              CONFIGURATION COMPLETED!                      ║
╚═════════════════════════════════════════════════════════════════╝"
    
    NEXT_STEPS="Next steps:"
    BASIC_USAGE="Basic usage:"
    USAGE_BASIC="  commit-ai               # Generate commit message in current directory"
    USAGE_DRY_RUN="  commit-ai --dry-run     # Just show the message without committing"
    USAGE_WATCH="  commit-ai --watch       # Continuous monitoring mode"
    
    MORE_INFO="For more information, visit: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="Thank you for setting up Commit-AI! 💻✨"
    
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
    ║        Asistente de Configuración             ║
    ╚═══════════════════════════════════════════════╝
    "
    
    WELCOME="¡Bienvenido al asistente de configuración de Commit-AI!"
    WIZARD_DESC="Este asistente lo ayudará a configurar Commit-AI para su entorno."
    
    STEP_PROVIDER="PASO 1: Elija el Proveedor de IA"
    PROVIDER_DESC="Commit-AI soporta varios proveedores de IA para generar mensajes de commit."
    PROVIDER_OPTIONS="Por favor, elija un proveedor de IA:"
    
    PROVIDER_OPENAI="OpenAI (requiere clave API)"
    PROVIDER_GEMINI="Google Gemini (requiere clave API, predeterminado)"
    PROVIDER_CLAUDE="Claude/Anthropic (requiere clave API)"
    PROVIDER_DEEPSEEK="DeepSeek (requiere clave API)"
    PROVIDER_OPENROUTER="OpenRouter (requiere clave API)"
    PROVIDER_GROK="Grok/xAI (requiere clave API)"
    PROVIDER_OLLAMA="Ollama (local, no requiere clave API)"
    
    PROVIDER_PROMPT="Introduce el número del proveedor [1-7] (predeterminado: 2):"
    
    STEP_API_KEY="PASO 2: Configurar Clave API"
    API_KEY_DESC="Ha seleccionado un proveedor que requiere una clave API."
    API_KEY_INSTRUCTIONS="Introduzca su clave API para"
    API_KEY_SKIP="Presione ENTER para omitir (puede configurarlo más tarde)"
    
    OLLAMA_SELECTED="¡Ha seleccionado Ollama como su proveedor. No se requiere clave API!"
    
    STEP_LANGUAGE="PASO 3: Elija el Idioma Preferido"
    LANGUAGE_DESC="Elija el idioma para los mensajes de commit generados:"
    
    LANG_OPTION_PT="Portugués (pt-br)"
    LANG_OPTION_EN="Inglés (en)"
    LANG_OPTION_ES="Español (es)"
    LANG_OPTION_FR="Francés (fr)"
    LANG_OPTION_DE="Alemán (de)"
    
    LANGUAGE_PROMPT="Introduzca el número del idioma [1-5] (predeterminado: 1):"
    
    STEP_REPO="PASO 4: Establecer Ruta Predeterminada del Repositorio"
    REPO_DESC="Puede establecer una ruta predeterminada para el repositorio Git:"
    REPO_CURRENT="Su directorio actual es:"
    REPO_PROMPT="Introduzca la ruta del repositorio predeterminada o deje en blanco para usar el directorio actual:"
    
    STEP_AUTO_COMMIT="PASO 5: Configurar Commit Automático"
    AUTO_COMMIT_DESC="Commit-AI puede hacer commits automáticamente después de generar mensajes."
    AUTO_COMMIT_PROMPT="¿Habilitar commit automático? [s/N]:"
    
    SAVING_CONFIG="Guardando configuración..."
    CONFIG_SAVED="¡Configuración guardada con éxito!"
    
    FINAL_HEADER="╔═════════════════════════════════════════════════════════════════╗
║            ¡CONFIGURACIÓN COMPLETADA!                     ║
╚═════════════════════════════════════════════════════════════════╝"
    
    NEXT_STEPS="Próximos pasos:"
    BASIC_USAGE="Uso básico:"
    USAGE_BASIC="  commit-ai               # Generar mensaje de commit en el directorio actual"
    USAGE_DRY_RUN="  commit-ai --dry-run     # Solo mostrar el mensaje sin hacer commit"
    USAGE_WATCH="  commit-ai --watch       # Modo de monitoreo continuo"
    
    MORE_INFO="Para más información, visite: https://github.com/eusoujuninho/commit-ai"
    
    THANK_YOU="¡Gracias por configurar Commit-AI! 💻✨"
    
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
    echo_color $BLUE "$LANGUAGE_SELECTION"
    echo
    echo_color $CYAN "1) $LANG_PT"
    echo_color $CYAN "2) $LANG_EN"
    echo_color $CYAN "3) $LANG_ES"
    echo
    echo_color $YELLOW "$LANG_PROMPT"
    read -r lang_choice
    
    case $lang_choice in
        1) LANGUAGE="pt-br" ;;
        2) LANGUAGE="en" ;;
        3) LANGUAGE="es" ;;
        *) LANGUAGE="pt-br" ;;
    esac
}

# Carregar strings iniciais para a seleção de idioma
load_pt_strings

# Função para exibir o banner
show_banner() {
    clear
    echo_color $BLUE "$BANNER"
    echo_color $BOLD "$WELCOME"
    echo_color $NC "$WIZARD_DESC"
    echo
}

# Mostrar banner inicial
show_banner

# Solicitar seleção de idioma
select_language

# Recarregar strings com base no idioma selecionado
case $LANGUAGE in
    "pt-br") load_pt_strings ;;
    "en") load_en_strings ;;
    "es") load_es_strings ;;
    *) load_pt_strings ;;
esac

# Exibir banner atualizado com o idioma selecionado
show_banner

# Passo 1: Escolher o provedor de IA
echo_color $BOLD "$STEP_PROVIDER"
echo_color $NC "$PROVIDER_DESC"
echo
echo_color $BLUE "$PROVIDER_OPTIONS"
echo
echo_color $CYAN "1) $PROVIDER_OPENAI"
echo_color $CYAN "2) $PROVIDER_GEMINI"
echo_color $CYAN "3) $PROVIDER_CLAUDE"
echo_color $CYAN "4) $PROVIDER_DEEPSEEK"
echo_color $CYAN "5) $PROVIDER_OPENROUTER"
echo_color $CYAN "6) $PROVIDER_GROK"
echo_color $CYAN "7) $PROVIDER_OLLAMA"
echo
echo_color $YELLOW "$PROVIDER_PROMPT"
read -r provider_choice

case $provider_choice in
    1) provider="openai" ;;
    3) provider="claude" ;;
    4) provider="deepseek" ;;
    5) provider="openrouter" ;;
    6) provider="grok" ;;
    7) provider="ollama" ;;
    *) provider="gemini" ;;
esac

# Passo 2: Configurar a chave API
echo
echo_color $BOLD "$STEP_API_KEY"

if [ "$provider" == "ollama" ]; then
    echo_color $GREEN "$OLLAMA_SELECTED"
    api_key=""
else
    echo_color $NC "$API_KEY_DESC"
    echo_color $YELLOW "$API_KEY_INSTRUCTIONS $provider ($API_KEY_SKIP):"
    read -r api_key
fi

# Passo 3: Escolher o idioma preferido
echo
echo_color $BOLD "$STEP_LANGUAGE"
echo_color $NC "$LANGUAGE_DESC"
echo
echo_color $CYAN "1) $LANG_OPTION_PT"
echo_color $CYAN "2) $LANG_OPTION_EN"
echo_color $CYAN "3) $LANG_OPTION_ES"
echo_color $CYAN "4) $LANG_OPTION_FR"
echo_color $CYAN "5) $LANG_OPTION_DE"
echo
echo_color $YELLOW "$LANGUAGE_PROMPT"
read -r lang_choice

case $lang_choice in
    2) lang="en" ;;
    3) lang="es" ;;
    4) lang="fr" ;;
    5) lang="de" ;;
    *) lang="pt-br" ;;
esac

# Passo 4: Definir o caminho padrão do repositório
echo
echo_color $BOLD "$STEP_REPO"
echo_color $NC "$REPO_DESC"
echo
current_dir=$(pwd)
echo_color $CYAN "$REPO_CURRENT $current_dir"
echo
echo_color $YELLOW "$REPO_PROMPT"
read -r repo_path

if [ -z "$repo_path" ]; then
    repo_path="$current_dir"
fi

# Passo 5: Configurar commit automático
echo
echo_color $BOLD "$STEP_AUTO_COMMIT"
echo_color $NC "$AUTO_COMMIT_DESC"
echo
echo_color $YELLOW "$AUTO_COMMIT_PROMPT"
read -r auto_commit

# Em português e espanhol, S/s é sim
if [[ "$LANGUAGE" == "pt-br" && "$auto_commit" =~ ^[Ss]$ ]] || 
   [[ "$LANGUAGE" == "es" && "$auto_commit" =~ ^[Ss]$ ]] || 
   [[ "$LANGUAGE" == "en" && "$auto_commit" =~ ^[Yy]$ ]]; then
    auto_commit="true"
else
    auto_commit="false"
fi

# Salvar configuração
echo
echo_color $BLUE "$SAVING_CONFIG"

# Use commit-ai para salvar as configurações
commit_ai_cmd="commit-ai --configure"

# Preparar configuração
config_data="{
  \"aiProvider\": \"$provider\",
  \"language\": \"$lang\",
  \"repoPath\": \"$repo_path\",
  \"autoCommit\": $auto_commit"

# Adicionar a chave API correspondente com base no provedor
if [ "$provider" == "openai" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"openaiKey\": \"$api_key\""
elif [ "$provider" == "gemini" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"geminiKey\": \"$api_key\""
elif [ "$provider" == "claude" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"claudeKey\": \"$api_key\""
elif [ "$provider" == "deepseek" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"deepseekKey\": \"$api_key\""
elif [ "$provider" == "openrouter" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"openrouterKey\": \"$api_key\""
elif [ "$provider" == "grok" ] && [ -n "$api_key" ]; then
    config_data="$config_data,
  \"grokKey\": \"$api_key\""
fi

# Fechar o JSON
config_data="$config_data
}"

# Caminho para o diretório de configuração
config_dir="$HOME/.config/commit-ai"
mkdir -p "$config_dir"
config_file="$config_dir/config.json"

# Salvar o JSON para o arquivo de configuração
echo "$config_data" > "$config_file"

echo_color $GREEN "$CONFIG_SAVED"
echo

# Exibir mensagem final
echo_color $BLUE "$FINAL_HEADER"
echo
echo_color $BOLD "$NEXT_STEPS"
echo
echo_color $GREEN "$BASIC_USAGE"
echo_color $NC "$USAGE_BASIC"
echo_color $NC "$USAGE_DRY_RUN"
echo_color $NC "$USAGE_WATCH"
echo
echo_color $YELLOW "$MORE_INFO"
echo
echo_color $BLUE "$THANK_YOU" 