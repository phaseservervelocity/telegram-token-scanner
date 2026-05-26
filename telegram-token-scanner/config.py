import yaml

def load_config(path):
    try:
        with open(path, "r") as f:
            return yaml.safe_load(f)
    except FileNotFoundError:
        return {
            "pattern": r"(\d{9,10}:[A-Za-z0-9_-]{35})",
            "targets": [],
            "output_file": "results.txt"
        }