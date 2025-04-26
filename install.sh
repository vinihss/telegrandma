#!/bin/bash

# Script para construir e instalar um projeto Go CLI localmente
# (para projetos já presentes no sistema de arquivos)

set -euo pipefail

# Cores para mensagens
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Verificar se o Go está instalado
if ! command -v go &> /dev/null; then
    echo -e "${RED}Erro: Go não está instalado.${NC}"
    echo -e "Instale o Go primeiro: https://golang.org/doc/install"
    exit 1
fi

# Verificar se estamos no diretório raiz do projeto Go
if [ ! -f "go.mod" ]; then
    echo -e "${RED}Erro: Não encontrado arquivo go.mod${NC}"
    echo -e "Execute este script do diretório raiz do projeto Go"
    exit 1
fi

# Obter nome do binário (usa nome do diretório atual por padrão)
DEFAULT_BINARY=$(basename "$(pwd)")
read -rp "Nome do binário [${DEFAULT_BINARY}]: " BINARY_NAME
BINARY_NAME=${BINARY_NAME:-$DEFAULT_BINARY}

# Diretórios de instalação (tenta system-wide primeiro)
INSTALL_DIRS=("/usr/local/bin" "${HOME}/.local/bin")

# Função para instalar/sobrescrever binário
install_binary() {
    local binary_path=$1
    local install_dir=$2
    
    echo -e "${YELLOW}Instalando em ${install_dir}...${NC}"
    
    # Criar diretório se não existir
    sudo mkdir -p "${install_dir}" 2>/dev/null || mkdir -p "${install_dir}" 2>/dev/null
    
    # Tentar instalar com sudo primeiro
    if sudo mv -f "$binary_path" "${install_dir}/${BINARY_NAME}" 2>/dev/null; then
        sudo chmod +x "${install_dir}/${BINARY_NAME}"
        echo -e "${GREEN}Binário instalado com sucesso em ${install_dir}/${BINARY_NAME}${NC}"
        return 0
    else
        # Tentar sem sudo
        if mv -f "$binary_path" "${install_dir}/${BINARY_NAME}" 2>/dev/null; then
            chmod +x "${install_dir}/${BINARY_NAME}"
            echo -e "${GREEN}Binário instalado com sucesso em ${install_dir}/${BINARY_NAME}${NC}"
            return 0
        fi
    fi
    
    return 1
}

# Construir o projeto
echo -e "${GREEN}Construindo o binário...${NC}"
if ! go build -o "${BINARY_NAME}" .; then
    echo -e "${RED}Falha ao construir o binário${NC}"
    exit 1
fi

# Tentar instalar nos diretórios possíveis
for dir in "${INSTALL_DIRS[@]}"; do
    if install_binary "./${BINARY_NAME}" "$dir"; then
        INSTALL_DIR="$dir"
        break
    fi
done

# Verificar se a instalação foi bem-sucedida
if [ -z "${INSTALL_DIR:-}" ]; then
    echo -e "${RED}Não foi possível instalar o binário automaticamente.${NC}"
    echo -e "O binário foi construído em: $(pwd)/${BINARY_NAME}"
    echo -e "Você pode movê-lo manualmente para um diretório no seu PATH"
    exit 1
fi

# Verificar se o diretório está no PATH
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo -e "${YELLOW}Aviso: ${INSTALL_DIR} não está no seu PATH. Adicione ao seu shell rc:${NC}"
    echo -e "export PATH=\"${INSTALL_DIR}:\$PATH\""
fi

echo -e "${GREEN}Instalação concluída!${NC}"
echo -e "Execute com: ${BINARY_NAME} --help"