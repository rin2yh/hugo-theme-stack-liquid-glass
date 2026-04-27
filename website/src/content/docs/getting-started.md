---
title: Getting Started
description: Install Stack Liquid Glass and configure a minimal Hugo site.
---

This page walks through installing the theme and writing the smallest possible `hugo.toml` to get a site running.

## Requirements

- Hugo **extended** `>= 0.146.0`. The theme uses Hugo Pipes (`resources.Concat`, `resources.Minify`, `resources.Fingerprint`) which require the extended build.

Verify your installation:

```bash
hugo version
# hugo v0.146.0+extended ...
```

## Install

### Option A — Git submodule (recommended)

From the root of your Hugo site:

```bash
git submodule add https://github.com/rin2yh/hugo-theme-stack-liquid-glass themes/stack-liquid-glass
```

### Option B — ZIP archive

Download the ZIP from GitHub and extract it into `themes/stack-liquid-glass/`.

## Enable the theme

In your site's `hugo.toml`:

```toml
theme = "stack-liquid-glass"
```

## Minimal `hugo.toml`

```toml
baseURL = "https://example.com/"
languageCode = "en"
title = "My Site"
theme = "stack-liquid-glass"
hasCJKLanguage = false

[outputs]
home = ["HTML", "RSS", "JSON"]

[permalinks]
post = "/post/:contentbasename/"

[markup.goldmark.renderer]
unsafe = true
[markup.tableOfContents]
ordered = true
startLevel = 2
endLevel = 4

[params]
mainSections = ["post"]
rssFullContent = true

[[params.widgets.homepage]]
type = "search"
[[params.widgets.homepage]]
type = "archives"
[[params.widgets.homepage]]
type = "categories"
[[params.widgets.homepage]]
type = "tag-cloud"

[[params.widgets.page]]
type = "toc"

[[menu.main]]
identifier = "home"
name = "Home"
url = "/"
weight = 1
[menu.main.params]
icon = "home"
```

`[outputs] home = ["HTML", "RSS", "JSON"]` is required — the JSON output powers the built-in full-text search.

## Required user-provided assets

`layouts/partials/head/head.html` references the following files **from your site's `static/` directory**. They are intentionally not bundled with the theme so that each site can use its own brand assets:

- `static/apple-touch-icon.png`
- `static/favicon-32x32.png`
- `static/favicon-16x16.png`
- `static/favicon.ico`
- `static/site.webmanifest`

If you do not provide these, the corresponding `<link>` tags in `<head>` will return 404 — the rest of the site still works.

## Run the dev server

```bash
hugo server
```

Visit http://localhost:1313/ to preview your site.

## Next

Read [Configuration](/configuration/) for the full reference of `[params]`, menus, and social links.
