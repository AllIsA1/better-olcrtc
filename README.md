# Better olcRTC

A plugin for the [LucX-UI](https://github.com/AlexeyLCP/lucx-ui) plugin system. **It does nothing now.**

It used to make the panel download the [ghostlane-project fork](https://github.com/ghostlane-project/olcrtc) of olcRTC (v0.12.1: VK Calls,
MTS Link, SaluteJazz, per-user keys, UDP, direct routes) instead of the archived original. Since **lucx.325**
([#157](https://github.com/AlexeyLCP/lucx-ui/issues/157)) the panel's release brings that fork itself, so the plugin has no job left (version
2.0.0 removed it: no core download, no permissions, no `requires`).

It is left installed for the time being and answers only the plugin system's health check. Remove it from the Plugins page whenever you like;
the olcRTC core of the panel stays as it is. A panel older than lucx.325 needs Better olcRTC **1.0.0**.

The `core-v0.12.1-1` release of this repository (the binaries the old version installed) stays where it is.

MIT, see `LICENSE`.
