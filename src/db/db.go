package db

import (
	"concierge-bot/src/common"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type DBConnect struct {
	db *gorm.DB
}

func (d DBConfig) getConnect() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.Host,
		d.User,
		d.Password,
		d.Name,
		d.Port)
}

func GetConnect(conf DBConfig) DBConnect {
	dsn := conf.getConnect()
	time.Sleep(5 * time.Second)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB:GetConnect: %w ", err)
	} else {
		log.Println("DB:GetConnect: connection succed")
	}
	return DBConnect{db: db}
}

func (c DBConnect) InitDB() {
	log.Println("DB:InitDB: init table")
	c.db.AutoMigrate(&ResidentilaComplex{})
	c.db.AutoMigrate(&Home{})
	c.db.AutoMigrate(&Resident{})
}

func (c DBConnect) GetResidentilaComplex() (common.ResidentilaComplex, error) {
	log.Println("DB:GetResidentilaComplex: create complex")
	complex := ResidentilaComplex{}
	result := c.db.Select("name", "description").Find(&complex)
	return residentilaComplexConvert(complex), result.Error
}

func (c DBConnect) CreateResidentilaComplex(
	name string,
	description string,
) error {
	log.Println("DB:CreateResidentilaComplex: create resident")
	complex := common.ResidentilaComplex{
		Name:        name,
		Description: description,
	}
	result := c.db.Create(&complex)
	return result.Error
}

func (c DBConnect) GetHomes() ([]common.Home, error) {
	log.Println("DB:GetHome: get homes")
	homes := []Home{}
	result := c.db.Find(&homes)
	return homesConvert(homes), result.Error
}

func (c DBConnect) GetHome(id string) (common.Home, error) {
	log.Println("DB:GetHome: get homes")
	home := Home{}
	result := c.db.Where("name = ?", id).Find(&home)
	return homeConvert(home), result.Error
}

func (c DBConnect) CreateHome(
	name string,
	description string,
	residentialComplex string,
	floors int,
	firstResidentialFloor int,
	apartments int,
	firstApartment int,
	entrances int,
) error {
	log.Println(fmt.Sprintf("DB:CreateHome: create home"))
	home := Home{
		Name:                  name,
		Description:           description,
		ResidentialComplex:    residentialComplex,
		Floors:                floors,
		FirstResidentialFloor: firstResidentialFloor,
		Apartments:            apartments,
		FirstApartment:        firstApartment,
		Entrances:             entrances,
	}
	result := c.db.Create(&home)
	return result.Error
}

func (c DBConnect) GetResident(id string) (common.Resident, error) {
	log.Println("DB:GetResident: get resident")
	resident := Resident{}
	result := c.db.Where("id = ?", id).Find(&resident)
	return residentConvert(resident), result.Error
}

func (c DBConnect) GetResidentsByHomeAndApartment(home string, apartment int) ([]common.Resident, error) {
	log.Println(fmt.Sprintf("DB:GetResident: get resident by home - %s and apartment - %d", home, apartment))
	residents := []Resident{}
	result := c.db.Where("home = ? and apartment = ?", home, apartment).Find(&residents)
	return residentsConvert(residents), result.Error
}

func (c DBConnect) GetResidentsByHome(home string) ([]common.Resident, error) {
	log.Println(fmt.Sprintf("DB:GetResident: get residents by home - %s", home))
	residents := []Resident{}
	result := c.db.Select("apartment").Where("home = ?", home).Find(&residents)
	return residentsConvert(residents), result.Error
}

func (c DBConnect) CreateResidentFull(
	id string,
	name string,
	description string,
	apartment int,
	home string,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:CreateResidentFull: create resident %s", id))
	resident := Resident{
		ID:          id,
		Name:        name,
		Description: description,
		Apartment:   apartment,
		Home:        home,
		Status:      status,
	}
	result := c.db.Create(&resident)
	return result.Error
}

func (c DBConnect) UpdateResidentFull(
	id string,
	name string,
	description string,
	apartment int,
	home string,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:UpdateResidentFull: update resident %s", id))
	resident := Resident{
		ID:          id,
		Name:        name,
		Description: description,
		Apartment:   apartment,
		Home:        home,
		Status:      status,
	}
	result := c.db.Save(&resident)
	return result.Error
}

func (c DBConnect) CreateResidentBase(
	id string,
	name string,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:CreateResidentBase: create resident %s", id))
	resident := Resident{
		ID:     id,
		Name:   name,
		Status: status,
	}
	result := c.db.Save(&resident)
	return result.Error
}

func (c DBConnect) UpdateResidentAddHome(
	id string,
	home string,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:UpdateResidentAddHome: update resident %s", id))
	resident := Resident{
		ID: id,
	}
	result := c.db.Model(&resident).Updates(Resident{
		Home:   home,
		Status: status,
	})
	return result.Error
}

func (c DBConnect) UpdateResidentAddApartment(
	id string,
	apartment int,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:UpdateResidentAddApartment: update resident %s", id))
	resident := Resident{
		ID: id,
	}
	result := c.db.Model(&resident).Updates(Resident{
		Apartment: apartment,
		Status:    status,
	})
	return result.Error
}

func (c DBConnect) UpdateResidentAddDescription(
	id string,
	description string,
	status string,
) error {
	log.Println(fmt.Sprintf("DB:UpdateResidentAddDescription: update resident %s", id))
	resident := Resident{
		ID:          id,
		Description: description,
		Status:      status,
	}
	result := c.db.Model(&resident).Updates(Resident{
		Description: description,
		Status:      status,
	})
	return result.Error
}
