# Commit messages: Conventional Commits

All commits **must** follow the [Conventional Commits](https://www.conventionalcommits.org/) specification. Releases and the `CHANGELOG.md` are generated automatically by [release-please](https://github.com/googleapis/release-please) from these commit messages, so the format is not optional — non-conforming commits are silently dropped from the changelog and version bumps.

## Format

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

- Use the imperative mood in the description (e.g. "add", not "added" / "adds").
- Keep the description short; lowercase, no trailing period.
- `scope` is optional but encouraged (e.g. `share`, `mobile`, `ogp`, `sidebar`, `search`, `docs`).

## Allowed types

These match the `changelog-sections` in `release-please-config.json`:

| Type       | Changelog section        | Notes                                  |
| ---------- | ------------------------ | -------------------------------------- |
| `feat`     | Features                 | Triggers a minor version bump          |
| `fix`      | Bug Fixes                | Triggers a patch version bump          |
| `perf`     | Performance Improvements |                                        |
| `revert`   | Reverts                  |                                        |
| `docs`     | Documentation            |                                        |
| `refactor` | Code Refactoring         |                                        |
| `ci`       | Continuous Integration   | Hidden from changelog                  |
| `build`    | Build System             | Hidden from changelog                  |
| `chore`    | Miscellaneous Chores     | Hidden from changelog                  |
| `test`     | Tests                    | Hidden from changelog                  |
| `style`    | Styles                   | Hidden from changelog (formatting)     |

## Breaking changes

Signal a breaking change with a `!` after the type/scope, and/or a `BREAKING CHANGE:` footer. This triggers a major version bump (or a minor bump while the theme is pre-1.0, per `bump-minor-pre-major`).

```
feat(share)!: drop the legacy share-url param

BREAKING CHANGE: the `shareUrl` param has been removed; use `params.shareIcon` instead.
```

## Examples

```
feat(mobile): add floating TOC + share buttons for small screens
fix(sidebar): stop sticky TOC from overlapping stacked widgets
docs(ogp): document link cards on the docs site instead of the README
refactor(mobile): dedup share URL, reuse TOC title, drop wasteful blur
chore: bump dependencies
```
