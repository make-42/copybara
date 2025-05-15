# Copybara

A **Wayland** clipboard automation tool for cleaning URLs by removing tracking query parameters and applying other regex rules like replacing links for twitter/instagram to fix embeds.

## Features
 - Remove tracking in URLs
 - Custom regex rules
 - (Soon: ability to toggle automations with [IPC](https://pkg.go.dev/github.com/james-barrow/golang-ipc))
 - More to come



## Config
Configuration files are located at `~/.config/ontake/copybara/config.yml`
An example configuration file is created on first launch.

Here's an example configuration file:
```yaml
notificationsonappliedautomations: true
enableregexautomations: true
enableurlcleaning: true
extraurlcleaningrulesandoverrides:
  exampleoverride:
    urlpattern: ""
    completeprovider: false
    rules: []
    referralmarketing: []
    exceptions: []
    rawrules: []
    redirections: []
    forceredirection: false
extraregexrules:
- isurlrule: true
  pattern: ^https?:\/\/(?:[a-z0-9-]+\.)*?instagram\.com\/reel
  exceptions: []
  replacewith: https://www.ddinstagram.com/reel
- isurlrule: true
  pattern: ^https?:\/\/(?:[a-z0-9-]+\.)*?x\.com
  exceptions:
  - ^https?:\/\/(?:[a-z0-9-]+\.)*?x\.com$
  replacewith: https://fxtwitter.com
```

## Nix Flake
There's a Nix flake for this project in this repository for easy installation.
