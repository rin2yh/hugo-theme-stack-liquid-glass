# Releasing

Notes for the maintainer. This project is not currently accepting external
contributions.

Releases are automated with [tagpr](https://github.com/Songmu/tagpr). tagpr
keeps a single open **release PR** up to date on every push to `main`, and
cutting a release is just merging that PR — no manual `git tag` or
`gh release create` step is needed.

## How versions are chosen

tagpr computes the next version from the **latest `vX.Y.Z` git tag**, not from
commit messages. By default it bumps the **patch** version. To bump minor or
major instead, label the open release PR:

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
on `main` (enforced on PR titles by `.github/workflows/pr-title-lint.yml`).
tagpr does not parse commit types for versioning, but the convention keeps the
generated `CHANGELOG.md` and PR history readable. See
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

### Tag format

Releases are tagged as plain `vX.Y.Z` (the standard Hugo theme / module
convention). The older release-please tags used a
`hugo-theme-stack-liquid-glass-v*` prefix; a matching `v0.5.0` tag was seeded
on the `0.5.0` release commit so tagpr continues the sequence from there.

## Configuration

tagpr is configured by [`.tagpr`](../.tagpr):

- `releaseBranch = main` — the branch tagpr watches.
- `vPrefix = true` — tag releases as `vX.Y.Z`.
- `versionFile = -` — the version lives only in git tags; there is no
  in-repo version file to bump (`theme.toml` has no version field).
- `release = true` — publish a GitHub Release on merge.
- `changelog = true`, `changelogFile = CHANGELOG.md` — maintain the changelog.

## Automation token

The workflow uses the default `GITHUB_TOKEN`. No workflow in this repo triggers
on tags or releases, and the release PR touches only `CHANGELOG.md` (no
`pull_request` path filters match it), so there is no need for a GitHub App
token to fan out downstream CI — unlike the previous release-please setup. The
default token also keeps the tagpr release PR (whose title is not a
Conventional Commit) from failing the `pr-title-lint` gate.

If you later add a workflow that must trigger on the release tag or the release
PR, swap `GITHUB_TOKEN` for a GitHub App token
([`actions/create-github-app-token`](https://github.com/actions/create-github-app-token)),
since events raised by the default token do not trigger further workflow runs.
