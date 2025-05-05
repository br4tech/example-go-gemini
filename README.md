# Gemini Go CLI

[![Linguagem](https://img.shields.io/badge/Go-informational?style=flat&logo=go)](https://go.dev/)
[![Licença](https://img.shields.io/badge/License-MIT-green)](LICENSE)

Um simples utilitário de linha de comando em Go para interagir com a API Gemini do Google Cloud e obter respostas para perguntas textuais.

## Sumário

* [Funcionalidades](#funcionalidades)
* [Como Começar](#como-começar)
    * [Pré-requisitos](#pré-requisitos)
    * [Instalação](#instalação)
    * [Configuração](#configuração)
    * [Execução](#execução)
* [Uso](#uso)

## Funcionalidades

Este aplicativo possui uma funcionalidade principal:

* **Enviar uma pergunta textual para o modelo Gemini e exibir a resposta.** Atualmente, a pergunta padrão é "Qual é a capital do Brasil?".

## Como Começar

Siga estas etapas para configurar e executar o Gemini Go CLI em seu ambiente.

### Pré-requisitos

* **Go instalado:** Certifique-se de ter o Go (versão 1.x ou superior) instalado em sua máquina. Você pode baixá-lo em [https://go.dev/doc/install](https://go.dev/doc/install).
* **Chave de API do Google Cloud:** Você precisará de uma chave de API do Google Cloud com acesso à Gemini API. Você pode criar e gerenciar suas chaves de API no [Google AI Studio](https://ai.google.dev/).

### Instalação

1.  Clone o repositório (você provavelmente já tem o código localmente, então pode pular este passo se não estiver compartilhando o projeto via Git):
    ```bash
    git clone [URL_DO_SEU_REPOSITORIO]
    cd seu_projeto
    ```
2.  Construa o executável Go:
    ```bash
    go build -o gemini-cli .
    ```
    Isso criará um arquivo executável chamado `gemini-cli` (ou `gemini-cli.exe` no Windows).

### Configuração

O Gemini Go CLI requer que você configure sua chave de API do Google Cloud como uma variável de ambiente.

1.  **Defina a variável de ambiente `GOOGLE_API_KEY`:**

    * **No Linux e macOS:**
        ```bash
        export GOOGLE_API_KEY="SUA_CHAVE_DE_API"
        ```
        Para tornar essa variável persistente, adicione a linha acima ao seu arquivo de configuração do shell (por exemplo, `.bashrc` ou `.zshrc`).

    * **No Windows (Prompt de Comando):**
        ```bash
        set GOOGLE_API_KEY="SUA_CHAVE_DE_API"
        ```
        Para tornar essa variável persistente, use o comando `setx`:
        ```bash
        setx GOOGLE_API_KEY "SUA_CHAVE_DE_API"
        ```
        Pode ser necessário reiniciar o terminal.

    * **No Windows (PowerShell):**
        ```powershell
        $env:GOOGLE_API_KEY = "SUA_CHAVE_DE_API"
        ```
        Para tornar essa variável persistente:
        ```powershell
        [Environment]::SetEnvironmentVariable('GOOGLE_API_KEY', 'SUA_CHAVE_DE_API', 'User')
        ```

    Substitua `"SUA_CHAVE_DE_API"` pela sua chave de API real do Google Cloud.

### Execução

Após a instalação e configuração, você pode executar o Gemini Go CLI a partir do seu terminal:

```bash
./gemini-cli
```

O aplicativo enviará a pergunta "Qual é a capital do Brasil?" para o modelo gemini-2.0-flash e exibirá a resposta no seu terminal.

### Uso

Atualmente, o aplicativo está configurado para fazer uma pergunta específica. Para modificar a pergunta, você precisará editar o arquivo main.go e alterar o valor da variável prompt antes de construir o executável novamente.

Exemplo de saída esperada:

```bash
Resposta do Gemini:
Brasília.
```