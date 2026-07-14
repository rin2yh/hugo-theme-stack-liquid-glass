# Commit messages: Conventional Commits

All commits **should** follow the [Conventional Commits](https://www.conventionalcommits.org/) specification. Releases are automated with [tagpr](https://github.com/Songmu/tagpr), which lists merged PRs in the generated `CHANGELOG.md` — so keeping titles in this format is what keeps the changelog and git history readable. This is a convention, not CI-enforced.

Note: the **version bump** is chosen by labelling the tagpr release PR (`minor` / `major`, default patch), not parsed from commit types — see [`docs/releasing.md`](../../docs/releasing.md). The type/scope convention below is about changelog and history hygiene, not version selection.

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

Use one of these types:

| Type       | Typical use              |
| ---------- | ------------------------ |
| `feat`     | Features                 |
| `fix`      | Bug Fixes                |
| `perf`     | Performance Improvements |
| `revert`   | Reverts                  |
| `docs`     | Documentation            |
| `refactor` | Code Refactoring         |
| `ci`       | Continuous Integration   |
| `build`    | Build System             |
| `chore`    | Miscellaneous Chores     |
| `test`     | Tests                    |
| `style`    | Styles (formatting)      |

## Breaking changes

Signal a breaking change with a `!` after the type/scope, and/or a `BREAKING CHANGE:` footer, so it stands out in the history. tagpr does not derive the bump from this marker — apply the `major` label (or `minor` while the theme is pre-1.0) to the release PR to bump accordingly.

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
