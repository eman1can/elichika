# Encapsulation Layers

The codebase is divided into layers of encapsulation. Each layer can access packages in the same layer or any earlier (lower) layer. The layers below are listed from closest to the database outward to the network.

Some older code predates this architecture and has not yet been refactored to conform — in particular, some game logic still lives in handler files rather than subsystems. When writing new features, follow the layer rules described here.

See [Package Overview](package.md) for the full current package listing.

---

## Primitive Layer

Contains definitions: types, configs, enums, and constants. No handling logic lives here, though field tags (`xorm`, `json`) and simple constructors/methods are allowed.

Key packages: `internal/config`, `internal/generic`, `internal/client`, `internal/client/request`, `internal/client/response`, `internal/enum`, `internal/item`.

---

## Database Layer

Handles initializing, reading, and writing the databases. There are several distinct databases:

### Userdata Database
Stores per-user progress (`internal/userdata/database`). Rules:
- Each table is created from exactly one type.
- Types should be `client` types plus relevant IDs (e.g., `user_id`), not mergers of multiple types.
- Arrays and maps may be split into separate tables and joined on read, or stored as JSON/blob if they are always read as a whole and are not large.

### Serverdata Database
Stores server-level settings: active exchanges, gacha banners, event schedules, login bonuses (`internal/serverdata`). Created from source code or initialization JSON files.

### Masterdata Database
The game's client database (`internal/clientdb`, `internal/gamedata`). Elichika does not modify it directly — modifications are made separately and applied as SQL scripts. See [Modifying the Database](../modify_database.md).

### Gamedata Database
The combined view of serverdata and masterdata (`internal/gamedata`). Represents the current game state. This is the primary database interface used by subsystems and handlers — they should not query masterdata or serverdata directly.

### Caching Database
Stores derived data: rankings, other-user profiles, etc. Rules:
- Must only contain data derived from the userdata database.
- Cached entries expire either on a time basis or when the relevant user makes a request.
- Currently not fully implemented; the policy is defined but not uniformly enforced.

---

## Asset Layer

Handles game assets (textures, sounds, packs) separately from the game databases (`internal/assetdata`, `internal/asset_manager`). Assets are served to clients via CDN or directly by the server.

---

## Subsystem Layer

Handles the game's subsystems — profile, gacha, missions, live, etc. All game logic should live here so it can be reused independently of the network layer.

There are 60 subsystems total. Each is its own package under `internal/subsystem/` and self-registers via `func init()`. The master `internal/subsystem/subsystem.go` blank-imports all of them to trigger registration.

The cascade model is central to how subsystems work: one action triggers others. For example:
- Finishing a live adds experience → may trigger a rank-up → rank-up sends a reward.
- Drawing gacha adds a card → may update bond limit → may unlock new bond board tiles.

Subsystems handle these chains internally, so the network layer only needs to invoke the top-level action.

---

## Network (Handler) Layer

Reads network requests, calls the subsystem layer to perform actions, and returns responses. There are 49 handler packages under `internal/handler/`, each corresponding to a URL path segment.

The handler layer should **not** contain game logic — it should identify what needs to happen and delegate to subsystems. For example:
- Clearing a live triggers bond level-up, user level-up, drops, etc.
- The handler only tells the subsystem "clear this live."
- The subsystem handles all consequences and returns what the handler needs to build the response.

**Registration:** Each handler file registers its endpoint via `server.AddHandler()` in `func init()`. The master `internal/handler/handler.go` blank-imports all handler packages to trigger registration. There is no dedicated middleware package — common request processing is handled in `internal/server/router.go` and shared utilities are in `internal/handler/common`.
