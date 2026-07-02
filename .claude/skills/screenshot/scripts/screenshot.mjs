// Capture a screenshot of the running theme demo with the pre-installed
// Chromium via Playwright. See ../SKILL.md for the full workflow.
//
// Usage:
//   node screenshot.mjs --url /post/typography-and-markdown/ --out out.png
//   node screenshot.mjs --url / --out home.png --viewport mobile --full
//   node screenshot.mjs --url /post/foo/ --out widgets.png \
//     --clip ".toc,.twitter-share.widget" --pad 28
//
// Flags:
//   --url <path>        Path on the dev server (default: /). Prepend the base
//                       yourself if you changed --baseURL.
//   --out <file>        Output PNG path (default: screenshot.png).
//   --base <url>        Dev server origin (default: http://localhost:1313).
//   --viewport d|m      desktop (1440x1400) or mobile (390x844). Default desktop.
//   --width <n>         Override viewport width.
//   --height <n>        Override viewport height.
//   --clip <sel[,sel]>  Clip to the union bounding box of these selectors
//                       (plus --pad margin) instead of the full viewport.
//   --pad <n>           Padding around the clip union in px (default: 28).
//   --full              Full-page screenshot (ignores --clip).
//   --theme light|dark  Force the color theme via ?theme or data-theme.
//   --measure           Also print each --clip selector's box to stdout as JSON
//                       (handy for verifying alignment without eyeballing).

import { createRequire } from 'node:module';
import { pathToFileURL } from 'node:url';

// Resolve `playwright` from the current working directory, not from this
// script's location — the browser is pre-installed globally but the npm
// package is typically installed ad hoc in a scratch dir the user runs from.
async function loadChromium() {
  try {
    return (await import('playwright')).chromium;
  } catch {
    const require = createRequire(pathToFileURL(process.cwd() + '/'));
    return require('playwright').chromium;
  }
}
const chromium = await loadChromium();

function arg(name, def) {
  const i = process.argv.indexOf(`--${name}`);
  if (i === -1) return def;
  const next = process.argv[i + 1];
  return next && !next.startsWith('--') ? next : true;
}

const base = arg('base', 'http://localhost:1313').replace(/\/$/, '');
const url = base + arg('url', '/');
const out = arg('out', 'screenshot.png');
const viewport = arg('viewport', 'desktop');
const isMobile = viewport === 'mobile' || viewport === 'm';
const width = Number(arg('width', isMobile ? 390 : 1440));
const height = Number(arg('height', isMobile ? 844 : 1400));
const clip = arg('clip', null);
const pad = Number(arg('pad', 28));
const full = arg('full', false) === true;
const theme = arg('theme', null);
const measure = arg('measure', false) === true;

// The web/CI container ships Chromium under PLAYWRIGHT_BROWSERS_PATH but the
// npm `playwright` build number may differ, so let Playwright resolve it first
// and only fall back to a discovered binary if that fails.
async function launch() {
  try {
    return await chromium.launch();
  } catch {
    const { execSync } = await import('node:child_process');
    const root = process.env.PLAYWRIGHT_BROWSERS_PATH || '/opt/pw-browsers';
    const found = execSync(
      `find ${root} -type f -name chrome -path '*chrome-linux*' 2>/dev/null | head -1`,
    ).toString().trim();
    if (!found) throw new Error('Could not find a Chromium binary under ' + root);
    return await chromium.launch({ executablePath: found });
  }
}

const browser = await launch();
const page = await browser.newPage({
  viewport: { width, height },
  deviceScaleFactor: 2,
  isMobile,
  // The theme is decided by an inline script at load time (localStorage, else
  // prefers-color-scheme), and the JS glass background initializes off it too —
  // so the theme must be right BEFORE navigation, not toggled afterwards.
  ...(theme ? { colorScheme: theme } : {}),
});

if (theme) {
  await page.addInitScript((t) => {
    try { localStorage.setItem('liquid-glass-theme', t); } catch (e) {}
  }, theme);
}

await page.goto(url, { waitUntil: 'networkidle' });

if (measure && clip) {
  const boxes = {};
  for (const sel of clip.split(',').map((s) => s.trim())) {
    const el = await page.$(sel);
    boxes[sel] = el ? await el.boundingBox() : null;
  }
  console.log(JSON.stringify(boxes, null, 2));
}

if (!full && clip) {
  let x0 = Infinity, y0 = Infinity, x1 = -Infinity, y1 = -Infinity;
  let any = false;
  for (const sel of clip.split(',').map((s) => s.trim())) {
    const el = await page.$(sel);
    if (!el) { console.error(`clip selector not found: ${sel}`); continue; }
    const b = await el.boundingBox();
    if (!b) continue;
    any = true;
    x0 = Math.min(x0, b.x); y0 = Math.min(y0, b.y);
    x1 = Math.max(x1, b.x + b.width); y1 = Math.max(y1, b.y + b.height);
  }
  if (!any) throw new Error('none of the --clip selectors matched a visible element');
  await page.screenshot({
    path: out,
    clip: { x: x0 - pad, y: y0 - pad, width: x1 - x0 + pad * 2, height: y1 - y0 + pad * 2 },
  });
} else {
  await page.screenshot({ path: out, fullPage: full });
}

await browser.close();
console.log(`saved ${out}`);
