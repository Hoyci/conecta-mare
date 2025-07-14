# 🚀 ConectaMaré

[![Status da Build](https://img.shields.io/github/actions/workflow/status/seu-usuario/seu-repo/go-test.yml?branch=main&label=Go%20Tests&style=for-the-badge)](https://github.com/seu-usuario/seu-repo/actions/workflows/go-test.yml)
[![Licença](https://img.shields.io/badge/licen%C3%A7a-MIT-blue.svg?style=for-the-badge)](https://github.com/seu-usuario/seu-repo/blob/main/server/LICENSE)
![Go](https://img.shields.io/badge/Go-1.22-00ADD8.svg?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB.svg?style=for-the-badge&logo=react)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-336791.svg?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Pronto-2496ED.svg?style=for-the-badge&logo=docker)

## 🎯 O Projeto

O ConectaMaré nasceu da necessidade de criar um ecossistema digital que fortaleça a comunidade do Complexo da Maré, no Rio de Janeiro. A plataforma oferece uma solução de ponta a ponta:

* **Para Clientes:** Uma maneira fácil e segura de encontrar, avaliar e contratar profissionais locais verificados para as mais diversas necessidades.
* **Para Profissionais:** Uma ferramenta poderosa para divulgar seus serviços, construir uma reputação online, gerenciar seu portfólio e alcançar uma base de clientes maior, tudo dentro da sua própria comunidade.

Este projeto vai além do código; é uma iniciativa para fomentar o empreendedorismo e a confiança na economia local.

## ✨ Principais Funcionalidades

### Para Clientes 🙋‍♀️
* 🔍 **Busca Inteligente:** Encontre profissionais por categoria, especialidade e localização.
* ⭐ **Avaliações Reais:** Tome decisões informadas com base em feedbacks de outros clientes.
* 🛡️ **Profissionais Verificados:** Mais segurança e confiança ao contratar.
* 📅 **Agendamento Facilitado:** Entre em contato e combine os serviços diretamente.

### Para Profissionais 👷‍♂️
* 📈 **Dashboard de Performance:** Acompanhe métricas essenciais como visualizações de perfil, taxas de conversão e desempenho de serviços.
* 📊 **Benchmarking:** Compare seu desempenho com a média de outros profissionais da sua área e identifique oportunidades de melhoria.
* 📝 **Onboarding Completo:** Um fluxo guiado para criar um perfil atraente, adicionando certificações, portfólio de projetos e descrição de serviços.
* 🖼️ **Portfólio Visual:** Mostre seu trabalho com galerias de imagens para cada projeto e serviço.
* 🔔 **Sistema de Notificações:** Receba alertas sobre novas avaliações, marcos de desempenho e dicas para otimizar seu perfil.

## 🛠️ Tecnologias Utilizadas

Este projeto foi construído utilizando um conjunto de tecnologias modernas, escaláveis e robustas, organizadas em uma arquitetura de monorepo.

| Categoria | Tecnologia |
| :--- | :--- |
| **Backend** | 🟢 **Go (Golang)**, **Chi** (Router), **SQLx** (DB Access), **PostgreSQL**, **JWT** (Autenticação) |
| **Frontend** | 🔵 **React**, **TypeScript**, **Vite** (SSR), **Tailwind CSS**, **Shadcn/ui**, **TanStack Query**, **Zustand** |
| **Infra & DevOps** | 🐳 **Docker**, **Docker Compose**, **Nginx** (Reverse Proxy), **LocalStack** (S3), **GitHub Actions** (CI/CD) |

## 🏗️ Arquitetura do Projeto

O ConectaMaré é estruturado como um monorepo, contendo três serviços principais que são orquestrados pelo Docker Compose:

1.  **`server` (Go):** Uma API robusta que gerencia toda a lógica de negócio, autenticação, dados de usuários, serviços e interações com o banco de dados.
2.  **`client` (React/Vite):** Uma aplicação frontend moderna com Server-Side Rendering (SSR) para melhor performance e SEO, oferecendo uma experiência de usuário rica e interativa.
3.  **`nginx` (Proxy Reverso):** Atua como a porta de entrada da aplicação, direcionando o tráfego para a API (`/api/v1`) ou para o frontend React, além de servir os assets estáticos de forma eficiente.

Essa arquitetura garante uma clara separação de responsabilidades, escalabilidade e um ambiente de desenvolvimento coeso.

## 🚀 Começando

É incrivelmente fácil colocar o projeto para rodar localmente, graças ao script de setup e ao Docker.

### Pré-requisitos

* [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/)
* [AWS CLI](https://aws.amazon.com/cli/) (utilizado para interagir com o LocalStack)
* `make` (opcional, para usar os atalhos do Makefile)

### Passos para Instalação

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/Hoyci/conecta-mare.git](https://github.com/Hoyci/conecta-mare.git)
    cd seu-repo
    ```

2.  **Execute o script de setup do ambiente:**
    Este script irá verificar as dependências, iniciar os contêineres e criar o bucket no LocalStack automaticamente.
    ```bash
    # Dê permissão de execução ao script
    chmod +x server/setup-dev-env.sh

    # Rode o script
    ./server/setup-dev-env.sh
    ```

3.  **Suba a aplicação completa com Docker Compose:**
    Se você pulou o passo 2, ou se os contêineres não estiverem rodando, use este comando na raiz do projeto:
    ```bash
    docker-compose up --build -d
    ```
    *O Docker Compose irá construir as imagens do backend, frontend e nginx, iniciar o banco de dados, o LocalStack e rodar as migrações automaticamente*.

4.  **Acesse a aplicação:**
    Abra seu navegador e acesse: **[http://localhost](http://localhost)**

Pronto! O ambiente completo do ConectaMaré está rodando na sua máquina.

## 🧰 Comandos Úteis (Makefile)

O `Makefile` na pasta `server` contém vários atalhos úteis para o desenvolvimento do backend:

| Comando | Descrição |
| :--- | :--- |
| `make build` | Compila o executável do servidor Go. |
| `make run` | Executa o servidor Go localmente (sem Docker). |
| `make test` | Roda todos os testes unitários do backend. |
| `make watch` | Inicia o servidor com live-reload usando `air`. |
| `make migrate-up` | Aplica todas as migrações pendentes no banco de dados. |
| `make migrate-down` | Reverte a última migração aplicada. |
| `make migrate name=nome_da_migration` | Cria novos arquivos de migração. |

## 📜 Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](server/LICENSE) para mais detalhes.

---
_Feito com ❤️ para fortalecer a comunidade e a economia local._
