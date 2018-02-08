package main

import (
	"bytes"
	"html/template"
	"log"
	"net/url"

	"github.com/zserge/webview"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func externalCallback(w webview.WebView, arg string) {
	log.Printf("External callback: \"%s\"\n", arg)
}

func main() {
	tmpl, err := template.New("index.html").Delims("[[", "]]").ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, nil)

	w := webview.New(webview.Settings{
		URL:                    `data:text/html,` + url.PathEscape(buffer.String()),
		Debug:                  true,
		Width:                  2000,
		Height:                 2000,
		ExternalInvokeCallback: externalCallback,
	})

	w.Dispatch(func() {
		w.InjectCSS(slurp("index.css"))
		checkErr(w.Eval(setBodyHTML(slurp("body.html"))))

		systemState := SystemState{}
		_, err = w.Bind("systemState", &systemState)
		checkErr(err)

		checkErr(w.Eval(slurp("https://unpkg.com/vue@2.5.13")))
		checkErr(w.Eval(slurp("index.js")))
	})

	w.Run()
}
