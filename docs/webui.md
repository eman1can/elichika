# WebUI Guide

The WebUI lets you interact with the server directly — outside of the game client — to manage your account and configure server behaviour.

To open the WebUI, navigate to the relevant address in any web browser.

---

## User WebUI

**Address:** `http://<server_address>:8080/webui/user`

For the embedded server, the default address is `http://127.0.0.1:8080/webui/user`. Keep the game client open and use your device's browser to access it.

### Logging In

Enter your **User ID** (visible on the title screen or in the profile menu) and your **password** (leave blank if you have not set one).

You must have completed the tutorial before the WebUI will allow access. Tutorial-skip is fine — just reach the main menu. This requirement exists to prevent account spam on multi-user servers.

### Locale

Select the correct client language before using the WebUI. Some content only exists for a specific locale (Japanese or Global), and using the wrong locale can make your account unplayable. The locale setting is primarily a server-side signal.

### Features

#### Account Builder
Automates common progression tasks. Includes a one-click button to produce a fully maxed-out account based on your selected locale.

#### Resource Helper
Grants items to your account so you can progress faster. Unlike the Account Builder, it adds resources but does not automate gameplay actions.

#### Reset Progress
Resets specific gameplay progress (e.g., DLP floors, live history) while keeping the rest of your account intact.

#### Import / Export Account
Export your account to `.db` (recommended) or `.json` format, and import from either format. Useful for:

- Backing up before an update
- Moving your account to another server
- Recovering from a pcap capture

See [Import & Export](import_export.md) for full details.

#### Other Features (Advanced)

> These features require detailed knowledge of the game's internals. Use them at your own risk — incorrect use can make your account unplayable. See [Advanced Usage](advanced_usage.md).

---

## Admin WebUI

**Address:** `http://<server_address>:8080/webui/admin`

For the embedded server, the default address is `http://127.0.0.1:8080/webui/admin`. Keep the game client open and use your device's browser to access it.

The Admin WebUI controls server-wide settings. The default admin password is empty — just press the login button. Changing the password is not recommended unless you are running a public server.

### Config Editor

| Option | Description |
|---|---|
| **Default item count** | Items given to a player when they first obtain that item type. Defaults to a generous amount; set to `0` for a more authentic experience. |
| **Mission progress multiplier** | Multiplies the progress earned per action in missions. Set higher to complete missions faster. Avoid extreme values. |
| **Resource config** | Controls how in-game resources behave (see below). |
| **Event frequency** | Controls how quickly events cycle. |

**Resource config modes:**

- `original` — Resources behave exactly as on the official server. All costs apply.
- `comfortable` *(default)* — Star gems, LP, and AP are unlimited. Daily song and tap bond limits are removed. Recommended for casual play.
- `free` — Resources can only increase, never decrease.

Note: accessories are not affected by resource config, though accessory *items* are.

Advanced config fields (server address, CDN address, etc.) are documented in the editor itself. Modifying them is considered [advanced usage](advanced_usage.md).

#### CDN Server Address

By default, assets are served from a public CDN. You can configure `elichika` to host its own CDN:

1. Place asset packages in `elichika/static/<package_name>`.
2. Set the CDN server address to the string `elichika` (or `elichika_tls` for HTTPS).

The server will then serve assets directly, using whatever address the client used to reach it.

### Event Selector

Choose the active event. Only events that have been fully prepared are shown.

### Event Scheduler

Schedule the next event to load automatically when the current one ends. Only prepared events are shown.

### Other Features (Advanced)

> Use at your own risk. Some features may not work correctly on all installations. See [Advanced Usage](advanced_usage.md).
