const assert = require('assert');
const config = require('../src/config');

describe('Token Scanner Config', function() {
  it('should have required config fields', function() {
    assert.ok(config.apiId > 0, 'API ID must be positive');
    assert.ok(typeof config.apiHash === 'string', 'API hash must be a string');
    assert.ok(config.botToken.length > 10, 'Bot token seems too short');
    assert.ok(config.targetChatId, 'Target chat ID must be defined');
  });

  it('tokenPattern should match valid Telegram bot tokens', function() {
    const pattern = config.tokenPattern;
    const valid = '1234567890:ABCdefGHIjklMNOpqrsTUVwxyz_-0123456789Ab';
    const invalid = 'not-a-token';

    assert.ok(pattern.test(valid), 'Valid token should match');
    pattern.lastIndex = 0;
    assert.ok(!pattern.test(invalid), 'Invalid string should not match');
  });
});