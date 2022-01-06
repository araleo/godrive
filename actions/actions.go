package actions

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"google.golang.org/api/drive/v3"
)

// printFilesResult receives an array of drive files and iterate through the array printing files names and ids.
func printFilesResult(files []*drive.File) {
	fmt.Println("Files:")
	if len(files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range files {
			fmt.Printf("%s (%s)\n\n", i.Name, i.Id)
		}
	}
}

// ListFiles receives a drive service and lists all files and folders in that drive.
func ListFiles(srv *drive.Service) {
	r, err := srv.Files.List().Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	printFilesResult(r.Files)
}

// QueryFiles receives a drive service and a query string and lists all files whose names contains the query (case insensitive).
func QueryFiles(srv *drive.Service, query string) {
	queryString := fmt.Sprintf("name contains '%s'", query)
	r, err := srv.Files.List().Q(queryString).Do()
	if err != nil {
		log.Fatal(err)
	}
	printFilesResult(r.Files)
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

	file, err := os.Create(path.Join("./downloads", meta.Name))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, data.Body)
	if err != nil {
		log.Fatal(err)
	}
}

// GetFromDrive receives a drive service and a fileId from a file in Google Docs and downloads the file to the ./downloads folder.
func GetFromDrive(srv *drive.Service, fileId string) {
	data, err := srv.Files.
		Export(fileId, "text/plain").
		Download()
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	file, err := os.Create(path.Join("./downloads", "teste.txt"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, data.Body)
	if err != nil {
		log.Fatal(err)
	}
}
