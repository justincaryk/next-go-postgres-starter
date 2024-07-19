### Go + Postgres + Nextjs Starter

### Getting started

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

