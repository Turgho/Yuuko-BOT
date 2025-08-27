#!/bin/bash

PROCESS_NAME="yuuko-bot"

echo "Compilando o bot..."
go build -o yuuko-bot cmd/main.go

echo "Procurando processos do bot..."
pkill -f "$PROCESS_NAME"

echo "Processos do bot finalizados."

sleep 2

echo "Iniciando o bot..."
./yuuko-bot &

echo "Bot reiniciado."
