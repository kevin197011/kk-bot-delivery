# Design

## Architecture
- **Language**: Golang
- **Entry Point**: `main.go`
- **Libs**: Standard library `flag` for arguments (or `pflag`/`cobra`), `net/http` and `multipart` for Telegram API interaction to avoid heavy dependencies if possible, or use `tgbotapi`. Given the requirement for simplicity, using a reputable library like `github.com/go-telegram-bot-api/telegram-bot-api` is often safer and faster.

## CLI Usage
```bash
./sender --token <BOT_TOKEN> --user <USER_ID> --file <FILE_PATH>
```

## Internal Logic
1. Parse flags.
2. Validate file existence.
3. Initialize Bot API with token.
4. Construct `NewDocument` upload request.
5. Send file to `USER_ID`.
6. Print success or error.
