package models

type RenderData struct {
	Html       *string `json:"html" example:"<b>Hello World</b>"`
	HeaderHtml string  `json:"headerHtml,omitempty" example:"<h1>Heading</h1>"`                                                                                                        // Optional html for header. If empty, the header html will be parsed from main html (<PdfHeader></PdfHeader>).
	FooterHtml string  `json:"footerHtml,omitempty" default:"<div class=\"default-footer\"><div><span class=\"pageNumber\"></span> of <span class=\"totalPages\"></span></div></div>"` // Optional html for footer. If empty, the footer html will be parsed from main html (<PdfFooter></PdfFooter>).

	RenderOptions RenderOptions `json:"options,omitempty"`
} // @name RenderData
