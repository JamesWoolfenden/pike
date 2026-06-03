#!/usr/bin/env python3
"""
Populate plan arrays in pike permission mapping files.

For each resource mapping JSON, reads the apply array and adds any
read-side permissions to the plan array that aren't already there.

GCP: read verbs are the last segment of "service.resource.verb"
Azure: permissions ending in /read
AWS: left unchanged (plan arrays are managed separately)

Run from the repo root:
    python3 tools/populate-plan-perms.py [--dry-run]
"""

import json
import os
import sys
import glob
import argparse

GCP_READ_VERBS = {
    "get", "list", "getIamPolicy", "fetch", "search", "query",
    "check", "access", "validate", "read", "view", "lookup",
    "describe", "export",
}


def is_gcp_read(perm: str) -> bool:
    parts = perm.split(".")
    return len(parts) >= 1 and parts[-1] in GCP_READ_VERBS


def is_azure_read(perm: str) -> bool:
    return perm.endswith("/read")


def process_file(path: str, filter_fn, dry_run: bool) -> bool:
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

        apply_perms = entry.get("apply", [])
        plan_perms  = entry.get("plan", [])

        if not isinstance(apply_perms, list) or not isinstance(plan_perms, list):
            continue

        read_from_apply = [p for p in apply_perms if filter_fn(p)]
        plan_set        = set(plan_perms)
        to_add          = [p for p in read_from_apply if p not in plan_set]

        if to_add:
            entry["plan"] = sorted(set(plan_perms) | set(to_add))
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
        "google":  (is_gcp_read,   "resource"),
        "azurerm": (is_azure_read, "resource"),
    }

    total_files   = 0
    total_changed = 0

    for provider, (filter_fn, kind) in providers.items():
        pattern = os.path.join(mapping_root, provider, kind, "**", "*.json")
        files   = sorted(glob.glob(pattern, recursive=True))
        changed = 0
        print(f"\n{provider} ({len(files)} files):")
        for path in files:
            if process_file(path, filter_fn, dry_run):
                changed += 1
        print(f"  {changed} file(s) {'would be ' if dry_run else ''}updated")
        total_files   += len(files)
        total_changed += changed

    print(f"\nTotal: {total_changed}/{total_files} files {'to update' if dry_run else 'updated'}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--dry-run", action="store_true",
                        help="Show what would change without writing files")
    args = parser.parse_args()

    repo_root    = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    mapping_root = os.path.join(repo_root, "src", "mapping")

    mode = "DRY RUN" if args.dry_run else "UPDATING FILES"
    print(f"populate-plan-perms — {mode}")
    print(f"mapping root: {mapping_root}")

    run(mapping_root, args.dry_run)
