# PdfTurtle Client Golang
Golang (Go) client library to use the [PdfTurtle](https://github.com/lucas-gaitzsch/pdf-turtle) service 

**HINT:** This lib is in the beta.

## How to use - Recommended way

See a working example in [examples/main.go](./examples/main.go).

### 1. Prepare project

Get the package and than you are ready to go.
```bash
go get github.com/lucas-gaitzsch/pdf-turtle-client-dotnet
```

```bash
// create your client
c := pdfturtleclient.NewPdfTurtleClient("https://pdfturtle.gaitzsch.dev")
```

### 2. Design your PDF in the playground
Go to [🐢PdfTurtle-Playground](https://pdfturtle.gaitzsch.dev/), put an example model as JSON and design your PDF.
Download the bundle as ZIP file and put it in your resources/assets.

### 3. Call the service with the client and your data
Call `RenderBundle` to render the pdf to a `io.Reader`.

```golang
pdf := c.RenderBundle([]io.Reader{ BUNDLE_AS_READER }, MODEL_AS_OBJECT)
```

**Done.**

### Hint: You can split your bundle
If you want to have the same header for all documents, you can create a ZIP file with with only the `header.html` file.
Now you can call the Service with multiple bundle files. The service will assemble the files together.

```golang
pdf := c.RenderBundle(
    []io.Reader{ BUNDLE_WITHOUT_HEADER_AS_READER, HEADER_BUNDLE_AS_READER },
    MODEL_AS_OBJECT,
)
```


## How to use - Alternative ways
### Without template (plain HTML)
If the described way does not match your expectations, you can use a template engine of your choice (for example `html/template`) and render HTML directly with PdfTurtle.

```golang
pdf := c.Render(RenderData{
    ...
})
```

### With template but no bundle
If you want to render a HTML template without any images or assets, you can use the `RenderTemplate` function.

```golang
pdf := c.RenderTemplateAsync(RenderTemplateData{
    ...
})
```


## Open TODOs
- [x] Working examples for all methods
- [x] Add documentation as comments
- [ ] Tests