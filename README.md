# gget

I needed a way to download file from GitHub without going to its raw version and I made this tiny program.

If there is a way to do it with a cURL, good but I didn't want to search for it online.

## Install

Requirements:

- Go installed

```bash
go install github.com/tsivinsky/gget@latest
```

If you don't have Go installed, you can install binary from latest release on [github.com/tsivinsky/gget/releases](https://github.com/tsivinsky/gget/releases)

## Usage

By default, gget will save file with the name from GitHub.

```bash
gget https://github.com/tsivinsky/gget/blob/master/README.md
```

But you can pass `-o path/to/file.txt` flag to change destination file.

```bash
gget -o file.txt https://github.com/tsivinsky/gget/blob/master/README.md
```
