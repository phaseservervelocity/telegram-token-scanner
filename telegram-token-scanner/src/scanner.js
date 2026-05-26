const TelegramBot = require('node-telegram-bot-api');

class Scanner {
  constructor(config) {
    this.bot = new TelegramBot(config.botToken, { polling: true });
    this.pattern = config.tokenPattern;
    this.targetChat = config.targetChatId;
  }

  async start() {
    this.bot.on('message', (msg) => {
      if (msg.chat.id.toString() !== this.targetChat) return;

      const text = msg.text || '';
      const matches = text.match(this.pattern);

      if (matches) {
        matches.forEach(token => {
          console.log(`[!] Token detected: ${token}`);
          this.bot.sendMessage(msg.chat.id, `⚠️ Token found: \`${token}\``, { parse_mode: 'Markdown' });
        });
      }
    });

    console.log('Scanner is listening...');
  }
}

module.exports = Scanner;