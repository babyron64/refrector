package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/kanatatsu64/refrector/auth"
)

var defaultPort = "8080"

// InitServer initialize the server to be ready for requests.
func InitServer(port string) {
	auth.InitAuth()

	http.HandleFunc("/", staticHandler)

	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// GetPort returns a port which the server should listen on.
func GetPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return defaultPort
}

// re is a regular expression used in splitPath to
// split a given path into three parts: directory name,
// file name and extension. It is safe to pass a path
// that does not have an extension.
//
// Directory name does not include the following slash,
// that is, does not end with a slash.
//
// File name does not include an extension and the
// preceeding dot.
var re, _ = regexp.Compile(`(.*)/(\w*)\.?(\w*)$`)

// splitPath splits the given path into a set of
// the attributes: directory name, file name and
// extension.
func splitPath(path string) (string, string, string) {
	res := re.FindStringSubmatch(path)
	return res[1], res[2], res[3]
}

// getFilePathFromURLPath constructs the path of the
// requested file in the server from the url path.
//
// When a url without an extension is provided,
// getFilePathFromURL takes it as a html file and
// add a html extension to the corresponding file path.
func getStaticFilePath(url *url.URL) string {
	urlpath := url.Path
	var dirname, filename, ext = splitPath(urlpath)

	if ext == "" {
		ext = "html"
	}

	switch ext {
	case "js":
	case "css":
	case "html":
	default:
		log.Fatalf("Unknown file url path %v", urlpath)
	}

	return fmt.Sprintf("%v/%v.%v", dirname, filename, ext)
}

// staticHandler handles requests for normal static files,
// such as javascript, css, html .etc, and returns the
// requested file.
func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := getStaticFilePath(r.URL)

	http.ServeFile(w, r, path)
}
