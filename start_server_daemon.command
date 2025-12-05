#!/bin/bash

# Change to the directory containing this script
cd "$(dirname "$0")"

SCRIPT_DIR="$(pwd)"
PID_FILE="$SCRIPT_DIR/server/server.pid"
LOG_FILE="$SCRIPT_DIR/server/server.log"

# Load environment variables from .env if it exists
if [ -f "$SCRIPT_DIR/.env" ]; then
    export $(grep -v '^#' "$SCRIPT_DIR/.env" | xargs)
fi

# Set defaults if not in .env
NODE_PATH=${NODE_PATH:-$(which node)}
SERVER_PORT=${SERVER_PORT:-23030}

# Check if server is already running
if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    if ps -p "$PID" > /dev/null 2>&1; then
        echo "✅ Server is already running (PID: $PID)"
        exit 0
    else
        # Stale PID file, remove it
        rm "$PID_FILE"
    fi
fi

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

# Start server in background
echo "🎅 Starting Advent of Code Auto-Fetcher (daemon mode)..."
echo "📍 Using Node.js: $NODE_PATH"
echo "🔌 Port: $SERVER_PORT"
nohup "$NODE_PATH" server.js > "$LOG_FILE" 2>&1 &
SERVER_PID=$!

# Save PID
echo "$SERVER_PID" > "$PID_FILE"

# Wait a moment to check if it started successfully
sleep 1

if ps -p "$SERVER_PID" > /dev/null 2>&1; then
    echo "✅ Server started successfully (PID: $SERVER_PID)"
    echo "📝 Logs: $LOG_FILE"
    echo ""
    echo "Commands:"
    echo "  ./stop_server.sh     - Stop the server"
    echo "  ./server_status.sh   - Check server status"
    echo "  ./server_logs.sh     - View server logs"
else
    echo "❌ Failed to start server"
    echo "Check logs: $LOG_FILE"
    rm "$PID_FILE"
    exit 1
fi
