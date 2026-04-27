---
title: "設定"
description: "params、メニュー、ソーシャルリンクのリファレンス。"
date: 2026-04-27
weight: 20
categories: ["ドキュメント"]
---

このページではテーマが解釈する設定キーを一覧にしています。明記がない限りすべて省略可能です。

## `[params]`

| キー | 型 | 説明 |
|---|---|---|
| `mainSections` | `[]string` | ホームページのフィードに表示するセクション。デフォルト `["post"]`。 |
| `rssFullContent` | `bool` | `true` のとき、RSS フィードに記事全文を含める。 |
| `description` | `string` | `<meta>` タグおよび RSS のフォールバックに使われるサイト説明。 |

### `[params.dateFormat]`

```toml
[params.dateFormat]
published = "2006/01/02"
lastUpdated = "2006/01/02 15:04 MST"
```

どちらも Go の time フォーマット文字列を受け付けます。

### `[params.footer]`

```toml
[params.footer]
since = 2026
customText = ""
```

`since` はコピーライト行に表示する年。`customText` を設定するとコピーライトのあとに描画されます。

### `[params.sidebar.avatar]`

```toml
[params.sidebar.avatar]
enabled = true
local = true
src = "image/avatar.webp"
```

`enabled = true` でサイドバー先頭にアバターが表示されます。`local = true` の場合、`src` は Hugo のアセットパイプラインで解決されます（サイトの `assets/` 配下にファイルを置きます）。`local = false` の場合は外部 URL として扱われます。

## ウィジェット

各ウィジェットの設定は [ウィジェット](/ja/widgets/) を参照してください。ウィジェットスロットは 2 種類あります:

- `[[params.widgets.homepage]]` — ホームページのサイドバーに表示。
- `[[params.widgets.page]]` — 単一ページのサイドバーに表示。

各エントリは `type` と必要に応じた `params` を持ちます。

## `[[menu.main]]`

トップレベルのナビゲーション項目。

```toml
[[menu.main]]
identifier = "home"
name = "ホーム"
url = "/"
weight = 1
[menu.main.params]
icon = "home"
```

`params.icon` は `assets/icons/<name>.svg` に解決されます。テーマは汎用アイコンセット（`home`、`search`、`archives`、`rss`、`tag`、`folder`、`user`、`clock`、`date`、`external`、`link`、`list`、`moon`、`sun`、`language`、`copyright`、`clipboard`、`donate-heart`、`qr-code`）を同梱しています。追加の SVG はサイト側の `assets/icons/` に置けば利用できます。

## `[[social]]`

サイドバー下部に表示されるソーシャルリンク。

```toml
[[social]]
identifier = "github"
name = "GitHub"
url = "https://github.com/your-handle"
weight = 1
[social.params]
icon = "brand-github"
```

サービス固有のブランドアイコン（GitHub、X/Twitter、Speaker Deck、Zenn など）はテーマには **同梱していません**。利用したい SVG をサイト側の `assets/icons/<name>.svg` に置き、`params.icon` から参照してください。

`params.icon` がファイルに解決されない場合、アイコン枠は空で描画されますが、リンク自体は機能します。

## `[outputs]`

```toml
[outputs]
home = ["HTML", "RSS", "JSON"]
```

JSON 出力は検索ウィジェットのデータソースです。これがないと検索結果が出ません。

## `[markup]`

推奨設定:

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

`unsafe = true` は記事内に生 HTML（埋め込みなど）がある場合に必要です。`noClasses = false` にするとインラインの色指定ではなくテーマの CSS でコードブロックがスタイリングされます。

## 多言語サイト

テーマには `i18n/en.toml` と `i18n/ja.toml` が同梱されています。多言語化するには:

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

そのうえで、コンテンツの日本語版ファイルは `.ja.md` の拡張子で配置します。
