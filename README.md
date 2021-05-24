# Dash - Simplistic Dashboard

### Configuration

```
dash -config path/to/config.yml
```

default config file location is ~/.dash.yml

```
port: ":3200"
server: localhost
seed: 238429834928374

users:
  - name: Maksim Kozhukh
    groups:
      - admin
    key: alu483hr98237

commands:
  - id: free-space
    name: "Free disk space"
    exec:
      - docker image prune
    groups:
      - admin

  - id: restart-db
    name: "Restart Docker"
    danger: true
    details: |
      Will result in small downtime ( ~15 seconds ) for ALL services, use with care
    exec:
      - service docker restart
    groups:
      - admin

info:
  - id: free-space
    name: free disk space
    exec: "df -h /dev/nvme0n1p2 | tail -1 | awk '{print $4}'"
```

