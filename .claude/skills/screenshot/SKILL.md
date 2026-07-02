---
name: screenshot
description: >-
  Render and screenshot this Hugo theme's demo (exampleSite) with the
  pre-installed Chromium + Playwright, so a layout/CSS/template change can be
  seen and visually verified instead of guessed at. Use this WHENEVER the user
  wants to look at, preview, screenshot, or visually confirm anything the theme
  renders — the sidebar, TOC, share widget, cards, article page, mobile FAB,
  light/dark theme — including terse asks like "スクショ撮って", "見せて",
  "take a screenshot", "show me the share widget on mobile", "before/after of
  this padding change", or "does this actually line up?". Also use it to MEASURE
  element positions (padding/alignment) when a change is supposed to line two
  things up. Prefer this over hand-rolling a Playwright script each time.
---

# Screenshot the theme demo

Capture what the theme actually renders. The flow is always: **serve the demo →
screenshot a page or component → look at the PNG** (and, for alignment work,
read the measured boxes). Two bundled scripts do the heavy lifting so you don't
re-derive them each time.

## Prerequisites (usually already met in web/CI containers)

- **Chromium**: pre-installed under `PLAYWRIGHT_BROWSERS_PATH` (`/opt/pw-browsers`).
  Do **not** run `playwright install`.
- **Playwright npm package**: the browser is present but the JS library may not
  be. If `node -e "require.resolve('playwright')"` fails, install it once in a
  scratch dir and run the script from there:
  ```bash
  cd "$SCRATCH" && npm init -y >/dev/null 2>&1 && npm install playwright >/dev/null 2>&1
  ```
  Use your session scratchpad dir as `$SCRATCH` so it stays out of the repo.
- **Hugo**: `scripts/serve.sh` installs a pinned version via `go` if it's missing.

## Step 1 — Serve the demo

```bash
.claude/skills/screenshot/scripts/serve.sh          # defaults to port 1313
```

It backgrounds `hugo server` against `exampleSite/`, waits until it answers, and
prints the URL. Re-running kills any previous server first, so it's safe to call
again after editing templates (Hugo live-reloads CSS/HTML anyway — just give it
~2s before re-shooting).

## Step 2 — Screenshot

Run the bundled script from wherever `playwright` resolves (your scratch dir):

```bash
node .claude/skills/screenshot/scripts/screenshot.mjs \
  --url /post/typography-and-markdown/ --out "$SCRATCH/share.png" \
  --clip ".toc,.twitter-share.widget"
```

Then **Read the PNG** to actually see it — a screenshot you don't look at proves
nothing.

### Common recipes

| Goal | Flags |
| --- | --- |
| Full article page | `--url /post/typography-and-markdown/ --full` |
| A component only | `--clip ".twitter-share.widget"` (union of several: `--clip ".toc,.twitter-share.widget"`) |
| Mobile view / FAB | `--viewport mobile` (sidebar widgets are `display:none` here — the FAB replaces them) |
| Force theme | `--theme light` or `--theme dark` (demo defaults to light) |
| Verify alignment | add `--measure` to print each `--clip` selector's box as JSON |

`--url` takes a path; the article with both a TOC and the share widget is
`/post/typography-and-markdown/`. See `exampleSite/content/post/` for others.

## Alignment / padding work

When a change is meant to line two things up, don't trust your eyes on a
downscaled PNG — **measure**. `--measure` prints each selector's bounding box;
compare the `x` (or content-left) values. Example that confirmed the share
widget lines up with the TOC:

```bash
node .claude/skills/screenshot/scripts/screenshot.mjs \
  --url /post/typography-and-markdown/ --out "$SCRATCH/a.png" \
  --clip ".toc__title,.twitter-share__title" --measure
# both titles reporting the same x => aligned
```

Note a subtle cascade gotcha this theme has: sidebar widgets get their padding
from the late-loaded `.widget` rule (`space-md` block / `space-lg` inline). A
more specific selector like `.twitter-share.widget` overrides it, so if you
restate padding there, match `.widget` or the content will inset differently
from the TOC.

## Before/after comparisons

To show a change's effect, screenshot the current state, revert the change,
re-serve/reshoot, then restore. To put the two PNGs side by side with labels,
render a tiny HTML page that `<img>`s both (as data URIs) and screenshot that
with the same script pointed at a `file://` or `data:` URL — or just send both
PNGs to the user with a caption. Keep all intermediate PNGs in `$SCRATCH`, never
in the repo.

## Cleanup

Stop the server when done so it doesn't linger:

```bash
pkill -f "hugo server" || true
```
