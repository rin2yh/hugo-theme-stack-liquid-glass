# Changelog

## [v0.5.4](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/v0.5.3...v0.5.4) - 2026-07-15

- perf(ogp): eager-load the first card image and downsize twimg sources by @rin2yh in https://github.com/rin2yh/hugo-theme-stack-liquid-glass/pull/76

## [v0.5.3](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/v0.5.2...v0.5.3) - 2026-07-15

- perf: skip off-screen glass surfaces with content-visibility by @rin2yh in https://github.com/rin2yh/hugo-theme-stack-liquid-glass/pull/74

## [v0.5.2](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/v0.5.1...v0.5.2) - 2026-07-14

- feat(fonts)!: ship no webfonts by default, own them from the site side by @rin2yh in https://github.com/rin2yh/hugo-theme-stack-liquid-glass/pull/72

## [v0.5.1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.5.0...v0.5.1) - 2026-07-14

- ci(release): replace release-please with Songmu/tagpr by @rin2yh in https://github.com/rin2yh/hugo-theme-stack-liquid-glass/pull/70

## [0.5.0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.4.0...hugo-theme-stack-liquid-glass-v0.5.0) (2026-07-14)


### Features

* **ogp:** draw site name on generated OGP images ([#68](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/68)) ([d3bb5ec](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d3bb5ece3495ac28d300182b5208d97246616197))

## [0.4.0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.3.0...hugo-theme-stack-liquid-glass-v0.4.0) (2026-07-02)


### Features

* **share:** align sidebar share widget with the TOC ([#58](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/58)) ([39c71f2](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/39c71f2b2102f06d38e19e19ce0b6df1e4e35044))


### Bug Fixes

* **toc:** keep mobile TOC close button legible in dark theme ([#59](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/59)) ([04c8619](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/04c86196af9937df8007f521c6cbbb4ef89c76ab))


### Documentation

* add Conventional Commits rules in CLAUDE.md ([eecdb70](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/eecdb70fa6356c6d94efd155bc3ac0fa950b3294))
* add Conventional Commits rules in CLAUDE.md ([ca3c705](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/ca3c70515e7aaf8d6c51ab22fa04f10b894af8b9))
* move Conventional Commits rules to .claude/rules ([e8297f2](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/e8297f2dcd4e9c4b1dea35c5a931a85d5da78f68))


### Code Refactoring

* **share:** drop the Twitter-named i18n key left over from genericization ([e63ed5b](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/e63ed5b1214a4ae4b4471294e2a2f412281ae1bd))
* **share:** unblock release-please and cut the owed 0.4.0 ([8ad5c06](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/8ad5c0699b3dd6ae5ce01df6e55c4d2a3287be4a))

## [0.3.0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.2.0...hugo-theme-stack-liquid-glass-v0.3.0) (2026-07-02)


### Features

* **mobile:** add floating TOC + share buttons for small screens ([bc6d731](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/bc6d7314bf644fae79972d91b816cc88633637dc))
* **ogp:** render Open Graph link cards from URLs in content ([12ec11a](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/12ec11a82ecd69fa8189c351fcbc28191042d3a3))
* **share:** add copy-link button to the share widget ([d503912](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d503912431f6664e405898031b48d577ca500869))
* **share:** lay share buttons out horizontally; showcase in demo ([9d460e9](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/9d460e9211ce480099a466c764813eec36e3b32e))
* **share:** stack the title above a horizontal row of buttons ([4647085](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/46470852c7a1e1591e8b546f50c3b36cf389fa5e))


### Bug Fixes

* **mobile:** make floating action buttons more opaque for legibility ([a6ce8d1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/a6ce8d17b5e14ebc393ab15ef08120f671244bcb))
* **share:** use page permalink for copy-link after share-url refactor ([de3cfcb](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/de3cfcb86746a8e9ab29cfbb812384bdd3b47a9c))
* **sidebar:** stop sticky TOC from overlapping stacked widgets ([211c6e8](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/211c6e8bbb735cbccbdc49ba519955271c3cc968))


### Documentation

* **ogp:** document link cards on the docs site instead of the README ([fbe2e6c](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/fbe2e6cfe1e915aa820a7672157ee0bc462ba135))


### Code Refactoring

* **mobile:** dedup share URL, reuse TOC title, drop wasteful blur ([a8aa52f](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/a8aa52ffd751fb304ad613995d1c9695cc0bea10))
* **mobile:** make floating actions a speed-dial with distinct TOC icon ([e5311a7](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/e5311a7996a52735546e13db5a96a89c9d55c097))
* **ogp:** simplify link-card helpers per cleanup review ([1d2cdbc](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/1d2cdbc0cd9152c564b5bace7b1ab2e0cf2f1b2a))
* **share:** drop dead CSS flagged by cleanup review ([7dbacf2](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/7dbacf279fec82e5ae5157ef9362ad865aa7dbe2))

## [0.2.0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.1.1...hugo-theme-stack-liquid-glass-v0.2.0) (2026-07-02)


### Features

* **ogp:** 記事タイトル入り OGP 画像をテーマ側で自動生成 ([ecd071a](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/ecd071a65e9ec533e6e16f691faedc3cd8f57f8b)), closes [#48](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/48)
* **website:** migrate docs site from Hugo to Astro Starlight ([075db96](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/075db960164617c6d1307274acfe9f4cb1f6caf5))
* **website:** migrate docs site from Hugo to Astro Starlight ([962c2c6](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/962c2c6b6fde5021d4feb3023f7974257d1ccd18))


### Bug Fixes

* **ci:** pin actions/setup-node to v6.4.0 (v5.0.0 SHA was invalid) ([d0e4a63](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d0e4a63b17a15f746e408d77a9e3f35dff753771))
* **ci:** use Node 22 for website job (Astro 6 requires &gt;=22.12.0) ([6df7b58](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/6df7b58b13fc0f1e1ac0b242ca9249eaf9fb195c))
* **release-please:** exclude website/ so docs-only changes don't bump theme version ([297efad](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/297efad9eed89b4d12176a74561765b6a4f263be))
* **release-please:** exclude website/ so docs-only changes don't bump theme version ([aa9370d](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/aa9370df87c644b7ee508c7c8e53dc6a4fef9c72))
* **release-please:** track theme paths only via include-paths allowlist ([5e89c31](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/5e89c31c5baf359e0ed831c7cc1606532feb254a))


### Documentation

* **ogp:** rewrite OGP README in English and register JPEG decoder ([d7281f0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d7281f096e632220fd57bb78e830a1b70e612b2b))
* **releasing:** tighten website-exclusion paragraph ([ea3289f](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/ea3289f6ac9b56dee2428701dc20407ed30f26af))
* relicense from GPL-3.0-or-later to MIT ([dc5ad56](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/dc5ad56516c9cf81025920a194b5fe0fc0f2356f))
* restore full Japanese feature summary in README ([2563026](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/2563026635e2d3eb887c57ab041a6db1d65d7151))
* trim README to typical Hugo theme scope ([0181f0b](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/0181f0bf02850ee5d8c8274bb86c5e500963c2f5))
* trim README to typical Hugo theme scope ([0e88b59](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/0e88b599b7a0968b586a7be0e27bfe9c4a4b9892))
* **website:** document OGP image generation ([3dcac5f](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/3dcac5fccc1c428cbfd5dbf8a5b7962a7a97fd9f))
* **website:** note that OGP images need no extra consumer CI steps ([9f93ebe](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/9f93ebe8871cda413261f454b9948cf6a390cb15))


### Code Refactoring

* **ogp:** simplify generator and title-wrapping partial ([a8e3664](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/a8e3664cea71879316b81e1e594198259a9cd1e6))
* **website:** split CI per workflow, prune deps, fix base-prefix on hero links ([af23e22](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/af23e2253c555a10396396d23248a126a453620b))
* **website:** unify hero/LinkCard hrefs as relative paths ([7a10431](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/7a10431b684f615f264a28ef6704ec3353bbda4e))

## [0.1.1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/compare/hugo-theme-stack-liquid-glass-v0.1.0...hugo-theme-stack-liquid-glass-v0.1.1) (2026-04-27)


### Bug Fixes

* **website:** add missing archives page so docs archive widget links resolve ([2bf1133](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/2bf11339cbd89f26cd9b5f47e2ff157254f863e6))
* **website:** add missing archives page so docs site archive widget links resolve ([6d9c792](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/6d9c79287df3d141b8404027039a040b285b7dc0))

## 0.1.0 (2026-04-27)


### Features

* add mermaid render hook and exampleSite showcase posts ([7c31d27](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/7c31d27077b9831b821c1d2af7c8ceee2949548a))
* add render hook for mermaid code fences ([d43b874](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d43b874ec44bb85f47fdb65d21e80a8cafd2447a))


### Bug Fixes

* **article:** pass SVG covers through unchanged ([91fe019](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/91fe01923b74ecb1760921f2c4482ec1cdc8f6dc)), closes [#9](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/9)
* **article:** pass SVG covers through unchanged ([#9](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/9)) ([27224a0](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/27224a0b99e86fd8e6e34025cee5bab4b06cf49a))
* **docs,theme:** pin actions to SHAs, ubuntu-24.04, and fix multilang URL ([4b4b56c](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/4b4b56ced4405961c5874725dee7addca6979c2a))
* emit lgI18n as a single JSON object literal ([083c575](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/083c575c8150fbce1b4155a9c75395e2351ff703))
* emit lgI18n as a single JSON object literal ([eba9e69](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/eba9e69582417e5541f4f291b79ef1e4e2dff589)), closes [#1](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/issues/1)
* **exampleSite:** drop SVG cover from typography post ([c75afea](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/c75afeae3a10bc10870e3074399952f03436fb5a))
* **release-please:** pin first release to 0.1.0 ([4472380](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/44723803d9ebbdde4f44ee83702567ebd73a85ad))
* **release-please:** pin first release to 0.1.0 ([5d3c2c4](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/5d3c2c44bcf76b60767ac6e0fa8787b155cf9caa))
* **website:** add missing search page so docs site search returns results, not 404 ([63e1606](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/63e1606ce31e79c10932531f251a5bceda68bcfe))
* **widgets:** fall back to external icon when brand-twitter is absent ([d48e276](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d48e276b444f8db5a7771b6d7b0c5ecd1e18cedd))


### Performance Improvements

* cache color-scheme partial per language ([59ba628](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/59ba628ff24cefd6a2739528cbe1c324716edbf9))


### Documentation

* add README and GPL-3.0 LICENSE ([b1678cb](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/b1678cb16894cbf1d30457760fcf4b4d6146e16f))
* add theme screenshots ([95df68e](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/95df68e1f9fcad436ee0d98e53da68d487a51dc6))
* **exampleSite:** add showcase posts and theme-image guide ([d8290ae](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d8290ae0de76c43f0e8713348bbd896e961eaa5a))
* publish bilingual usage docs on GitHub Pages ([e9e6f25](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/e9e6f255e4fb8f5fb1a85d066f9fb09d6a7a7906))
* publish bilingual usage documentation on GitHub Pages ([6934ef5](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/6934ef5f0e57a9943d9054b9077aa1e9623b8a05))
* split CONTRIBUTING.md into Versioning + docs/releasing.md ([a5aa9c7](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/a5aa9c7a9593c96da23f66f7b13348fc3fe4ebe9))
* tighten release-please config and CONTRIBUTING ([3f50c0c](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/3f50c0cd1bd6589d634c5015de680c0973e09b87))
* trim duplicated content from workflow comment and releasing.md ([d4bdc77](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/d4bdc774470e40a90e6b77b8a6508a8371b5d8a3))


### Code Refactoring

* **docs:** rename docs/ → website/ to free up docs/ for plain Markdown ([7d32816](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/7d328160872e54f999704b94929b314abe9e967f))
* drop service-specific icons from theme ([cd4d759](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/cd4d75931883d0e0a3126ad2b306cf741b3476f4))
* drop service-specific icons from theme ([032338b](https://github.com/rin2yh/hugo-theme-stack-liquid-glass/commit/032338b46df721b9b67652c50490d6e068373467))
