---
title: User Guide
layout: default
nav_order: 2
has_children: true
---

# User Guide

This is section contain basic information about how `lore-fetcher` is used.

The `lore-fetcher` CLI commands can be broken down into three different components:

``` lore-fetcher <option-flag> <parameter-flags>```

To use `lore-fetcher` you must run the `lore-fetcher` command with an `<option flag>`, which invokes
one of the `lore-fetcher` utility commands (search for new patches, send a result, etc.) to be used. Notice
you must provide exactly one `<option-flag>`.

Then, you must provide parameter information the `lore-fetcher` option requires to be executed correctly.
Those are primarily provided on the `lore-fetcher`'s config file, but you can use the `<parameter-flags>`
to override specific parameters.

For example, if the config file sets `mailing-list = all`, the `lore-fetcher --fetch` feature will search
for patches in the `all` mailing list. However, if one uses `lore-fetcher --fetch --mailing-list=amd-gfx` with
the same config file, the config file parameter will be overriden and patches will be tracked on `amd-gfx`.
When using `lore-fetcher` commands, ensure that the parameters are given with either config files or flags.


