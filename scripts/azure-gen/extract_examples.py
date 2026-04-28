#!/usr/bin/env python3
"""
Extract HCL example blocks from terraform-provider-azurerm docs and write
one .tf file per resource/datasource.

Usage:
  ./extract_examples.py <outdir>           # only for generated (new) mappings
  ./extract_examples.py <outdir> --all     # every doc file
"""
import re
import sys
from pathlib import Path

HERE = Path(__file__).parent
DOCS = HERE / "docs"
GEN_RES = HERE / "generated"
GEN_DATA = HERE / "generated-data"

HCL_BLOCK = re.compile(r"```(?:hcl|terraform)\s*\n(.*?)```", re.DOTALL)
REF = re.compile(r"\b(?:data\.)?[a-z][a-z0-9_]*\.[a-z][a-z0-9_-]*(?:\.[a-zA-Z0-9_\[\]\.]+)?\b")


def isolate_block(hcl, kind, tf_type):
    """Return only the `<kind> "<tf_type>" "..." { ... }` block, brace-matched."""
    pat = re.compile(rf'^{kind}\s+"{re.escape(tf_type)}"\s+"[^"]+"\s*{{', re.M)
    m = pat.search(hcl)
    if not m:
        return None
    start = m.start()
    i = m.end()
    depth = 1
    while i < len(hcl) and depth:
        c = hcl[i]
        if c == "{":
            depth += 1
        elif c == "}":
            depth -= 1
        i += 1
    return hcl[start:i]


QUOTED = re.compile(r'"[^"]*"')


def stub_refs(block):
    """Replace cross-resource references with literal "pike" so the block
    stands alone without its example dependencies."""
    out_lines = []
    for line in block.splitlines():
        if "=" in line and not line.lstrip().startswith(("#", "//")):
            lhs, rhs = line.split("=", 1)
            # protect quoted literals so dotted strings inside them survive
            saved = []
            def keep(m):
                saved.append(m.group(0))
                return f"\x00{len(saved)-1}\x00"
            tmp = QUOTED.sub(keep, rhs)
            tmp = REF.sub('"pike"', tmp)
            for i, s in enumerate(saved):
                tmp = tmp.replace(f"\x00{i}\x00", s)
            out_lines.append(f"{lhs}={tmp}")
        else:
            out_lines.append(line)
    return "\n".join(out_lines)


def extract(doc_path, kind, tf_type):
    text = doc_path.read_text(encoding="utf-8")
    idx = text.find("## Example Usage")
    search = text[idx:] if idx >= 0 else text
    for m in HCL_BLOCK.finditer(search):
        block = isolate_block(m.group(1), kind, tf_type)
        if block:
            block = stub_refs(block)
            block = re.sub(
                rf'^{kind}\s+"{re.escape(tf_type)}"\s+"[^"]+"',
                f'{kind} "{tf_type}" "pike_gen"',
                block,
                count=1,
                flags=re.M,
            )
            return block + "\n"
    return None


def wanted_names(all_mode):
    if all_mode:
        res = {("r", p.stem.replace(".html", "")) for p in (DOCS / "r").glob("*.markdown")}
        data = {("d", p.stem.replace(".html", "")) for p in (DOCS / "d").glob("*.markdown")}
        return res, data
    res = {("r", p.stem[len("azurerm_"):]) for p in GEN_RES.rglob("azurerm_*.json")}
    data = {("d", p.stem[len("azurerm_"):]) for p in GEN_DATA.rglob("azurerm_*.json")}
    return res, data


def main():
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(1)
    outdir = Path(sys.argv[1])
    outdir.mkdir(parents=True, exist_ok=True)
    all_mode = "--all" in sys.argv

    res, data = wanted_names(all_mode)
    ok = nodoc = nohcl = 0

    for dockind, blockkind, items, prefix in [
        ("r", "resource", res, ""),
        ("d", "data", data, "data."),
    ]:
        for _, short in sorted(items):
            doc = DOCS / dockind / f"{short}.html.markdown"
            if not doc.exists():
                nodoc += 1
                continue
            hcl = extract(doc, blockkind, f"azurerm_{short}")
            if not hcl:
                nohcl += 1
                continue
            (outdir / f"{prefix}azurerm_{short}.tf").write_text(hcl)
            ok += 1

    print(f"wrote {ok} .tf files to {outdir}")
    print(f"missing doc: {nodoc}, no hcl block: {nohcl}")


if __name__ == "__main__":
    main()
