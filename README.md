# csn

## Description

CLI version of chiasenhac.com, written in Golang

## Usage

```
$ csn search "Co khi nao roi xa"
Search: Co khi nao roi xa
=======
Có Khi Nào Rời Xa - Bích Phương | 4:25Lossless | Download: 147.063
Song play link: http://chiasenhac.com/mp3/vietnam/v-pop/co-khi-nao-roi-xa~bich-phuong~1110936.html
Download link:
MP3 128kbps: http://data19.chiasenhac.com/downloads/1111/2/1110936-9de8168c/128/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong%20%5BMP3%20128kbps%5D.mp3
MP3 320kbps: http://data19.chiasenhac.com/downloads/1111/2/1110936-9de8168c/320/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong%20%5BMP3%20320kbps%5D.mp3
M4A 32kbps: http://data19.chiasenhac.com/downloads/1111/2/1110936-9de8168c/32/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong%20%5BM4A%2032kbps%5D.m4a


Có Khi Nào Rời Xa - Bích Phương; Nukan Trần Tùng Anh | 4:37Lossless | Download: 24.503
Song play link: http://chiasenhac.com/mp3/vietnam/v-pop/co-khi-nao-roi-xa~bich-phuong-nukan-tran-tung-anh~1015655.html
Download link:
MP3 320kbps: http://data16.chiasenhac.com/downloads/1016/2/1015655-c01d39a9/320/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong_%20Nukan%20T%20%5BMP3%20320kbps%5D.mp3
M4A 32kbps: http://data16.chiasenhac.com/downloads/1016/2/1015655-c01d39a9/32/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong_%20Nukan%20T%20%5BM4A%2032kbps%5D.m4a
MP3 128kbps: http://data16.chiasenhac.com/downloads/1016/2/1015655-c01d39a9/128/Co%20Khi%20Nao%20Roi%20Xa%20-%20Bich%20Phuong_%20Nukan%20T%20%5BMP3%20128kbps%5D.mp3
```


```
$ csn help
NAME:
   csn - ChiaSeNhac.com CLI

USAGE:
   csn [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   tunglam14 <dangtunglam14@gmail.com>

COMMANDS:
    search	Search a song. Usage: csn search <song_name>
    get		Download a song. Usage: csn get <song_id>

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/tunglam14/gocsn
```

## Contribution

1. Fork ([https://github.com/tunglam14/gocsn/fork](https://github.com/tunglam14/gocsn/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[tunglam14](https://github.com/tunglam14)
