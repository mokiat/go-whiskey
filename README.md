# Whiskey Game Engine in Go

[![](https://travis-ci.org/mokiat/go-whiskey.svg?branch=master)](https://travis-ci.org/mokiat/go-whiskey)

```
_______ _______      _     _ __   __ ___ _______ ___   _ _______ __   __
|       |       |    | | _ | |  | |  |   |       |   | | |       |  | |  |
|    ___|   _   |____| || || |  |_|  |   |  _____|   |_| |    ___|  |_|  |
|   | __|  | |  |____|       |       |   | |_____|      _|   |___|       |
|   ||  |  |_|  |    |       |       |   |_____  |     |_|    ___|_     _|
|   |_| |       |    |   _   |   _   |   |_____| |    _  |   |___  |   |  
|_______|_______|    |__| |__|__| |__|___|_______|___| |_|_______| |___|  

```
_(ASCII generated via [http://patorjk.com/software/taag/](http://patorjk.com/software/taag/))_

A simple game engine (framework) written in Go.

The idea behind this library is to provide a set of APIs through which one could develop a game. This should allow developers to abstract themselves from the complexities of low-level APIs (e.g. OpenGL) or algorithms (e.g. A*). At this point in time, there is no plan to provide any editors or tools that would allow one to easily model and script a whole game, though it should not be difficult for anyone to implement them on top.

**Note: This is in active development so expect changes to the APIs and backward incompatibilities.**

## Demo

One can see how some of the APIs can be used by looking at the **[go-whiskey-demo](https://github.com/mokiat/go-whiskey-demo)** example application.

## Plugins

Though most of the APIs provided in this library can be used stand-alone, some APIs are merely interfaces that don't have an actual implementation. One needs to include a specific implementation libraries for their exact use-case.

* **[go-whiskey-android](https://github.com/mokiat/go-whiskey-android)** - Contains API implementations for Android

## User's Guide

This repository does not provide any executables so the most one needs is to import it as a dependency inside their project and consume the desired APIs.

Still, you can download the library through the following command.

```bash
go get github.com/mokiat/go-whiskey
```

## Developer's Guide

This project uses the **[gostub](https://github.com/mokiat/gostub)** tool to generate mocks of interfaces, which allows for more thorough testing.

```bash
go get github.com/mokiat/gostub
```

The interface mock implementations are shipped with the code. However, if you ever need to regenerate those, you can do it with the following command.

```bash
go generate ./...
```

The `ginkgo` tool is used to write and run all BDD-style unit tests. You can check the **[Ginkgo](https://github.com/onsi/ginkgo)** repository for detailed documentation.

```bash
go get github.com/onsi/ginkgo/ginkgo
```

You can run all the tests with the following command.

```bash
ginkgo -r --race
```
