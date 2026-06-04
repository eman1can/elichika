# The `assets/` Folder

The `assets/` directory is a **git submodule** (the "Harasho" repository) containing all static data the server needs: masterdata databases, SQL migrations, server configuration, event data, beatmaps, and daily theater scripts. Changes to game data are tracked independently from server code.

---

## Directory Overview

```
assets/
├── db/                  ← client masterdata and asset manifest databases
│   ├── gl/              ← Global (EN/KO/ZH) databases
│   └── jp/              ← Japanese databases
├── sql/                 ← SQL migrations applied on top of the raw databases
│   ├── gl/              ← GL-specific migrations
│   └── jp/              ← JP-specific migrations
├── server/              ← server initialization data (gacha, events, login bonuses)
│   ├── event/
│   ├── gacha/
│   └── login_bonus/
├── serverdata/          ← trade/shop definitions and word filter lists
├── event/               ← per-event assets (rewards, missions, board layouts)
│   ├── marathon/
│   └── mining/
├── daily_theater/       ← daily theater episode scripts (EN/JA/KO/ZH)
└── stages/              ← beatmap data (one JSON file per live difficulty)
```

---

## Database List

The `assets/db/` directory contains the game's **masterdata** — the read-only ruleset that describes every card, live, mission, gacha pool, etc. There are separate databases for each locale.

### Global (`assets/db/gl/`)

| File | Description |
|---|---|
| `masterdata.db` | Main GL masterdata (EOS state). The authoritative game rules database for GL. |
| `masterdata_a_en` | Encrypted GL masterdata for Android EN clients (generated from `masterdata.db`) |
| `masterdata_i_en` | Encrypted GL masterdata for iOS EN clients |
| `masterdata_a_ko`, `masterdata_i_ko` | Encrypted KO client databases |
| `masterdata_a_zh`, `masterdata_i_zh` | Encrypted ZH client databases |
| `asset_a_en.db` | Asset manifest for Android EN — maps asset paths to pack names, offsets, and encryption keys |
| `asset_i_en.db` | Asset manifest for iOS EN |
| `asset_a_ko.db`, `asset_i_ko.db` | KO asset manifests |
| `asset_a_zh.db`, `asset_i_zh.db` | ZH asset manifests |
| `auxinfo_a`, `auxinfo_i` | Auxiliary info files delivered to the Android/iOS client on startup |
| `dictionary_en_*.db` | Localised text strings for each platform/sub-type combination (android, ios, k, m, s, v, petag) |
| `dictionary_ko_*.db`, `dictionary_zh_*.db` | KO and ZH text dictionaries |

### Japanese (`assets/db/jp/`)

| File | Description |
|---|---|
| `masterdata.db` | Main JP masterdata (EOS state) |
| `masterdata_a_ja`, `masterdata_i_ja` | Encrypted JP client databases |
| `asset_a_ja.db`, `asset_i_ja.db` | JP asset manifests |
| `auxinfo_a`, `auxinfo_i` | JP auxiliary info files |
| `dictionary_ja_*.db` | JP text dictionaries (android, ios, k, m, s, v, petag) |

### Encrypted vs. Unencrypted

The `.db` files (SQLite) are the source of truth — the server reads these and modifies them. The files *without* a `.db` extension (e.g., `masterdata_a_en`) are encrypted copies that the server delivers to clients. **Never edit the encrypted files directly** — they are regenerated automatically by the server from the SQLite databases on startup (or when launched with `rebuild_assets`).

---

## SQL Migrations

### Purpose

Rather than shipping a pre-modified database binary, all changes to the EOS masterdata are expressed as numbered SQL scripts. This keeps the database history auditable, lets changes be applied incrementally, and makes it easy to re-apply modifications after a database update.

### Location

```
assets/sql/
├── {NNN}.masterdata.db.sql       ← locale-agnostic migrations (applied to both GL and JP)
├── gl/
│   ├── {NNN}.masterdata.db.sql   ← GL-only masterdata changes
│   ├── {NNN}.dictionary_en_k.db.sql
│   ├── {NNN}.asset_a_en.db.sql
│   └── ...                       ← one file per target database, per migration step
└── jp/
    ├── {NNN}.masterdata.db.sql
    ├── {NNN}.dictionary_ja_m.db.sql
    ├── {NNN}.asset_a_ja.db.sql
    └── ...
```

The `{NNN}` prefix is a zero-padded integer that controls execution order. Each file name after the number identifies the target database (e.g., `masterdata.db`, `asset_a_en.db`, `dictionary_en_k.db`).

### How They Are Applied

On startup, `internal/clientdb` discovers and applies migrations:

1. All `.sql` files under `assets/sql/{locale}/` and `assets/sql/` are collected and sorted by their numeric prefix.
2. The server checks whether the target database has changed since migrations were last applied (using git metadata). If the database is unmodified, migrations are skipped.
3. Each SQL file is executed line-by-line against its target database.

### Writing a Migration

Migrations are plain SQL — typically `INSERT OR REPLACE`, `UPDATE`, or `DELETE` statements against masterdata tables. Migrations are applied in numeric order, so pick the next available number for your locale:

```sql
-- assets/sql/gl/029.masterdata.db.sql
-- Example: add a missing live entry
INSERT OR REPLACE INTO m_live VALUES (42030, 0, 2030, 'music_2030', ...);
INSERT OR REPLACE INTO m_live_difficulty VALUES (...);
```

Locale-agnostic migrations (files directly under `assets/sql/`, not in a subdirectory) are applied to **both** GL and JP. Use a subdirectory if the change only applies to one locale.

> **Tip:** Save any manual changes you make to `masterdata.db` as a numbered SQL script so they survive database updates. After pulling a new version of the assets submodule, re-apply your scripts.

---

## Server Init Data

The files under `assets/server/` and `assets/serverdata/` seed the server's own `serverdata.db` on startup. Unlike masterdata, these define server-side behaviours that have no counterpart in the client database.

### Event Definitions (`assets/server/event/`)

| File | Description |
|---|---|
| `m_event.json` | Master list of events: `event_id`, `event_type` (1=marathon, 2=mining), `release_order`, `available`, optional `gacha_master_id` |
| `m_event_marathon.json` | Marathon-specific event metadata |
| `m_event_mining.json` | Mining-specific event metadata |

Only events listed here with `"available": true` can be selected or scheduled in the Admin WebUI.

### Gacha Definitions (`assets/server/gacha/`)

One JSON5 file per gacha banner, named by `gacha_master_id`:

```json5
{
  gacha_master_id: 2230615,
  gacha_type: 1,
  gacha_draw_type: 1,
  title: "Wonderful Rush",
  banner_image_asset: "...",
  gacha_draws: [
    { draw_count: 1, cost_type: 1, cost_amount: 150 },
    { draw_count: 10, cost_type: 1, cost_amount: 1500 }
  ],
  gacha_appeals: [
    { card_master_id: 1234, is_pickup: true }
  ]
}
```

To add a new gacha banner:
1. Create a new file named `{gacha_master_id}.json5` in `assets/server/gacha/`.
2. Add the banner to `m_event.json` if it is tied to an event, or reference it from a login bonus.
3. Restart the server (or run with `rebuild_assets`).

### Login Bonus Definitions (`assets/server/login_bonus/`)

One JSON5 file per login bonus sequence, named by `login_bonus_id`:

```json5
{
  login_bonus_id: 1000001,
  login_bonus_type: 1,
  background_id: 1,
  login_bonus_handler: "limited_login_bonus",
  rewards: [
    { day: 1, content_type: 9, content_id: 0, content_amount: 1 },
    { day: 2, content_type: 1, content_id: 0, content_amount: 50 }
  ]
}
```

### Trade/Shop Definitions (`assets/serverdata/`)

| File | Description |
|---|---|
| `trade.json` | Primary exchange shop products |
| `trade_2.json` | Secondary exchange shop products |
| `wordlist_gl.json` | GL profanity / bad word filter list |
| `wordlist_jp.json` | JP bad word filter list |

Each trade product specifies the source currency (`source_content_type`, `source_content_id`, `source_amount`), stock limits, reset schedule, and reward contents.

#### Content Type Reference

The `content_type` field used in trade products, login bonuses, and rewards maps to:

| Value | Content |
|---|---|
| 1 | Star gems (premium currency) |
| 4 | EXP |
| 9 | Gacha tickets |
| 10 | Gold |
| 16 | Card / member |
| 17 | Items (e.g., Show Candy) |
| 28 | Special items |
| 30 | Memory Keys |

---

## Event Assets

Each implemented event has a subdirectory under `assets/event/marathon/` or `assets/event/mining/`. The directory name is the `event_id`.

```
assets/event/marathon/30001/
├── main.json                ← event metadata
├── board.csv                ← marathon board layout
├── mission.csv              ← event missions
├── point_reward.csv         ← rewards at point thresholds
├── ranking_reward.csv       ← rewards by final ranking
├── ranking_topic_reward.csv
├── total_topic_reward.csv
└── bonus_popup_order.csv    ← character display order
```

**`main.json`** fields of note:
- `event_id`, `event_name` (multilingual object)
- `booster_item_id` — item that grants a score bonus
- `gacha_master_id` — event gacha banner, if any
- Image/BGM asset paths for the event UI

To add a new event, create the directory and populate all required files, then add the event to `assets/server/event/m_event.json` with `"available": true`.

---

## Other Asset Files

### Daily Theater (`assets/daily_theater/`)

2,490 episode scripts, one JSON file per episode per language, named `{locale}-{episode_id}.json`:

```json
{
  "language": "en",
  "daily_theater_id": 1000001,
  "year": 2020, "month": 2, "day": 3,
  "title": "Let's scatter beans!",
  "detail_text": "<:th_ch0201/>Hey, today is Setsubun..."
}
```

Supported locales: `en`, `ja`, `ko`, `zh`. Missing locales for an episode fall back gracefully.

### Beatmaps (`assets/stages/`)

2,824 beatmap files, one per live difficulty, named `{live_difficulty_id}.json`. Each file contains the full note chart (`live_notes`), wave/health-gate settings (`live_wave_settings`), and note/stage gimmicks. These are lazy-loaded by `internal/gamedata` on demand — the file for a given difficulty is only read the first time that difficulty is played.

---

## Forcing a Rebuild

On startup the server checks whether the databases are up to date and skips re-generating them if nothing has changed. To force a full rebuild of all databases and serverdata:

```bash
./elichika rebuild_assets
```

This is necessary after:
- Adding or modifying files under `assets/server/` or `assets/serverdata/`
- Pulling a new version of the assets submodule
- Manually editing a `.db` file in `assets/db/`
