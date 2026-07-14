# Releasing

Notes for the maintainer. This project is not currently accepting external
contributions.

Releases are automated with [tagpr](https://github.com/Songmu/tagpr). tagpr
keeps a single open **release PR** up to date on every push to `main`, and
cutting a release is just merging that PR — no manual `git tag` or
`gh release create` step is needed.

## How versions are chosen

tagpr tracks the current version in the [`VERSION`](../VERSION) file (kept in
sync with the latest `vX.Y.Z` git tag) and does **not** parse commit messages.
By default it bumps the **patch** version. To bump minor or major instead,
label the open release PR:

| Label on the release PR | Effect                                 |
| ----------------------- | -------------------------------------- |
| _(none)_                | patch bump (e.g. `0.5.0` → `0.5.1`)    |
| `minor`                 | minor bump (e.g. `0.5.0` → `0.6.0`)    |
| `major`                 | major bump (e.g. `0.5.0` → `1.0.0`)    |

While the project is in `0.x`, apply `minor` for new features or breaking
changes (there is no `0.x` "breaking = minor" auto-rule like release-please
had — you choose the bump via the label). After `1.0.0`, use `major` for
breaking changes per standard SemVer.

Because the bump is label-driven, **decide the version when you merge the
release PR**, not per commit. If you forget a label the release goes out as a
patch; re-label and the next push refreshes the proposed version.

## Commit conventions

We still follow [Conventional Commits](https://www.conventionalcommits.org/)
on `main` as a convention (it is no longer CI-enforced). tagpr does not parse
commit types for versioning, but keeping the format makes the generated
`CHANGELOG.md` and PR history readable. See
[`.claude/rules/conventional-commits.md`](../.claude/rules/conventional-commits.md).

## Cutting a release

The `.github/workflows/tagpr.yml` workflow runs on every push to `main` and
maintains the open release PR. To cut a release:

1. Land normal changes via PRs into `main`.
2. tagpr opens (or updates) a PR titled `Release for vX.Y.Z` containing the
   proposed `CHANGELOG.md` diff (a list of the PRs merged since the last tag).
3. Review the changelog. If the next version should be a minor or major bump,
   add the `minor` or `major` label to the PR; tagpr refreshes the proposed
   version on the next push.
4. Merge the release PR. tagpr then:
   - creates the git tag (`vX.Y.Z`),
   - publishes the GitHub Release with the changelog excerpt,
   - closes the release PR.

### Tag format and the first release

Releases are tagged as plain `vX.Y.Z` (the standard Hugo theme / module
convention). The older release-please tags used a
`hugo-theme-stack-liquid-glass-v*` prefix that tagpr cannot read, so there is
no `vX.Y.Z` tag for tagpr to start from. Rather than hand-seed one, the
[`VERSION`](../VERSION) file carries the current version (`0.5.1`), which tagpr
uses to propose the first release — so the first release PR is **`v0.5.1`** and
every release after it continues from the tag tagpr just cut. No manual tagging
is ever needed.

> **One-time caveat.** Because no `vX.Y.Z` tag exists yet, the *first* release
> PR's generated changelog spans the whole history (tagpr has no earlier tag to
> diff against). Just trim that first `CHANGELOG.md` entry down to the relevant
> changes in the release PR before merging; from the second release on, each
> changelog covers only the commits since the previous tag.

## Configuration

tagpr is configured by [`.tagpr`](../.tagpr):

- `releaseBranch = main` — the branch tagpr watches.
- `vPrefix = true` — tag releases as `vX.Y.Z`.
- `versionFile = VERSION` — the current version is stored in the `VERSION`
  file; tagpr reads it to pick the next version and bumps it in the release PR.
- `release = true` — publish a GitHub Release on merge.
- `changelog = true`, `changelogFile = CHANGELOG.md` — maintain the changelog.

## Automation token

The workflow uses the built-in `GITHUB_TOKEN` — no GitHub App or PAT needed.

One repo setting is required: **Settings → Actions → General → Workflow
permissions → "Allow GitHub Actions to create and approve pull requests"** must
be enabled, otherwise `GITHUB_TOKEN` is not allowed to open the release PR (a
`GitHub Actions is not permitted to create ... pull requests` error on the first
run). Most repos already have it on; if the first tagpr run fails to open a PR,
that toggle is why.

The old release-please GitHub App and its `RELEASE_PLEASE_APP_ID` /
`RELEASE_PLEASE_PRIVATE_KEY` secrets are no longer used and can be removed.
