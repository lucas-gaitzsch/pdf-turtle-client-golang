package pdfturtleclient

import (
	"context"
	"io"

	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-dotnet/models/dto"
)

type PdfTurtleClientInterface interface {
	Render(renderData models.RenderData)(io.ReadCloser, error)
	RenderWithContext(ctx context.Context, renderData models.RenderData)(io.ReadCloser, error)

	RenderTemplate(renderTemplateData models.RenderTemplateData)(io.ReadCloser, error)
	RenderTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData)(io.ReadCloser, error)

	TestTemplate(renderTemplateData models.RenderTemplateData)(*dto.TemplateTestResult, error)
	TestTemplateWithContext(ctx context.Context, renderTemplateData models.RenderTemplateData)(*dto.TemplateTestResult, error)

	RenderBundle(bundles []io.Reader, model any)(io.ReadCloser, error)
	RenderBundleWithContext(ctx context.Context, bundles []io.Reader, model any)(io.ReadCloser, error)
}
