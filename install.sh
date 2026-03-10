#!/bin/bash

ZED_SNIPPETS="/mnt/c/Users/ramosmg/AppData/Roaming/Zed/snippets"
VSCODE_SNIPPETS="/mnt/c/Users/ramosmg/AppData/Roaming/Code/User/snippets"

mkdir -p "$ZED_SNIPPETS"
mkdir -p "$VSCODE_SNIPPETS"

cp snippets/go.json "$ZED_SNIPPETS/go.json"
cp snippets/go.json "$VSCODE_SNIPPETS/go.json"

echo "✓ Snippets instalados en Zed y VSCode"
