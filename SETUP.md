# Advent of Code Auto-Fetcher Setup Guide

This guide will help you set up the Advent of Code auto-fetcher daemon on your machine.

## Prerequisites

- Node.js installed (any recent version)
- macOS (for daemon/launchd features)
- Advent of Code account and session cookie

## Quick Setup

### 1. Get Your Session Cookie

1. Go to https://adventofcode.com and log in
2. Open DevTools (F12 or Cmd+Option+I)
3. Go to Application → Cookies
4. Find the `session` cookie and copy its value
5. Create a `.session` file in this directory with just the cookie value:

```bash
echo "your_session_cookie_value_here" > .session
```

### 2. Configure Environment

Copy the example environment file and customize it:

```bash
cp .env.example .env
```

Edit `.env` and update these values:

```bash
# Find your Node.js path
which node

# Update .env with your actual path
NODE_PATH=/path/to/your/node  # e.g., /opt/homebrew/bin/node
SERVER_PORT=23030              # Change if port conflicts
```

**Common Node.js locations:**
- Homebrew (Intel): `/usr/local/bin/node`
- Homebrew (Apple Silicon): `/opt/homebrew/bin/node`
- nvm: `~/.nvm/versions/node/vX.X.X/bin/node`

### 3. Install Dependencies

```bash
cd server
npm install
cd ..
```

### 4. Test the Server

Start the server manually first to ensure everything works:

```bash
./start_server_daemon.command
```

You should see:
```
🎅 Starting Advent of Code Auto-Fetcher (daemon mode)...
📍 Using Node.js: /path/to/node
🔌 Port: 23030
✅ Server started successfully (PID: xxxxx)
```

Visit http://localhost:23030 to confirm it's running.

### 5. (Optional) Enable Auto-Start on Login

To have the server start automatically when you log in:

#### macOS (using launchd)

1. Create the launch agent plist:

```bash
mkdir -p ~/Library/LaunchAgents
```

2. Create `~/Library/LaunchAgents/com.yourusername.advent-of-code-server.plist`:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.yourusername.advent-of-code-server</string>

    <key>ProgramArguments</key>
    <array>
        <string>/full/path/to/advent-of-code/start_server_daemon.command</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/full/path/to/advent-of-code/server/launchd.log</string>

    <key>StandardErrorPath</key>
    <string>/full/path/to/advent-of-code/server/launchd.error.log</string>

    <key>WorkingDirectory</key>
    <string>/full/path/to/advent-of-code</string>
</dict>
</plist>
```

3. Load the agent:

```bash
launchctl load ~/Library/LaunchAgents/com.yourusername.advent-of-code-server.plist
```

**Note:** Replace `/full/path/to/advent-of-code` with the actual absolute path.

## Usage

### Manual Control

```bash
# Start server
./start_server_daemon.command

# Stop server
./stop_server.command

# Check status
./server_status.command

# View logs
./server_logs.command           # Last 50 lines
./server_logs.command -f         # Follow in real-time
./server_logs.command 100        # Last 100 lines
```

### launchd Control (if using auto-start)

```bash
# Start
launchctl load ~/Library/LaunchAgents/com.yourusername.advent-of-code-server.plist

# Stop
launchctl unload ~/Library/LaunchAgents/com.yourusername.advent-of-code-server.plist

# Check status
launchctl list | grep advent-of-code
```

## Configuration Reference

### .env Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `NODE_PATH` | Full path to Node.js binary | `$(which node)` |
| `SERVER_PORT` | HTTP port for the server | `23030` |

### Files to Keep Private

**Never commit these files to git:**
- `.env` - Your environment configuration
- `.session` - Your Advent of Code session cookie
- `server/server.pid` - Process ID file
- `server/server.log` - Server logs
- `server/launchd.log` - launchd stdout
- `server/launchd.error.log` - launchd stderr

These are already in `.gitignore`.

## Troubleshooting

### Server won't start

1. Check Node.js path is correct:
   ```bash
   cat .env | grep NODE_PATH
   # Then verify:
   /path/from/env --version
   ```

2. Check if port is already in use:
   ```bash
   lsof -i :23030
   ```

3. Check logs for errors:
   ```bash
   ./server_logs.command
   ```

### Session cookie expired

If you get authentication errors:

1. Log in to Advent of Code again
2. Get a fresh session cookie
3. Update `.session` file
4. Restart server:
   ```bash
   ./stop_server.command
   ./start_server_daemon.command
   ```

### launchd not starting server

1. Check launchd logs:
   ```bash
   cat server/launchd.error.log
   ```

2. Ensure paths in plist are absolute (not relative)

3. Test manual start first to isolate the issue

## Sleep/Wake Behavior

When you close your MacBook:
- Server processes are **suspended** (not terminated)
- On wake, processes **resume automatically**
- If server crashes during wake, launchd will restart it

This means the server should survive sleep/wake cycles in most cases.

## Port Conflicts

If port 23030 is in use by another service:

1. Choose a new port (e.g., 23031)
2. Update `.env`:
   ```bash
   SERVER_PORT=23031
   ```
3. Restart server

## Security Notes

- `.session` file contains your Advent of Code authentication
- Keep `.session` private - don't commit to git
- Session cookies can expire - you may need to refresh them
- Server runs locally - not accessible from network by default

## Contributing

When sharing this project or contributing:

1. Never commit `.env` or `.session` files
2. Always use `.env.example` as a template
3. Document any new configuration options
4. Test on a fresh clone to ensure setup works

## Support

For issues or questions:
1. Check logs: `./server_logs.command`
2. Try manual start to see detailed errors
3. Verify Node.js version: `node --version`
4. Ensure `.session` file exists and is valid
