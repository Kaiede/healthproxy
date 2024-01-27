# Healthproxy

A very simple container tool meant to expose out an HTTP style health check endpoint for services that expose UDP publically, but TCP internally.

Example would be game servers like Ark, Astroneer, Palworld, etc. These provide RCON functionality over TCP, but it's not recommended to expose these ports openly. But you might want to use something like Uptime Robot to check uptime of the game server. 

It's recommended to hide these behind some sort of reverse proxy like nginx proxy manager so that multiple services can be checked independently.

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