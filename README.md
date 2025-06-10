# Copybara

A **Wayland** clipboard automation tool for cleaning URLs by removing tracking query parameters and applying other regex rules like replacing links for twitter/instagram to fix embeds.

(this demo shows copybara in action as well as its toggle function)

https://github.com/user-attachments/assets/3e1e7432-ca74-4f23-a031-af9100b51a9c


## Features
 - Remove tracking in URLs
 - Custom regex rules
 - Ability to toggle it on and off by running `copybara --toggle` (so that it can be keybound)
 - More to come



## Install
### Nix/NixOS
There's a Nix flake for this project in this repository for easy installation.

You can even try `copybara` out by running `nix run github:make-42/copybara`
### Arch
There is also an AUR package for Arch Linux users: `ontake-copybara-git`

## Config
Configuration files are located at `~/.config/ontake/copybara/config.yml`
An example configuration file is created on first launch.

Here's an example configuration file:
```yaml
enableregexautomations: true
enableurlcleaning: true
notificationsonappliedautomations: true
extraurlcleaningrulesandoverrides:
  exampleoverride:
    urlpattern: ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}
    completeprovider: false
    rules:
      - ved
      - bi[a-z]*
      - gfe_[a-z]*
      - ei
      - source
      - gs_[a-z]*
      - site
      - oq
      - esrc
      - uact
      - cd
      - cad
      - gws_[a-z]*
      - atyp
      - vet
      - _u
      - je
      - dcr
      - ie
      - sei
      - sa
      - dpr
      - btn[a-z]*
      - usg
      - cd
      - cad
      - uact
      - aqs
      - sourceid
      - sxsrf
      - rlz
      - i-would-rather-use-firefox
      - pcampaignid
      - sca_(?:esv|upv)
      - iflsig
      - fbs
      - ictx
    referralmarketing:
      - referrer
    exceptions:
      - ^https?:\/\/mail\.google\.com\/mail\/u\/
      - ^https?:\/\/accounts\.google\.com\/o\/oauth2\/
      - ^https?:\/\/accounts\.google\.com\/signin\/oauth\/
      - ^https?:\/\/(?:docs|accounts)\.google(?:\.[a-z]{2,}){1,}
      - ^https?:\/\/([a-z0-9-\.])*(chat|drive)\.google\.com\/videoplayback
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}(?:\/upload)?\/drive\/
      - ^https?:\/\/news\.google\.com.*\?hl=.
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/s\?tbm=map.*?gs_[a-z]*=.
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/(?:complete\/search|setprefs|searchbyimage)
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/(?:appsactivity|aclk\?)
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/safe[-]?browsing\/([^&]+)
    rawrules: []
    redirections:
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/url\?.*?(?:url|q)=(https?[^&]+)
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/.*?adurl=([^&]+)
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?google(?:\.[a-z]{2,}){1,}\/amp\/s\/([^&]+)
extraregexrules:
  - isurlrule: true
    pattern: ^https?:\/\/(?:[a-z0-9-]+\.)*?instagram\.com\/reel
    exceptions: []
    replacewith: https://www.ddinstagram.com/reel
  - isurlrule: true
    pattern: ^https?:\/\/(?:[a-z0-9-]+\.)*?x\.com
    exceptions:
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?x\.com$
      - ^https?:\/\/(?:[a-z0-9-]+\.)*?x\.com/$
    replacewith: https://fxtwitter.com
```

## Build
Run
```
go mod download
go build
```
in the root directory of the project.
The `./copybara` file can be installed anywhere on your system.

## License and LGPL Library Notice
This project is licensed under the MIT License, with the exception of the following third-party asset:
 - The file `./urlclean/assets/data.min.json` is derived from the [ClearURLs](https://github.com/ClearURLs/Rules), which is licensed under the [GNU Lesser General Public License (LGPL) v3](https://www.gnu.org/licenses/lgpl-3.0.html).
###  Embedded LGPL Asset
- This file is **embedded into the binary at compile time**.
- In accordance with section 4 of the LGPL v3:
  - You are permitted to modify or replace `data.min.json` under the terms of the LGPL.
  - To support this, we provide either:
    - **Build instructions** to allow replacing `data.min.json` before compilation, or
    - **Object files or source code** to recompile the application with a modified version of the file.

#### Rebuilding with a Modified Version
To use a modified version of the embedded ClearURLs rules:

1. Replace the `./urlclean/assets/data.min.json` file with your own modified version.
2. Rebuild the project from source using the standard build instructions (see [Build](#build)).

This ensures compliance with the LGPL’s requirements for user modification and replacement.
