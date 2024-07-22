# Repository for ppacer notifiers implementations

This repository contains implementations of ppacer external notification
senders. In technical terms those are concrete implementations of ppacer
`ppacer/core/notify.Sender` interface.

Notifiers implementations are usually wrapped in a separate Go package, to
limit number of pulled dependencies by users.


## Available implementations

### Discord

Discord notifier sends notifications onto configured Discord channel.

Installation:

```
go get github.com/ppacer/notifiers/discord

```

### Telegram

Telegram notifier sends notifications onto Telegram chat.

Installation:

```
go get github.com/ppacer/notifiers/telegram

```
