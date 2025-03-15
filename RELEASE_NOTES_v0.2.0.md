# Commit-AI v0.2.0 - Notas de Lançamento

Temos o prazer de anunciar o lançamento da versão 0.2.0 do Commit-AI, que traz várias funcionalidades novas e melhorias significativas!

## 🌟 Novos Recursos

### Modo Watcher

A principal adição nesta versão é o **Modo Watcher** que permite monitoramento contínuo do seu repositório:

- **Monitoramento em tempo real** - Detecta alterações de arquivos automaticamente
- **Commits automáticos** - Realiza commits periodicamente quando detecta mudanças
- **Configuração flexível** - Defina intervalos personalizados e requisitos mínimos de alterações
- **Padrões de exclusão** - Ignore arquivos específicos ou diretórios com padrões glob
- **Modo silencioso** - Opções para reduzir o output no console

Acesse facilmente o modo watcher com:
```bash
commit-ai --watch
```

### Scripts de Instalação e Configuração

Para tornar a experiência do usuário mais amigável, adicionamos:

- **Script de instalação automática** (`install.sh`)
  - Detecta automaticamente seu sistema operacional
  - Instala o binário apropriado globalmente
  - Adiciona ao PATH do seu sistema
  - Guia pela configuração inicial

- **Assistente de configuração interativo** (`setup-wizard.sh`)
  - Interface colorida e amigável
  - Explicações detalhadas de todas as opções
  - Seleção simplificada de provedores de IA e idiomas
  - Configuração visual e guiada
  - Dicas de uso após a configuração

## 🔧 Melhorias

- Melhor detecção de repositórios Git
- Manipulação aprimorada de arquivos ignorados
- Documentação expandida com exemplos para todas as novas funcionalidades
- Resolução de caminhos de repositório aprimorada
- Feedback mais claro durante a execução
- Referências atualizadas para o repositório GitHub correto

## 📦 Binários Atualizados

Todos os binários pré-compilados foram atualizados para a versão 0.2.0:

- **Linux**: `bin/commit-ai-linux-amd64`
- **macOS**: `bin/commit-ai-darwin-amd64`
- **Windows**: `bin/commit-ai-windows-amd64.exe`

## 🚀 Como Atualizar

Se você já tem o Commit-AI instalado, você pode atualizar simplesmente baixando a nova versão e substituindo o binário existente, ou usar o script de instalação:

```bash
# Clone o repositório
git clone https://github.com/eusoujuninho/commit-ai.git

# Navegue para o diretório
cd commit-ai

# Execute o script de instalação
./install.sh
```

## 🤔 Feedback

Adoraríamos ouvir sua opinião sobre estas novas funcionalidades! Por favor, crie uma issue no [repositório GitHub](https://github.com/eusoujuninho/commit-ai/issues) se encontrar algum problema ou tiver sugestões.

Agradecemos seu apoio contínuo! 