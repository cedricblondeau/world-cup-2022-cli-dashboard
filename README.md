![screenshot](https://raw.githubusercontent.com/cedricblondeau/world-cup-2022-cli-dashboard/main/screenshot.png)

# World Cup 2022 CLI Dashboard

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/kinda-sfw.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Install

### Method 1: Docker 🐳

Requirements:
- Docker
- Git

```bash
docker build https://github.com/cedricblondeau/world-cup-2022-cli-dashboard.git#main -t world-cup-2022-cli-dashboard && \
docker run -ti -e TZ=America/Toronto world-cup-2022-cli-dashboard
```

Replace `America/Toronto` with the desired timezone.

### Method 2: Go package

Requirements:
- Go 1.19+ (with `$PATH` properly set up)
- Git

```bash
go install github.com/cedricblondeau/world-cup-2022-cli-dashboard@latest
world-cup-2022-cli-dashboard
```

## How does it work?

UI is powered by [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

Data is sourced from [worldcupjson.net](https://worldcupjson.net/). Matches get updated every minute.

## LICENSE

MIT
