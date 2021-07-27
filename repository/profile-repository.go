package repository

import (
	"github.com/BhasmaSur/histalkergo/entity"
	"gorm.io/gorm"
)

//BookRepository is a ....
type ProfileRepository interface {
	InsertProfile(p entity.Profile) entity.Profile
	UpdateProfile(p entity.Profile) entity.Profile
	FindProfile(profileID uint64) entity.Profile
	FindProfileByID(profileID uint64) entity.Profile
	FindProfileByUserID(userID uint64) entity.Profile
}

type profileConnection struct {
	connection *gorm.DB
}

func NewProfileRepository(dbConn *gorm.DB) ProfileRepository {
	return &profileConnection{
		connection: dbConn,
	}
}
func (db *profileConnection) InsertProfile(p entity.Profile) entity.Profile {
	db.connection.Save(&p)
	db.connection.Preload("Profile").Find(&p)
	return p
}

func (db *profileConnection) UpdateProfile(p entity.Profile) entity.Profile {
	db.connection.Save(&p)
	db.connection.Preload("Profile").Find(&p)
	return p
}

func (db *profileConnection) FindProfileByID(profileID uint64) entity.Profile {
	var profile entity.Profile
	db.connection.Preload("User").Find(&profile, profileID)
	return profile
}

func (db *profileConnection) FindProfile(profileID uint64) entity.Profile {
	var profile entity.Profile
	db.connection.Preload("User").Find(&profile, profileID)
	return profile
}

func (db *profileConnection) FindProfileByUserID(userID uint64) entity.Profile {
	var profile entity.Profile
	db.connection.Preload("Profile").Where("user_id = ?", userID).First(&profile)
	return profile
}
