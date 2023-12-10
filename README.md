# kbot
DevOps application from scratch

## Simple password generator


t.me/autonibit_bot

1. Make sure Go is installed — `go version`, otherwise https://go.dev/doc/install
2. Install dependencies github.com/spf13/cobra and gopkg.in/telebot.v3 — `go get`
3. Import token — `read -s TELE_TOKEN, CTRL+V Enter, export TELE_TOKEN`
4. Build — `go build -ldflags "-X="github.com/autonibit/kbot/cmd.appVersion=v1.0.3`
5. Run `./kbot start`

### Commands:

- `/start` — greeting
- `/generate` — generate a password with a default length
- `/generate N` — generate a password with a specified length
- `/help` — FAQ


![Screenshot from 2023-11-13 20-51-22](https://github.com/autonibit/kbot/assets/150580646/71f6b172-3168-4852-9956-be101d8348f8)

![workfow](img/workflow.svg)
