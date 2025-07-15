# repl-template

a minimal go REPL scaffold with dynamic build-time config.
This is a template for making CLIs. You can rename it at build time to something like 'myrepl' 😎 etc., and it will have the corresponding prompt + exit message.

⸻

## features

- define REPL name at build time (e.g. `make myrepl`)
- injects prompt + messages into generated config.go
- runs as interactive REPL (`go run .`) or binary (`./myrepl`)

## Makefile commands:

- `make <name>` → build binary
- `make config NAME=<name>` → generate config only
- `make clean` → remove binary + config

⸻

## usage

build and run as binary:

```
make myrepl
./bin/myrepl
```

run without building (just generate config)

```
make config NAME=myrepl
go run .
```

cleanup:

```
make clean
```

⸻

structure

```
.
├── bin/ # built binaries (auto-created)
├── config.go # generated config (name, prompt, exit msg)
├── gen/
│ └── genconfig.go # generates config.go from name
├── main.go # REPL entrypoint
├── Makefile
```
