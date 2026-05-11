# GUI
This package implement GUIs relevant to elichika:

- This would contain programs that are "assets managers".
- So programs that make preparing content for elichika (and SIFAS generally) easier.

Also, this is only tested on Windows 10.

## Requirements

### CGO and a C compiler

The GUI programs use CGO, so you need a C compiler available on your PATH. On Windows, install [MinGW-w64](https://www.mingw-w64.org/) (e.g. via [MSYS2](https://www.msys2.org/)) and make sure `gcc` is on your PATH.

### CSFML

The programs use SFML via the [go-sfml](https://github.com/telroshan/go-sfml) binding, which requires **CSFML 2.5** (the C binding of SFML) to be installed.

1. Download the CSFML 2.5 binaries for Windows from https://www.sfml-dev.org/download/csfml/ (pick the MinGW 64-bit build to match your compiler).
2. Extract the archive. You will get an `include/` directory and a `lib/` directory.
3. When building, point CGO at these directories:
   ```
   set CGO_CFLAGS=-I"C:\path\to\csfml\include"
   set CGO_LDFLAGS=-L"C:\path\to\csfml\lib"
   ```
4. At runtime the DLLs (`csfml-graphics-2.dll`, `csfml-window-2.dll`, `csfml-system-2.dll`) must be reachable. The simplest approach is to copy them next to the built executable, or add the CSFML `bin/` directory to your `PATH`.

---

## Texture Viewer (`app/texture_viewer`)

The texture viewer lets you inspect and export game textures stored in the elichika asset packs.

### Data requirements

Before running the texture viewer you need two things in the elichika root directory:

**`texture.db`** — a SQLite database that contains the texture index. It is expected to have:
- Tables named `texture_en`, `texture_ja`, `texture_zh`, `texture_ko`, `texture_th` — one per game language — each with columns `asset_path`, `pack_name`, `head`, `size`, `key1`, `key2`.
- Optionally a `raw_texture` table (columns: `pack_name`, `head`, `size`, `key1`, `key2`, `width`, `height`, `error`) used for startup dimension pre-calculation. If this table is absent the viewer prints a warning and skips that step — loading individual textures by asset path still works normally.
- Optionally a `static_file_dir` table with a `dir` column listing paths to search for asset packs. If this table is absent the viewer falls back to the two hardcoded paths `static/b66ec2295e9a00aa/` and `static/2d61e7b4e89961c7/`.

**Asset pack files** — the raw encrypted pack files referenced in `texture.db`. By default the viewer looks for them relative to the elichika root under `static/b66ec2295e9a00aa/<pack_name>` and `static/2d61e7b4e89961c7/<pack_name>`. You can override these search paths by populating the `static_file_dir` table in `texture.db`.

### Building

Run from the **elichika root** (so the `elichika` Go module is in scope):

```bat
set CGO_CFLAGS=-I"C:\path\to\csfml\include"
set CGO_LDFLAGS=-L"C:\path\to\csfml\lib"
go build -o texture_viewer.exe ./gui/app/texture_viewer/
```

This produces `texture_viewer.exe` in the elichika root.

### Running

The executable must be run from the **elichika root** directory so it can find `texture.db` and the `static/` asset packs at their relative paths:

```bat
.\texture_viewer.exe
```

On startup the viewer automatically calculates and caches width/height metadata for any textures in `raw_texture` that have not been processed yet (width = 0, height = 0). This can take a moment the first time.

### UI overview

The viewer window is 1800×1000 and shows controls across the top, with the loaded texture image displayed below them.

**Asset path row**
- **Asset path** text box — type a game asset path (e.g. `ui/texture/card/m_sr_kasu_04_t.png`) and press Enter, or click the **Load asset** button to display that texture.
- **Load asset** button — loads the asset path currently in the text box.
- **Extract** button — saves the currently displayed texture to `texture_output.png` in the working directory.

**SQL row**
- **SQL** text box — enter a raw SQL query against `texture.db` and press Enter or click **Run SQL** to execute it. The query should return rows with columns `asset_path`, `pack_name`, `head`, `size`, `key1`, `key2`. If you type a very short string (3 characters or fewer) it is treated as an asset path and wrapped in `SELECT * FROM texture_pack WHERE asset_path = "..."` automatically.
- **Run SQL** button — executes the current SQL and loads the first result.
- **Next SQL** / **Previous SQL** buttons — cycle through the result set returned by the last SQL query. The status bar at the top of the window shows the current asset path and its index in the result set.

**Status bar** — the first row of the window; updated automatically to reflect the last operation or error message.

**Screenshot** — press Ctrl+C at any time to save a screenshot of the current window to `screenshot.png` in the working directory.

**Paste** — Ctrl+V pastes clipboard text into the focused text box.

### Troubleshooting

- *Panic at startup* — make sure `texture.db` exists in the directory you are running the executable from. The viewer panics immediately if it cannot open the database. If the database exists but lacks a `raw_texture` table, the viewer now prints a warning and continues normally instead of panicking.
- *"asset pack doesn't exist"* — the pack file referenced in `texture.db` was not found under any of the configured `static/` directories. Check that your asset packs are present and that the paths in `static_file_dir` (or the hardcoded defaults) are correct.
- *"can't find asset path"* — the typed path does not exist in any of the five language tables. Check spelling and try a different language's table if relevant.
- *Window does not appear / crashes on startup* — make sure the CSFML DLLs are next to the executable or on your `PATH`.

