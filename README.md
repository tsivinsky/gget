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

By default, gget will output file's content to stdout.

```bash
gget https://github.com/tsivinsky/gget/blob/master/README.md
```

But you can pass `-o file.txt` flag to set destination file.

```bash
gget -o file.txt https://github.com/tsivinsky/gget/blob/master/README.md
```

If url has fragment with selected lines, gget will get only those lines

Single line

```bash
gget https://github.com/tsivinsky/gget/blob/master/Makefile#L1
```

Range of lines

```bash
gget https://github.com/tsivinsky/gget/blob/master/Makefile#L1-L5
```
