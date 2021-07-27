package service

import (
	"github.com/BhasmaSur/histalkergo/dto"
	"github.com/BhasmaSur/histalkergo/entity"
)

type (
	PicService interface {
		UpdatePics(userPics dto.PicUploadDTO) entity.Profile
		GetAllPics() []string
	}

	picService struct {
	}
)

func NewPicService() PicService {
	return &picService{}
}

func (service *picService) UpdatePics(userPics dto.PicUploadDTO) entity.Profile {
	profileToUpdate := entity.Profile{}
	return profileToUpdate
}

func (service *picService) GetAllPics() []string {
	return nil
}
