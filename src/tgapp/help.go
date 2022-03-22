package tgapp

import "concierge-bot/src/common"

type DBConnector interface {
	UpdateResidentAddDescription(
		id string,
		description string,
		status string,
	) error
	UpdateResidentAddApartment(
		id string,
		apartment int,
		status string,
	) error
	UpdateResidentAddHome(
		id string,
		home string,
		status string,
	) error
	CreateResidentBase(
		id string,
		name string,
		status string,
	) error
	GetHomes() ([]common.Home, error)
	GetHome(id string) (common.Home, error)
	GetResidentilaComplex() (common.ResidentilaComplex, error)
	GetResidentsByHomeAndApartment(home string, apartment int) ([]common.Resident, error)
	GetResidentsByHome(home string) ([]common.Resident, error)
	GetResident(id string) (common.Resident, error)
}
