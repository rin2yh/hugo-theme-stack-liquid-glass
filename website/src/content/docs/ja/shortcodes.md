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
