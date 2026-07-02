---
title: ショートコード
description: テーマが提供するカスタムショートコード。
---

Hugo 組み込みのショートコードに加え、テーマは以下を追加します。

## `qr`

任意のテキストの QR コードを描画します。画像は `https://api.qrserver.com/v1/create-qr-code/` から取得します。

### パラメータ

| 名前 | 必須 | 説明 |
|---|---|---|
| `text` | はい | QR コードにエンコードする文字列。 |
| `alt` | いいえ | 描画される `<img>` の alt 属性。省略時は `text` の値。 |
| `title` | いいえ | 画像の下に表示するキャプション。省略時は `text` の値。 |

### 使い方

```markdown
{{< qr text="https://example.com/" >}}

{{< qr text="https://example.com/" alt="example.com の QR" title="読み取って訪問" >}}
```

ショートコードは `figure.qr-card.glass` 要素内に描画されるため、QR 画像にもサイト全体と同じガラス質感のスタイルが適用されます。

### 補足

- 画像は遅延読み込み（`loading="lazy"`）されます。
- テキストは `urlquery` でエスケープされるため、クエリ文字列を含む URL も安全です。
- QR 生成を完全にセルフホストしたい場合は、サイト側の `layouts/shortcodes/qr.html` に上書き定義を置いてください。サイトレベルのショートコードはテーマよりも優先されます。

## `ogp`

URL の Open Graph プレビューカードを描画します。情報はビルド時に取得されます。

### パラメータ

| 名前 | 必須 | 説明 |
|---|---|---|
| `url` | はい | プレビューする URL。位置引数でも `url=` でも指定できます。 |

### 使い方

```markdown
{{< ogp "https://example.com/" >}}

{{< ogp url="https://example.com/" >}}
```

単独の行に貼った素の URL も自動的に同じカードになるため、このショートコードは本文中の任意の位置にカードを置きたい場合にのみ必要です。詳しい挙動・フォールバック・`params.ogpCard.enabled` の切り替えについては [リンクカード](/ja/link-cards/) を参照してください。
