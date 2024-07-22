// Copyright 2023 The ppacer Authors.
// Licensed under the Apache License, Version 2.0.
// See LICENSE file in the project root for full license information.

// Package telegram implements ppacer notifier for Telegram chats.
package telegram

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ppacer/core/notify"
)

// Notifier implements ppacer/core/notify.Sender interface for sending
// notifications onto Telegram chat. Bot needs to be invited to the channel.
type Notifier struct {
	botToken   string
	channelId  int64
	httpClient *http.Client
}

// NewNotifier craetes new instance of Telegram Notifier.
func NewNotifier(botToken string, channelId int64, client *http.Client) *Notifier {
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}
	return &Notifier{
		botToken:   botToken,
		channelId:  channelId,
		httpClient: client,
	}
}

// Send sends templated message onto Telegram chat.
func (n *Notifier) Send(
	ctx context.Context, tmpl notify.Template, data notify.MsgData,
) error {
	var msgBytes bytes.Buffer
	writeErr := tmpl.Execute(&msgBytes, data)
	if writeErr != nil {
		return writeErr
	}
	url := n.sendMessageUrl(msgBytes.String())
	req, rErr := http.NewRequestWithContext(ctx, "GET", url, nil)
	if rErr != nil {
		return rErr
	}
	resp, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status %d, got: %d", http.StatusOK,
			resp.StatusCode)
	}
	return nil
}

func (c *Notifier) sendMessageUrl(text string) string {
	const urlTmpl = "https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s"
	encodedText := url.QueryEscape(text)
	return fmt.Sprintf(urlTmpl, c.botToken, c.channelId, encodedText)
}
