import json, os, re, sys, urllib.request, urllib.error
from concurrent.futures import ThreadPoolExecutor

def get_latest_action(action_name):
    base = action_name.split('@')[0]
    headers = {"Accept": "application/vnd.github.v3+json", "User-Agent": "Python-StdLib"}
    try:
        req = urllib.request.Request(f"https://api.github.com/repos/{base}/releases/latest", headers=headers)
        with urllib.request.urlopen(req) as res:
            data = json.load(res)
            tag, url = data['tag_name'], data['html_url']

        req_commit = urllib.request.Request(f"https://api.github.com/repos/{base}/commits/{tag}", headers=headers)
        with urllib.request.urlopen(req_commit) as res:
            sha = json.load(res)['sha']
        return base, {"tag": tag, "sha": sha, "url": url}
    except Exception:
        return base, None

def process_workflows(path):
    files = []
    if os.path.isfile(path):
        files.append(path)
    elif os.path.isdir(path):
        for root, _, filenames in os.walk(path):
            for f in filenames:
                if f.endswith(('.yml', '.yaml')):
                    files.append(os.path.join(root, f))
    
    if not files:
        print(f"No workflow files found at: {path}")
        return

    all_action_refs = set()
    file_contents = {}

    for file_path in files:
        try:
            with open(file_path, 'r') as f:
                content = f.read()
            file_contents[file_path] = content
            refs = re.findall(r'uses:\s*([\w\-\./]+@?[\w\.-]*)', content)
            all_action_refs.update(refs)
        except Exception as e:
            print(f"Error reading {file_path}: {e}")

    clean_refs = [ref for ref in all_action_refs if ref.startswith('actions/')]
    if not clean_refs:
        print("No 'actions/' references found to update.")
        return

    print(f"Checking latest versions for {len(clean_refs)} unique actions...")
    with ThreadPoolExecutor() as executor:
        results = dict(filter(None, executor.map(get_latest_action, clean_refs)))

    for file_path, content in file_contents.items():
        new_content = content
        file_updated = False
        
        for base, data in results.items():
            if not data: continue
            pattern = rf'uses:\s*{re.escape(base)}@[^\r\n]*'
            
            if re.search(pattern, new_content):
                replacement = f'uses: {base}@{data["sha"]} # {data["tag"]}'
                new_content = re.sub(pattern, replacement, new_content)
                file_updated = True
                print(f"[{os.path.basename(file_path)}] Updated {base} to {data['tag']}")

        if file_updated:
            with open(file_path, 'w') as f:
                f.write(new_content)

if __name__ == "__main__":
    if len(sys.argv) > 1: 
        process_workflows(sys.argv[1])
    else:
        print("Usage: python update.py <file_or_directory_path>")
