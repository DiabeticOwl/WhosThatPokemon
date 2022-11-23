# Resources generation

This package is used for embedding static information required for the main game
to run. This is meant to be run once while building the application. The assets
description and method of embedding are described here.

## Global Pokemon information

Pokemon data is extracted from the [PokéApi](https://pokeapi.co/). After the
recollection everything is stored in json files for further use.

This is a list of what data is extracted:

* [Species](https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_(species))
* [Generations](https://bulbapedia.bulbagarden.net/wiki/Generation)
* [Types](https://bulbapedia.bulbagarden.net/wiki/Type)
* [Games versions](https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9mon_games)

## Embedding process

Ebitengine allows us to embed files by converting them into slices of bytes and
then casting them inside of the game as needed. The conversion is done using the
[file2byteslice](https://github.com/hajimehoshi/file2byteslice/) Go tool.

In order to use this tool a file called "[generate.go](https://github.com/DiabeticOwl/WhosThatPokemon/blob/main/resources/generate/generate.go)"
is created with the purpose of being used with [Go generate](https://go.dev/blog/generate)
command. The contents to be generated are the json files previously described
and any other resources required to be embed for the game execution.

## Building the game

In order to build the game for the first time (without any of the files generated)
or for updating the resources the following steps should be followed:

* Run this package with `go run` for creating the json and generate files.
* Execute the `go generate` command for storing the slices of bytes in their
respective packages.
* Build the main game.
