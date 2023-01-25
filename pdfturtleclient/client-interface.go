package pdfturtleclient

import (
	"context"
	"io"

	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/models/dto"
)

type PdfTurtleClientInterface interface {
	// Returns PDF file generated from HTML of body, header and footer
	Render(renderData models.RenderData) (io.ReadCloser, error)
	// Returns PDF file generated from HTML of body, header and footer
	RenderWithContext(ctx context.Context, renderData models.RenderData) (io.ReadCloser, error)

	// Returns PDF file generated from HTML template plus model of body, header and footer
	RenderTemplate(renderTemplateData models.RenderTemplateData) (io.ReadCloser, error)
	// Returns PDF file generated from HTML template plus model of body, header and footer
	RenderTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData) (io.ReadCloser, error)

	// Returns information about matching model data to template
	TestTemplate(renderTemplateData models.RenderTemplateData) (*dto.TemplateTestResult, error)
	// Returns information about matching model data to template
	TestTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData) (*dto.TemplateTestResult, error)

	// Returns PDF file generated from bundle (Zip-File) of HTML or HTML template of body, header, footer and assets. The index.html file in the Zip-Bundle is required.
	RenderBundle(bundles map[string]io.Reader, model any) (io.ReadCloser, error)
	// Returns PDF file generated from bundle (Zip-File) of HTML or HTML template of body, header, footer and assets. The index.html file in the Zip-Bundle is required.
	RenderBundleWithContext(ctx context.Context, bundles map[string]io.Reader, model any) (io.ReadCloser, error)
}
