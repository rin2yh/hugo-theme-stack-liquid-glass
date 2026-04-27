---
title: "Configuration"
description: "Reference for params, menus, and social links."
date: 2026-04-27
weight: 20
categories: ["Docs"]
---

This page lists the configuration knobs the theme understands. All keys are optional unless marked otherwise.

## `[params]`

| Key | Type | Description |
|---|---|---|
| `mainSections` | `[]string` | Sections shown on the homepage feed. Default: `["post"]`. |
| `rssFullContent` | `bool` | If `true`, the RSS feed includes full post bodies. |
| `description` | `string` | Site description used in `<meta>` tags and as a fallback for RSS. |

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
