---
title: Shortcodes
description: Custom shortcodes the theme provides.
---

Beyond Hugo's built-in shortcodes, the theme adds the following:

## `qr`

Renders a QR code for arbitrary text. The image is fetched from `https://api.qrserver.com/v1/create-qr-code/`.

### Parameters

| Name | Required | Description |
|---|---|---|
| `text` | yes | The string encoded into the QR code. |
| `alt` | no | Alt text for the rendered `<img>`. Defaults to the value of `text`. |
| `title` | no | Caption text rendered below the image. Defaults to the value of `text`. |

### Usage

```markdown
{{< qr text="https://example.com/" >}}

{{< qr text="https://example.com/" alt="QR for example.com" title="Scan to visit" >}}
```

The shortcode renders inside a `figure.qr-card.glass` element so the QR image picks up the same glass surface treatment as the rest of the site.

### Notes

- The image is loaded lazily (`loading="lazy"`).
- The encoded text is `urlquery`-escaped, so URLs with query strings are safe.
- If you need to fully self-host the QR generation, override this shortcode in your site at `layouts/shortcodes/qr.html` — site-level shortcodes take precedence over theme ones.

## `ogp`

Renders an Open Graph preview card for a URL, fetched at build time.

### Parameters

| Name | Required | Description |
|---|---|---|
| `url` | yes | The URL to preview. May be passed positionally or as `url=`. |

### Usage

```markdown
{{< ogp "https://example.com/" >}}

{{< ogp url="https://example.com/" >}}
```

A bare URL on its own line renders the same card automatically, so the
shortcode is only needed when you want a card inline or in a specific spot. See
[Link cards](/link-cards/) for the full behavior, fallbacks, and the
`params.ogpCard.enabled` toggle.
