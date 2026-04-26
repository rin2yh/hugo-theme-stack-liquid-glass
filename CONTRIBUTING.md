# Contributing

Thanks for considering a contribution. This document covers the developer-facing
workflow for the theme: local setup, commit conventions, versioning policy, and
the release process.

## Local development

The repository ships an `exampleSite/` whose `themes/stack-liquid-glass` entry
is a symlink back to the repo root, so Hugo resolves the theme without extra
flags.

```bash
cd exampleSite
hugo server
```

CI builds `exampleSite/` against the documented minimum Hugo version on every
push and pull request — see `.github/workflows/ci.yml`.

## Commit conventions

This repository uses [Conventional Commits](https://www.conventionalcommits.org/).
The release tooling parses commit messages to determine the next version and
generate the changelog, so accurate types and scopes matter.

| Type        | Effect on version (pre-1.0)         | Notes                                     |
| ----------- | ----------------------------------- | ----------------------------------------- |
| `feat:`     | minor bump (e.g. `0.1.0` → `0.2.0`) | Listed under "Features" in the changelog. |
| `fix:`      | patch bump (e.g. `0.1.0` → `0.1.1`) | Listed under "Bug Fixes".                 |
| `perf:`     | patch bump                          | Listed under "Performance Improvements".  |
| `docs:`     | no version bump on its own          | Listed under "Documentation".             |
| `refactor:` | no version bump on its own          | Listed under "Code Refactoring".          |
| `ci:`, `build:`, `chore:`, `test:`, `style:` | no version bump | Hidden from the changelog. |

A breaking change is signalled by either a `!` after the type/scope or a
`BREAKING CHANGE:` footer:

```
feat(params)!: rename params.sidebar.avatar.local to params.sidebar.avatar.bundled

BREAKING CHANGE: sites that set params.sidebar.avatar.local must rename the key.
```

While the project is in `0.x`, breaking changes bump the **minor** version
(e.g. `0.2.0` → `0.3.0`) rather than jumping to `1.0.0`. After `1.0.0`,
breaking changes will bump the major version per standard SemVer.

## Versioning policy

The theme follows [Semantic Versioning](https://semver.org/). The "public API"
that the version number guarantees stability for is:

- **`params.*` keys** read by the theme's templates.
- **Layout files** that downstream sites are expected to override under
  `layouts/` (e.g. `partials/article/header.html`, `_default/baseof.html`).
- **i18n keys** under `i18n/*.toml`.

The following are explicitly **not** part of the public API and may change in
patch or minor releases without a breaking-change marker:

- CSS class names and design tokens (`assets/scss/`).
- JavaScript globals (e.g. `window.lgI18n`) and inline script internals.
- Internal partial helpers under `layouts/partials/helper/`.
- File names of bundled icons under `assets/icons/`.

Sites that depend on those internals should pin to a specific tag.

## Release process

Releases are fully automated by
[release-please](https://github.com/googleapis/release-please-action). The
`.github/workflows/release-please.yml` workflow runs on every push to `main`
and maintains an open "release PR" that accumulates pending changes.

The maintainer's role is to merge that PR when ready to cut a release:

1. Land normal changes via PRs into `main` using Conventional Commits.
2. release-please opens (or updates) a PR titled
   `chore(main): release <next-version>` containing:
   - the proposed `CHANGELOG.md` diff for the next version,
   - the bumped `.release-please-manifest.json`.
3. Review the changelog entries. Adjust commit messages on `main` if anything
   reads incorrectly (the release PR will refresh on the next push).
4. Merge the release PR. release-please then:
   - creates the git tag (`vX.Y.Z`),
   - publishes the GitHub Release with the changelog excerpt,
   - closes the release PR.

No manual tagging or `gh release create` is required.

### Forcing a specific version

To override the auto-computed version (e.g. graduate `0.x` → `1.0.0`), add a
`Release-As: 1.0.0` footer to a commit message on `main`, or set `release-as`
in `release-please-config.json` for a one-off bump. See the
[release-please docs](https://github.com/googleapis/release-please) for
details.

## Pull request checklist

- [ ] CI is green (`exampleSite/` builds against the documented minimum Hugo).
- [ ] Commit messages follow Conventional Commits.
- [ ] Breaking changes to `params`, override-target layouts, or i18n keys are
      marked with `!` or `BREAKING CHANGE:`.
- [ ] User-visible changes are described in a way that will read well in the
      changelog.
