#!/usr/bin/env node

/**
 * Advent of Code Auto-Fetcher with Claude SDK
 * - Auto-fetches problems and inputs at 4:00 PM
 * - Uses Claude to extract test cases from problem HTML
 * - Re-fetches Part 2 when Part 1 completes
 */

import { query } from '@anthropic-ai/claude-agent-sdk';
import http from 'http';
import https from 'https';
import { HttpsProxyAgent } from 'https-proxy-agent';
import TurndownService from 'turndown';
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const YEAR = 2025;
const UNLOCK_HOUR = 16; // 4:00 PM
const UNLOCK_MINUTE = 0;
const UNLOCK_SECOND = 5;
const CHECK_INTERVAL = 1000; // 1 second
const HTTP_PORT = 3030; // Port for harness to trigger Part 2 fetch

const REPO_ROOT = path.join(__dirname, '..');
const STATE_FILE = path.join(__dirname, 'state.json');

// Load or initialize state
function loadState() {
  try {
    if (fs.existsSync(STATE_FILE)) {
      return JSON.parse(fs.readFileSync(STATE_FILE, 'utf8'));
    }
  } catch (err) {
    console.error('Warning: Error loading state file:', err.message);
  }
  return {};
}

function saveState(state) {
  try {
    fs.writeFileSync(STATE_FILE, JSON.stringify(state, null, 2), 'utf8');
  } catch (err) {
    console.error('Warning: Error saving state file:', err.message);
  }
}

function hasFetched(state, year, day, part) {
  return state[year]?.[day]?.[`part${part}`] === true;
}

function markFetched(state, year, day, part) {
  if (!state[year]) state[year] = {};
  if (!state[year][day]) state[year][day] = {};
  state[year][day][`part${part}`] = true;
  saveState(state);
}

const state = loadState();

console.log('Advent of Code Auto-Fetcher Started!');
console.log(`Year: ${YEAR}`);
console.log(`Unlock Time: ${UNLOCK_HOUR.toString().padStart(2, '0')}:${UNLOCK_MINUTE.toString().padStart(2, '0')}:${UNLOCK_SECOND.toString().padStart(2, '0')}\n`);

// Get session cookie
const session = getSessionCookie();
if (!session) {
  console.error('Error: No session cookie found in .session file!');
  console.error('Please add your Advent of Code session cookie to .session');
  process.exit(1);
}

console.log('Success: Session cookie loaded');

// HTTP server for harness to trigger Part 2 fetch
const server = http.createServer(async (req, res) => {
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'POST, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');

  if (req.method === 'OPTIONS') {
    res.writeHead(200);
    res.end();
    return;
  }

  if (req.url === '/fetchPart2' && req.method === 'POST') {
    let body = '';
    req.on('data', chunk => {
      body += chunk.toString();
    });

    req.on('end', async () => {
      try {
        const data = JSON.parse(body);
        const { day } = data;

        if (!day) {
          res.writeHead(400, { 'Content-Type': 'application/json' });
          res.end(JSON.stringify({ success: false, error: 'Missing day parameter' }));
          return;
        }

        console.log(`\nComplete: Part 1 completed for Day ${day}! Fetching Part 2...`);

        await processDay(day, 2);

        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ success: true }));
      } catch (err) {
        console.error('Error: Error processing Part 2 fetch:', err.message);
        res.writeHead(500, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ success: false, error: err.message }));
      }
    });
    return;
  }

  if (req.url === '/populate' && req.method === 'POST') {
    let body = '';
    req.on('data', chunk => {
      body += chunk.toString();
    });

    req.on('end', async () => {
      try {
        const data = JSON.parse(body);
        const { day, part, prompt } = data;

        if (!day || !part || !prompt) {
          res.writeHead(400, { 'Content-Type': 'application/json' });
          res.end(JSON.stringify({ success: false, error: 'Missing parameters' }));
          return;
        }

        console.log(`\nAI: Populating test cases for Day ${day} Part ${part}...`);

        await populateTestsWithPrompt(day, prompt);

        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ success: true }));
      } catch (err) {
        console.error('Error: Error populating tests:', err.message);
        res.writeHead(500, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ success: false, error: err.message }));
      }
    });
    return;
  }

  res.writeHead(404, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({ success: false, error: 'Not found' }));
});

server.listen(HTTP_PORT, () => {
  console.log(`HTTP server listening on port ${HTTP_PORT}`);
  console.log('Waiting for puzzle unlock time...\n');
});

let lastCheckedDay = 0;

// Main loop
setInterval(async () => {
  const now = new Date();
  const currentDay = getCurrentDay(now);

  // Check if it's unlock time
  if (isUnlockTime(now) && currentDay > 0 && currentDay !== lastCheckedDay) {
    if (!hasFetched(state, YEAR, currentDay, 1)) {
      console.log(`\nAlert: Unlock time reached! Processing Day ${currentDay}...`);
      await processDay(currentDay, 1);
      lastCheckedDay = currentDay;
    }
  }

  // Heartbeat every minute
  if (now.getSeconds() === 0) {
    const time = now.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });
    console.log(`${time} - Waiting for unlock...`);
  }
}, CHECK_INTERVAL);

function getSessionCookie() {
  try {
    // Walk up directory tree to find .session
    let dir = REPO_ROOT;
    while (dir !== '/') {
      const sessionPath = path.join(dir, '.session');
      if (fs.existsSync(sessionPath)) {
        return fs.readFileSync(sessionPath, 'utf8').trim();
      }
      dir = path.dirname(dir);
    }
  } catch (err) {
    return null;
  }
  return null;
}

function getCurrentDay(now) {
  // Advent of Code runs Dec 1-25
  if (now.getMonth() !== 11 || now.getDate() > 25) {
    return 0;
  }
  return now.getDate();
}

function isUnlockTime(now) {
  return now.getHours() === UNLOCK_HOUR &&
         now.getMinutes() === UNLOCK_MINUTE &&
         now.getSeconds() >= UNLOCK_SECOND;
}

async function processDay(day, part) {
  console.log(`\nFetching Day ${day} Part ${part}...`);

  try {
    // Fetch problem description
    const problemHTML = await fetchProblem(day);
    if (problemHTML) {
      saveProblem(day, problemHTML);
    }

    // Fetch personal input
    const input = await fetchInput(day);
    if (input) {
      saveInput(day, input);
    }

    // Use Claude to populate tests.txt (reads problem.md)
    await populateTests(day, part);

    // Mark as fetched in state
    markFetched(state, YEAR, day, part);

    console.log(`Success: Day ${day} Part ${part} complete!`);
  } catch (err) {
    console.error(`Error: Error processing day ${day}:`, err.message);
  }
}

async function fetchProblem(day) {
  const url = `https://adventofcode.com/${YEAR}/day/${day}`;
  return await fetchURL(url);
}

async function fetchInput(day) {
  const url = `https://adventofcode.com/${YEAR}/day/${day}/input`;
  return await fetchURL(url);
}

async function fetchURL(url) {
  return new Promise((resolve, reject) => {
    const options = {
      headers: {
        'Cookie': `session=${session}`,
        'User-Agent': 'github.com/trob9/advent-of-code auto-fetcher'
      }
    };

    // Use proxy if configured
    const proxy = process.env.https_proxy || process.env.HTTPS_PROXY;
    if (proxy) {
      options.agent = new HttpsProxyAgent(proxy);
    }

    https.get(url, options, (response) => {
      if (response.statusCode !== 200) {
        reject(new Error(`HTTP ${response.statusCode}: ${response.statusMessage}`));
        return;
      }

      let data = '';
      response.on('data', chunk => data += chunk);
      response.on('end', () => resolve(data));
    }).on('error', (err) => {
      console.error(`Error: Error fetching ${url}:`, err.message);
      resolve(null);
    });
  });
}

function saveProblem(day, content) {
  const dayDir = path.join(REPO_ROOT, `${YEAR}`, `day${day}`);

  try {
    // Extract article content from HTML
    const articleRegex = /<article[^>]*class="day-desc"[^>]*>(.*?)<\/article>/gs;
    const articles = [];
    let match;

    while ((match = articleRegex.exec(content)) !== null) {
      articles.push(match[1]);
    }

    if (articles.length === 0) {
      console.log('Warning: No article content found in HTML');
      return;
    }

    // Convert to markdown
    const turndownService = new TurndownService({
      headingStyle: 'atx',
      codeBlockStyle: 'fenced'
    });

    let markdown = '';
    articles.forEach((article, index) => {
      const articleMarkdown = turndownService.turndown(article);
      markdown += articleMarkdown;
      if (index < articles.length - 1) {
        markdown += '\n\n---\n\n';
      }
    });

    // Save only as markdown (no HTML file)
    const problemPath = path.join(dayDir, 'problem.md');
    fs.writeFileSync(problemPath, markdown);
    console.log(`Success: Saved problem to ${problemPath}`);
  } catch (err) {
    console.error(`Error: Error saving problem:`, err.message);
  }
}

function saveInput(day, content) {
  const dayDir = path.join(REPO_ROOT, `${YEAR}`, `day${day}`);
  const inputPath = path.join(dayDir, 'input.txt');

  try {
    fs.writeFileSync(inputPath, content);
    console.log(`Success: Saved input to ${inputPath}`);
  } catch (err) {
    console.error(`Error: Error saving input:`, err.message);
  }
}

async function populateTests(day, part) {
  const dayDir = path.join(REPO_ROOT, `${YEAR}`, `day${day}`);
  const problemPath = path.join(dayDir, 'problem.md');

  // Read the markdown file
  let problemMarkdown = '';
  try {
    problemMarkdown = fs.readFileSync(problemPath, 'utf8');
  } catch (err) {
    console.error(`Warning: Could not read problem.md:`, err.message);
    return;
  }

  console.log(`AI: Using Claude to extract test cases for Part ${part}...`);

  let prompt;
  if (part === 1) {
    prompt = `
Extract the example test case from this Advent of Code problem and write it to testcases.txt.

Problem (Markdown):
${problemMarkdown}

CRITICAL FORMAT REQUIREMENTS:
The testcases.txt file is parsed by automated testing. You MUST follow this EXACT format with NO deviations:

*** Part 1 ***
input:
<raw example input - no code blocks, no formatting, just the literal text>
expected:
<just the number or answer - nothing else>

Rules:
1. Do NOT add markdown code blocks (\`\`\`) around the input
2. Do NOT add explanations, comments, or extra text
3. Do NOT add quotes around the input or expected value
4. The input should be the EXACT example input from the problem (usually shown in a code block in the problem)
5. The expected value should be JUST the answer (a number or string), nothing more
6. Do NOT add blank lines between the sections
7. After "input:" put the literal example input on the next line(s)
8. After "expected:" put ONLY the expected answer on the next line

Example of CORRECT format:
*** Part 1 ***
input:
L68
L30
R48
expected:
3

The file already exists but is empty. Write the Part 1 section to it.
`.trim();
  } else {
    prompt = `
Extract the Part 2 example test case and APPEND it to the existing testcases.txt file.

Problem (Markdown):
${problemMarkdown}

CRITICAL FORMAT REQUIREMENTS:
The testcases.txt file already contains Part 1 test cases. You MUST:
1. Read the existing file
2. Keep Part 1 EXACTLY as is
3. Add Part 2 AFTER Part 1

EXACT format to append:

*** Part 2 ***
input:
<raw example input - no code blocks, no formatting, just the literal text>
expected:
<just the number or answer - nothing else>

Rules:
1. Do NOT modify or remove the existing Part 1 section
2. Do NOT add markdown code blocks (\`\`\`) around the input
3. Do NOT add explanations, comments, or extra text
4. Do NOT add quotes around the input or expected value
5. The input should be the EXACT Part 2 example input from the problem
6. The expected value should be JUST the Part 2 answer (a number or string), nothing more
7. Do NOT add blank lines between the sections
8. After "input:" put the literal example input on the next line(s)
9. After "expected:" put ONLY the expected answer on the next line

Example of CORRECT full file after adding Part 2:
*** Part 1 ***
input:
L68
L30
R48
expected:
3

*** Part 2 ***
input:
L68
L30
R48
expected:
6

Note: Part 2 often uses the same example input as Part 1, but with a different expected answer. Check the Part 2 section carefully.
`.trim();
  }

  try {
    const result = query({
      prompt,
      options: {
        cwd: dayDir,
        systemPrompt: {
          type: 'preset',
          name: 'default'
        }
      }
    });

    // Stream Claude's response
    for await (const message of result) {
      if (message.type === 'stream_event') {
        const event = message.streamEvent;

        if (event.type === 'content_block_delta' && event.delta?.type === 'text_delta') {
          process.stdout.write(event.delta.text);
        }
      }
    }

    console.log(`\nSuccess: Tests populated for Part ${part}`);
  } catch (err) {
    console.error(`Error: Error populating tests:`, err.message);
  }
}

async function populateTestsWithPrompt(day, customPrompt) {
  const dayDir = path.join(REPO_ROOT, `${YEAR}`, `day${day}`);

  console.log(`AI: Using Claude with custom prompt...`);
  console.log(`AI: Working directory: ${dayDir}`);
  console.log(`AI: Settings should be loaded from: ${REPO_ROOT}/.claude/settings.json`);

  try {
    const result = query({
      prompt: customPrompt,
      options: {
        cwd: dayDir,
        systemPrompt: {
          type: 'preset',
          preset: 'claude_code'
        },
        env: {
          CLAUDE_CODE_USE_VERTEX: '1',
          ANTHROPIC_VERTEX_PROJECT_ID: process.env.ANTHROPIC_VERTEX_PROJECT_ID || '',
          CLOUD_ML_REGION: process.env.CLOUD_ML_REGION || '',
          ...process.env
        },
        includePartialMessages: true,
        canUseTool: async (toolName, input, options) => {
          console.log(`[AUTO-APPROVE] ${toolName}`);
          return { behavior: 'allow', updatedInput: input };
        }
      }
    });

    let sawToolUse = false;
    let currentThinkingText = '';
    let currentToolName = null;
    let currentToolInput = '';
    let finalMessage = null;

    for await (const message of result) {
      // Handle stream events (real-time updates)
      if (message.type === 'stream_event') {
        const event = message.event;

        // CONTENT BLOCK START - Tool use or text begins
        if (event.type === 'content_block_start' && event.content_block) {
          // If starting new block and we have accumulated thinking text, send it now
          if (currentThinkingText.trim()) {
            console.log(currentThinkingText);
            currentThinkingText = '';
          }

          if (event.content_block.type === 'tool_use') {
            sawToolUse = true;
            currentToolName = event.content_block.name;
            currentToolInput = '';
            console.log(`\nTool: Using ${currentToolName}...`);
          } else if (event.content_block.type === 'text') {
            currentThinkingText = '';
          }
        }

        // CONTENT BLOCK DELTA - Streaming text or tool input
        if (event.type === 'content_block_delta' && event.delta) {
          if (event.delta.type === 'text_delta') {
            const text = event.delta.text;
            currentThinkingText += text;
            process.stdout.write(text); // Stream to console in real-time
          }

          if (event.delta.type === 'input_json_delta') {
            currentToolInput += event.delta.partial_json;
          }
        }

        // CONTENT BLOCK STOP - Tool or text complete
        if (event.type === 'content_block_stop') {
          if (currentThinkingText.trim()) {
            console.log(''); // Newline after text block
            currentThinkingText = '';
          }

          if (currentToolName && currentToolInput) {
            try {
              const parsedInput = JSON.parse(currentToolInput);
              console.log(`Tool input:`, JSON.stringify(parsedInput, null, 2).substring(0, 200));
            } catch (e) {
              console.log(`Tool input:`, currentToolInput.substring(0, 200));
            }
            currentToolName = null;
            currentToolInput = '';
          }
        }

        // MESSAGE STOP - Complete assistant message
        if (event.type === 'message_stop') {
          console.log(''); // Newline after message
        }
      }

      // ASSISTANT MESSAGE (complete)
      if (message.type === 'assistant' && message.content) {
        finalMessage = message;
      }

      // USER MESSAGE - Contains tool results
      if (message.type === 'user' && message.content) {
        if (Array.isArray(message.content)) {
          for (const block of message.content) {
            if (block.type === 'tool_result') {
              const resultText = typeof block.content === 'string'
                ? block.content
                : JSON.stringify(block.content);
              console.log(`Tool result:`, resultText.substring(0, 300));
            }
          }
        }
      }
    }

    if (!sawToolUse) {
      console.log(`\n[WARNING] Claude did not use any tools! Check if prompt is directive enough.`);
    }

    console.log(`\nSuccess: Tests populated`);
  } catch (err) {
    console.error(`Error: Error populating tests:`, err.message);
    throw err;
  }
}

