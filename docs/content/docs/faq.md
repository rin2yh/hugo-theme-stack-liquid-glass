---
title: "FAQ"
description: "Common questions and gotchas."
date: 2026-04-27
weight: 50
categories: ["Docs"]
---

## Why are my favicon links 404?

The theme's `<head>` references `/apple-touch-icon.png`, `/favicon-32x32.png`, `/favicon-16x16.png`, `/favicon.ico`, and `/site.webmanifest` from your site's `static/` directory. They are intentionally **not** bundled with the theme — provide your own under `static/`. The site renders correctly without them, but the browser console will log 404s.

## Why are some menu icons missing?

`params.icon` on `[[menu.main]]` and `[[social]]` resolves to `assets/icons/<name>.svg`. The theme ships a generic icon set (`home`, `search`, `archives`, `rss`, `tag`, `folder`, `user`, `clock`, `date`, `external`, `link`, `list`, `moon`, `sun`, `language`, `copyright`, `clipboard`, `donate-heart`, `qr-code`).

Service-specific brand icons (GitHub, X/Twitter, Speaker Deck, Zenn, etc.) are intentionally not bundled. Place the SVG you want to use at `assets/icons/<name>.svg` in your own site. Hugo's asset pipeline resolves your project's `assets/` before the theme's, so site-provided icons take precedence and supplement the theme set.

If `params.icon` resolves to no file, the icon slot renders empty but the link itself still works.

## Does the theme require JavaScript?

The theme has a near-zero JS pipeline. The pieces that **do** run JS are:

- The dark-mode toggle.
- The full-text search widget (loads `/index.json` lazily).
- Code-block "copy" buttons.
- Mermaid diagrams (loaded on demand from a CDN, only on pages that use them).

Pages without search and Mermaid render with no JS at all.

## Where does Mermaid load from?

`layouts/_default/baseof.html` loads Mermaid from `https://cdn.jsdelivr.net/npm/mermaid@10`. If you need a fully self-hosted setup, replace the script tag in your site's override layout.

## How do I add a Twitter / Zenn / Speaker Deck icon?

1. Save the SVG at `assets/icons/<name>.svg` in your **site** (not in the theme).
2. Reference it from a `[[social]]` entry:
   ```toml
   [[social]]
   identifier = "twitter"
   name = "Twitter"
   url = "https://twitter.com/your-handle"
   weight = 2
   [social.params]
   icon = "brand-twitter"
   ```

## Is the search index automatic?

Yes — `[outputs] home = ["HTML", "RSS", "JSON"]` produces `/index.json` containing every regular page. The `search` widget consumes that file. No external indexer is required.

## What's the license?

GPL-3.0-or-later. See [LICENSE](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/LICENSE) in the repository for the full text.

The theme is based on [Hugo Theme Stack](https://github.com/CaiJimmy/hugo-theme-stack) by Jimmy Cai (MIT-licensed).
