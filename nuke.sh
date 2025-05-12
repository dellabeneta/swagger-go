#!/bin/bash
# nuke.sh - Remove absolutamente tudo do Docker, incluindo containers, imagens, volumes, networks e cache.

set -e

echo "Parando todos os containers..."
docker stop $(docker ps -aq) 2>/dev/null || true

echo "Removendo todos os containers..."
docker rm -f $(docker ps -aq) 2>/dev/null || true

echo "Removendo todas as imagens..."
docker rmi -f $(docker images -aq) 2>/dev/null || true

echo "Removendo todos os volumes..."
docker volume rm -f $(docker volume ls -q) 2>/dev/null || true

echo "Removendo todas as networks customizadas..."
docker network rm $(docker network ls -q | grep -v '^bridge$\|^host$\|^none$') 2>/dev/null || true

echo "Limpando cache do builder..."
docker builder prune -af --filter "until=0h" 2>/dev/null || true

echo "Limpando system prune..."
docker system prune -af --volumes 2>/dev/null || true

echo "Docker completamente limpo!"