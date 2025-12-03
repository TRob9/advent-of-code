#!/bin/bash

# Change to the directory containing this script
cd "$(dirname "$0")"

echo "🎅 Starting Advent of Code Auto-Fetcher..."
echo ""

# Check if .session file exists
if [ ! -f ".session" ]; then
    echo "❌ Error: .session file not found!"
    echo ""
    echo "Please create a .session file with your Advent of Code session cookie."
    echo "Instructions:"
    echo "  1. Go to https://adventofcode.com and log in"
    echo "  2. Open DevTools (F12) → Application → Cookies"
    echo "  3. Copy the 'session' cookie value"
    echo "  4. Paste it into .session file in this directory"
    echo ""
    read -p "Press Enter to exit..."
    exit 1
fi

# Change to server directory
cd server

# Install dependencies if needed
if [ ! -d "node_modules" ]; then
    echo "📦 Installing dependencies..."
    npm install
    echo ""
fi

# Run the server
echo "Starting server..."
echo ""
node server.js
