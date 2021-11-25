# extrude

Analyse binaries for missing security features, information disclosure and more.

:construction: Extrude is in the early stages of development, and currently only supports ELF and MachO binaries. PE (Windows) binaries will be supported soon.

![Screenshot](screenshot.png)

## Usage

```
Usage:
  extrude [flags] [file]

Flags:
  -a, --all               Show details of all tests, not just those which failed.
  -w, --fail-on-warning   Exit with a non-zero status even if only warnings are discovered.
  -h, --help              help for extrude

```

## Docker

You can optionally run extrude with docker via:

```
docker run -v `pwd`:/blah -it ghcr.io/liamg/extrude /blah/targetfile
```

## Supported Checks

### ELF

- PIE
- RELRO
- BIND NOW
- Fortified Source
- Stack Canary
- NX Stack

### MachO

- PIE
- Stack Canary
- NX Stack
- NX Heap
- ARC

### Windows

_Coming soon..._

## TODO

- Add support for PE
- Add secret scanning
- Detect packers

