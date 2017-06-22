gen-fakecert
==========

Automatically generates self signed certificate.

# Download

- Release page: [https://github.com/noritama/gen-fakecert/releases](https://github.com/noritama/gen-fakecert/releases)
- Download binary file(version 0.0.3): [https://github.com/noritama/gen-fakecert/releases/download/v0.0.3/gen-fakecert](https://github.com/noritama/gen-fakecert/releases/download/v0.0.3/gen-fakecert)

# install

```sh
$ go get github.com/noritama/gen-fakecert
```

# Use

```sh
$ gen-fakecert -key /tmp/server.key -crt /tmp/server.crt -country JP -organization Ex
Output private-key file: /tmp/server.key
Output cert file: /tmp/server.crt

# Options

```
$ gen-fakecert -h

Usage of gen-fakecert:
   gen-fakecert [OPTIONS] ARGS...

Options  -out="/Users/noritama/repository/github/gen-fakecert/keypair_gen.go": output file path
  -pkgname="main": package name
  
```

# go generate

```go
//go:generate gen-fakecert -pkgname gen -out ./gen/keypair_gen.go
```

> go1.4

# build

```sh
$ make # => ./gen-fakecert
```
