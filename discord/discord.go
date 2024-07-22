// Copyright 2023 The ppacer Authors.
// Licensed under the Apache License, Version 2.0.
// See LICENSE file in the project root for full license information.

// Package discord implements ppacer notifier for Discord channels.
package discord

import (
	"bytes"
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/ppacer/core/notify"
)

// Notifier implements ppacer/core/notify.Sender interface for sending
// notifications onto Discord channel.
type Notifier struct {
	token     string
	channelId string
}

// NewNotifier creates new instance of discord Notifier.
func NewNotifier(botToken, channelId string) *Notifier {
	return &Notifier{
		token:     botToken,
		channelId: channelId,
	}
}

// Send sends templated message as a regular Discord message.
func (n *Notifier) Send(
	ctx context.Context, tmpl notify.Template, data notify.MsgData,
) error {
	session, sessErr := discordgo.New("Bot " + n.token)
	if sessErr != nil {
		return sessErr
	}
	session.Open()
	defer session.Close()

	var msgBytes bytes.Buffer
	writeErr := tmpl.Execute(&msgBytes, data)
	if writeErr != nil {
		return writeErr
	}

	_, sendErr := session.ChannelMessageSend(n.channelId, msgBytes.String())
	if sendErr != nil {
		return sendErr
	}
	return nil
}
