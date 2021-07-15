# Go server for getting the boot duration

This project implements a go server that returns the boot duration using
`systemd-analyze`.

## Running

To launch the server, simply run:

```bash
./ucr-e1
```

Get the version:

```bash
> curl -Ss localhost:8080/version
v0.0.1
```

Get the boot duration:

```bash
> curl -Ss localhost:8080/duration
3min 34.249s
```

## Building

First, clone the repository:

```bash
git clone https://github.com/gc-plp/ex1
```

And build it:

```bash
cd ex1 && go build
```
