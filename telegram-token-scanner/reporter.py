import datetime

class Reporter:
    def __init__(self, config):
        self.output_file = config.get("output_file", "results.txt")
        
    def report(self, results):
        with open(self.output_file, "a") as f:
            f.write(f"\n--- Scan at {datetime.datetime.now()} ---\n")
            for r in results:
                f.write(f"Token: {r['token']} | Source: {r['source']}\n")
        print(f"Results written to {self.output_file}")