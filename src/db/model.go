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
	Name        string `gorm:"primary_key" ,json:"name" ,yaml:"name"`
	Description string `json:"description" ,yaml:"description"`
	Homes       []Home `gorm:"foreignKey:ResidentialComplex" ,json:"homes" ,yaml:"homes"`
}

type Home struct {
	Name                  string `gorm:"primary_key" ,json:"name" ,yaml:"name"`
	Description           string `json:"description" ,yaml:"description"`
	ResidentialComplex    string `json:"residentilaComplex" ,yaml:"residentilaComplex"`
	Floors                int    `json:"floors" ,yaml:"floors"`
	FirstResidentialFloor int    `json:"firstResidentialFloor" ,yaml:"firstResidentialFloor"`
	Apartments            int    `json:"apartments" ,yaml:"apartments"`
	FirstApartment        int    `json:"firstApartment" ,yaml:"firstApartment"`
	Entrances             int    `json:"entrances" ,yaml:"entrances"`
	//Residents             []Resident `gorm:"foreignKey:Home"`
}
