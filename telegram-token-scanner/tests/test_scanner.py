import unittest
from scanner import TokenScanner

class TestTokenScanner(unittest.TestCase):
    def test_scan_with_mock(self):
        config = {
            "pattern": r"(\d{9,10}:[A-Za-z0-9_-]{35})",
            "targets": []
        }
        scanner = TokenScanner(config)
        result = scanner.scan()
        self.assertEqual(result, [])

if __name__ == "__main__":
    unittest.main()