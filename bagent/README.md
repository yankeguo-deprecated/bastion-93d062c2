# bagent

Bastion Agent

## Requirements

These executables are required

* `/etc/passwd`
* `/bin/bash`
* `/bin/false`
* `chsh`
* `passwd`
* `gpasswd`

## Installation

## Settings

Settings are passed with environments

* `BASTION_HOST` host for bastion server, with leading `http://` or `https://`, without tailing `/`
* `BASTION_TOKEN` access token for this server
* `BASTION_HOME` home directory for bastion created accounts, default to `/home`

* `BASTION_DEBUG` if set, operations will **NOT** be executed and generated script will be printed