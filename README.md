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

By default, gget will print file content to standard output.

If you want to use file as an output, you can add `-o` flag

```bash
gget https://github.com/tsivinsky/gget/blob/master/main.go -o gget.go
```
