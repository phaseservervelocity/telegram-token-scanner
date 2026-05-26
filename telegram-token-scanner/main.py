import sys
from scanner import TokenScanner
from reporter import Reporter
from config import load_config

def main():
    config = load_config("config.yaml")
    scanner = TokenScanner(config)
    reporter = Reporter(config)
    
    results = scanner.scan()
    if results:
        reporter.report(results)
    else:
        print("No tokens found in scan.")

if __name__ == "__main__":
    main()