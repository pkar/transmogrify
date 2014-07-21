# Transmogrify

A simple text encoder that transforms an input stream of text to an encoded version.
The encoding is based on a keyboard layout that is modified through commands. See transmogrify.go

```bash
$ export GOPATH=`pwd`
$ go get github.com/pkar/transmogrify

$ go run src/github.com/pkar/transmogrify/cmd/main.go -h
Usage of /var/folders/r_/h7h04myn1277s7dp__w7gxfj1_0y00/T/go-build783350671/command-line-arguments/_obj/exe/main:
  -cmds="": transform commands H,V,int... If -trans provided, this is ignored
  -trans="": path to file with transform commands H,V,int... If provided -cmds is ignored
  -text="STDIN": path to text to encode, default is stdin

$ # run with either an input file and the -text option, there is also a -trans
$ go run src/github.com/pkar/transmogrify/cmd/main.go -cmds=H,-1,V,4,H -text=src/github.com/pkar/transmogrify/cmd/text.txt

$ # or with a piped in stream
$ tail -f /var/log/syslog | go run src/github.com/pkar/transmogrify/cmd/main.go -cmds=H,-1,V,4,H

$ # there are also two binaries in ./bin mac and linux versions
./src/github.com/pkar/transmogrify/bin/transmogrify_mac -cmds=H,-1,V,4,H -text=src/github.com/pkar/transmogrify/cmd/text.txt
./src/github.com/pkar/transmogrify/bin/transmogrify_linux -cmds=H,-1,V,4,H -text=src/github.com/pkar/transmogrify/cmd/text.txt
```
