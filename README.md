# estafette-cli
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

Development versions get pushed to github for each commit and have their own tap repository. You can install it via

```bash
brew install estafette/dev/estafette-dev
```

And then use it with

```bash
estafette-dev help
```