const Scanner = require('./scanner');
const config = require('./config');

const main = async () => {
  console.log('Telegram Token Scanner starting...');
  const scanner = new Scanner(config);
  await scanner.start();
};

main().catch(err => {
  console.error('Fatal error:', err);
  process.exit(1);
});