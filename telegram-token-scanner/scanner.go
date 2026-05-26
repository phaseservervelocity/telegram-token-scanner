package main

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Scanner struct {
	addressPattern *regexp.Regexp
}

func NewScanner() *Scanner {
	return &Scanner{
		addressPattern: regexp.MustCompile(`\b(0x[a-fA-F0-9]{40})\b`),
	}
}

func (s *Scanner) HandleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil || update.Message.Text == "" {
		return
	}

	matches := s.addressPattern.FindAllString(update.Message.Text, -1)
	if len(matches) == 0 {
		return
	}

	var response strings.Builder
	response.WriteString("🔍 *Token Addresses Detected:*\n\n")

	for i, addr := range matches {
		response.WriteString(s.formatAddressInfo(i+1, addr))
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      response.String(),
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (s *Scanner) formatAddressInfo(index int, address string) string {
	return strings.Join([]string{
		"",
		"━━━━━━━━━━━━━━━━━━━",
		"",
		"📌 *Token #", string(rune('0'+index)), "*",
		"",
		"`", address, "`",
		"",
		"💎 [Etherscan](https://etherscan.io/address/", address, ")",
		"📊 [DexScreener](https://dexscreener.com/ethereum/", address, ")",
		"🦎 [CoinGecko](https://www.coingecko.com/en/coins/", address, ")",
		"",
	}, "")
}