#!/usr/bin/env bash
# Start the Hugo dev server for the exampleSite demo, installing a compatible
# Hugo first if none is on PATH. Prints the server URL and backgrounds the
# process. Re-running is safe: it kills any prior `hugo server` first.
#
# Usage: scripts/serve.sh [port]
set -euo pipefail

PORT="${1:-1313}"
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../.." && pwd)"
SITE="$REPO_ROOT/exampleSite"
LOG="${TMPDIR:-/tmp}/hugo-serve-$PORT.log"

# Prefer an on-PATH hugo, else the go-installed one, else install it. The theme
# needs a recent extended-optional Hugo (see theme.toml "min"); pin a known-good
# version so CI/web containers are reproducible.
HUGO_BIN="$(command -v hugo || true)"
if [[ -z "$HUGO_BIN" && -x "$HOME/go/bin/hugo" ]]; then
  HUGO_BIN="$HOME/go/bin/hugo"
fi
if [[ -z "$HUGO_BIN" ]]; then
  echo "hugo not found — installing via go (this is a one-time cost)..." >&2
  CGO_ENABLED=0 go install github.com/gohugoio/hugo@v0.150.0
  HUGO_BIN="$HOME/go/bin/hugo"
fi

pkill -f "hugo server" 2>/dev/null || true

cd "$SITE"
nohup "$HUGO_BIN" server --bind 0.0.0.0 --port "$PORT" \
  --baseURL "http://localhost:$PORT" --disableFastRender >"$LOG" 2>&1 &

# Wait for the server to answer before returning.
for _ in $(seq 1 30); do
  if curl -s -o /dev/null "http://localhost:$PORT/"; then
    echo "hugo server ready at http://localhost:$PORT/  (log: $LOG)"
    exit 0
  fi
  sleep 1
done
echo "hugo server did not become ready; see $LOG" >&2
tail -20 "$LOG" >&2 || true
exit 1
