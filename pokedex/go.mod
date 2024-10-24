module github.com/Victor-AvestaAI/pokedex

go 1.23.2

replace github.com/Victor-AvestaAI/pokedex/pokeapi v0.0.0 => ../pokeapi

replace github.com/Victor-AvestaAI/pokedex/pokecache v0.0.0 => ../pokecache

require (
	github.com/Victor-AvestaAI/pokedex/pokeapi v0.0.0
	github.com/Victor-AvestaAI/pokedex/pokecache v0.0.0
)
