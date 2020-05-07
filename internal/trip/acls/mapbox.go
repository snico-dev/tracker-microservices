package acls

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type mapBoxACL struct {
	token string
	url   string
}

//NewMapBoxACL contructor
func NewMapBoxACL() IMapACL {
	return &mapBoxACL{
		token: "pk.eyJ1IjoibHVjYXNkYW50YXNuZCIsImEiOiJjazQ4djNxd2IwcDBoM2tvY2k2bG96aHAxIn0.0M8vijA_74sXTv4VbCUabA",
		url:   "https://api.mapbox.com/geocoding/v5",
	}
}

func toString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func (mp *mapBoxACL) GetAddressName(latitude float64, longitude float64) (string, error) {
	endpoint := mp.url + "/mapbox.places/" + toString(longitude) + "," + toString(latitude) + ".json?access_token=" + mp.token
	resp, err := http.Get(endpoint)

	if err != nil {
		return "Endereço não encontrado", err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var mapboxPlacesDTO MapboxPlacesDTO
	json.Unmarshal(bodyBytes, &mapboxPlacesDTO)

	if mapboxPlacesDTO.Type == "" || len(mapboxPlacesDTO.Features) == 0 {
		return "Endereço não encontrado", nil
	}

	addressName := mapboxPlacesDTO.Features[0].PlaceName

	return addressName, err
}
