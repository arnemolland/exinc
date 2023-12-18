# exinc

Utility tool to merge and extract given intervals from a set of include intervals.

## Usage

```bash
# json input
exinc -f <file>
# stdin
cat <file> | exinc
```

## Building

```bash
# fmt, vet, test and build
make all
# build only
make bin/exinc
```
