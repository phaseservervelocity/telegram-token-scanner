require('dotenv').config();

module.exports = {
  apiId: parseInt(process.env.TELEGRAM_API_ID, 10),
  apiHash: process.env.TELEGRAM_API_HASH,
  botToken: process.env.BOT_TOKEN,
  targetChatId: process.env.TARGET_CHAT_ID,
  tokenPattern: /(\d{5,10}:[A-Za-z0-9_-]{35,45})/g
};