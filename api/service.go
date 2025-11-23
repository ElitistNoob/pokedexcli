package api

func (c *Client) GetLocations(endpoint *string) (Locations, error) {
	target := BaseUrl + LocationArea
	if endpoint != nil {
		target = *endpoint
	}

	var locations Locations
	if err := c.fetchJSON(target, &locations); err != nil {
		return Locations{}, err
	}

	return locations, nil
}

func (c *Client) ExploreArea(area string) (*AreaPokemon, error) {
	target := BaseUrl + LocationArea + area

	var areaResponse AreaResponse
	if err := c.fetchJSON(target, &areaResponse); err != nil {
		return nil, err
	}

	var pokemons []Pokemon
	for _, p := range areaResponse.PokemonEncounters {
		pokemons = append(pokemons, p.Pokemon)
	}

	return &AreaPokemon{
		Count:   len(pokemons),
		Results: pokemons,
	}, nil
}

func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	target := BaseUrl + PokemonUrl + name

	var pokemon Pokemon
	if err := c.fetchJSON(target, &pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}
