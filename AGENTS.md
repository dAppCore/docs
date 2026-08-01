# docs Agent Notes

This repository follows the core/go v0.9.0 consumer layout.

- Keep Go source under `go/`.
- Keep dependency source under `external/` submodules.
- Do not edit `.core/` runtime configuration.
- Markdown documentation lives under `docs/` (217+ files, YAML frontmatter).
- Site build orchestration lives in `zensical.toml` (Python toolchain).
- Run the v0.9.0 audit from the repository root before finishing.
