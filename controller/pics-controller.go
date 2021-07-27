package controller

import (
	"fmt"
	"net/http"

	"github.com/BhasmaSur/histalkergo/helper"
	"github.com/BhasmaSur/histalkergo/service"
	"github.com/gin-gonic/gin"
)

type PicsController interface {
	GetUploadedPics(context *gin.Context)
	PerformOperationAndUpload(context *gin.Context)
	ExtractFaceDescriptors(context *gin.Context)
	StoreOnStorage(context *gin.Context)
}

type picsController struct {
	picService service.PicService
	jwtService service.JWTService
}

func NewPicController(picService service.PicService, jwtService service.JWTService) PicsController {
	return &picsController{
		picService: picService,
		jwtService: jwtService,
	}
}

func (p *picsController) GetUploadedPics(context *gin.Context) {
	//fmt.Fprintf(context.Request.Response, "pic sent is of : ")
	r := context.Request
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error recieving file from server")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("uploaded file : %+v\n", handler.Filename)
	fmt.Printf("file size : %+v\n", handler.Size)
	fmt.Printf("Mime header : %+v\n", handler.Header)
	response := helper.BuildResponse(true, "OK", handler.Filename)
	context.JSON(http.StatusCreated, response)
}

func (p *picsController) ExtractFaceDescriptors(context *gin.Context) {

}

func (p *picsController) StoreOnStorage(context *gin.Context) {

}

func (p *picsController) PerformOperationAndUpload(context *gin.Context) {

	// errDTO := context.ShouldBind(&picUploadDTO)
	// if errDTO != nil {
	// 	res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	// 	context.JSON(http.StatusBadRequest, res)
	// } else {
	// 	userID := service.NewCommonService(p.jwtService).GetUserId(context)
	// 	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Println("working", convertedUserID)
	// 	response := helper.BuildResponse(true, "OK", picUploadDTO)
	// 	context.JSON(http.StatusCreated, response)
	// }
}
