package models

type RenderTemplateData struct {
	HtmlTemplate       *string `json:"htmlTemplate"`
	HeaderHtmlTemplate string  `json:"headerHtmlTemplate,omitempty"` // Optional template for header. If empty, the header template will be parsed from main template (<PdfHeader></PdfHeader>).
	FooterHtmlTemplate string  `json:"footerHtmlTemplate,omitempty"` // Optional template for footer. If empty, the footer template will be parsed from main template (<PdfFooter></PdfFooter>).

	Model any `json:"model,omitempty" swaggertype:"object"`

	TemplateEngine string `json:"templateEngine,omitempty" default:"golang" enums:"golang,handlebars,django"`

	RenderOptions RenderOptions `json:"options,omitempty"`
} // @name RenderTemplateData
