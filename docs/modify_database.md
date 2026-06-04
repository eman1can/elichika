# Modifying the Database

> **Advanced usage.** See [Advanced Usage](advanced_usage.md) for what is expected before attempting this.

---

## Why Modify the Database?

The server ships with databases as they were at EOS, with minor modifications to enable server features. You can edit these databases to customize the game in ways that go beyond what the WebUI offers, for example:

- Include all songs in the daily rotation instead of the usual 3 per day
- Raise the skip ticket limit above 20
- Add JP-exclusive content to the Global version (or vice versa)
- Perform model swaps or other visual changes

---

## How It Works

The server reads its masterdata database from:

- `elichika/assets/db/jp` — Japanese version
- `elichika/assets/db/gl` — Global version

The client downloads this database from the server. However, the client expects an encrypted database, not raw SQLite, so the server automatically encrypts it before sending. You only need to edit the unencrypted SQLite files — the server handles encryption on the fly.

---

## Editing the Database

You can modify the SQLite database using any SQLite-compatible tool:

- **[DB Browser for SQLite](https://sqlitebrowser.org/)** — visual GUI, good for browsing and manual edits
- **SQL scripts** — repeatable and version-controllable; can be run via DB Browser or any SQLite tool
- **File replacement** — replace the database files entirely with files obtained elsewhere

After making changes, restart the server. It will regenerate the encrypted files automatically. Log in or move around in the client to trigger a database update on the client side.

---

## Important Notes

- **Back up the database files before making any changes.**
- This only modifies game data. Modifying or adding assets requires additional encryption steps (not documented here).
- If your database changes are inconsistent with the client's expectations, either the client or server may behave incorrectly.
- **Save modifications as SQL scripts.** Server updates will reset the database, so you will need to reapply your changes after updating — scripts make this repeatable.
- **Update first, then apply modifications.** Applying your changes before updating will likely cause conflicts.
