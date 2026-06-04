# Asset Management

The game client downloads assets (textures, audio, animations, etc.) from the server on demand. Elichika manages those
assets through a two-tier system: a local pack cache under `static/packs/`, and a proxy CDN that automatically fetches
anything not already cached.

---

## How Assets Are Stored

All asset packs are stored under `static/packs/` using their pack name as the filename:

```
static/
└── packs/
    ├── abcd12          ← individual pack
    ├── efgh34
    ├── bhjak1          ← metapack (contains several packs concatenated)
    └── ...
```

**Packs** are the base unit of download — a small binary file identified by a short alphanumeric name (e.g., `abcd12`).
The client requests packs by name.

**Metapacks** bundle several related packs into one file to reduce round-trips. When the server tells the client to
download a pack that lives inside a metapack, it gives the client the metapack filename along with a byte offset and
length so the client can extract the specific pack. Which files are metapacks is determined by the
`m_asset_package_mapping` table in the masterdata.db database. Metapack files are named just the same as normal files,
so it can be hard to find them without the lookup table.

### Downloading content from the WebArchive downloads / other zipped downloads

If you download the data from the [internet archive](https://archive.org/details/ll-sifas-cdn-data), the files will be
pre-pended by `pack<c>` or `meta<c>` where `<c>` is the first letter of the packs or metapacks that folder contains.
This is how the game internally stores the packs so that the filesystem doesn't slow down due to so many files in one
directory.

Due to the fact that many of the original assets are missing, and we commonly need to serve assets from the jp client to
the gl client or vice-versa, and the fact that pack names are unique across both gl and en, the elichika server stores
all packs under `/static/packs/` regardless of the name. This is the same for the official CDN.

Thus, if you download the assets from the web archive, you will need to extract the files from their folders into one
file list. A simple python script to do this can be found below.

```python
import os
import shutil
for folder is on.listdir("static/data"):
    for file in os.listdir(os.path.join("static/data", folder)):
        shutil.copy(os.path.join("static/data", folder, file), os.path.join("static/packs/", file))
```

This code assumes that you have put the data from the web archive into `static/data/` and that you are running the
script from the root of the elichika project.

---

## The Proxy CDN

By default, Elichika does **not** require you to pre-download any assets. When the client requests a pack that is not
yet in `static/packs/`, the server fetches it from an upstream CDN, saves it to `static/packs/<packname>`, and streams
it to the client — all transparently.

This is controlled by the `static_proxy_cdn` config option (default: `https://llsifas.imsofucking.gay`):

| Value           | Behaviour                                                  |
|-----------------|------------------------------------------------------------|
| A URL (default) | Missing packs are fetched from that URL and cached locally |
| Empty string    | Proxy disabled; missing packs return an error              |

The first time a pack is requested it is downloaded and cached. Subsequent requests are served from disk.

---

## Extraction Utilities

Three utilities under `cmd/` extract specific asset types from the pack files into more accessible formats. These are
useful for development, debugging, and previewing content — they are **not** required for normal server operation.

All three utilities expect the asset databases to be present at `assets/db/` and write output to subdirectories of
`static/`.

---

### `extract-sounds`

Extracts all audio from the asset packs and converts it to WAV.

**What it does:**

1. Reads `asset_a_en.db` to find all sound sheets and their pack names.
2. Extracts the ACB (audio container) and AWB (audio wave bank) for each sheet from `static/packs/`.
3. Decrypts the HCA audio frames inside the AWB.
4. Converts HCA to WAV.
5. Caches already-extracted files so re-runs are fast.

**Output:**

```
static/sounds/
├── acb/    ← audio container files
├── awb/    ← audio wave bank files
└── wav/    ← converted WAV files (one per sound sheet)
```

WAV files are also cached here by the WebUI when it plays audio previews — so running `extract-sounds` up front means
the WebUI does not have to decode audio on the fly.

**Run:**

```bash
go run ./cmd/extract-sounds
```

---

### `extract-movies`

Extracts USM video files from the asset packs.

**What it does:**

1. Reads `asset_a_en.db` for the movie list and their pack mappings.
2. For each movie, reads the raw pack data from `static/packs/` (using stored offsets).
3. Writes the USM (video container) file to disk.

Processing is parallelised across 20 goroutines.

**Output:**

```
static/sounds/usm/
└── <pack_name>.usm    ← one USM file per movie
```

**Run:**

```bash
go run ./cmd/extract-movies
```

---

### `extract-timelines`

Extracts and decrypts navigator animation timeline bundles.

**What it does:**

1. Reads `masterdata.db` to map timeline IDs to asset paths.
2. Reads `asset_a_en.db` for pack locations and per-asset encryption keys.
3. For each timeline, reads the encrypted data from the relevant pack in `static/packs/` and decrypts it.
4. Writes the decrypted UnityFS asset bundle to disk.

**Output:**

```
static/timelines/
└── <id>.unityfs    ← one UnityFS bundle per timeline
```

**Run:**

```bash
go run ./cmd/extract-timelines
```

---

## WebUI Asset Previews

The Admin WebUI can display images and play audio directly, without running the extraction utilities first. It decrypts
and decodes assets on demand using the same keys stored in the asset databases:

- **Images** are decrypted in memory using per-asset keys and returned as raw image data.
- **Audio** is decrypted and converted to WAV on first access, then cached to `static/sounds/wav/` for subsequent
  requests.

This means the extraction utilities speed up the WebUI (pre-warmed cache) but are not required for it to work.

---

## Directory Reference

```
static/
├── packs/           ← raw pack files (auto-populated by proxy, or manually)
├── sounds/
│   ├── acb/         ← audio containers (extract-sounds)
│   ├── awb/         ← audio wave banks (extract-sounds)
│   ├── wav/         ← decoded WAV audio (extract-sounds / WebUI cache)
│   └── usm/         ← video files (extract-movies)
└── timelines/       ← UnityFS animation bundles (extract-timelines)

assets/
└── db/
    ├── gl/
    │   ├── asset_a_en.db    ← Global asset metadata (textures, sounds, etc.)
    │   ├── asset_i_en.db
    │   ├── asset_a_ko.db, asset_a_zh.db, asset_i_ko.db, asset_i_zh.db
    └── jp/
        ├── asset_a_ja.db    ← JP asset metadata
        └── asset_i_ja.db
```
