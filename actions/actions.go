package actions

import (
	"fmt"
	"log"
	"os"
	"path"

	"google.golang.org/api/drive/v3"
)

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
