# Specification workflow

Sauris uses OpenSpec 1.11.0 with the `spec-driven` schema and project-local Codex skills.

- `config.yaml` supplies project context and artifact rules.
- `specs/` contains architectural requirements and, as features land, their maintained behavior contracts. Requirements do not imply implementation is complete.
- `changes/` holds proposed changes and implementation tasks; completed changes are archived through OpenSpec.
- `.agents/skills/openspec-*` at the repository root contains generated Codex workflows. Update those with the OpenSpec CLI rather than hand-editing generated instructions.

## Start and complete a change

In Codex, use `$openspec-propose` with a concrete feature description. For manual artifact authoring:

```bash
openspec new change <change-name>
openspec status --change <change-name>
openspec instructions proposal --change <change-name>
```

Follow the schema to write the proposal, design, capability delta specs, and tasks. Use `$openspec-apply-change` for implementation and keep task status accurate. Update plans when the scope changes.

```bash
openspec validate --all --strict --no-interactive
```

Use `$openspec-archive-change` only when implementation and relevant verification are complete. Preserve outstanding tasks and unresolved decisions instead of marking a proposal complete prematurely.

## Bootstrap baseline

The `platform-foundation` specification captures architectural constraints agreed during project preparation. Runtime capabilities remain unimplemented. No feature change is marked as delivered by this scaffold.
