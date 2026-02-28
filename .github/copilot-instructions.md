# Copilot Repository Instructions

In this repository, follow these mandatory rules when generating any prompts or documentation files:

1. Always create or update prompt/document files only under `aigc/prompts/` (or its subfolders).
2. Never create prompt/document files in the repository root.
3. If a requested path is outside `aigc/prompts/`, automatically redirect it to `aigc/prompts/`.
4. Use lowercase kebab-case filenames, and use `.zh-CN.md` suffix for Chinese variants when needed.
5. After writing files, return the actual written paths.

## Open-source collaboration constraints

This repository is open source. When generating comments, documentation, or prompts, follow these rules:

1. Use neutral, discussable, and verifiable language; avoid absolute wording.
2. Base conclusions on currently visible repository code/docs, and state scope/assumptions when needed.
3. Never include sensitive information (keys, credentials, private addresses, personal data, internal system details).
4. Prefer reproducible steps and validation methods so contributors can verify results.
5. Keep notes understandable to external contributors without relying on internal context.
