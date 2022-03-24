package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func (c DBConnect) FillDB() {
	log.Println("DB:FillDB: fill table")

	file, err := ioutil.ReadFile("db.json")

	if err != nil {
		log.Println("DB:FillDB: error read config %w", err)
		return
	}

	data := ResidentilaComplex{
		Homes: []Home{},
	}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		log.Println("DB:FillDB: error convert config %w", err)
		return
	}

	err = c.CreateResidentilaComplex(data.Name, data.Description)

	if err != nil {
		log.Println("DB:FillDB: error create complex %w", err)
		return
	}

	for _, home := range data.Homes {
		err := c.CreateHome(
			home.Name,
			home.Description,
			home.ResidentialComplex,
			home.Floors,
			home.FirstResidentialFloor,
			home.Apartments,
			home.FirstApartment,
			home.Entrances,
		)
		if err != nil {
			log.Println("DB:FillDB: error create homes %w", err)
		}
	}
}
