# gojuon

gojuon is a little tool help Japanese amateur learn gojunon

## install 

```bash
go get -u github.com/ZhangYet/gojuon
```

## usage

```bash
$ gojuon 
NAME:
   gojuon - help japanese amateur learn gojuon.

USAGE:
   gojuon [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     reference, r  print gojuon list
     gen, g        print gojuon test
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## examples

### list gana of s line

List lines of hiragana, katagana, and rome.

```bash
$ gojuon r s

sa: さ さ	shi: し し	su: す す	se: せ せ	so: そ そ
```

### gen test of s and m line

Generate random hiragana, katagana, and rome list.

````bash
$ gojuon g --type roma s m

shi, mu, so, sa, mo, su, se, ma, me, mi

$ gojuon g --type hira s m

し, む, そ, さ, も, す, せ, ま, め, み

$ gojuon g --type kata s m

シ, ム, ソ, サ, モ, ス, セ, マ, メ, ミ
````