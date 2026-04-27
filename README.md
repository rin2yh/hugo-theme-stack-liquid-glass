# Stack Liquid Glass

A glassmorphism-flavored Hugo theme — a liquid-glass redesign of [Hugo Theme Stack](https://github.com/CaiJimmy/hugo-theme-stack), with a near-zero JavaScript pipeline, light/dark mode, and built-in JSON search.

> 日本語: Hugo Theme Stack をベースに、ガラス質感（glassmorphism）の UI とほぼ JS なしの配信パイプラインに作り変えたテーマです。ライト/ダーク切替、`/index.json` ベースの全文検索、TOC、Mermaid、外部 URL 投稿、日本語 i18n に対応しています。

![Stack Liquid Glass](images/screenshot.png)

## Documentation

Full usage documentation (English / 日本語) is published at:

<https://rin2yh.github.io/hugo-theme-stack-liquid-glass/>

## Features

- Glassmorphism / liquid-glass design
- Light & dark mode (system preference aware)
- Near-zero JS pipeline via Hugo Pipes
- `/index.json` powered full-text search
- TOC, Mermaid, related posts, external URL posts
- Japanese i18n out of the box

## Requirements

- Hugo **extended** `>= 0.146.0`

## Installation

As a git submodule:

```bash
git submodule add https://github.com/rin2yh/hugo-theme-stack-liquid-glass themes/stack-liquid-glass
```

Then enable the theme in your site's `hugo.toml`:

```toml
theme = "stack-liquid-glass"
```

See the [documentation](https://rin2yh.github.io/hugo-theme-stack-liquid-glass/) for configuration, widgets, shortcodes, and required static assets.

## Credits

- Based on [Hugo Theme Stack](https://github.com/CaiJimmy/hugo-theme-stack) by Jimmy Cai (MIT).
- Icons inspired by [Tabler Icons](https://tabler.io/icons) and [Feather Icons](https://feathericons.com/) (MIT).

## License

GPL-3.0-or-later.

- Copyright (C) 2026 rin2yh
- Copyright (C) 2020 Jimmy Cai (original Hugo Theme Stack)

See [LICENSE](LICENSE) for details.
