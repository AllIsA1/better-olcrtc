# Better olcRTC

A plugin for the [LucX-UI](https://github.com/AlexeyLCP/lucx-ui) plugin system. It replaces the **olcRTC core** of the panel (the
program an olcRTC inbound runs) by the **[ghostlane-project fork](https://github.com/ghostlane-project/olcrtc)**, release **v0.12.1**.

## What the fork adds

- Three more providers: **SaluteJazz**, **VK Calls** and **MTS Link** (besides Jitsi, Yandex Telemost and WB Stream), and the
  provider/transport pairs they allow (SaluteJazz: datachannel; VK Calls and MTS Link: vp8channel).
- Several keys on one server (`crypto.keys`): a key per person next to the room key.
- The lossy UDP relay (`udp`), a fixed DTLS ClientHello profile (`dtls.profile`), and for a client the destinations dialed directly
  (`route.direct`) and the server's `/stats`.
- More video settings (QR size and error correction, tile module and Reed–Solomon share), and `auth.token` for WB Stream.

The inbound form that sets all of this, and the `olcrtc://…$name` link, are part of the panel build `v3.10.2-lucx.320-plugins.1` and later;
the plugin only brings the program those settings need. A setting left at its default is not written into the config, so the
old core still runs what it ran before; the three new providers need this plugin's core.

## What it does

- The panel starts every olcRTC inbound with a program file it keeps in its `bin` folder. The panel's own release ships the original
  olcRTC; this plugin makes the panel download the reviewed fork build and check its hash. It is the panel that downloads (its
  Settings → Cores mechanism), over HTTPS from a release of this repository, and the download is refused unless the sha256 in the
  plugin matches.
- It checks at start and every six hours. It replaces a core it did not install (the one that came with the panel) once, and leaves
  alone the one it installed, so a file you upload yourself on Settings → Cores stays until a newer plugin version brings a newer core.
- The inbounds that run the file are stopped when it is replaced and started again by the panel on its next pass (panel build
  `v3.10.2-lucx.320-plugins.1` and later; on an older build they pick the new file up at their next restart).
- Switch the plugin off or remove it and the core stays as it is. It has no page and no sidebar entry.

**Clients.** The fork keeps the original's wire format (the OLC2 record layer and OLVC v5 frames), and a server of this version is
documented to serve every client of v0.11 and v0.12 of the fork. Whether a client built on the original olcRTC connects to it is not
verified here: try one client before moving people over. The three new providers need a client that has them.

## Requires

Better olcRTC declares `"requires": ["better-cores"]`: the plugin system (host-5 and later) installs and turns it on only while
[Better Cores](https://github.com/AllIsA1/better-cores) is installed and on, and will not turn that one off or remove it while this
is installed. Better Cores shows the fork as the project the olcRTC core comes from and its last release.

## Permissions

The permission `api:olcrtc:write` is new in the plugin system **host-7**; an older host refuses the plugin (and the catalog) until it is updated.

`api:cores:read` (is a core there: the status of the tunnel cores, read only) and `api:olcrtc:write` (download the olcRTC core by a
link and a hash, remove it): the narrow permissions of the plugin system for exactly these panel routes and no other. **Write is a
dangerous permission**: the panel runs this program as root, so the operator confirms it when installing. The plugin has no network
permission: the panel downloads the file itself.

## The core

`core.go` pins the version, the fork's commit and the sha256 of each architecture. The binaries are built by
`.github/workflows/core.yml` (`scripts/build-core.sh`: the fork's tag must point at exactly the named commit; static, CGO off,
`-trimpath`, Go 1.26.8, so the same commit gives the same bytes) and published as the release `core-<version>-<n>` of this repository.
olcRTC is Apache-2.0 software of its authors (https://github.com/ghostlane-project/olcrtc); the release names the source commit and
builds it unchanged.

## Build

`make build` writes `bin/<arch>/betterolcrtc` (commit `bin/`: the panel downloads the repository as it is). `make test` runs the tests.
The build is reproducible (`-trimpath -buildvcs=false`, CGO off): with Go 1.26.8, `make build` leaves `bin/` unchanged, which is how a reviewer checks that the
committed binaries are this source.

MIT, see `LICENSE`. The olcRTC program keeps its own license.
