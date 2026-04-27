---
title: FAQ
description: よくある質問と注意点。
---

## favicon のリンクが 404 になります

テーマの `<head>` はサイト側の `static/` ディレクトリから `/apple-touch-icon.png`、`/favicon-32x32.png`、`/favicon-16x16.png`、`/favicon.ico`、`/site.webmanifest` を参照します。これらはテーマには **同梱していません** — サイトの `static/` に各自で配置してください。これらが無くてもサイト本体は正常に動作しますが、ブラウザのコンソールに 404 が出ます。

## メニューのアイコンが一部表示されません

`[[menu.main]]` と `[[social]]` の `params.icon` は `assets/icons/<name>.svg` に解決されます。テーマは汎用アイコンセット（`home`、`search`、`archives`、`rss`、`tag`、`folder`、`user`、`clock`、`date`、`external`、`link`、`list`、`moon`、`sun`、`language`、`copyright`、`clipboard`、`donate-heart`、`qr-code`）を同梱しています。

サービス固有のブランドアイコン（GitHub、X/Twitter、Speaker Deck、Zenn など）は同梱していません。サイト側の `assets/icons/<name>.svg` に SVG を置いてください。Hugo のアセットパイプラインはサイトの `assets/` をテーマよりも先に解決するため、サイト側のアイコンが優先され、テーマのアイコンセットを拡張する形になります。

`params.icon` がファイルに解決されない場合、アイコン枠は空で描画されますが、リンク自体は機能します。

## テーマに JavaScript は必要ですか？

テーマは「ほぼ JS なし」のパイプラインです。実際に JS を使うのは:

- ダークモード切替。
- 全文検索ウィジェット（`/index.json` を遅延読み込み）。
- コードブロックのコピーボタン。
- Mermaid ダイアグラム（使用しているページでのみ CDN から遅延読み込み）。

検索と Mermaid を使わないページは JS ゼロで描画されます。

## Mermaid はどこから読み込まれますか？

`layouts/_default/baseof.html` が `https://cdn.jsdelivr.net/npm/mermaid@10` から読み込みます。完全にセルフホストしたい場合は、サイト側のオーバーライドレイアウトで `<script>` タグを差し替えてください。

## Twitter / Zenn / Speaker Deck のアイコンを追加するには

1. **サイト側** の `assets/icons/<name>.svg` に SVG を保存します（テーマ側ではありません）。
2. `[[social]]` エントリから参照します:
   ```toml
   [[social]]
   identifier = "twitter"
   name = "Twitter"
   url = "https://twitter.com/your-handle"
   weight = 2
   [social.params]
   icon = "brand-twitter"
   ```

## 検索インデックスは自動で作られますか？

はい。`[outputs] home = ["HTML", "RSS", "JSON"]` を設定すると、すべての通常ページを含む `/index.json` が生成されます。`search` ウィジェットがこのファイルを読み込みます。外部のインデクサは不要です。

## ライセンスは？

MIT ライセンスです。完全な条文はリポジトリの [LICENSE](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/blob/main/LICENSE) を参照してください。

このテーマは Jimmy Cai 氏の [Hugo Theme Stack](https://github.com/CaiJimmy/hugo-theme-stack)（こちらも MIT ライセンス）をベースにしています。
