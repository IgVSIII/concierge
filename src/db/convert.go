package db

import "concierge-bot/src/common"

func homesConvert(homes []Home) []common.Home {
	convert := []common.Home{}
	for _, value := range homes {
		convert = append(convert, homeConvert(value))
	}
	return convert
}

func homeConvert(home Home) common.Home {
	return common.Home{
		Name:                  home.Name,
		Description:           home.Description,
		ResidentialComplex:    home.ResidentialComplex,
		Floors:                home.Floors,
		FirstResidentialFloor: home.FirstResidentialFloor,
		Apartments:            home.Apartments,
		Entrances:             home.Entrances,
	}
}

func residentsConvert(residents []Resident) []common.Resident {
	convert := []common.Resident{}
	for _, value := range residents {
		convert = append(convert, residentConvert(value))
	}
	return convert
}

func residentConvert(resident Resident) common.Resident {
	return common.Resident{
		ID:          resident.ID,
		Name:        resident.Name,
		Description: resident.Description,
		Apartment:   resident.Apartment,
		Home:        resident.Home,
		Status:      resident.Status,
	}
}

func residentilaComplexsConvert(complexs []ResidentilaComplex) []common.ResidentilaComplex {
	convert := []common.ResidentilaComplex{}
	for _, value := range complexs {
		convert = append(convert, residentilaComplexConvert(value))
	}
	return convert
}

func residentilaComplexConvert(complex ResidentilaComplex) common.ResidentilaComplex {
	return common.ResidentilaComplex{
		Name:        complex.Name,
		Description: complex.Description,
	}
}
