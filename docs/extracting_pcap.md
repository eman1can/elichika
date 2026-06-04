# Extracting Game Data from a PCAP

> **Advanced usage.** You need to know what a PCAP is and have a valid capture from before EOS. See [Advanced Usage](advanced_usage.md).

This guide assumes you used **PCAPdroid** to capture the network data. The steps are mostly the same for other capture tools.

---

## What You Need

A valid capture produces two files:

- `PCAPdroid_<date>_<time>.pcap` — the raw encrypted network data
- `sslkeylog.txt` (or similar) — the TLS session keys needed to decrypt it

The key file should look like this:
```
CLIENT_RANDOM <some_data> <some_other_data>
CLIENT_RANDOM <some_other_data> <yet_more_data>
...
```

**Both files are required.** Without the key file, the capture cannot be decrypted in any reasonable amount of time.

> The decrypted capture will contain sensitive account data, including your SIF2 transfer password. Only share it with people you fully trust.

---

## Step 1: Install Wireshark

Download Wireshark from [wireshark.org](https://www.wireshark.org/download.html). It is available for Windows and macOS natively. On Linux, install from your package manager or build from source.

---

## Step 2: Load the TLS Key

1. Open Wireshark.
2. Go to **Edit → Preferences** (`Ctrl+Shift+P` on Windows).
3. In the left panel, expand **Protocols** and scroll to **TLS**.
4. Set the **Master-Secret log filename** field to your `sslkeylog.txt` file.

![TLS preferences panel](images/pcap_1.png)

---

## Step 3: Open the PCAP File

Go to **File → Open** (`Ctrl+O`) and select your `.pcap` file. If the key was loaded correctly, you should see decrypted HTTP/2 traffic.

![Wireshark with decrypted traffic](images/pcap_2.png)

Click the **Apply a display filter** bar and type `json`. This filters to only JSON traffic, which is the game data.

---

## Step 4: Find the Login Packet

Go to **Edit → Find Packet** (`Ctrl+F`). Set the options to:

- **Search in:** Packet list
- **String type:** Narrow & Wide
- **Case sensitive:** unchecked
- **Search for:** `login/login`

![Find packet dialog](images/pcap_3.png)

The matching packet will be highlighted.

![Highlighted login packet](images/pcap_4.png)

> If you see two `login/login` packets (from logging in on a new device), ignore the one with the smaller length — it does not contain account data.

---

## Step 5: Extract the Login Response

Select the packet **immediately after** the highlighted login request. Right-click it and choose **Follow → HTTP/2 Stream**.

A new window opens showing the stream data.

![HTTP/2 stream window](images/pcap_6.png)

At the bottom left, click **Entire conversation** and choose whichever side is larger. You know you have the right side when all the displayed text turns blue. Then click **Save As** and save the file — for example, as `login.json`.

---

## Step 6: Clean Up the JSON File

Open `login.json` in a text editor that can handle large files (Notepad or Notepad++ work well).

**Remove the HTTP headers at the top.** Delete everything from the beginning of the file up to and including the first blank line (the line with `status: 200`, the `x-amz-...` header line, etc.).

**Remove the framing data at the start.** Your data will look something like:
```
[16xxxxxxxxxxx,"2d61e7b4e89961c7",0,{"session_key":...
```
Delete everything before `{"session_key":`, leaving:
```
{"session_key":...
```

**Remove the framing data at the end.** The file will end with something like:
```
..."check_maintenance":true,"repro_info":{"group_no":1}},"somerandomhexstring"]
```
Remove the trailing `,"somerandomhexstring"]` to leave:
```
..."check_maintenance":true,"repro_info":{"group_no":1}}
```

Save the file. You now have a valid `login.json` that can be imported using the [WebUI import feature](import_export.md).
