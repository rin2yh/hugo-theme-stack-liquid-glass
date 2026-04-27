# Releasing

Notes for the maintainer. This project is not currently accepting external
contributions.

## Commit conventions

Releases and changelog entries are derived from commit messages by
[release-please](https://github.com/googleapis/release-please-action), so
every commit on `main` must follow
[Conventional Commits](https://www.conventionalcommits.org/).

| Type        | Effect on version (pre-1.0)         | Changelog placement       |
| ----------- | ----------------------------------- | ------------------------- |
| `feat:`     | minor bump (e.g. `0.1.0` → `0.2.0`) | Features                  |
| `fix:`      | patch bump (e.g. `0.1.0` → `0.1.1`) | Bug Fixes                 |
| `perf:`     | patch bump                          | Performance Improvements  |
| `revert:`   | follows the reverted commit's bump  | Reverts                   |
| `docs:`     | no version bump on its own          | Documentation             |
| `refactor:` | no version bump on its own          | Code Refactoring          |
| `ci:`, `build:`, `chore:`, `test:`, `style:` | no version bump | hidden       |

`exclude-paths` in `release-please-config.json` skips commits that touch
only `website/`, the Astro Starlight docs site shipped separately to
GitHub Pages — not part of the theme that consumers vendor, so its
changes should not bump the theme version. A commit that also touches
theme files still counts as a theme-affecting commit.

A breaking change is signalled by either a `!` after the type/scope or a
`BREAKING CHANGE:` footer:

```
feat(params)!: rename params.sidebar.avatar.local to params.sidebar.avatar.bundled

BREAKING CHANGE: sites that set params.sidebar.avatar.local must rename the key.
```

While the project is in `0.x`, breaking changes bump the **minor** version
(e.g. `0.2.0` → `0.3.0`) rather than jumping to `1.0.0`. After `1.0.0`,
breaking changes will bump the major version per standard SemVer.

## Cutting a release

The `.github/workflows/release-please.yml` workflow runs on every push to
`main` and maintains an open release PR. Cutting a release is just merging
that PR:

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

No manual `git tag` or `gh release create` step is needed.

### Forcing a specific version

To override the auto-computed version (e.g. graduate `0.x` → `1.0.0`), add a
`Release-As: 1.0.0` footer to a commit message on `main`, or set `release-as`
in `release-please-config.json` for a one-off bump. See the
[release-please docs](https://github.com/googleapis/release-please) for
details.

> **Historical note.** When the manifest was first bootstrapped at `0.0.0`,
> release-please treated the next release as a "graduation" and proposed
> `1.0.0`, even with `bump-minor-pre-major: true` (which only governs
> breaking-change behaviour _within_ `0.x`). The first release was pinned
> to `0.1.0` via a one-off `release-as: "0.1.0"` on the package, and that
> override was removed once `v0.1.0` shipped. Future bootstraps of similar
> projects can use the same workaround — and must remove `release-as`
> immediately after the override release, or every subsequent release-please
> PR will keep proposing the pinned version.

## Release automation setup (one-time)

The workflow runs as a GitHub App rather than with the default `GITHUB_TOKEN`,
because PRs and tags created by the default token do not trigger downstream
workflows — that would skip CI on the release PR. Setup steps:

1. Create a GitHub App at <https://github.com/settings/apps/new> with
   permissions: **Contents: Read & write**, **Pull requests: Read & write**.
   (Subscribe to no events; the App is only used as a token issuer.)
2. Install the App on this repository.
3. Add two repository secrets:
   - `RELEASE_PLEASE_APP_ID` — the App's Client ID (passed to the action's `client-id` input).
   - `RELEASE_PLEASE_PRIVATE_KEY` — the contents of the `.pem` private key.
4. Push to `main`. The first run will open the initial release PR.
