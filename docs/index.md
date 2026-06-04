# Elichika

Note: This is a development fork of [arina999999997/elichika](https://github.com/arina999999997/elichika).

> Expect bugs, breaking features, and differing implementation details / defaults from the official fork.
> Refer to the official fork for playing the game!

Elichika is a private server implementation for *School Idol Festival ALL STARS* (SIFAS). It is a fork
of [YumeMichi/elichika](https://github.com/YumeMichi/elichika) with substantial additional features and fixes.

The easiest way to play is with the **embedded Android client**, which bundles the server directly inside the game APK —
no separate server setup needed. Alternatively, you can host your own server on Android (via Termux), PC, or Docker.

---

## I want to…

### Play the game

See the [Installing the Game](getting-started.md) guide for Android and iOS setup instructions.

### Host my own server

See the [Hosting Your Own Server](hosting.md) guide, or the [Docker Deployment](docker.md) guide if you prefer
containers.

### Use the WebUI

See the [WebUI Guide](webui.md) to manage your account, back up data, and configure the server.

### Back up or move my account

See [Import & Export](import_export.md) for exporting to `.db` or `.json` format and restoring on any compatible server.

### Understand how the server works

See the [Implementation Progress](progress.md) for a feature status overview, the [API Endpoints](endpoints.md)
reference, or the [Developer Docs](design/encapsulation_layers.md).

---

## Implementation Status

Most core gameplay is fully working. See the [Implementation Progress](progress.md) page for the full feature checklist.
Key highlights:

- **Working:** Live shows, story, training, bond system, DLP, gacha, daily theater, social/friends, shop, missions
- **Partial:** Events (marathon working; mining/SBL/voltage ranking not yet complete), membership tracking
- **Not planned:** Billing/IAP, co-op live, GPS presents

---

## Credits

Special thanks to the LL Hax community for archiving databases and assets, and specifically:

- **arina999999997** — significant development and management of the primary project fork
- **YumeMichi** — original elichika
- **triangle** — database encoding/decoding scripts, iOS client patches, daily theater logs, databases across all
  versions, missing asset files
- **gam** — missing asset files
- **[SIFAStheatre](https://twitter.com/SIFAStheatre)** and **[Idol Story](https://twitter.com/idoldotst)** — daily
  theater English translations and Japanese transcripts
- **ethan** — hosting resources and a public testing server, Docker help
- **[yunimoo](https://github.com/yunimoo)** — Docker help and resolving TODOs
- **rayfirefirst, cppo** — cryptographic keys
- **tungnotpunk** — iOS client and network structure
- **Suyooo** — the [SIFAS wiki](https://suyo.be/sifas/wiki/), accurate stage data, bad word lists
- **sarah** — public Internet CDN hosting
- **AuahDark** — embedded client development
- **Caret** — LL Hax Discord

---

## Disclaimer

This repository is designed for official SIFAS contents only. The authors do not endorse unofficial modifications to
assets, servers, or clients beyond what is already included in this repository. All such modifications are outside the
authors' control.
