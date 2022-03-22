package tgapp

import (
	"concierge-bot/src/common"
	"strings"
)

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

func rightPad(s string, padStr string, pLen int) string {
	countRepeat := pLen - len(s)
	if countRepeat < 0 {
		countRepeat = 0
	}
	return strings.Repeat(padStr, countRepeat) + s
}
