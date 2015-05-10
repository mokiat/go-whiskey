# Whiskey Game Engine in Go

![](https://travis-ci.org/momchil-atanasov/go-whiskey.svg?branch=master)

A simple game engine written in Go. It will target both 2D and 3D use cases.

## User's Guide

TODO


## Developer's Guide

You should clone the repository according to Go's standards.

```bash
go get github.com/momchil-atanasov/go-whiskey
```

The `counterfeiter` tool is used to generate fakes/stubs of interfaces. Use the following command to download it.

```bash
go get github.com/maxbrunsfeld/counterfeiter
```

Once you have the tool downloaded, should you need to, you can regenerate all the fakes with the following command.

```bash
go generate ./...
```

The `ginkgo` tool is used to run BDD-style unit tests. You can check the [Ginkgo](https://github.com/onsi/ginkgo) repository for detailed documentation. Run the following command to download the tool.

```bash
go get github.com/onsi/ginkgo/ginkgo
```

Once you have `ginkgo` installed, you can run all the unit tests with the following command.

```bash
ginkgo -r --race
```
