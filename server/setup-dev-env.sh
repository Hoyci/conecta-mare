#!/bin/bash

# Encerra o script imediatamente se um comando falhar.
set -e

# --- Variáveis de Cor para Melhor Visualização ---
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # Sem Cor

# --- Variáveis de Configuração ---
BUCKET_NAME="conecta-mare"
S3_ENDPOINT="http://localhost:4566"
DOCKER_COMPOSE_PROJECT_NAME="conecta-mare-server" # Nome do projeto para evitar conflitos

#==============================================================================
# FUNÇÃO 1: Checar e Instalar o Docker
#==============================================================================
check_install_docker() {
  echo -e "${YELLOW}Verificando a instalação do Docker...${NC}"
  if command -v docker &>/dev/null; then
    echo -e "${GREEN}Docker já está instalado.${NC}"
  else
    echo "Docker não encontrado. Tentando instalar a versão mais recente..."
    # Utiliza o script oficial do Docker para uma instalação genérica
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    rm get-docker.sh
    echo -e "${GREEN}Docker instalado com sucesso.${NC}"

    # Adiciona o usuário atual ao grupo do Docker para evitar o uso de 'sudo'
    sudo usermod -aG docker "$USER"
    echo -e "${YELLOW}AVISO: Você precisa fazer logout e login novamente para usar o Docker sem 'sudo'.${NC}"
    echo -e "${YELLOW}Alternativamente, você pode executar o seguinte comando para aplicar as alterações de grupo imediatamente:${NC}"
    echo -e "newgrp docker"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 2: Checar e Instalar a AWS CLI
#==============================================================================
check_install_aws_cli() {
  echo -e "${YELLOW}Verificando a instalação da AWS CLI...${NC}"
  if command -v aws &>/dev/null; then
    echo -e "${GREEN}AWS CLI já está instalada.${NC}"
  else
    echo "AWS CLI não encontrada. Tentando instalar a versão mais recente..."

    # Checa e instala o 'unzip', necessário para a instalação da AWS CLI
    if ! command -v unzip &>/dev/null; then
      echo "'unzip' não encontrado. Tentando instalar..."
      if command -v apt-get &>/dev/null; then
        sudo apt-get update && sudo apt-get install -y unzip
      elif command -v yum &>/dev/null; then
        sudo yum install -y unzip
      elif command -v dnf &>/dev/null; then
        sudo dnf install -y unzip
      else
        echo -e "${RED}Não foi possível instalar o 'unzip' automaticamente. Por favor, instale-o e execute o script novamente.${NC}"
        exit 1
      fi
    fi

    # Utiliza o método de instalação oficial e genérico para Linux
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    sudo ./aws/install

    # Limpa os arquivos de instalação
    rm -rf aws awscliv2.zip
    echo -e "${GREEN}AWS CLI instalada com sucesso.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 3: Subir os contêineres do Docker Compose
#==============================================================================
start_docker_compose() {
  echo -e "${YELLOW}Verificando o status dos contêineres Docker...${NC}"

  # Tenta usar 'docker compose' (v2) e fallback para 'docker-compose' (v1)
  if docker compose version &>/dev/null; then
    DOCKER_COMPOSE_CMD="docker compose"
  else
    DOCKER_COMPOSE_CMD="docker-compose"
  fi

  # Verifica se os contêineres do projeto já estão rodando
  if [ ! -z "$($DOCKER_COMPOSE_CMD -p $DOCKER_COMPOSE_PROJECT_NAME ps -q)" ]; then
    echo -e "${GREEN}Os contêineres do projeto '$DOCKER_COMPOSE_PROJECT_NAME' já estão em execução.${NC}"
  else
    echo "Contêineres não estão rodando. Iniciando com '$DOCKER_COMPOSE_CMD up'..."
    $DOCKER_COMPOSE_CMD -p "$DOCKER_COMPOSE_PROJECT_NAME" up --build -d
    echo -e "${GREEN}Contêineres iniciados em modo detached.${NC}"
  fi

  # Aguarda o PostgreSQL ficar saudável
  echo "Aguardando o PostgreSQL ficar disponível..."
  local retries_pg=20
  local count_pg=0
  until $DOCKER_COMPOSE_CMD -p "$DOCKER_COMPOSE_PROJECT_NAME" exec postgres pg_isready -U "${DB_USERNAME}" -q; do
    count_pg=$((count_pg + 1))
    if [ $count_pg -gt $retries_pg ]; then
      echo -e "${RED}O PostgreSQL não ficou disponível a tempo. Verifique os logs do contêiner.${NC}"
      exit 1
    fi
    sleep 2
  done
  echo -e "${GREEN}PostgreSQL está pronto!${NC}"

  # Aguarda o LocalStack ficar pronto
  echo "Aguardando o LocalStack (S3) ficar disponível..."
  local retries_s3=20
  local count_s3=0
  until AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test AWS_DEFAULT_REGION=us-east-1 aws --endpoint-url="$S3_ENDPOINT" s3 ls &>/dev/null; do
    count_s3=$((count_s3 + 1))
    if [ $count_s3 -gt $retries_s3 ]; then
      echo -e "${RED}O LocalStack não ficou disponível a tempo. Verifique os logs do contêiner.${NC}"
      exit 1
    fi
    sleep 2
  done
  echo -e "${GREEN}LocalStack está pronto!${NC}"
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 4: Criar o banco de dados para o RudderStack
#==============================================================================
create_rudder_database() {
  echo -e "${YELLOW}Verificando a existência do banco de dados para o RudderStack ('$RUDDER_DB_NAME')...${NC}"

  local db_exists=$($DOCKER_COMPOSE_CMD -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T postgres psql -U "$DB_USERNAME" -tAc "SELECT 1 FROM pg_database WHERE datname='$RUDDER_DB_NAME'")

  if [ "$db_exists" = "1" ]; then
    echo -e "${GREEN}O banco de dados '$RUDDER_DB_NAME' para o RudderStack já existe.${NC}"
  else
    echo "Banco de dados '$RUDDER_DB_NAME' não encontrado. Criando..."
    $DOCKER_COMPOSE_CMD -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T postgres psql -U "$DB_USERNAME" -c "CREATE DATABASE \"$RUDDER_DB_NAME\";"
    echo -e "${GREEN}Banco de dados '$RUDDER_DB_NAME' criado com sucesso.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 5: Criar o bucket S3 no LocalStack se não existir
#==============================================================================
create_s3_bucket() {
  echo -e "${YELLOW}Verificando a existência do bucket S3 '$BUCKET_NAME' no LocalStack...${NC}"

  local AWS_ENV_VARS="AWS_ACCESS_KEY_ID=test AWS_SECRET_ACCESS_KEY=test AWS_DEFAULT_REGION=us-east-1"

  if env $AWS_ENV_VARS aws --endpoint-url="$S3_ENDPOINT" s3api head-bucket --bucket "$BUCKET_NAME" &>/dev/null; then
    echo -e "${GREEN}O bucket '$BUCKET_NAME' já existe.${NC}"
  else
    echo "Bucket não encontrado. Criando bucket '$BUCKET_NAME'..."
    env $AWS_ENV_VARS aws --endpoint-url="$S3_ENDPOINT" s3 mb "s3://$BUCKET_NAME"
    echo -e "${GREEN}Bucket '$BUCKET_NAME' criado com sucesso.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO PRINCIPAL (MAIN)
#==============================================================================
main() {
  echo -e "${GREEN}--- Iniciando Script de Configuração do Ambiente de Desenvolvimento ---${NC}"

  echo -e "${YELLOW}Carregando variáveis do arquivo .env...${NC}"
  if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
    echo -e "${GREEN}Variáveis carregadas.${NC}"
  else
    echo -e "${RED}ERRO: Arquivo .env não encontrado! Crie um arquivo .env com as variáveis necessárias (DB_USERNAME, RUDDER_DB_NAME, etc).${NC}"
    exit 1
  fi
  echo "--------------------------------------------------"

  check_install_docker
  check_install_aws_cli
  start_docker_compose
  create_rudder_database
  create_s3_bucket
  echo -e "${GREEN}--- Configuração do Ambiente Finalizada com Sucesso! ---${NC}"
}

# Executa a função principal
main
