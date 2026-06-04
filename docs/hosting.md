# Hosting Your Own Server

> **Advanced usage:** Hosting your own server is not required to play the game. If you just want to play, use the
> embedded Android client described in [Installing the Game](getting-started.md). Self-hosting can range from
> straightforward to complex depending on your setup — see [Advanced Usage](advanced_usage.md) for expectations.

---

## Prerequisites

Before installing, make sure you have the following depending on your platform:

| Platform      | Requirements                                                                                                                  |
|---------------|-------------------------------------------------------------------------------------------------------------------------------|
| Android       | Termux (from [F-Droid](https://f-droid.org/en/packages/com.termux/) or [GitHub](https://github.com/termux/termux-app#github)) |
| Windows       | Git, Go (both must be in `%PATH%`), and a bash shell (Git Bash works)                                                         |
| Linux / macOS | Git, Go                                                                                                                       |
| Docker        | Docker + Docker Compose                                                                                                       |


If you are hosting on your PC (Windows/Linux/macOS/Docker), the default asset CDN will download all files to `/static/packs/`, taking around 31GB of disk space. Make sure you have enough, or follow the instructions in [Assets](assets.md) to use only the public CDN. If you are hosting on PC, you will also need a compatible client. Check client availability before proceeding.

---

## Installing the Server

### Android (Termux) and Linux/macOS

Run the install script, which fetches and sets everything up automatically:

```bash
curl -L {{ repo_raw }}/master/bin/install.sh | bash
```

It is highly recommended to use only the public CDN in this case, to avoid a total download size of ~50GB.

### Windows (Manual)

Install Git and Go, ensure both are on your `%PATH%`, then run the same install script using Git Bash:

```bash
curl -L {{ repo_raw }}/master/bin/install.sh | bash
```

The script leaves some temporary files behind. If you prefer a clean setup, clone the repository and build manually —
the install script documents the necessary steps.

### Docker

See the dedicated [Docker Deployment](docker.md) guide. Docker skips several of the manual steps below and manages
configuration through a volume-mounted `data/config.json`.

---

## Running the Server

You must have the server running whenever you want to play.

**Android/Linux (installed via script):**

```bash
~/run_elichika
```

**Windows/Linux with a GUI:** run the compiled executable directly.

If you close Termux or the terminal, the server stops — reopen it and run the command again before playing.

---

## Configuration

On first startup, the server creates a `data/config.json` file with default settings. Most settings are manageable
through the [Admin WebUI](webui.md#admin-webui) at `http://<server_address>:8080/webui/admin`.

---

## Updating the Server

It is recommended to back up `userdata.db` (and optionally `serverstate.db`) before updating. The WebUI's import/export
feature can also serve as a backup — see [Import & Export](import_export.md).

### Basic Update (Recommended for large version gaps)

This backs up your data, reinstalls the server, and restores your data. It is slower but works from any version:

```bash
curl -L {{ repo_raw }}/master/bin/basic_update.sh | bash
```

Or via the installed shortcut:

```bash
~/basic_update_elichika
```

### Normal Update (Recommended for regular updates)

Faster than the basic update, but may break if your version is very old:

```bash
curl -L {{ repo_raw }}/master/bin/update.sh | bash
```

Or via the installed shortcut:

```bash
~/update_elichika
```
