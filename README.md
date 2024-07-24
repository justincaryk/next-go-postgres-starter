## Go + Postgres + Nextjs Starter

### Requirements

- [Install Docker](https://docs.docker.com/engine/install/)
- [Install go](https://go.dev/doc/install) \*This app requires v1.22.4
- [Install air](https://github.com/air-verse/air)
- [Install node](https://nodejs.org/en/download/prebuilt-installer/current)
- [Install node version manager (nvm)](https://github.com/nvm-sh/nvm) \*Optional

### Documentation

Go Docs:

- [Main Docs](https://go.dev/doc/)
- [Effective Go Docs](https://go.dev/doc/effective_go)

The project architecture is cribbed from [this dev.to article](https://dev.to/francescoxx/go-typescript-full-stack-web-app-with-nextjs-postgresql-and-docker-42ln).

### Getting started

1. Install frontend dependencies (from root):

```bash
cd frontend
yarn
```

2. Install go packages (from root):

```bash
cd backend
go install github.com/air-verse/air@latest
go mod tidy
```

3. Docker Profiles have been set up for the frontend and backend directories. The db is always configured to run. Choose which best suits your development requirements:

- To spin up nextapp and db only: `docker compose --profile nextapp up`
- To spin up goapp and db only: `docker compose --profile goapp up`
- To spin up the whole app: `docker compose --profile nextapp --profile goapp up`

If you want to run these in a detached state, you can append the `-d` flag to the end of the commands above, which will free up the terminal window.

4. Spin up your working directory in a separate terminal

- For frontend dev: `cd frontend && yarn dev`
- For backend dev: `cd backend && air`

### Testing

```bash
cd backend && go test -v
cd frontend && yarn test
```

### TODO

1. go: env variables 