package common

type Resident struct {
	ID          string
	Name        string
	Description string
	Apartment   int
	Home        string
	Status      string
}

type ResidentilaComplex struct {
	Name        string
	Description string
}

type Home struct {
	Name                  string
	Description           string
	ResidentialComplex    string
	Floors                int
	FirstResidentialFloor int
	Apartments            int
	FirstApartment        int
	Entrances             int
}
