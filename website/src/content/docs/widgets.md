---
title: Widgets
description: Sidebar widgets available on the homepage and on single pages.
---

The theme exposes two widget slots: `[[params.widgets.homepage]]` for the homepage sidebar, and `[[params.widgets.page]]` for single page sidebars. Each widget is added as a TOML array entry with a `type` and an optional `params` table.

## Available widgets

- `search`
- `archives`
- `categories`
- `tag-cloud`
- `toc`
- `twitter-share`
- `profile`

## `search`

Renders a client-side full-text search box backed by the site's `/index.json` output.

```toml
[[params.widgets.homepage]]
type = "search"
```

The widget loads `/index.json` lazily on first focus. Make sure your `[outputs] home` array includes `"JSON"`.

## `archives`

Lists recent posts grouped by date.

```toml
[[params.widgets.homepage]]
type = "archives"
[params.widgets.homepage.params]
limit = 10
```

`limit` controls how many entries are shown. Defaults to 5 if omitted.

## `categories`

Lists categories ranked by post count.

```toml
[[params.widgets.homepage]]
type = "categories"
[params.widgets.homepage.params]
limit = 10
```

## `tag-cloud`

Renders the most-used tags weighted by frequency.

```toml
[[params.widgets.homepage]]
type = "tag-cloud"
[params.widgets.homepage.params]
limit = 10
```

## `toc`

Renders the page's table of contents (the headings inside the post body). Most useful as a per-page widget.

```toml
[[params.widgets.page]]
type = "toc"
```

The TOC level range is taken from `[markup.tableOfContents]` in your `hugo.toml`. By default the theme expects `startLevel = 2` and `endLevel = 4`.

## `twitter-share`

A "Share on Twitter" button rendered next to the article. Best used as a per-page widget.

```toml
[[params.widgets.page]]
type = "twitter-share"
```

The share URL pre-fills the page title and absolute URL.

## `profile`

A short author profile card. Read from front matter / site `params`. Place it on the homepage:

```toml
[[params.widgets.homepage]]
type = "profile"
```

Combine with `[params.sidebar.avatar]` to render the avatar above the profile text.
