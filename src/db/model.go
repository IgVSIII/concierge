package db

type Resident struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
	Apartment   int
	Home        string
	Status      string
}

type ResidentilaComplex struct {
	Name        string `gorm:"primary_key"`
	Description string
	Homes       []Home `gorm:"foreignKey:ResidentialComplex"`
}

type Home struct {
	Name                  string `gorm:"primary_key"`
	Description           string
	ResidentialComplex    string
	Floors                int
	FirstResidentialFloor int
	Apartments            int
	Entrances             int
	//Residents             []Resident `gorm:"foreignKey:Home"`
}
