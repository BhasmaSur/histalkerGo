package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BhasmaSur/histalkergo/dto"
	"github.com/BhasmaSur/histalkergo/entity"
	"github.com/BhasmaSur/histalkergo/helper"
	"github.com/BhasmaSur/histalkergo/service"
	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	Insert(context *gin.Context)
	Update(context *gin.Context)
	UserProfile(context *gin.Context)
	Test(context *gin.Context)
}

type profileController struct {
	profileService service.ProfileService
	jwtService     service.JWTService
}

func NewProfileController(profileService service.ProfileService, jwtService service.JWTService) ProfileController {
	return &profileController{
		profileService: profileService,
		jwtService:     jwtService,
	}
}

func (c *profileController) Test(context *gin.Context) {
	firestoreService := service.NewFirestoreService()
	response := helper.BuildResponse(true, "OK", firestoreService.SaveImage(context))
	context.JSON(http.StatusCreated, response)
	fmt.Sprintln("its fine")
}
func (c *profileController) Insert(context *gin.Context) {
	var profileCreateDTO dto.ProfileCreateDTO
	errDTO := context.ShouldBind(&profileCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		userID := service.NewCommonService(c.jwtService).GetUserId(context)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			profileCreateDTO.UserID = convertedUserID
		}
		result := c.profileService.Insert(profileCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}
func (c *profileController) Update(context *gin.Context) {
	var ProfileUpdateDTO dto.ProfileUpdateDTO
	errDTO := context.ShouldBind(&ProfileUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userID := service.NewCommonService(c.jwtService).GetUserId(context)
	if c.profileService.IsAllowedToEdit(userID, ProfileUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			ProfileUpdateDTO.UserID = id
		}
		result := c.profileService.Update(ProfileUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *profileController) UserProfile(context *gin.Context) {
	strId := service.NewCommonService(c.jwtService).GetUserId(context)
	id, err := strconv.ParseUint(strId, 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var profile entity.Profile = c.profileService.FindByUserID(id)
	if (profile == entity.Profile{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with givern id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", profile)
		context.JSON(http.StatusOK, res)
	}
}
