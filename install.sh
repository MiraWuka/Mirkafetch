#!/bin/bash
set -e

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
esac

REPO="MiraWuka/Mirkafetch"
BINARY_NAME="Mirkafetch"
INSTALL_DIR="/usr/local/bin"

echo "🚀 Installing $BINARY_NAME for $OS-$ARCH..."

URL="https://github.com/$REPO/releases/latest/download/${BINARY_NAME}-${OS}-${ARCH}"

curl -L "$URL" -o "$BINARY_NAME"
chmod +x "$BINARY_NAME"

# Пытаемся установить в PATH
if [ -w "$INSTALL_DIR" ]; then
    sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
    echo "[+] Installed to $INSTALL_DIR/$BINARY_NAME"
else
    echo "[-] Please move $BINARY_NAME to your PATH manually"
fi

echo "🎉 Installation complete! Run: $BINARY_NAME --help"