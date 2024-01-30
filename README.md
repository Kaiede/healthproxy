# Healthproxy

A very simple container tool meant to expose out an HTTP style health check endpoint for services that expose UDP publically, but have some sort of TCP admin endpoint. The goal is to avoid exposing the admin port to the world, while sitll being able to use it to monitor server health.

Game servers that support RCON are good examples, such as Ark, Astroneer, Palworld, etc. These provide RCON functionality over TCP, and in principle you could use this port with a tool like Uptime Robot or Uptime Kuma. However, exposing this unencrypted admin port isn't ideal and opens up many risks.

It's highly recommended to hide these behind some sort of reverse proxy like nginx proxy manager so that multiple services can be checked independently.

## Example Docker Compose

```yaml
version: '3.7'

services:
  healthcheck:
    image: kaiede/healthproxy:latest
    restart: always
    environment:
      ADDRESS: "gameserver:25575"
    ports:
      - 8086:8086/tcp

  gameserver:
    # Your Service Settings Here
    expose:
      # Expose the RCON port only to other services in this file
      - 25575
    ports:
      # Expose the game port publically
      - 8211:8211/udp
```
