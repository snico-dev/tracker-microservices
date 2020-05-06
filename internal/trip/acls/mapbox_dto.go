package acls

//MapboxPlacesDTO struct api response
type MapboxPlacesDTO struct {
	Type     string    `json:"type"`
	Query    []float64 `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			Accuracy  string `json:"accuracy"`
			ShortCode string `json:"short_code"`
			Wikidata  string `json:"wikidata"`
		} `json:"properties"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			Text      string `json:"text"`
			Wikidata  string `json:"wikidata"`
			ShortCode string `json:"short_code"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}
