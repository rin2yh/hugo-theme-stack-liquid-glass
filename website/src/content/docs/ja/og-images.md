---
title: OGP 画像
description: テーマが OGP / Twitter Card 画像を選択・自動生成する仕組み。
---

テーマはすべてのページに Open Graph / Twitter Card の `<meta>` タグを出力し、画像が
指定されていない記事ページについては **ビルド時にタイトル画像を自動生成** します。これに
より、リンクのプレビューが常にブランドの効いた見た目になります。

## 画像の選択順

`og:image` / `twitter:image` は次の順で解決され、最初に見つかったものが使われます。

1. ページ front matter の `image`
2. ページ front matter の `cover.image`
3. **生成されるタイトル画像** — タイトルを持つ単一ページ（`.IsPage`）のみ
4. サイト設定の `params.defaultImage.opengraph.src`

明示的なページ画像が常に優先されるため、`image` または `cover.image` を設定した
ページは生成画像の対象外になります。

```yaml
---
title: 記事タイトル
image: cover.png        # og:image / twitter:image にそのまま使われる
# または:
cover:
  image: cover.png
---
```

## 生成されるタイトル画像

ページが手順 3 に到達すると、テーマは固定のテーマ配色背景（`assets/ogp/base.png`）に
ページタイトルを Hugo の `images.Text` フィルターと同梱フォント `NotoSansJP-Bold` で
合成し、その結果を `og:image` / `twitter:image` として出力します。日本語と英数字が
混在するタイトルも適切に折り返り、`PlantUML` のような ASCII 単語は途中で割れず、長い
タイトルは `…` で省略されます。

設定は不要で、そのまま動作します。一覧ページ・ホームページ・タクソノミーページは対象外で、
単一の記事ページだけが生成画像を持ちます。

## 背景を差し替える

同梱の `assets/ogp/base.png` はブランド中立なデフォルト（テーマ配色のグラデーション＋
グラスパネル＋アクセントバー）です。独自ブランドにするには、**サイト側** の
`assets/ogp/base.png` に `1200×630` の PNG を置いてください。Hugo の union ファイル
システムはサイトの `assets/` をテーマより先に解決するため、サイト側のファイルが
デフォルトを上書きします。タイトルが読みやすいよう、タイトル領域（左上のアクセントバーの
下・左下のブランディングの上）は空けておきます。

テーマにはデフォルト背景を生成した小さな Go 製ジェネレータ（`cmd/ogp`）も同梱しており、
ブランディングを付与できます。

```sh
# テーマのチェックアウト内で
cd cmd/ogp
go run . -site "My Blog"                        # サイト名を入れる
go run . -site "My Blog" -avatar /path/avatar.png  # アバターも入れる（png/jpg/webp）
```

詳細はテーマの [`assets/ogp/README.md`](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/assets/ogp/README.md)
を参照してください。再生成は任意で、多くのサイトでは `base.png` を直接差し替えるだけで
十分です。

## デプロイ（CI）

追加の CI ステップは不要です。画像は通常の `hugo` ビルド中に Hugo が生成するため、
別の生成ステップはなく、デプロイ時に Go ツールチェインも **不要** です（デフォルトの
`base.png` とフォントはテーマに同梱されています）。テーマをビルドする既存のワークフロー
（Hugo extended `>= 0.146.0`）であれば、OGP 画像は自動的に生成されます。

唯一の要件は、テーマのファイル（同梱の `assets/ogp/` を含む）がチェックアウトされて
いることです。git submodule で導入している場合は `actions/checkout` に
`submodules: true` を指定してください。Hugo Modules の場合はモジュール取得に含まれます。

```yaml
# .github/workflows/deploy.yml（抜粋）
- uses: actions/checkout@v4
  with:
    submodules: true          # テーマ本体と assets/ogp/ を取得
- uses: peaceiris/actions-hugo@v3
  with:
    hugo-version: '0.146.0'
    extended: true
- run: hugo --gc --minify      # ここで OGP 画像が生成される（追加ステップ不要）
```

## 関連設定

```toml
[params.defaultImage.opengraph]
src = "img/og-default.png"   # 他に画像がないときの最終フォールバック

[params.opengraph.twitter]
card = "summary_large_image" # デフォルト。生成される 1200×630 画像に最適
site = "your-handle"         # twitter:site として出力（"@" は任意）
```

| キー | 型 | 説明 |
|---|---|---|
| `params.defaultImage.opengraph.src` | `string` | サイト全体のフォールバック画像。生成画像の手順のあとにのみ使われる。`absURL` で解決。 |
| `params.opengraph.twitter.card` | `string` | `twitter:card` の種類。デフォルトは `summary_large_image`。 |
| `params.opengraph.twitter.site` | `string` | `twitter:site` のハンドル。先頭の `@` は無ければ付与される。 |

## フォントライセンス

同梱の `NotoSansJP-Bold.ttf` は Google Fonts の Noto Sans JP を `wght=700` の静的
インスタンスにしたもので、
[SIL Open Font License 1.1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/assets/ogp/NotoSansJP-LICENSE.txt)
のもとで提供されています。
