# estafette

The estafette command line interface to run tons of commands locally for Estafette CI


# Install

## With Homebrew

First install Homebrew:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

And docker:

```
brew install --cask docker
```

Then install the `estafette` cli with

```
brew install estafette/stable/estafette
```

Updating to the latest can be done with

```
brew upgrade
```

## With GoFish

First install [GoFish](https://gofi.sh/):

_MacOS/Linux_

```
curl -fsSL https://raw.githubusercontent.com/fishworks/gofish/main/scripts/install.sh | bash
gofish init
```

_Windows_

```
Set-ExecutionPolicy Bypass -Scope Process -Force
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/fishworks/gofish/main/scripts/install.ps1'))
gofish init
```

Then install the `estafette` cli with

```
gofish rig add https://github.com/estafette/fish-food
gofish install estafette
```

If you already have it installed upgrade to the latest and greatest with

```
gofish upgrade
```

## From source

```
go install github.com/estafette/estafette
```


# Usage

Run `estafette help` to see what commands are available.

## Validate

To check whether your .estafette.yaml manifest file is valid run

```bash
estafette validate
```

## Build

To build the .estafette.yaml manifest locally run

```bash
estafette build
```

This will only run the first stage; if you want to run more stages pass the stages names as a comma separated list

```bash
estafette build restore,build,test
```

Do realize your local build doesn't have access to the same credentials the Estafette CI server has so you won't be able to build all stages.


# Development

For local development when running `go build .` the generated binary can be used with

```bash
./estafette help
```

Development versions get released as a pre-release version on github for each commit and have their brew formular updates in a development tap repository. You can install it via

```bash
brew install estafette/dev/estafette-dev
```

And then use it with

```bash
estafette-dev help
```

## Releases

Official releases are performed by creating and pushing a branch with the same version as specified in `version.semver.releaseBranch`.