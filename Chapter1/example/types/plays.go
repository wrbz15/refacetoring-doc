package types

import (
	"encoding/json"
	"io/ioutil"
)

type Plays map[string]Play
type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func LoadPlays(path string) (Plays, error) {
	playData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	plays := make(map[string]Play)
	err = json.Unmarshal(playData, &plays)
	if err != nil {
		return nil, err
	}
	return plays, nil
}
