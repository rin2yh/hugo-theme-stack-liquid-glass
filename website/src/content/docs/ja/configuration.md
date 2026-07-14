---
title: 設定
description: params、メニュー、ソーシャルリンクのリファレンス。
---

このページではテーマが解釈する設定キーを一覧にしています。明記がない限りすべて省略可能です。

## `[params]`

| キー | 型 | 説明 |
|---|---|---|
| `mainSections` | `[]string` | ホームページのフィードに表示するセクション。デフォルト `["post"]`。 |
| `rssFullContent` | `bool` | `true` のとき、RSS フィードに記事全文を含める。 |
| `description` | `string` | `<meta>` タグおよび RSS のフォールバックに使われるサイト説明。 |
| `ogpCard.enabled` | `bool` | 本文中の素の URL を [リンクカード](/ja/link-cards/) にするか。デフォルト `true`。`false` で通常のリンクとして描画。 |

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

### フォント

**テーマはデフォルトでウェブフォントを一切読み込みません。** すべて CSS 変数（`--font-display` / `--font-ui` / `--font-body` / `--font-mono`）のシステムフォントにフォールバックします。これは最速・最もプライバシーに配慮した選択肢で、PageSpeed Insights などが Google Fonts のスタイルシートを大きな「未使用の CSS」として指摘する問題も回避できます（特に CJK のスタイルシートは数百のサブセット `@font-face` 宣言を列挙します）。

フォントはサイト側で、次の 2 ステップで自由に指定します。

**1. フォントファイルを読み込む** — サイトに `layouts/partials/head/custom.html` を作成します。`<head>` の末尾付近で描画されるので、任意の `<link>` / `<style>` / preload をここに置きます:

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;700&display=swap">
```

（このフックは汎用で、preload や認証用 `<meta>` タグなどにも使えます。）

**2. 変数を指し向ける** — `[params.fonts]` で各キーが 1 つの CSS 変数を上書きします。CSS の編集は不要です:

```toml
[params.fonts]
display = "'Playfair Display', serif"        # 見出し / ブランド
ui      = "'Inter', system-ui, sans-serif"
body    = "'Inter', 'Noto Sans JP', sans-serif"
mono    = "'Fira Code', monospace"
```

| キー | 型 | 説明 |
|---|---|---|
| `display` | `string` | `--font-display`（見出し・ブランド・カードタイトル）を上書き。 |
| `ui` | `string` | `--font-ui`（ボタン・タグ・ラベル）を上書き。 |
| `body` | `string` | `--font-body`（本文）を上書き。 |
| `mono` | `string` | `--font-mono`（コード）を上書き。 |

#### 元のバンドルフォントに戻す

上記を行わずにテーマ本来の見た目（Italiana・DM Sans・Noto Sans JP・JetBrains Mono）が欲しい場合は、バンドル済みの Google Fonts リクエストを有効にします:

```toml
[params.fonts]
googleFonts = true       # テーマ本来のウェブフォントを有効化
japaneseBodyFont = true  # Noto Sans JP を含める。false で重い CJK フォントを除外
```

| キー | 型 | 説明 |
|---|---|---|
| `googleFonts` | `bool` | テーマ本来のウェブフォント一式を Google Fonts から読み込む（非同期・非ブロッキング）。デフォルト `false`。 |
| `japaneseBodyFont` | `bool` | `googleFonts = true` のとき、**Noto Sans JP** を含めるか。これはペイロードの大部分を占めます。日本語を使わないサイトは `false` で除外でき、本文は `--font-body` のシステム日本語フォントにフォールバックします。デフォルト `true`。 |

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
