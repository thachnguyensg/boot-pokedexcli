package pokeapi

type NamedResp struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}
