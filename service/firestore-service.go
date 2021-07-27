package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/BhasmaSur/histalkergo/dto"
	"github.com/BhasmaSur/histalkergo/helper"
	"github.com/BhasmaSur/histalkergo/pool"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

const (
	projectId      string = "histalker-b9672"
	collectionName string = "users"
)

var (
	client storage.Client = pool.GetInstance().GetStorageClient()
)

type FirestoreService interface {
	InitialiseFireBase(context *gin.Context) (*firestore.Client, error)
	SaveImage(context *gin.Context) string
	SaveUserImagesOnStorage(fileInput [7][]byte, userId string) string
	RetreiveAllImagesOfUser(userId string, countOfPics int) string
	convertImageIntoBytes(imgName string) []byte
	convertBytesIntoImage(imgByte []byte, picName string) string
	convertFileIntoBytes(fileName string) []byte
}

type firestoreService struct{}

func NewFirestoreService() FirestoreService {
	return &firestoreService{}
}

func (f *firestoreService) InitialiseFireBase(ctx *gin.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("./histalker-b9672-firebase-adminsdk-vt2os-c47d930be0.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	} else {
		client, err := app.Firestore(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		return client, nil
	}
}
func (f *firestoreService) SaveImage(context *gin.Context) string {
	img1 := f.convertImageIntoBytes("assets/images/20.jpeg")
	img2 := f.convertImageIntoBytes("assets/images/15.jpeg")
	var picUploadDTO dto.PicDownloadDTO
	picUploadDTO.Pics[0] = img1
	picUploadDTO.Pics[1] = img2
	picUploadDTO.PicCount = 2
	response := helper.BuildResponse(true, "OK", picUploadDTO)
	context.JSON(http.StatusCreated, response)
	// img3 := f.convertImageIntoBytes("assets/images/21.jpeg")
	// img4 := f.convertImageIntoBytes("assets/images/16.jpeg")
	// img5 := f.convertImageIntoBytes("assets/images/me.jpeg")
	// img6 := f.convertImageIntoBytes("assets/images/52.jpeg")
	// var allImages [7][]byte
	// allImages[0] = img1
	// allImages[1] = img2
	// allImages[2] = img3
	// allImages[3] = img4
	// allImages[4] = img5
	// allImages[5] = img6
	// allImages[6] = nil
	// f.SaveUserImagesOnStorage(allImages, "mradul")
	//result := f.SaveUserImagesOnStorage(allImages, "mradul")
	// if result == "successfull" {
	// 	return f.RetreiveAllImagesOfUser("mradul", 6)
	// }
	//Users/mradulmishra/Desktop/Projects/HiStalkerGo/
	// privateKey := f.convertFileIntoBytes("histalker-b9672-29235ae90a29.p12")
	// lol := f.RetreiveAllImagesURLOfUser(privateKey)
	return "lol"
}

func (f *firestoreService) SaveUserImagesOnStorage(fileInput [7][]byte, userId string) string {
	id := uuid.New()
	ctx := context.Background()

	bucket := client.Bucket("histalker-b9672.appspot.com")
	for i := 0; i < 7; i++ {
		if fileInput[i] == nil {
			break
		}
		object := bucket.Object(collectionName + "/" + userId + "/" + "pic" + strconv.Itoa(i) + ".jpeg")
		writer := object.NewWriter(ctx)
		writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
		defer writer.Close()

		_, errN := io.Copy(writer, bytes.NewReader(fileInput[i]))
		if errN != nil {
			return errN.Error()
		}
	}
	return "successfull"
}

func (f *firestoreService) RetreiveAllImagesOfUser(userId string, countOfPics int) string {
	ctx := context.Background()
	// client, err := storage.NewClient(ctx)
	for i := 0; i < countOfPics; i++ {
		rc, err := client.Bucket("histalker-b9672.appspot.com").Object("users/" + userId + "/pic" + strconv.Itoa(i) + ".jpeg").NewReader(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()
		body, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		f.convertBytesIntoImage(body, "pic"+strconv.Itoa(i)+".jpeg")
	}
	return "successfull"
}

func (f *firestoreService) RetreiveAllImagesURLOfUser(privateKey []byte) string {
	//ctx := context.Background()
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	opts := storage.SignedURLOptions{
		GoogleAccessID: "firebase-adminsdk-vt2os@histalker-b9672.iam.gserviceaccount.com",
		PrivateKey:     privateKey,
		Method:         "GET",
		Expires:        tomorrow,
	}
	lol, err := storage.SignedURL("histalker-b9672.appspot.com", "users/mradul/pic1.jpeg", &opts)
	if err != nil {
		return err.Error()
	}
	fmt.Println(lol)
	// client, err := storage.NewClient(ctx)
	//rc, err := client.Bucket("histalker-b9672.appspot.com").Object("users/" + userId + "/pic" + strconv.Itoa(i) + ".jpeg").
	// for i := 0; i < countOfPics; i++ {
	// 	rc, err := client.Bucket("histalker-b9672.appspot.com").Object("users/" + userId + "/pic" + strconv.Itoa(i) + ".jpeg").NewReader(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer rc.Close()
	// 	body, err := ioutil.ReadAll(rc)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	f.convertBytesIntoImage(body, "pic"+strconv.Itoa(i)+".jpeg")
	// }
	return "successfull"
}

func (*firestoreService) convertImageIntoBytes(imgName string) []byte {
	new_image, err := os.Open(imgName)
	fmt.Sprintln("here", new_image)
	if err != nil {
		return nil
	}
	img, fmtName, err := image.Decode(new_image)
	if err != nil {
		fmt.Sprintln(fmtName)
		return nil
	}
	buf := new(bytes.Buffer)
	newErr := jpeg.Encode(buf, img, nil)
	if newErr != nil {
		return nil
	}
	send_s3 := buf.Bytes()
	return send_s3

}

func (*firestoreService) convertBytesIntoImage(imgByte []byte, picName string) string {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./" + picName)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}
	return "successfull"
}

func (*firestoreService) convertFileIntoBytes(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

// ctx := context.Background()
// client, err := firestore.NewClient(ctx, projectId)
// //client, err := secretmanager.NewClient(context.Background(), option.WithCredentialsFile("histalker-b9672-firebase-adminsdk-vt2os-7293908177.json"))
// if err != nil {
// 	log.Fatalf("Failed to create a firestore client: %v", err)
// 	return "nil"
// }
// defer client.Close()
// _, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
// 	"pic1": "anything testing",
// 	"pic2": "anything testing",
// 	"pic3": "anything testing",
// 	"pic4": "anything testing",
// 	"pic5": "anything testing",
// 	"pic6": "anything testing",
// 	"pic7": "anything testing",
// })
// if err != nil {
// 	log.Fatalf("Failed to add pics to firestore client: %v", err)
// 	return "nil"
// }
// return "random"

//client, err := storage.NewClient(ctx)
//client, err := secretmanager.NewClient(context.Background(), option.WithCredentialsFile("histalker-b9672-firebase-adminsdk-vt2os-7293908177.json"))
// if err != nil {
// 	log.Fatalf("Failed to create a storage client: %v", err)
// 	return "nil"
// }
// defer client.Close()

//object := bucket.Object(collectionName + "/" + userId + "/" + fileName)
//object := bucket.Object(collectionName + "/" + userId + "/" + fileName)
// writer := object.NewWriter(ctx)
// writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
// defer writer.Close()

// _, errN := io.Copy(writer, bytes.NewReader(fileInput))
// if errN != nil {
// 	return errN.Error()
// }
