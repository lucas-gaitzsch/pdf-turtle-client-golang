package pdfturtleclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/models/dto"
)

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

func (c *PdfTurtleClient) Render(renderData models.RenderData) (io.ReadCloser, error) {
	return c.RenderWithContext(context.Background(), renderData)
}
func (c *PdfTurtleClient) RenderWithContext(ctx context.Context, renderData models.RenderData) (io.ReadCloser, error) {
	json, err := json.Marshal(renderData)
	if err != nil {
		return nil, err
	}

	return c.sendRenderRequest(ctx, "/api/pdf/from/html/render", bytes.NewReader(json), "application/json")
}

func (c *PdfTurtleClient) RenderTemplate(renderTemplateData models.RenderTemplateData) (io.ReadCloser, error) {
	return c.RenderTemplateWithContext(context.Background(), renderTemplateData)
}

func (c *PdfTurtleClient) RenderTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData) (io.ReadCloser, error) {
	json, err := json.Marshal(renderTemplateData)
	if err != nil {
		return nil, err
	}

	return c.sendRenderRequest(ctx, "/api/pdf/from/html-template/render", bytes.NewReader(json), "application/json")
}

func (c *PdfTurtleClient) TestTemplate(renderTemplateData models.RenderTemplateData) (*dto.TemplateTestResult, error) {

	return c.TestTemplateWithContext(context.Background(), renderTemplateData)
}
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

func (c *PdfTurtleClient) RenderBundle(bundles []io.Reader, model any) (io.ReadCloser, error) {
	return c.RenderBundleWithContext(context.Background(), bundles, model)
}
func (c *PdfTurtleClient) RenderBundleWithContext(ctx context.Context, bundles []io.Reader, model any) (io.ReadCloser, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// bundles
	for i, b := range bundles {
		part, err := writer.CreateFormFile("bundle", filepath.Base(fmt.Sprintf("bundle_%d.zip", i)))
		if err != nil {
			return nil, err
		}
		io.Copy(part, b)
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
