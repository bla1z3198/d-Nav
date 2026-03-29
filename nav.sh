#!/bin/bash

# --- Настройки ---
APP_NAME="nav"           # Название будущего бинарника
SOURCE_FILE="main.go"      # Твой главный файл Go
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# 1. Приветствие
echo -e "${CYAN}=======================================${NC}"
echo -e "${CYAN}      Trajectories engine d-Nav        ${NC}"
echo -e "${CYAN}=======================================${NC}"

# 2. Проверка и сборка бинарника
if [ ! -f "$APP_NAME" ]; then
    echo -e "${YELLOW}Warning!:${NC} Binary file '$APP_NAME' not found..."
    
    if [ -f "$SOURCE_FILE" ]; then
        echo -e "${GREEN}==>${NC} Starting build from - > $SOURCE_FILE..."
        # Компилируем
        go build -o "$APP_NAME" "$SOURCE_FILE"
        
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}Success:${NC} Build completed!"
        else
            echo -e "${RED}Error:${NC} Unable to build =("
            exit 1
        fi
    else
        echo -e "${RED}Error:${NC} '$SOURCE_FILE' not found =("
        exit 1
    fi
fi

# 3. Запуск бинарника с передачей всех аргументов
# $@ — это специальная переменная, которая передает все аргументы скрипта в программу
echo -e "${GREEN}Starting...${NC}\n"

./"$APP_NAME" "$@"

# Проверка статуса выхода после работы программы
if [ $? -ne 0 ]; then
    echo -e "\n${RED}Crashed...${NC}"
fi
