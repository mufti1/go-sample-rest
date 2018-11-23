package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mufti1/simple-rest/lib/responsewriter"
)

type pokedexInput struct {
	Name    string   `bson:"name" json:"name"`
	Height  float64  `bson:"height" json:"height"`
	Ability []string `bson:"ability" json:"ability"`
}

// InsertPokedex insert pokedex into databases
func (h *Handler) InsertPokedex(w http.ResponseWriter, r *http.Request) {
	rf := responsewriter.ResponseFormat{}

	decoder := json.NewDecoder(r.Body)
	input := pokedexInput{}

	err := decoder.Decode(&input)
	if err != nil {
		rf.ResponseNOK(http.StatusBadRequest, "Bad Request", w)
		return
	}

	pokedex, err := h.mongodb.InsertPokedex(input.Name, input.Height, input.Ability)

	if err != nil {
		rf.ResponseNOK(501, err, w)
		return
	}

	rf.ResponseOK(http.StatusOK, pokedex, w)
	return
}
