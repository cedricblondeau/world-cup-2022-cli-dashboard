![screenshot](https://raw.githubusercontent.com/cedricblondeau/world-cup-2022-cli-dashboard/main/screenshot.png)

# World Cup 2022 CLI Dashboard [![lint](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/workflows/lint/badge.svg)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/actions) [![test](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/workflows/test/badge.svg)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/actions) [![release](https://badgen.net/github/release/cedricblondeau/world-cup-2022-cli-dashboard)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/releases)

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/kinda-sfw.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## Features

- ⚽ Live matches (goals, bookings, substitutions)
- 🗒️ Team lineups
- 📅 Scheduled and past matches
- 📒 Standings & bracket
- 📊 Player stats (goals, yellow cards, red cards)

## Install

### Method 1: Homebrew 🍺

Install:
```bash
brew tap cedricblondeau/cedricblondeau
brew install world-cup-2022-cli-dashboard
```

Run:
```bash
world-cup-2022-cli-dashboard
```

### Method 2: Docker 🐳

Build from the `main` branch:
```bash
docker build --no-cache https://github.com/cedricblondeau/world-cup-2022-cli-dashboard.git#main -t world-cup-2022-cli-dashboard
```

Run it:
```bash
docker run -ti -e TZ=America/Toronto world-cup-2022-cli-dashboard
```

Replace `America/Toronto` with the desired timezone.

### Method 3: Go package

Requirements:
- Go 1.19+ (with `$PATH` properly set up)
- Git

```bash
go install github.com/cedricblondeau/world-cup-2022-cli-dashboard@latest
world-cup-2022-cli-dashboard
```

### Method 4: Pre-compiled binaries

Pre-compiled binaries are available on the [releases page](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/releases).

## Data

Data can be sourced from:
1. [worldcupjson.net](https://worldcupjson.net/)
2. [football-data.org](https://www.football-data.org/)

By default, the dashboard uses `worldcupjson.net` - which is an awesome free and open source project but with limited availability and accuracy guarantees.

To use `football-data.org` instead, you'll need to [register](https://www.football-data.org/client/register) and get an API token (it's easy and free). Then, start the dashboard with an env variable:
```bash
FOOTBALLDATA_API_TOKEN=my_fake_token world-cup-2022-cli-dashboard
```

Or with Docker:
```bash
docker run -ti -e TZ=America/Toronto -e FOOTBALLDATA_API_TOKEN=my_fake_token world-cup-2022-cli-dashboard
```

Note that the _free_ `football-data.org` plan comes with less features than `worldcupjson.net`.

|              | worldcupjson.net | football-data.org |
|--------------|:----------------:|:-----------------:|
| Live scores  |         ✅        |         ✅         |
| Schedule     |         ✅        |         ✅         |
| Standings    |         ✅        |         ✅         |
| Bracket      |         ✅        |         ✅         |
| Lineups      |         ✅        |         ❌         |
| Goal scorers |         ✅        |         ❌         |

The data source gets polled every minute.

## UI

UI is powered by [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

For optimal results, it's recommended to use a terminal with:
- True Color (24-bit) support;
- at least 102 columns and 35 rows.
## LICENSE

MIT
