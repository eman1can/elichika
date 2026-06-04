# Subsystem

Subsystems handle a specific part of the game and the cascading consequences of actions within it.

For example:
- Finishing a live adds experience → may trigger a rank-up → rank-up sends a reward.
- Drawing gacha adds a card → may raise bond limit → may unlock new bond board tiles.

The network handler only needs to invoke the top-level action. The subsystem handles all downstream effects and returns what the handler needs to build its response.

## Registration

Each subsystem is its own package under `internal/subsystem/`. Subsystems self-register via `func init()`. This file (`subsystem.go`) blank-imports all subsystem packages to trigger those registrations at startup.

## Adding a New Subsystem

1. Create a new package under `internal/subsystem/<name>`.
2. Implement the subsystem logic.
3. Add a blank import for the new package in `subsystem.go`.

## Subsystem List

See [docs/design/package.md](../../docs/design/package.md) for the full listing of all 60 subsystems and their descriptions.
