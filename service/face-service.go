package service

type (
	FaceService interface {
		AddImagesToFaceDescriptor()
		GetFaceDescriptor()
		SaveFaceDescriptor()
		FindFaceInDiscriptor()
	}

	faceService struct {
	}
)
