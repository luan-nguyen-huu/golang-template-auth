#!/bin/bash

# ─────────────────────────────────────────────
#  dev.sh — Hot Reload runner using Air
#  Usage: ./dev.sh
# ─────────────────────────────────────────────

set -e

GREEN="\033[0;32m"
YELLOW="\033[1;33m"
CYAN="\033[0;36m"
RED="\033[0;31m"
RESET="\033[0m"

echo -e "${CYAN}"
echo "╔══════════════════════════════════════╗"
echo "║     🔥 Go Hot Reload — Air           ║"
echo "╚══════════════════════════════════════╝"
echo -e "${RESET}"

if ! command -v air &> /dev/null; then
  echo -e "${YELLOW}⚠️  'air' not found. Installing...${RESET}"
  go install github.com/air-verse/air@latest

  export PATH="$PATH:$(go env GOPATH)/bin"

  if ! command -v air &> /dev/null; then
    echo -e "${RED}❌ Failed to install air. Make sure GOPATH/bin is in your PATH.${RESET}"
    echo -e "   Run: ${YELLOW}export PATH=\$PATH:\$(go env GOPATH)/bin${RESET}"
    exit 1
  fi

  echo -e "${GREEN}✅ air installed successfully!${RESET}"
else
  echo -e "${GREEN}✅ air is already installed: $(air -v 2>&1 | head -1)${RESET}"
fi

mkdir -p tmp

if [ ! -f ".air.toml" ]; then
  echo -e "${YELLOW}⚠️  .air.toml not found. Using default config.${RESET}"
fi

echo -e "\n${GREEN}🚀 Starting hot reload... Edit any .go file to trigger rebuild.${RESET}\n"
air