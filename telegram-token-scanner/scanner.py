import re
import requests

class TokenScanner:
    def __init__(self, config):
        self.pattern = config.get("pattern", r"(\d{9,10}:[A-Za-z0-9_-]{35})")
        self.targets = config.get("targets", [])
        
    def scan(self):
        found = []
        for url in self.targets:
            try:
                resp = requests.get(url, timeout=5)
                matches = re.findall(self.pattern, resp.text)
                for token in set(matches):
                    found.append({"token": token, "source": url})
            except Exception as e:
                print(f"Error scanning {url}: {e}")
        return found