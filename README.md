### Go + Postgres + Nextjs Starter

### Getting started

https://dev.to/francescoxx/go-typescript-full-stack-web-app-with-nextjs-postgresql-and-docker-42ln

Go Docs:
[Main Docs](https://go.dev/doc/)
[Effective Go Docs](https://go.dev/doc/effective_go)

#### Init Postgres in Docker

1. Run:

```bash
docker compose up -d
```

Retrieve image and start container
`-d` flag tells docker to run in a detached state (freeing up terminal)

2. Confirm the container is running:

```bash
docker ps -a
```

3. Step into db container to confirm no relations

```bash
docker exec -it db psql -U postgres
```

Run `\l` and `\dt` cmds
