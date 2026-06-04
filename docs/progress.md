# Implementation Progress

This page tracks what is working, partially working, or not yet implemented. For a technical view of which API endpoints are handled, see [API Endpoints](endpoints.md).

## Feature Status

- [x] **Account creation / startup**
    - New accounts are created on first login or via the transfer system.
    - New accounts trigger the opening MV and tutorial, which fully work (some minor areas could be improved).
- [x] **Login**
    - Normal login, idol birthday login bonus, and new player login bonus all work.
- [x] **Profile**
    - Profile customization works.
    - Birthday can be set during the tutorial or changed via the WebUI.
- [x] **Live show**
    - Normal live, skip ticket, and 3DMV modes are all fully working.
    - Bond points are correctly awarded.
    - Own-partner guest selection works.
    - Drops are handled correctly.
- [x] **Story**
    - All story types are fully working.
    - Story songs work.
    - Starting fresh and progressing through the story correctly unlocks content.
- [ ] **Gacha**
    - [x] Working gacha with one banner per group.
    - [ ] Scouting tickets are not yet implemented.
- [x] **Training**
    - Training works and drops skills similarly to the real server. See [SIFAS-Lesson-Data](https://github.com/eman1can/SIFAS-Lesson-Data) for more details.
    - "Adjusted" mode (configurable) makes passive skills harder to obtain, increasing meta difficulty.
    - Items drop according to the real server; rally megaphones drop during channel live.
- [x] **Member bond**
    - Bond system, bond board, bond stories (unlocked by level), and bond songs all work.
- [x] **Bond ranking**
    - Working, though may be slow with a large number of accounts.
- [ ] **Membership (subscription)**
    - [x] Membership info is preserved for imported accounts.
    - [x] Default membership is applied to new accounts.
    - [ ] No tracking or veteran reward system.
- [x] **Shop** — returns fixed data; no further implementation planned.
- [x] **Exchange** — fully working; default data reflects the global server at EOS.
- [x] **School idol / Practice** — card grade up, level up, and practice all work.
- [x] **Accessories** — power up, drops from live, shop exchange, and WebUI addition all work.
- [x] **Channel** — working with ranking rewards.
- [x] **Present box** — fully working; all expected items should be present.
- [x] **Goal list**
    - Daily and weekly goals reset correctly.
    - Free goals available at EOS are tracked.
    - [ ] Some event-exclusive goals are not implemented and may be added later.
- [x] **Notices / news** — always empty; returns fixed data. No current plan to populate.
- [x] **Social (friends)** — working friend system with bad word checker.
- [x] **Title list** — titles stored in database, obtainable via goals.
- [x] **Datalink (account transfer)** — working. Passwords stored with bcrypt.
- [x] **Daily theater**
    - Server code and Global client integration fully working.
    - Japanese uses network logs or transcripts; English uses community translations.
    - [ ] Korean and Chinese (zh) translations not available.
- [x] **User model** — LP and AP recovery work correctly in original resource setting.
- [x] **DLP** — working with voltage ranking tracking; progress can be reset via WebUI.
- [ ] **Events**
    - [x] Marathon (story/point reward) events work; the first marathon event is available.
    - [ ] Event goals and event gacha not available.
    - [ ] Other marathon events have missing assets that need to be recreated.
    - [ ] Some limitations exist due to current design.
    - [ ] Mining (exchange) event not implemented.
    - [ ] SBL event not implemented.
    - [ ] Voltage ranking event not implemented.

---

## Embedded Client

For details on how the embedded Android client works, see the [elichika_embedded project](https://github.com/arina999999997/elichika_embedded).
