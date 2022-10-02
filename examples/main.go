package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/pdfturtleclient"
)

func main() {

	c := pdfturtleclient.NewPdfTurtleClient("https://pdfturtle.gaitzsch.dev")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<!DOCTYPE html>"))
		w.Write([]byte("<a href='/html-bundle-example'>   HTML bundle example   </a><br>"))
		w.Write([]byte("<a href='/html-example'>          HTML example          </a><br>"))
		w.Write([]byte("<a href='/html-template-example'> HTML template example </a><br>"))
	})

	http.HandleFunc("/html-bundle-example", func(w http.ResponseWriter, r *http.Request) {
		bodyBundle, _ := os.ReadFile("pdf-turtle-bundle-body.zip")
		headerBundle, _ := os.ReadFile("pdf-turtle-bundle-header.zip")

		resp, _ := c.RenderBundle([]io.Reader{
			bytes.NewReader(bodyBundle),
			bytes.NewReader(headerBundle),
		},
			map[string]any{
				"title":   "PdfTurtle _🐢_ TestReport",
				"heading": "Sales Overview",
				"summery": map[string]any{
					"totalSales":       32993,
					"salesPerWeek":     82,
					"performanceIndex": 5.132,
					"salesVolume":      848932,
				},
			})
		defer resp.Close()

		io.Copy(w, resp)
	})

	http.HandleFunc("/html-example", func(w http.ResponseWriter, r *http.Request) {
		html := "<b>test</b>"

		resp, _ := c.Render(models.RenderData{
			Html:       &html,
			HeaderHtml: "test header",
		})
		defer resp.Close()

		io.Copy(w, resp)
	})

	http.HandleFunc("/html-template-example", func(w http.ResponseWriter, r *http.Request) {
		html := "<b>hi {{.Name}}</b>"

		resp, _ := c.RenderTemplate(models.RenderTemplateData{
			HtmlTemplate:       &html,
			HeaderHtmlTemplate: "<div>test header: {{.Title}}</div>",
			Model: struct {
				Name  string
				Title string
			}{Name: "Timo", Title: "Test-Titel"},
		})
		defer resp.Close()

		io.Copy(w, resp)
	})

	fmt.Printf("Starting server at port 8888\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
