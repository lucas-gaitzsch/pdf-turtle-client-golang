package pdfturtleclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/models/dto"
)

// Creates a new PdfTurtle http client with the given PdfTurtle service baseUrl
func NewPdfTurtleClient(baseUrl string) PdfTurtleClientInterface {
	return &PdfTurtleClient{
		baseUrl: strings.TrimRight(baseUrl, "/"),
		client:  &http.Client{},
	}
}

type PdfTurtleClient struct {
	baseUrl string
	client  *http.Client
}


// Returns PDF file generated from HTML of body, header and footer
func (c *PdfTurtleClient) RenderWithContext(ctx context.Context, renderData models.RenderData) (io.ReadCloser, error) {
	json, err := json.Marshal(renderData)
	if err != nil {
		return nil, err
	}

	return c.sendRenderRequest(ctx, "/api/pdf/from/html/render", bytes.NewReader(json), "application/json")
}
// Returns PDF file generated from HTML of body, header and footer
func (c *PdfTurtleClient) Render(renderData models.RenderData) (io.ReadCloser, error) {
	return c.RenderWithContext(context.Background(), renderData)
}

// Returns PDF file generated from HTML template plus model of body, header and footer
func (c *PdfTurtleClient) RenderTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData) (io.ReadCloser, error) {
	json, err := json.Marshal(renderTemplateData)
	if err != nil {
		return nil, err
	}

	return c.sendRenderRequest(ctx, "/api/pdf/from/html-template/render", bytes.NewReader(json), "application/json")
}
// Returns PDF file generated from HTML template plus model of body, header and footer
func (c *PdfTurtleClient) RenderTemplate(renderTemplateData models.RenderTemplateData) (io.ReadCloser, error) {
	return c.RenderTemplateWithContext(context.Background(), renderTemplateData)
}

// Returns information about matching model data to template
func (c *PdfTurtleClient) TestTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData) (*dto.TemplateTestResult, error) {
	jb, err := json.Marshal(renderTemplateData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getUrlOfMethod("/api/pdf/from/html-template/test"), bytes.NewReader(jb))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bytes))
	}

	ttr := &dto.TemplateTestResult{}

	jd := json.NewDecoder(resp.Body)
	jd.Decode(ttr)

	return ttr, nil
}
// Returns information about matching model data to template
func (c *PdfTurtleClient) TestTemplate(renderTemplateData models.RenderTemplateData) (*dto.TemplateTestResult, error) {
	return c.TestTemplateWithContext(context.Background(), renderTemplateData)
}

// Returns PDF file generated from bundle (Zip-File) of HTML or HTML template of body, header, footer and assets. The index.html file in the Zip-Bundle is required.
func (c *PdfTurtleClient) RenderBundleWithContext(ctx context.Context, bundles map[string]io.Reader, model any) (io.ReadCloser, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// bundles
	for fileName, bundle := range bundles {
		part, err := writer.CreateFormFile("bundle", filepath.Base(fileName))
		if err != nil {
			return nil, err
		}
		io.Copy(part, bundle)
	}

	// model
	if model != nil {
		part, err := writer.CreateFormField("model")
		if err != nil {
			return nil, err
		}
		jd := json.NewEncoder(part)
		err = jd.Encode(model)
		if err != nil {
			return nil, err
		}
	}
	writer.Close()

	return c.sendRenderRequest(ctx, "/api/pdf/from/html-bundle/render", body, writer.FormDataContentType())
}
// Returns PDF file generated from bundle (Zip-File) of HTML or HTML template of body, header, footer and assets. The index.html file in the Zip-Bundle is required.
func (c *PdfTurtleClient) RenderBundle(bundles map[string]io.Reader, model any) (io.ReadCloser, error) {
	return c.RenderBundleWithContext(context.Background(), bundles, model)
}

func (c *PdfTurtleClient) sendRenderRequest(ctx context.Context, method string, body io.Reader, contentType string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getUrlOfMethod(method), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return errorFromResponse(resp)
	}

	return resp.Body, nil
}

func (c *PdfTurtleClient) getUrlOfMethod(method string) string {
	return c.baseUrl + method
}

func errorFromResponse(resp *http.Response) (io.ReadCloser, error) {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(string(bytes))
}
