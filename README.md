# KK Bot Delivery CLI

A simple and efficient command-line tool written in Go to send files to a specific Telegram user (the bot owner) using a Telegram Bot Token.

## Features

- 🚀 **Fast & Lightweight**: Built with Go, producing a single binary with no runtime dependencies.
- 🔒 **Secure**: Uses the official Telegram Bot API.
- 🖥️ **Cross-Platform**: Support for macOS, Linux, and Windows.
- 🛠️ **DevOps Friendly**: Perfect for backup scripts, CI/CD pipelines, or server monitoring notifications.

## Installation

### Download Binary
Download the latest release for your platform from the [Releases Page](https://github.com/kevin197011/kk-bot-delivery/releases).

### Build from Source
Ensure you have Go 1.25+ installed.

```bash
git clone https://github.com/kevin197011/kk-bot-delivery.git
cd kk-bot-delivery
go mod tidy
go build -o sender main.go
```

## Usage

Run the tool with the required flags:

```bash
./sender --token <YOUR_BOT_TOKEN> --user <TARGET_USER_ID> --file <PATH_TO_FILE>
```

### Arguments

| Flag | Description | Required | Example |
|------|-------------|:--------:|---------|
| `--token` | Your Telegram Bot Token (from @BotFather) | ✅ | `123456:ABC-DEF1234ghIkl-zyx57W2v1u1` |
| `--user` | The target Telegram User ID (Chat ID) | ✅ | `987654321` |
| `--file` | Absolute or relative path to the file to upload | ✅ | `./backup.tar.gz` |

### Example

```bash
./sender \
  --token "123456:ABC-DEF1234ghIkl-zyx57W2v1u1" \
  --user 987654321 \
  --file "/var/backups/daily_report.pdf"
```

## Development

### Prerequisites
- Go 1.25 or higher

### Running Locally
```bash
go run main.go --token "..." --user 123 --file "test.txt"
```

### Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
[MIT](LICENSE)
