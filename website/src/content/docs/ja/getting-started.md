---
title: はじめに
description: Stack Liquid Glass をインストールし、最小構成の Hugo サイトを動かすまで。
---

このページではテーマのインストール手順と、サイトを動かすための最小限の `hugo.toml` を解説します。

## 必要環境

- Hugo **extended** `>= 0.146.0`。テーマは Hugo Pipes（`resources.Concat`、`resources.Minify`、`resources.Fingerprint`）を使用するため、extended 版が必要です。

バージョン確認:

```bash
hugo version
# hugo v0.146.0+extended ...
```

## インストール

### 方法 A — Git submodule（推奨）

Hugo サイトのルートで:

```bash
git submodule add https://github.com/rin2yh/hugo-theme-stack-liquid-glass themes/stack-liquid-glass
```

### 方法 B — ZIP アーカイブ

GitHub から ZIP をダウンロードし、`themes/stack-liquid-glass/` に展開します。

## テーマの有効化

サイトの `hugo.toml` で:

```toml
theme = "stack-liquid-glass"
```

## 最小構成の `hugo.toml`

```toml
baseURL = "https://example.com/"
languageCode = "ja"
title = "My Site"
theme = "stack-liquid-glass"
hasCJKLanguage = true

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
name = "ホーム"
url = "/"
weight = 1
[menu.main.params]
icon = "home"
```

`[outputs] home = ["HTML", "RSS", "JSON"]` は必須です。JSON 出力が組み込み全文検索のデータソースになります。

## サイト側で用意する必要があるアセット

`layouts/partials/head/head.html` は以下のファイルを **サイト側の `static/` ディレクトリ** から参照します。各サイトが独自のブランドアセットを使えるようにするため、テーマには同梱していません:

- `static/apple-touch-icon.png`
- `static/favicon-32x32.png`
- `static/favicon-16x16.png`
- `static/favicon.ico`
- `static/site.webmanifest`

これらを置かなくても他の機能は動作しますが、`<head>` 内の対応する `<link>` タグが 404 になります。

## 開発サーバーの起動

```bash
hugo server
```

http://localhost:1313/ にアクセスしてプレビューを確認します。

## 次に読む

[設定](/ja/configuration/) で `[params]`、メニュー、ソーシャルリンクの完全リファレンスを参照してください。
