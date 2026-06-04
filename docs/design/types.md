# Type System

The server manages many kinds of types: types for client communication, types for storing user data, and types for internal handling. These are divided into categories with clear rules about how they may depend on each other. Types in a given category may only use types from the same category or a simpler (lower) one.

---

## Client Types

Client types mirror the types used by the game client. They are defined in the `client` package and used to read and write network data.

**Rules:**

- The type name must match the client's name exactly, including capitalisation of every letter except the first (Go requires the first letter to be uppercase for exported types).
- Field types must match precisely: `int32` must be `int32`, `int64` must be `int64`, etc.
- Do not use anonymous types for subfields — replicate the type and name it.
- Use the `Nullable` generic for fields that can be `null` in JSON.
- Use the `Dictionary` generic for `Dictionary` fields.
- Fields with an `enum` type require an `enum` struct tag naming the enum, even though it is not enforced at runtime yet.
- The type must round-trip through `json.Marshal` / `json.Unmarshal` without loss. Write a custom marshaler/unmarshaler if necessary.
- Field order is not important, but array element order is — preserve it.
- Client types must not be modified to assist handling logic. If you need to attach extra data, use an embedded type or wrapper.
- Each type lives in its own file. File names use `snake_case` derived from the type name.

### Request / Response Types

Request and response types are client types and follow the same rules, but they go into `client/request` and `client/response` respectively. Any subtype used exclusively by a request or response type still belongs in `client`, not in `client/request` or `client/response`.

---

## Gamedata Types

Gamedata types represent server-side game state: which event is active, which gacha banners are available, etc. They are less strictly constrained than client types, but should:

- Follow the naming and field conventions of the corresponding client types.
- Be loadable from and savable to the database.

Gamedata types are defined in the `gamedata` package alongside their loaders.

---

## Userdata Types

Userdata types store per-user progress. They should be built on top of `client` types:

- The general pattern is a `UserIdWrapper` around the relevant `User*` client type.
- If no matching client type exists, follow the naming used by the parts of the codebase that use the data.
- One table per type — do not merge unrelated data into the same table even if it is technically possible.
- Do not store derived or aggregated data unless it directly mirrors a `client` type. Calculate aggregates from userdata at read time and cache them separately.

Userdata types are defined in `userdata/database`.

---

## Handling Types

Types defined and used within handler packages. These can be almost anything, but should follow the same naming conventions as the rest of the codebase. Types used across multiple handlers should live in a shared utilities package rather than being copied.
