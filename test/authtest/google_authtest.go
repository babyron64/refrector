package authtest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kanatatsu64/refrector/auth"
)

// GoogleAuthTestHandler is a test handler for google drive api authorization.
func GoogleAuthTestHandler(w http.ResponseWriter, r *http.Request) {
	srv, err := auth.CreateGoogleDriveClient()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	res, err := srv.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(res.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range res.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
