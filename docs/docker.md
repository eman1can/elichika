# Docker Deployment

> **Advanced usage:** This guide assumes you are familiar with Docker and Docker Compose.
> See [Advanced Usage](advanced_usage.md) for general expectations around self-hosting.

The Docker image provides a lightweight, cross-platform deployment using a `golang:alpine` base image.
Install [Docker and Docker Compose](https://docs.docker.com/engine/install/) before proceeding.

---

## Deploying

Navigate to the `docker` directory in the repository and run:

```bash
docker compose build
docker compose up -d
```

The container exposes port `8080`. Once running, the Admin WebUI is accessible at
`http://<server_address>:8080/webui/admin`.

### Deploying from a Specific Branch

```bash
# Build image from a specific branch
docker build --build-arg BRANCH=<BRANCH_NAME> -t llas .

# Start the container
docker compose up -d
```

---

## Configuration

All config options are stored in `data/config.json`, which is created on first startup. The recommended way to edit
settings is through the [Admin WebUI](webui.md#admin-webui).

The `data/` directory (including `userdata.db` and `serverstate.db`) should be mounted as a volume so your data persists
across container rebuilds. Check the included [`docker-compose.yml`](../docker/docker-compose.yml) for the volume
configuration.

---

## Data Persistence

User data is stored in `userdata.db` inside the container at `/elichika/userdata.db`. **Always back this file up before
rebuilding the container image.**

You can also export accounts via the [WebUI](webui.md) as an alternative backup method.

---

## Updating

### Option 1: Rebuild the Image

This produces the cleanest update. Back up your data first, then:

```bash
# Copy user data out of the container
docker container cp llas:/elichika/userdata.db .

# Rebuild
docker compose down
docker rmi llas:latest
docker compose build
docker compose up -d

# Restore user data
docker container cp userdata.db llas:/elichika

# Restart to pick up the restored data
docker container restart llas
```

### Option 2: In-place Update

Faster, but less clean than a full rebuild:

```bash
docker container exec -it llas bash /root/update_elichika
docker container restart llas
```
