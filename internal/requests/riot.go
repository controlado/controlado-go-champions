package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/controlado/league-champions-go/pkg/league"
)

const (
	baseURL = "https://raw.communitydragon.org" // Dados crus da Riot.
)

// GetUnits retorna as unidades (campeões e skins) de forma crua.
func GetUnits(region string) (units map[string]league.Unit, err error) {
	endpoint := fmt.Sprintf("/latest/plugins/rcp-be-lol-game-data/global/%s/v1/skins.json", region)
	response, err := http.Get(baseURL + endpoint)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("requesting champions: %v", err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&units)
	if err != nil {
		return nil, fmt.Errorf("json decoding: %v", err)
	}
	return units, nil
}
