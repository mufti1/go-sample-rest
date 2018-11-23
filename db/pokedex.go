package db

import (
	"github.com/mufti1/simple-rest/models"
	"gopkg.in/mgo.v2/bson"
)

const pokedexCollection = "pokedex"

// GetPokedex return the pokedex from mongo
func (m *MongoDB) GetPokedex() ([]models.Pokedex, error) {
	result := []models.Pokedex{}
	err := m.session.session.DB(m.dbname).C(pokedexCollection).Find((bson.M{})).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// InsertPokedex implement function to insert data to mongodb
func (m *MongoDB) InsertPokedex(name string, height float64, ability []string) (*models.Pokedex, error) {
	pokedex := &models.Pokedex{
		ID:        bson.NewObjectId(),
		Name:      name,
		Height:    height,
		Abilities: ability,
	}

	err := m.session.session.DB(m.dbname).C(pokedexCollection).Insert(pokedex)

	if err != nil {
		return nil, err
	}

	return pokedex, nil
}
