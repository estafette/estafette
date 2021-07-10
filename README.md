# estafette

The estafette command line interface to run tons of commands locally for Estafette CI

## Usage

To install this cli on your mac run:

```bash
brew install estafette/stable/estafette
```

Then run `estafette help` to see what commands are available.

### > estafette manifest validate

The check whether your .estafette.yaml manifest file is valid runL

```bash
estafette manifest validate
```

## Development

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