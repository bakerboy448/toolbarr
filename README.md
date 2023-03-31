
# <img style="height:400px" src="frontend/src/assets/images/logo.png">

Provides a suite of utilities to fix problems with Starr applications.
Toolbarr allows you to perform various actions against your Starr apps and their SQLite3 databases.

# Features

The first release provides only 1 feature:

- **Updating Root Folders**: Use this when you need to migrate your starr application to another host, and the file system paths to your media are changing.

# How?

Run this application on _your_ computer (mac/windows).
Configure it with your starr app details (url/api key), and an optional sqlite3 database file.

# Installation

There's a DMG for mac users and an EXE installer for Windows users on the [Releases](https://github.com/Notifiarr/toolbarr/releases) page.
Download the latest file for your OS and install it. Once installed, you can use the `About` menu to check for updates.

Linux users probably won't have great luck since the binaries provided are only going to work on a specific version of `glibc`.
Prepare to compile it yourself, but you can try installing it from the repo first:

```shell
curl -s https://golift.io/repo.sh | sudo bash -s - toolbarr
```

# Caution

This app may be destructive. Make backups. Do not connect it to a live SQLite database file; use a backup copy!

# I want more..

We do too. What other things do you want to do Sonarr or Radarr that you just can't do easily some other way?

- [Let us know how we can make your life easier](https://github.com/Notifiarr/toolbarr/issues/new).