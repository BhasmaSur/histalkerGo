package main

import (
	"github.com/BhasmaSur/histalkergo/config"
	"github.com/BhasmaSur/histalkergo/controller"
	"github.com/BhasmaSur/histalkergo/middleware"
	"github.com/BhasmaSur/histalkergo/repository"
	"github.com/BhasmaSur/histalkergo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	profileRepository repository.ProfileRepository = repository.NewProfileRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserService(userRepository)
	picService        service.PicService           = service.NewPicService()
	profileService    service.ProfileService       = service.NewProfileService(profileRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	firestoreService  service.FirestoreService     = service.NewFirestoreService()
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	profileController controller.ProfileController = controller.NewProfileController(profileService, jwtService)
	picController     controller.PicsController    = controller.NewPicController(picService, jwtService)
)

func main() {
	//os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "histalker-b9672-firebase-adminsdk-vt2os-7293908177.json")
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	profileRoutes := r.Group("api/profile", middleware.AuthorizeJWT(jwtService))
	{
		profileRoutes.POST("/create", profileController.Insert)
		profileRoutes.GET("/user", profileController.UserProfile)
		profileRoutes.PUT("/update", profileController.Update)
		profileRoutes.GET("/test", profileController.Test)
		profileRoutes.POST("/update-pics", picController.GetUploadedPics)
	}

	r.Run()
}
