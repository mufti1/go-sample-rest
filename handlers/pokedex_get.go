package handlers

import (
	"net/http"

	"github.com/mufti1/simple-rest/lib/responsewriter"
)

// GetPokedex get pokedex from mongodb
func (h *Handler) GetPokedex(w http.ResponseWriter, r *http.Request) {
	rf := &responsewriter.ResponseFormat{}
	data, err := h.mongodb.GetPokedex()
	if len(data) <= 0 {
		rf.ResponseOK(404, data, w)
		return
	}
	if err != nil {
		rf.ResponseNOK(501, err, w)
		return
	}
	rf.ResponseOK(200, data, w)
	return
}
