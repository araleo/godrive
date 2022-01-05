package actions

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"google.golang.org/api/drive/v3"
)

// ListFiles receives a drive service and lists all files and folders in that drive.
func ListFiles(srv *drive.Service) {
	r, err := srv.Files.List().PageSize(50).
		Fields("nextPageToken, files(id, name, kind, parents)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s - %s (%s)\n%s\n\n", i.Kind, i.Name, i.Id, i.Parents)
		}
	}
}

// UploadFile receives a drive service and a local filepath and uploads the file to the root of the drive.
func UploadFile(srv *drive.Service, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := &drive.File{
		Name: path.Base(filepath),
	}

	r, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - %s\n", r.Id, r.Name)
}

// GetFile receives a drive service and a drive file ID and downloads the file to the ./downloads folder.
func GetFile(srv *drive.Service, fileId string) {
	meta, err := srv.Files.Get(fileId).Do()
	if err != nil {
		log.Fatal(err)
	}

	data, err := srv.Files.Get(fileId).Download()
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	file, err := os.Create(meta.Name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, data.Body)
	if err != nil {
		log.Fatal(err)
	}
}
