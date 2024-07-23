### Go + Postgres + Nextjs Starter

### Documentation


Go Docs:
[Main Docs](https://go.dev/doc/)
[Effective Go Docs](https://go.dev/doc/effective_go)

The project architecture is cribbed from [this dev.to article](https://dev.to/francescoxx/go-typescript-full-stack-web-app-with-nextjs-postgresql-and-docker-42ln).

### Requirements

[Install Docker](https://docs.docker.com/engine/install/)
[Install go](https://go.dev/doc/install) \*This app requires v1.22.4
[Install node](https://nodejs.org/en/download/prebuilt-installer/current)
[Install node version manager (nvm)](https://github.com/nvm-sh/nvm) *Optional

#### Getting started

1. Install frontend dependencies (from root):

```bash
cd frontend
yarn
```

2. Install go packages (from root):

```bash
cd backend
go mod tidy
```

3. Profiles have been set up for Nextjs and Go. The db is always configured to run. Choose which best suits your development requirements:

* To spin up nextapp and db only: `docker compose --profile nextapp up`
* To spin up goapp and db only: `docker compose --profile goapp up`
* To spin up the whole app: `docker compose --profile nextapp --profile goapp up`

If you want to run these in a detached state, you can append the `-d` flag to the end of the commands above, which will free up the terminal window.

4. Spin up your working directory in a separate terminal

* For frontend dev: `cd frontend && yarn dev`
* For backend dev: **TODO (see reflex todo below)**

#### Init Postgres in Docker


2. Confirm the container is running:

```bash
docker ps -a
```

### Init Next frontend

1. In a separate terminal run:

```bash
yarn
yarn dev
```

2. See the `README.md` in `[root]/frontend` and follow any subsequent instructions

3. Open a `http://localhost:3000/` in your browser of choice.


### TODO

1. Install reflex for go:

Run `go install github.com/cespare/reflex@latest`

Create a `reflex.conf` file:
```
# Watch .go files and run docker-compose up when changes are detected
-r '\.go$' docker-compose up --build
```

Run `reflex -c backend/reflex.conf`


2. Install test suite for backend (go)

3. Install test suite for frontend (nextjs)