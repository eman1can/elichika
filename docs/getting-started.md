# Installing the Game

For a more comprehensive guide, see the [LL Hax docs](https://carette.codeberg.page/ll-hax-docs/sifas/easy-install/).

---

## Android

There are three ways to play on Android, depending on how much control you want over the server.

### Option 1: Public Server (Easiest)

Use a pre-built client pointed at a community-hosted server — no setup required.

1. Download the APK for your preferred version:
    - **Global:
      ** [ALL_STARS_3.12.0_gl.apk](https://ethanthesleepy.one/public/lovelive/sifas/public_server/ALL_STARS_3.12.0_gl.apk)
    - **Japanese:
      ** [ALL_STARS_3.12.0_ja.apk](https://ethanthesleepy.one/public/lovelive/sifas/public_server/ALL_STARS_3.12.0_ja.apk)
2. Install the APK ([how to install an APK](https://www.wikihow.com/Install-APK-Files-on-Android)) and launch the game.

### Option 2: Embedded Server

The embedded client bundles its own server inside the APK — no separate machine or setup needed.

> **Note:** Embedded clients are 64-bit only. If your device is 32-bit, use a public server or self-host instead.

1. Download the latest embedded APK from the [releases page]({{ config.repo_url }}/releases/tag/embedded) and install
   it.
2. Open the game and wait until the title screen appears:
    - **Japanese:** just wait for the title screen.
    - **Global:** select your language, then wait for the title screen.
3. Close the game, then open it again — asset extraction runs at this point and may take a few minutes. Once the title
   screen appears again, extraction is complete.
    - On some devices, extraction may start on the very first run instead.
4. On subsequent launches, wait for the startup popups to finish before playing. Once the popups stop, the embedded
   server is ready.

### Option 3: Self-Hosted Server

Run the server yourself on a PC, a separate Android device (via Termux), or Docker. This gives you full control over
server configuration. See [Hosting Your Own Server](hosting.md) to get started.
Make sure to install the patched client from the [releases page]({{ config.repo_url }}/releases/tag/embedded) and
install it.

### Updating (Embedded or Self-Hosted)

1. Optionally, [back up your account](webui.md#features) before updating.
2. Download the latest APK from the [releases page]({{ config.repo_url }}/releases/tag/embedded) and install it.
3. Install it **on top of the existing app** — do **not** uninstall first.
4. Open the game normally. The first run after an update will extract some data, so expect a brief wait.

---

## iOS

Follow the [LL Hax iOS setup guide](https://carette.codeberg.page/ll-hax-docs/sifas/easy-install/#ios_setup) to install
the patched client.

Once the client is installed, open the app settings and set the server address:

### Public Server

Enter the server URL: `https://sifas.ethanthesleepy.one/`

Leave the public key field empty and play.

### Self-Hosted Server

Point the server URL at your own server. See [Hosting Your Own Server](hosting.md) for setup instructions.
