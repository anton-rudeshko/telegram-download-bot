# Telegram Download Bot

The bot downloads received [_documents_][telegram-documents] (not to be confused with photo, audio, video) to a specified location which can be optionally filtered by MIME type.
The bot only handles messages from allowed users and if a document is successfully downloaded the bot will send a reply message.

It is meant to be used with other server software waiting for new files at the destination location.

[telegram-documents]: https://core.telegram.org/bots/api#document

## Usage

1. Get a **Telegram API** bot token:
    * Use an existing one or
    * speak with [@BotFather] to create a new bot, see [docs][botfather-docs].
1. Create a config file (see below).
1. Download or build a binary for a platform on which the bot will be used:

    ```console
    $ git version
    git version 2.21.0
    $ git clone git@github.com:anton-rudeshko/telegram-download-bot.git
    $ cd telegram-download-bot
    $ go version
    go version go1.12.1 darwin/amd64
    $ GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
    ```

1. Start the bot:

    ```console
    $ ./telegram-download-bot /path/to/config.json
    ```

[@BotFather]: https://t.me/BotFather
[botfather-docs]: https://core.telegram.org/bots#6-botfather

## Configuration

Configuration is read from a JSON file and you need to provide a path to one via command line arguments:

```console
$ ./telegram-download-bot /path/to/config.json
2021/11/03 18:21:49 Starting up...
2021/11/03 18:21:49 Reading config from "/path/to/config.json"
2021/11/03 18:21:49 token is "NNNNNNNNN:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
2021/11/03 18:21:49 Using proxy "socks5://user:pass@host.tld:port"
2021/11/03 18:21:50 Authorized as your_awesome_bot
…
```

At the very least a configuration file must include `bot_token`, `location`, and `allowed_user_ids` fields:

```json
{
  "bot_token": "NNNNNNNNN:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "location": "/download/location",
  "allowed_user_ids": [
    12345678,
    23456789
  ]
}
```

Required fields:

* `bot_token` (`string`) — Surely you need a bot token.
* `location` (`string`) — Store received files to this location.

    * The file will be named same as the received document.
    * Files are limited in size by Telegram API to a maximum of 20 MB.
    * Existing files will be overwritten without a warning.

* `allowed_user_ids` (`[]int`) — Filter incoming messages by sender's Telegram user ID.

    * This is required because no one likes to store documents from random users around the world.
    * To determine your user ID you may want to use [@userinfobot] ([source][@userinfobot-source]).

[@userinfobot]: https://t.me/userinfobot
[@userinfobot-source]: https://github.com/nadam/userinfobot

Optional fields:

* `bot_debug` (`bool`; `false` by default) — The value of this field is passed to underlying Telegram API client library ([go-telegram-bot-api/telegram-bot-api]).
* `bot_poll_timeout` (`int`; `60` by default) — Timeout in seconds for API requests [long polling].
* `proxy_url` (`string`; empty by default) — Use a proxy for all interaction with Telegram API.

    * Expecting a URL in form of `(http|https|socks5)://user:pass@host.tld:port`.

* `mime_whitelist` (`[]string`; empty slice by default) — Filter incoming documents by their MIME type.
* `success_text` (`string`; `"✅ Done."` by default) — This message will be sent as the reply to the received message with the original document.

[go-telegram-bot-api/telegram-bot-api]: https://github.com/go-telegram-bot-api/telegram-bot-api

[long polling]: https://en.wikipedia.org/wiki/Push_technology#Long_polling

## License

MIT
