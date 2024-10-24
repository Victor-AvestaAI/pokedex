module example.com/username/pokedex

go 1.23.2

replace example.com/username/pokeapi v0.0.0 => ../pokeapi

require example.com/username/pokeapi v0.0.0

replace example.com/username/pokecache v0.0.0 => ../pokecache

require example.com/username/pokecache v0.0.0
