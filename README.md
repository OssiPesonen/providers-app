# Project providers.app

This is a hobby application that has taken it's inspiration from the lights you see next to an exam room (green, yellow, red). The idea is to provide a platform for service providers that take on clients to let people know if they have capacity, if they might have capacity (can be contacted) or they are fully booked. Users should be able to list service providers in a specific location, for a specific service.

> [!WARNING]  
> Given that I am someone who likes to do prototyping in the beginning, while doing upfront design (and also to figure out how some things actually work), I am continuously refactoring this app while incrementally adding features. Things might change drastically. Things might also still be in a broken state, as I have not gone through edge cases.

## Motivation

I wanted to find something to build with Go and this is a problem I faced with my son who needed the services of a speech therapist. The healthcare system in our country has an antiquated process that requires parents to call through a list of accepted service providers (which is in a PDF file) to see if they have capacity to offer their services. During the span of 18 months we called through that list four times until someone actually said yes.

The idea here is service providers could register to this platform and simply let everyone know if they have capacity to take on new clients, patients or customers. People could pick the city they live in and see a list of service providers for their current need and if they have capacity, after which they could simply contact one that is available.

## Development 

### Getting Started

Start off by copying .env.example to .env and filling out the values. If you are running this app locally, no changes are necessary. 

If you wish to compile the protocol buffer files for the UI, you need to follow the installation guide at [grpc-web](https://github.com/grpc/grpc-web) repo.

You can find helper commands below that are run via `make`. To run the app in watch mode for development, you need to install [air-verse/air](https://github.com/air-verse/air)

    go install github.com/air-verse/air@latest

### Makefile

Here are helpful utility commands by using make, which assist you in running the app in development mode, building it, testing it etc.

If you need `make` for Windows, the recommendation is to install [Chocolatey](https://chocolatey.org/install) and then run  `choco install make` in your CLI.

Build the application binary (Unix)
```bash
make build-api
```

Build the application binary (Windows)
```bash
make build-api-win
```

Compile protocol buffer messages for client (UI)
```bash
make gen-proto-client
```

Compile protocol buffer messages for API (Go)
```bash
make gen-proto-api
```

Run the application
```bash
make dev-api
```
Start database and grpc proxy containers
```bash
make docker-up
```

Shutdown containers
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest-api
```

Live reload the application (Unix):
```bash
make watch-api
```

Live reload the application (Windows):
```bash
make watch-api-win
```

Run the test suite:
```bash
make test-api
```

Clean up binary from the last build:
```bash
make clean-api
```
