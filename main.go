package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/darkcl/skeleton-go-desktop/helpers"
	webview "github.com/darkcl/webview"
	"github.com/leaanthony/mewn"
)

func handleRPC(w webview.WebView, data string) {
	switch {
	case strings.HasPrefix(data, "openlink: "):
		url := strings.TrimPrefix(data, "openlink: ")
		helpers.OpenBrowser(url)
	default:
		panic("Not Implemented")
	}
}

func main() {
	js := mewn.String("./ui/dist/bundle.min.js")
	indexHTML := mewn.String("./ui/dist/index.html")

	dir, err := ioutil.TempDir("", "skeleton")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer os.RemoveAll(dir)
	tmpIndex := filepath.Join(dir, "index.html")
	if err := ioutil.WriteFile(tmpIndex, []byte(indexHTML), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}
	abs, err := filepath.Abs(tmpIndex)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	tmpJs := filepath.Join(dir, "bundle.min.js")
	if err := ioutil.WriteFile(tmpJs, []byte(js), 0666); err != nil {
		log.Fatal(err)
		panic(err)
	}

	w := webview.New(webview.Settings{
		Title:                  "Skeleton",
		URL:                    "file://" + abs,
		Resizable:              true,
		Width:                  1024,
		Height:                 768,
		ExternalInvokeCallback: handleRPC,
		Debug:                  true,
	})
	defer w.Exit()

	w.Run()
}
