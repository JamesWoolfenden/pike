#!/usr/bin/env python3
"""
Populate modify and destroy arrays in pike permission mapping files.

Copies update-style and delete-style permissions from the apply array into
the modify and destroy arrays. Permissions are NOT removed from apply —
apply retains all permissions needed for terraform apply.

GCP:
  destroy — verbs: delete, remove, cancel, abandon
  modify  — verbs: update, patch, set, setIamPolicy, enable, disable,
                   resize, move, import, restore

Azure:
  destroy — suffix: /delete
  modify  — suffix: /write

AWS is left unchanged.

Run from the repo root:
    python3 tools/populate-modify-destroy.py [--dry-run]
"""

import json, os, sys, glob, argparse

GCP_MODIFY_VERBS = {
    "update", "patch", "set", "setIamPolicy", "enable", "disable",
    "resize", "move", "import", "restore",
}

GCP_DESTROY_VERBS = {
    "delete", "remove", "cancel", "abandon",
}


def classify_gcp(perm: str):
    verb = perm.split(".")[-1]
    if verb in GCP_DESTROY_VERBS:
        return "destroy"
    if verb in GCP_MODIFY_VERBS:
        return "modify"
    return None


def classify_azure(perm: str):
    p = perm.lower()
    if p.endswith("/delete"):
        return "destroy"
    if p.endswith("/write"):
        return "modify"
    return None


def process_file(path: str, classify_fn, dry_run: bool) -> bool:
    try:
        with open(path, encoding="utf-8") as f:
            data = json.load(f)
    except Exception as e:
        print(f"  SKIP {path}: {e}", file=sys.stderr)
        return False

    if not isinstance(data, list) or not data:
        return False

    changed = False
    for entry in data:
        if not isinstance(entry, dict):
            continue

        apply_in   = list(entry.get("apply",   []) or [])
        modify_in  = set(entry.get("modify",   []) or [])
        destroy_in = set(entry.get("destroy",  []) or [])

        new_modify  = set(modify_in)
        new_destroy = set(destroy_in)

        for perm in apply_in:
            bucket = classify_fn(perm)
            if bucket == "destroy":
                new_destroy.add(perm)
            elif bucket == "modify":
                new_modify.add(perm)

        if new_modify == modify_in and new_destroy == destroy_in:
            continue

        # apply is unchanged — permissions are copied, not moved
        entry["modify"]  = sorted(new_modify)
        entry["destroy"] = sorted(new_destroy)
        changed = True

    if changed:
        rel = os.path.relpath(path)
        if dry_run:
            print(f"  would update: {rel}")
        else:
            with open(path, "w", encoding="utf-8", newline="\n") as f:
                json.dump(data, f, indent=2)
                f.write("\n")
            print(f"  updated: {rel}")

    return changed


def run(mapping_root: str, dry_run: bool):
    providers = {
        "google":  (classify_gcp,   "resource"),
        "azurerm": (classify_azure, "resource"),
    }
    total_files = total_changed = 0
    for provider, (classify_fn, kind) in providers.items():
        pattern = os.path.join(mapping_root, provider, kind, "**", "*.json")
        files   = sorted(glob.glob(pattern, recursive=True))
        changed = 0
        print(f"\n{provider} ({len(files)} files):")
        for path in files:
            if process_file(path, classify_fn, dry_run):
                changed += 1
        print(f"  {changed} file(s) {'would be ' if dry_run else ''}updated")
        total_files   += len(files)
        total_changed += changed
    print(f"\nTotal: {total_changed}/{total_files} files {'to update' if dry_run else 'updated'}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--dry-run", action="store_true")
    args = parser.parse_args()
    repo_root    = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    mapping_root = os.path.join(repo_root, "src", "mapping")
    print(f"populate-modify-destroy — {'DRY RUN' if args.dry_run else 'UPDATING FILES'}")
    print(f"mapping root: {mapping_root}")
    run(mapping_root, args.dry_run)
