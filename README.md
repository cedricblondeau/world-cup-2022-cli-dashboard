![screenshot](https://raw.githubusercontent.com/cedricblondeau/world-cup-2022-cli-dashboard/main/demo.png)

# World Cup 2022 CLI Dashboard

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/kinda-sfw.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Install

#### Method 1: Docker

Build the image:
```bash
git clone git@github.com:cedricblondeau/world-cup-2022-cli-dashboard.git
cd world-cup-2022-cli-dashboard
docker build . -t world-cup-2022-cli-dashboard
```

Run a container:
```
docker run -ti -e TZ=America/Toronto world-cup-2022-cli-dashboard
```

Replace `America/Toronto` with the desired timezone.

#### Method 2: Go package

Install:
```bash
go install github.com/cedricblondeau/world-cup-2022-cli-dashboard@latest
```

Run (assumes `$PATH` is properly set up):
```bash
world-cup-2022-cli-dashboard
```

## How does it work?

UI is powered by [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

Data is sourced from [worldcupjson.net](https://worldcupjson.net/). Matches get updated every minute.

## LICENSE

MIT
