# Pokedex CLI

A command-line Pokedex interface using the PokeAPI.

## Features

- Explore Pokemon locations
- Catch Pokemon
- View your caught Pokemon
- Inspect Pokemon details
- Navigation through location areas

## Installation

```bash
go install github.com/fatkungfu/pokedexcli@latest
```

## Usage

Start the application:

```bash
pokedexcli
```

### Available Commands

- `help` - Displays a help message
- `map` - Displays the next 20 location areas
- `mapb` - Displays the previous 20 location areas
- `explore <location_name>` - Explore a location
- `catch <pokemon_name>` - Attempt to catch a pokemon
- `inspect <pokemon_name>` - View information about a caught pokemon
- `pokedex` - View all the Pokemon you have caught
- `exit` - Exit the Pokedex

## Requirements

- Go 1.21 or higher