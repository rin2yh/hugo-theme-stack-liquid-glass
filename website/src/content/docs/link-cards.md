---
title: Link cards
description: Turn bare URLs in content into Open Graph preview cards.
---

A bare URL on its own line in article content is rendered as a rich **link
preview card** instead of a plain link. The theme reads the target page's Open
Graph tags at build time (via Hugo's `resources.GetRemote`) and shows its
image, title, description, and site name.

This works out of the box — **link previews are always on by default**. There is
nothing to enable.

## Usage

Put a URL on its own line:

```markdown
https://gohugo.io/
```

Or insert one explicitly with the [`ogp` shortcode](/shortcodes/#ogp):

```markdown
{{< ogp "https://gohugo.io/" >}}
```

Both render the same card.

## What becomes a card

Only a **bare, standalone URL** (an autolink whose visible text equals the
destination) becomes a card. Everything else is left untouched:

- `[Hugo](https://gohugo.io/)` — a labelled link stays an ordinary link.
- Relative and internal links, anchors, and `mailto:` links are unchanged.

## Metadata and fallbacks

The card is built from the target's `<head>`:

| Field | Source | Fallback |
|---|---|---|
| Title | `og:title` | `<title>` |
| Description | `og:description` | `<meta name="description">` |
| Image | `og:image` | omitted if absent |
| Site name | `og:site_name` | the URL's host |

Relative `og:image` URLs (`/cover.png`, `//cdn/…`) are resolved against the
target. HTML entities in the title and description are decoded.

If the target exposes **no usable Open Graph data**, or is **unreachable at
build time**, the URL falls back to a plain link — a build is never broken by a
missing preview. Each unique URL is fetched once per build (results are cached).

## Disabling it

Link previews are on by default. To turn them off site-wide — falling back to
plain links everywhere — set [`params.ogpCard.enabled`](/configuration/#params)
to `false`:

```toml
[params.ogpCard]
enabled = false
```

## Notes

- Card images are loaded lazily (`loading="lazy"`) and are served from the
  target's own host — the browser fetches them at view time.
- Because previews are fetched during `hugo build`, the build host needs
  outbound network access to the URLs you link. In restricted CI, unreachable
  targets simply fall back to plain links.
- This is separate from [OG images](/og-images/), which control how *your own*
  pages preview when shared elsewhere.
