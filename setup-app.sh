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
DOCKER_COMPOSE_PROJECT_NAME="conecta-mare-server"

# --- Verificação inicial ---
echo -e "${YELLOW}Verificando a instalação do Docker Compose...${NC}"
if ! docker compose version &>/dev/null; then
  echo -e "${RED}ERRO: 'docker compose' não foi encontrado no seu ambiente.${NC}"
  echo -e "${YELLOW}Certifique-se de que o Docker está instalado e configurado corretamente.${NC}"
  exit 1
fi
echo -e "${GREEN}Docker Compose detectado.${NC}"
echo "--------------------------------------------------"

#==============================================================================
# FUNÇÃO 1: Checar AWS CLI
#==============================================================================
check_install_aws_cli() {
  echo -e "${YELLOW}Verificando a instalação da AWS CLI...${NC}"
  if command -v aws &>/dev/null; then
    echo -e "${GREEN}AWS CLI já está instalada.${NC}"
  else
    echo "AWS CLI não encontrada. Tentando instalar..."
    if ! command -v unzip &>/dev/null; then
      echo "'unzip' não encontrado. Tentando instalar..."
      if command -v apt-get &>/dev/null; then
        sudo apt-get update && sudo apt-get install -y unzip
      else
        echo -e "${RED}Não foi possível instalar 'unzip'. Por favor, instale-o manualmente.${NC}"
        exit 1
      fi
    fi
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    sudo ./aws/install
    rm -rf aws awscliv2.zip
    echo -e "${GREEN}AWS CLI instalada com sucesso.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 2: Subir a INFRAESTRUTURA do Docker Compose
#==============================================================================
start_infra_services() {
  echo -e "${YELLOW}Subindo serviços de infraestrutura (Postgres, Localstack, etc)...${NC}"
  docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" up --build -d

  echo "Aguardando o PostgreSQL ficar disponível..."
  local retries_pg=20
  local count_pg=0
  until docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec postgres pg_isready -U "${DB_USERNAME}" -q; do
    count_pg=$((count_pg + 1))
    if [ $count_pg -gt $retries_pg ]; then
      echo -e "${RED}O PostgreSQL não ficou disponível a tempo.${NC}"
      exit 1
    fi
    sleep 2
  done
  echo -e "${GREEN}PostgreSQL está pronto!${NC}"

  echo "Aguardando o LocalStack (S3) ficar disponível..."
  local retries_s3=20
  local count_s3=0
  until docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T localstack aws --endpoint-url="$S3_ENDPOINT" s3 ls &>/dev/null; do
    count_s3=$((count_s3 + 1))
    if [ $count_s3 -gt $retries_s3 ]; then
      echo -e "${RED}O LocalStack não ficou disponível a tempo.${NC}"
      exit 1
    fi
    sleep 3
  done
  echo -e "${GREEN}LocalStack está pronto!${NC}"
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 3: Criar o banco de dados para o RudderStack
#==============================================================================
create_rudder_database() {
  echo -e "${YELLOW}Verificando banco de dados do RudderStack ('$RUDDER_DB_NAME')...${NC}"
  local db_exists=$(docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T postgres psql -U "$DB_USERNAME" -tAc "SELECT 1 FROM pg_database WHERE datname='$RUDDER_DB_NAME'")

  if [ "$db_exists" = "1" ]; then
    echo -e "${GREEN}O banco de dados '$RUDDER_DB_NAME' já existe.${NC}"
  else
    echo "Criando banco de dados '$RUDDER_DB_NAME'..."
    docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T postgres psql -U "$DB_USERNAME" -c "CREATE DATABASE \"$RUDDER_DB_NAME\";"
    echo -e "${GREEN}Banco de dados '$RUDDER_DB_NAME' criado.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 4: Criar o bucket S3 no LocalStack se não existir
#==============================================================================
create_s3_bucket() {
  echo -e "${YELLOW}Verificando bucket S3 '$BUCKET_NAME' no LocalStack...${NC}"

  if docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T localstack aws --endpoint-url="$S3_ENDPOINT" s3api head-bucket --bucket "$BUCKET_NAME" >/dev/null 2>&1; then
    echo -e "${GREEN}O bucket '$BUCKET_NAME' já existe.${NC}"
  else
    echo "Criando bucket '$BUCKET_NAME'..."
    docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T localstack aws --endpoint-url="$S3_ENDPOINT" s3 mb "s3://$BUCKET_NAME"

    echo "Tornando o bucket '$BUCKET_NAME' público para leitura..."
    docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" exec -T localstack aws --endpoint-url="$S3_ENDPOINT" s3api put-bucket-acl --bucket "$BUCKET_NAME" --acl public-read

    echo -e "${GREEN}Bucket '$BUCKET_NAME' criado com sucesso.${NC}"
  fi
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO 5: Subir os serviços de APLICAÇÃO do Docker Compose
#==============================================================================
start_app_services() {
  echo -e "${YELLOW}Subindo serviços de aplicação (Rudder, Server, Nginx, etc)...${NC}"
  docker compose -p "$DOCKER_COMPOSE_PROJECT_NAME" --profile apps up --build -d
  echo -e "${GREEN}Serviços de aplicação iniciados.${NC}"
  echo "--------------------------------------------------"
}

#==============================================================================
# FUNÇÃO PRINCIPAL (MAIN)
#==============================================================================
main() {
  echo -e "${GREEN}--- Iniciando Script de Configuração do Ambiente ---${NC}"

  echo -e "${YELLOW}Carregando variáveis do arquivo .env...${NC}"
  if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
    echo -e "${GREEN}Variáveis carregadas.${NC}"
  else
    echo -e "${RED}ERRO: Arquivo .env não encontrado!${NC}"
    exit 1
  fi
  echo "--------------------------------------------------"

  check_install_aws_cli

  start_infra_services
  create_rudder_database
  create_s3_bucket
  start_app_services

  echo -e "${GREEN}--- Configuração do Ambiente Finalizada com Sucesso! ---${NC}"
  echo -e "${YELLOW}Sua aplicação deve estar disponível em http://localhost ${NC}"
}

# Executa a função principal
main
