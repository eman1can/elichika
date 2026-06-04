# Package Overview

All packages live under `internal/`. The module name is `elichika`, so full import paths are `elichika/internal/<package>`.

To get a current list of packages:
```
go list ./...
```

For a conceptual overview of how packages relate to each other, see [Encapsulation Layers](encapsulation_layers.md).

---

## Entry Point

| Path | Description |
|---|---|
| `cmd/api-server/` | Main executable. Imports handler and subsystem packages to trigger `init()` registration, then starts the server. |

---

## Primitive Layer

| Package | Description |
|---|---|
| `internal/config` | Config and runtime config definitions |
| `internal/generic` | Generic utility types: arrays, dictionaries, `Nullable` wrapper, `UserIdWrapper`, drop mechanics, rankings |
| `internal/client` | Client types mirrored from the game |
| `internal/client/request` | Client types for network requests |
| `internal/client/response` | Client types for network responses |
| `internal/enum` | Enum values used by client and server (attributes, rarities, etc.) |
| `internal/item` | Common item and content definitions (English naming convention) |

---

## Database Layer

| Package | Description |
|---|---|
| `internal/db` | Generic database utilities shared across layers |
| `internal/clientdb` | Masterdata database keying and lookup |
| `internal/serverdata` | Server-level settings: active events, gacha banners, login bonuses, trades |
| `internal/gamedata` | Aggregated view of masterdata + serverdata; the primary interface handlers use to read game state |
| `internal/userdata` | Mixed subsystem and database code — user session, import/export, delete, authentication |
| `internal/userdata/database` | Userdata table definitions (one type per table: cards, events, live partners, tower, trade products, etc.) |

> **Note:** Dictionary functionality is not a standalone package. It is integrated into `gamedata`, `serverdata`, and `internal/generic` (`dictionary.go`).

---

## Asset Layer

| Package | Description |
|---|---|
| `internal/assetdata` | Asset metadata: textures, sounds, packs |
| `internal/asset_manager` | Asset management and caching: database, import/export, CDN checks |

---

## Support / Utility Layer

| Package | Description |
|---|---|
| `internal/encrypt` | Encryption and decryption utilities |
| `internal/utils` | Shared utility code |
| `internal/parser` | CSV and JSON parsing utilities |
| `internal/serverstate` | Server state management |
| `internal/locale` | Request locale handling |
| `internal/account` | Account export and import (with card stat reconstruction) |

---

## Subsystem Layer

| Package | Description |
|---|---|
| `internal/subsystem` | Master loader — blank-imports all 60 subsystem packages to trigger `init()` registration |
| `internal/subsystem/<name>` | Individual subsystem implementations (60 total — see below) |

### Non-User Subsystems

| Package | Description |
|---|---|
| `subsystem/banner` | Bootstrap banner responses |
| `subsystem/cache` | Caching system |
| `subsystem/event` | Event management: scheduling, active events, marathon, mining, story completion |
| `subsystem/pickup_info` | Event pickup information |
| `subsystem/reset_progress` | Progress reset functionality |
| `subsystem/time` | Server time management |
| `subsystem/voltage_ranking` | Voltage ranking system |

### User Subsystems

| Package | Description |
|---|---|
| `subsystem/user_accessory` | Accessory management |
| `subsystem/user_account` | Account management |
| `subsystem/user_authentication` | Authentication |
| `subsystem/user_beginner_challenge` | Beginner challenge system |
| `subsystem/user_bootstrap` | Bootstrap data generation |
| `subsystem/user_card` | Card inventory, level-up, drops, play counts |
| `subsystem/user_content` | Content viewing |
| `subsystem/user_custom_background` | Custom backgrounds |
| `subsystem/user_daily_theater` | Daily theater |
| `subsystem/user_emblem` | Emblem collection and equipping |
| `subsystem/user_event` | User event progress |
| `subsystem/user_expired_item` | Expired item handling |
| `subsystem/user_gacha` | Gacha draw mechanics |
| `subsystem/user_gps_present` | GPS-based presents |
| `subsystem/user_info_trigger` | Info trigger popups |
| `subsystem/user_lesson` | Lesson content |
| `subsystem/user_lesson_deck` | Lesson deck management |
| `subsystem/user_live` | Live performance (resume state) |
| `subsystem/user_live_deck` | Live deck composition |
| `subsystem/user_live_difficulty` | Live difficulty levels |
| `subsystem/user_live_mv` | Live MV |
| `subsystem/user_live_party` | Live party setup |
| `subsystem/user_login` | Login sequences |
| `subsystem/user_login_bonus` | Login bonuses |
| `subsystem/user_love_ranking` | Bond/love ranking |
| `subsystem/user_member` | Member profile management |
| `subsystem/user_member_guild` | Member guild functionality |
| `subsystem/user_mission` | Mission tracking and completion |
| `subsystem/user_new_badge` | New-content badge tracking |
| `subsystem/user_play_list` | Play history |
| `subsystem/user_present` | Present/gift management |
| `subsystem/user_profile` | User profile data |
| `subsystem/user_reference_book` | Reference book/codex |
| `subsystem/user_review_request_process_flow` | Review request flow |
| `subsystem/user_rule_description` | Game rule descriptions |
| `subsystem/user_scene_tips` | Scene tips and tutorials |
| `subsystem/user_sif_2_data_link` | SIF2 data linking |
| `subsystem/user_sifid_reward_mission` | SIFID reward missions |
| `subsystem/user_social` | Social features (friends) |
| `subsystem/user_status` | User status and stats |
| `subsystem/user_steady_voltage_ranking` | Steady voltage ranking |
| `subsystem/user_story_event_history` | Story event completion history |
| `subsystem/user_story_linkage` | Story linkage connections |
| `subsystem/user_story_main` | Main story progression |
| `subsystem/user_story_member` | Member story progression |
| `subsystem/user_story_side` | Side story progression |
| `subsystem/user_subscription_status` | Subscription management |
| `subsystem/user_suit` | Suit/costume management |
| `subsystem/user_tower` | Tower/floor challenge |
| `subsystem/user_trade` | Trading system |
| `subsystem/user_training_tree` | Training tree progression |
| `subsystem/user_tutorial` | Tutorial progression |
| `subsystem/user_unlock_scene` | Scene unlock tracking |
| `subsystem/user_voice` | Voice/audio preferences |

---

## Network Layer

| Package | Description |
|---|---|
| `internal/server` | Network system registration, routing (Gin), and WebUI endpoint definitions |
| `internal/webui` | WebUI system: admin, user, agnostic, and login endpoints |
| `internal/handler` | Master loader — blank-imports all 49 handler packages to trigger `init()` registration |
| `internal/handler/<name>` | Individual endpoint handlers (49 packages — named after their URL path segment) |
| `internal/handler/common` | Shared response utilities used across handlers |

> **Note:** There is no dedicated `middleware` package. Common request processing (authentication, session setup, etc.) is handled in `internal/server/router.go` and shared utilities in `internal/handler/common`.
