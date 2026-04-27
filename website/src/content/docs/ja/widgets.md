---
title: ウィジェット
description: ホームページと単一ページのサイドバーに表示できるウィジェット。
---

テーマは 2 種類のウィジェットスロットを持ちます。`[[params.widgets.homepage]]` はホームページのサイドバー、`[[params.widgets.page]]` は単一ページのサイドバーに表示されます。各ウィジェットは `type` と必要に応じた `params` テーブルを持つ TOML 配列エントリとして追加します。

## 利用可能なウィジェット

- `search`
- `archives`
- `categories`
- `tag-cloud`
- `toc`
- `twitter-share`
- `profile`

## `search`

サイトの `/index.json` 出力をデータソースにしたクライアントサイド全文検索ボックスを描画します。

```toml
[[params.widgets.homepage]]
type = "search"
```

ウィジェットは初回フォーカス時に `/index.json` を遅延読み込みします。`[outputs] home` 配列に `"JSON"` が含まれていることを確認してください。

## `archives`

日付でグループ化された最近の投稿を表示します。

```toml
[[params.widgets.homepage]]
type = "archives"
[params.widgets.homepage.params]
limit = 10
```

`limit` は表示するエントリ数。省略時は 5。

## `categories`

投稿数の多い順にカテゴリを表示します。

```toml
[[params.widgets.homepage]]
type = "categories"
[params.widgets.homepage.params]
limit = 10
```

## `tag-cloud`

頻出タグを使用頻度の重みづけ付きで表示します。

```toml
[[params.widgets.homepage]]
type = "tag-cloud"
[params.widgets.homepage.params]
limit = 10
```

## `toc`

ページ内見出しの目次を描画します。各ページのサイドバーで使うのが最適です。

```toml
[[params.widgets.page]]
type = "toc"
```

TOC のレベル範囲は `hugo.toml` の `[markup.tableOfContents]` から取得します。テーマは `startLevel = 2` / `endLevel = 4` を前提としています。

## `twitter-share`

記事横に「Twitter で共有」ボタンを描画します。各ページのサイドバーで使うのが最適です。

```toml
[[params.widgets.page]]
type = "twitter-share"
```

共有 URL はページタイトルと絶対 URL を自動で埋め込みます。

## `profile`

短い著者プロフィールカード。フロントマター / サイト `params` から読み込まれます。ホームページに配置するには:

```toml
[[params.widgets.homepage]]
type = "profile"
```

`[params.sidebar.avatar]` と組み合わせると、プロフィール文の上にアバターを描画できます。
