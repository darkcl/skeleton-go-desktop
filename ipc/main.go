package ipc

import (
	"encoding/json"
	"fmt"
	"sync"
	"text/template"

	"github.com/darkcl/webview"
)

// Main - Main Process IPC
type Main struct {
	w webview.WebView
}

var once sync.Once
var (
	instance *Main
)

// SharedMain - Get Shared IPC Instance
func SharedMain() *Main {
	once.Do(func() {
		instance = &Main{}
	})

	return instance
}

// SetView -  Set Webview to main
func (m *Main) SetView(view webview.WebView) {
	m.w = view
}

// On - Handle renderer incoming messagin
func (m *Main) On(event string, value string) {

}

// Send - Send a message to renderer
func (m *Main) Send(event string, value interface{}) {
	fmt.Println("Send Event")
	jsonString, err := json.Marshal(value)

	if err != nil {
		fmt.Printf("Error on sending value: %v\n", err)
	}

	jsString := fmt.Sprintf(`window.renderer.trigger("%s", "%s")`, event, template.JSEscapeString(string(jsonString)))
	m.w.Eval(jsString)
}
