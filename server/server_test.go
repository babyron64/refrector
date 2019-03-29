package server

import (
	"net/url"
	"testing"
)

func TestSplitPath1(t *testing.T) {
	dirname, filename, ext := splitPath("/static/javascript/test.js")
	if dirname == "/static/javascript" &&
		filename == "test" &&
		ext == "js" {
		return
	}
	t.Logf("\n"+
		"dir name: %v\n"+
		"file name: %v\n"+
		"extension: %v",
		dirname, filename, ext)
	t.Fail()
}
func TestSplitPath2(t *testing.T) {
	dirname, filename, ext := splitPath("/home")
	if dirname == "" &&
		filename == "home" &&
		ext == "" {
		return
	}
	t.Logf("\n"+
		"dir name: %v\n"+
		"file name: %v\n"+
		"extension: %v",
		dirname, filename, ext)
	t.Fail()
}

func TestGetFilePathFromURLPath1(t *testing.T) {
	url, _ := url.Parse("http://example.com/static/javascript/test.js")
	path := getStaticFilePath(url)
	expect := "/static/javascript/test.js"
	if path != expect {
		t.Logf("file path: %v", path)
		t.Fail()
	}
}

func TestGetFilePathFromURLPath2(t *testing.T) {
	url, _ := url.Parse("http://example.com/home")
	path := getStaticFilePath(url)
	expect := "/home.html"
	if path != expect {
		t.Logf("file path: %v", path)
		t.Fail()
	}
}
