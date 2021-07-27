package service

import (
	"fmt"
	"log"

	"github.com/BhasmaSur/histalkergo/dto"
	"github.com/BhasmaSur/histalkergo/entity"
	"github.com/BhasmaSur/histalkergo/repository"
	"github.com/mashingan/smapping"
)

type ProfileService interface {
	Insert(profile dto.ProfileCreateDTO) entity.Profile
	Update(profile dto.ProfileUpdateDTO) entity.Profile
	FindByID(profileID uint64) entity.Profile
	FindByUserID(userID uint64) entity.Profile
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type profileService struct {
	profileRepository repository.ProfileRepository
}

//NewUserService creates a new instance of ProfileService
func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepo,
	}
}

func (service *profileService) Insert(p dto.ProfileCreateDTO) entity.Profile {
	profile := entity.Profile{}
	err := smapping.FillStruct(&profile, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.profileRepository.InsertProfile(profile)
	return res
}

func (service *profileService) Update(profile dto.ProfileUpdateDTO) entity.Profile {
	profileToUpdate := entity.Profile{}
	err := smapping.FillStruct(&profileToUpdate, smapping.MapFields(&profile))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedProfile := service.profileRepository.UpdateProfile(profileToUpdate)
	return updatedProfile
}

func (service *profileService) Profile(profileID uint64) entity.Profile {
	return service.profileRepository.FindProfile(profileID)
}

func (service *profileService) FindByID(profileID uint64) entity.Profile {
	return service.profileRepository.FindProfileByID(profileID)
}

func (service *profileService) IsAllowedToEdit(userID string, profileID uint64) bool {
	p := service.profileRepository.FindProfileByID(profileID)
	id := fmt.Sprintf("%v", p.UserID)
	return userID == id
}

func (service *profileService) FindByUserID(userID uint64) entity.Profile {
	p := service.profileRepository.FindProfileByUserID(userID)
	return p
}
