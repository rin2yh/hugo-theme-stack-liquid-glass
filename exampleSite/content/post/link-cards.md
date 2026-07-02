---
title: "Link Cards from URLs"
date: 2026-04-28T00:00:00+09:00
categories: ["link"]
tags: ["ogp"]
description: "Bare URLs in the article body turn into Open Graph preview cards."
---

Paste a URL on its own line and the theme fetches the page's Open Graph
metadata at build time, rendering a rich preview card instead of a plain link.

https://gohugo.io/

You can also use the shortcode explicitly:

{{< ogp "https://gohugo.io/" >}}

Inline links such as [the Hugo docs](https://gohugo.io/documentation/) keep
their normal styling — only a bare, standalone URL becomes a card. If the
target has no Open Graph data, or `params.ogpCard.enabled` is `false`, the URL
falls back to an ordinary link.
