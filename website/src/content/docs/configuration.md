---
title: Configuration
description: Reference for params, menus, and social links.
---

This page lists the configuration knobs the theme understands. All keys are optional unless marked otherwise.

## `[params]`

| Key | Type | Description |
|---|---|---|
| `mainSections` | `[]string` | Sections shown on the homepage feed. Default: `["post"]`. |
| `rssFullContent` | `bool` | If `true`, the RSS feed includes full post bodies. |
| `mobileBackButton` | `bool` | Whether the mobile header shows a back button left of the site name on non-home pages. Default: `true`; set to `false` to hide it. |
| `description` | `string` | Site description used in `<meta>` tags and as a fallback for RSS. |
| `ogpCard.enabled` | `bool` | Whether bare URLs in content become [link cards](/link-cards/). Default: `true`; set to `false` to render plain links instead. |

### `[params.dateFormat]`

```toml
[params.dateFormat]
published = "2006/01/02"
lastUpdated = "2006/01/02 15:04 MST"
```

Both keys accept Go time-format layout strings.

### `[params.footer]`

```toml
[params.footer]
since = 2026
customText = ""
```

`since` is the year shown in the copyright line. `customText` is rendered after the copyright if set.

### `[params.sidebar.avatar]`

```toml
[params.sidebar.avatar]
enabled = true
local = true
src = "image/avatar.webp"
```

When `enabled = true`, the avatar appears at the top of the sidebar. With `local = true`, `src` is resolved through Hugo's asset pipeline (place the file under your site's `assets/` directory). With `local = false`, `src` is treated as an external URL.

### Fonts

**The theme loads no webfonts by default.** Everything falls back to the system-font stacks in the CSS variables (`--font-display`, `--font-ui`, `--font-body`, `--font-mono`), which is the fastest and most private option — and it stops tools like PageSpeed Insights from flagging a large "unused CSS" Google Fonts stylesheet (the CJK stylesheet in particular enumerates hundreds of subset `@font-face` rules).

You choose the fonts from your own site, in two steps.

**1. Load the font files** — create `layouts/partials/head/custom.html` in your site. It is rendered near the end of `<head>`, so put any `<link>`, `<style>`, or preload there:

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;700&display=swap">
```

(This hook is general-purpose — you can also use it for preloads, verification `<meta>` tags, etc.)

**2. Point the variables at them** — with `[params.fonts]`. Each key overrides one CSS variable, no CSS editing required:

```toml
[params.fonts]
display = "'Playfair Display', serif"        # headings / brand
ui      = "'Inter', system-ui, sans-serif"
body    = "'Inter', 'Noto Sans JP', sans-serif"
mono    = "'Fira Code', monospace"
```

| Key | Type | Description |
|---|---|---|
| `display` | `string` | Overrides `--font-display` (headings, brand, card titles). |
| `ui` | `string` | Overrides `--font-ui` (buttons, tags, labels). |
| `body` | `string` | Overrides `--font-body` (body copy). |
| `mono` | `string` | Overrides `--font-mono` (code). |

#### The theme's original typography

The CSS variables already name the theme's designed families first (Italiana, DM Sans, Noto Sans JP, JetBrains Mono), so to get the original look you only need to **load** those fonts — no variable overrides required. Drop this into `layouts/partials/head/custom.html`:

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="stylesheet" media="print" onload="this.media='all'"
  href="https://fonts.googleapis.com/css2?family=Italiana&family=DM+Sans:wght@400;500&family=Noto+Sans+JP:wght@400;500;700&family=JetBrains+Mono:wght@400;600&display=swap">
<noscript><link rel="stylesheet"
  href="https://fonts.googleapis.com/css2?family=Italiana&family=DM+Sans:wght@400;500&family=Noto+Sans+JP:wght@400;500;700&family=JetBrains+Mono:wght@400;600&display=swap"></noscript>
```

The `media="print"` → `onload` swap loads it asynchronously so it never blocks first paint. Non-Japanese sites can drop `&family=Noto+Sans+JP:…` to skip the heavy CJK font (body text then falls back to the system Japanese fonts in `--font-body`). This is exactly what the [exampleSite](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/exampleSite/layouts/partials/head/custom.html) does.

## Widgets

See [Widgets](/widgets/) for the per-widget configuration. There are two widget slots:

- `[[params.widgets.homepage]]` — appears on the homepage sidebar.
- `[[params.widgets.page]]` — appears on single page sidebars.

Each entry has a `type` and optional `params`.

## `[[menu.main]]`

Top-level navigation items.

```toml
[[menu.main]]
identifier = "home"
name = "Home"
url = "/"
weight = 1
[menu.main.params]
icon = "home"
```

`params.icon` resolves to `assets/icons/<name>.svg`. The theme ships a generic icon set (`home`, `search`, `archives`, `rss`, `tag`, `folder`, `user`, `clock`, `date`, `external`, `link`, `list`, `moon`, `sun`, `language`, `copyright`, `clipboard`, `donate-heart`, `qr-code`). Place additional SVGs under your site's `assets/icons/` to extend it.

## `[[social]]`

Bottom-of-sidebar social links.

```toml
[[social]]
identifier = "github"
name = "GitHub"
url = "https://github.com/your-handle"
weight = 1
[social.params]
icon = "brand-github"
```

Service-specific brand icons (GitHub, X/Twitter, Speaker Deck, Zenn, etc.) are intentionally **not** bundled with the theme. Place the SVG you want to use at `assets/icons/<name>.svg` in your own site and reference it from `params.icon`.

If `params.icon` resolves to no file, the icon slot renders empty but the link itself still works.

## `[outputs]`

```toml
[outputs]
home = ["HTML", "RSS", "JSON"]
```

The JSON output is what the search widget consumes. Without it, search returns no results.

## `[markup]`

Recommended markup settings:

```toml
[markup.goldmark.renderer]
unsafe = true

[markup.tableOfContents]
ordered = true
startLevel = 2
endLevel = 4

[markup.highlight]
noClasses = false
codeFences = true
guessSyntax = true
tabWidth = 2
```

`unsafe = true` is required if your posts contain raw HTML (e.g. embeds). `noClasses = false` lets the theme's CSS style code blocks instead of inlining colors.

## Multilingual sites

The theme ships with `i18n/en.toml` and `i18n/ja.toml`. To enable multiple languages:

```toml
defaultContentLanguage = "en"
hasCJKLanguage = true

[languages]
  [languages.en]
    languageName = "English"
    weight = 1
  [languages.ja]
    languageName = "日本語"
    weight = 2
```

Then suffix content files with `.ja.md` for Japanese variants of `.md` content.
