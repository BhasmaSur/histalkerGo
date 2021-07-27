package dto

type PicUploadDTO struct {
	Pics     [2][]byte `json:"pics" form:"pics" binding:"required"`
	PicCount int       `json:"pic_count" form:"pic_count" binding:"required"`
}

type PicDownloadDTO struct {
	Pics     [7][]byte `json:"pics" form:"pics" binding:"required"`
	PicCount int       `json:"pic_count" form:"pic_count" binding:"required"`
}
