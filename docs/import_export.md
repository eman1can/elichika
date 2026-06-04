# Import & Export

You can export and import your account using the [User WebUI](webui.md#user-webui). Two formats are supported: `.db` and `.json`.

---

## DB Format (Recommended)

The `.db` export is a SQLite database containing all data for a single account. Exporting and re-importing produces an identical account with two exceptions:

- **Friend data is not included.** Friends are server-scoped — they remain tied to the original server and do not transfer. Importing to a different server uses that server's friend list for your user ID.
- **Credential data is not included** for the same reason.

Use this format for backups and for migrating between compatible Elichika instances.

---

## JSON Format

The `.json` export is based on the login response from the server, which contains almost all account-relevant data. It is useful for:

- Recovering an account from captured network data (see [Extracting from PCAP](extracting_pcap.md))
- Importing accounts from older server versions
- Cross-server migration to non-Elichika implementations

### Automatic Backups

The server generates a JSON backup of your account on every login. Backups are stored in `elichika/backup/` on the server machine.

### Limitations

The login response does not contain every piece of account data. Where possible, missing data is reconstructed:

- Card practice data is reconstructed from card stats (the specific tile layout may differ from the original).
- Member stats (card count, training tree fill) are reconstructed.
- Usage statistics (times a card was used, skill activations) are **not** recoverable — this data is only present in profile captures for up to 6 cards.

These limitations have minimal impact on the core gameplay experience.
