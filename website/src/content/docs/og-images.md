---
title: OG images
description: How the theme picks and auto-generates OGP / Twitter Card images.
---

The theme sets Open Graph and Twitter Card `<meta>` tags for every page and, for
content pages that have no image of their own, **generates a title image at build
time** so links always preview with something branded.

## How the image is chosen

`og:image` and `twitter:image` resolve through the following chain, stopping at
the first match:

1. `image` in the page front matter
2. `cover.image` in the page front matter
3. **A generated title image** — only for a single page (`.IsPage`) that has a title
4. `params.defaultImage.opengraph.src` from the site config

An explicit page image always wins, so setting `image` or `cover.image` opts a
page out of the generated one.

```yaml
---
title: My post
image: cover.png        # used verbatim for og:image / twitter:image
# or:
cover:
  image: cover.png
---
```

## The generated title image

When a page reaches step 3, the theme composites the page title onto a fixed
themed background (`assets/ogp/base.png`) using Hugo's `images.Text` filter and a
bundled `NotoSansJP-Bold` font, then emits the result as `og:image` /
`twitter:image`. Titles mixing Japanese and alphanumerics wrap sensibly, ASCII
words such as `PlantUML` are kept intact, and long titles are truncated with `…`.

No configuration is required — this works out of the box. List pages, the home
page, and taxonomy pages are skipped; only single content pages get a generated
image.

## Customizing the background

The bundled `assets/ogp/base.png` is a brand-neutral default (themed gradient +
glass panel + accent bar). To use your own branding, drop a `1200×630` PNG at
`assets/ogp/base.png` **in your own site** — Hugo's union filesystem resolves your
project's `assets/` before the theme's, so your file overrides the default. Keep
the title area clear (below the top-left accent bar, above the bottom-left
branding) so the overlaid title stays legible.

The theme also ships a small Go generator (`cmd/ogp`) that produced the default
background and can add branding:

```sh
# from a checkout of the theme
cd cmd/ogp
go run . -site "My Blog"                        # add a site name
go run . -site "My Blog" -avatar /path/avatar.png  # add an avatar too (png/jpg/webp)
```

See the theme's [`assets/ogp/README.md`](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/assets/ogp/README.md)
for details. Regenerating is optional — replacing `base.png` directly is enough
for most sites.

## Deployment (CI)

No extra CI steps are needed. The images are produced by Hugo during the normal
`hugo` build — there is no separate generation step, and the Go toolchain is
**not** required at deploy time (the default `base.png` and font ship with the
theme). Any workflow that already builds the theme (Hugo extended `>= 0.146.0`)
generates the OG images automatically.

The only requirement is that the theme's files — including its bundled
`assets/ogp/` — are checked out. For a git submodule install, set
`submodules: true` on `actions/checkout`; with Hugo Modules the module fetch
already includes them.

```yaml
# .github/workflows/deploy.yml (excerpt)
- uses: actions/checkout@v4
  with:
    submodules: true          # pulls the theme + its assets/ogp/ files
- uses: peaceiris/actions-hugo@v3
  with:
    hugo-version: '0.146.0'
    extended: true
- run: hugo --gc --minify      # OG images are generated here, no extra step
```

## Related config

```toml
[params.defaultImage.opengraph]
src = "img/og-default.png"   # final fallback when no other image resolves

[params.opengraph.twitter]
card = "summary_large_image" # default; the generated 1200×630 image suits this
site = "your-handle"         # rendered as twitter:site, "@" optional
```

| Key | Type | Description |
|---|---|---|
| `params.defaultImage.opengraph.src` | `string` | Site-wide fallback image, used only after the generated image step. Resolved with `absURL`. |
| `params.opengraph.twitter.card` | `string` | `twitter:card` type. Defaults to `summary_large_image`. |
| `params.opengraph.twitter.site` | `string` | Handle for `twitter:site`. A leading `@` is added if missing. |

## Font license

The bundled `NotoSansJP-Bold.ttf` is a `wght=700` static instance of Google
Fonts' Noto Sans JP, licensed under the
[SIL Open Font License 1.1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/assets/ogp/NotoSansJP-LICENSE.txt).
