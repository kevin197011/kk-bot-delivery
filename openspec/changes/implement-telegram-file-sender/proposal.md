# Implement Telegram File Sender CLI

## Summary
Create a Golang CLI tool that sends a specified file to a Telegram user (the bot's owner) using a Bot Token.

## Motivation
The user needs a quick and automated way to send files to their Telegram account for delivery or backup purposes directly from the command line.

## Scope
- Create a new Golang application/script.
- Support command-line arguments for:
  - Bot Token
  - Target User ID (Owner Channel/Chat ID)
  - File Path
- Graceful error handling and success confirmation.
