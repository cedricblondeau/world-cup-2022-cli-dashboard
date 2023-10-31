![screenshot](https://raw.githubusercontent.com/cedricblondeau/world-cup-2022-cli-dashboard/main/screenshot.png)

# World Cup 2022 CLI Dashboard [![lint](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/workflows/lint/badge.svg)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/actions) [![test](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/workflows/test/badge.svg)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/actions) [![release](https://badgen.net/github/release/cedricblondeau/world-cup-2022-cli-dashboard)](https://github.com/cedricblondeau/world-cup-2022-cli-dashboard/releases)

[![forthebadge](https://raw.githubusercontent.com/BraveUX/for-the-badge/master/src/images/badges//built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://raw.githubusercontent.com/BraveUX/for-the-badge/master/src/images/badges/kinda-sfw.svg)](https://forthebadge.com) [![forthebadge](https://raw.githubusercontent.com/BraveUX/for-the-badge/master/src/images/badges/made-with-go.svg)](https://forthebadge.com)

Featured in 📹 [Charm in the Wild | December 2022](https://www.youtube.com/watch?v=XuTb7Ao27w4&t=252s) ❤️.

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

## UI

UI is powered by [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss).

For optimal results, it's recommended to use a terminal with:
- True Color (24-bit) support;
- at least 160 columns and 50 rows.

## LICENSE

MIT
