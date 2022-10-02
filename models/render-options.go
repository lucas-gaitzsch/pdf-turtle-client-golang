package models

type RenderOptionsMargins struct {
	// margin top in mm
	Top int `json:"top,omitempty" default:"25"`
	// margin right in mm
	Right int `json:"right,omitempty" default:"25"`
	// margin bottom in mm
	Bottom int `json:"bottom,omitempty" default:"20"`
	// margin left in mm
	Left int `json:"left,omitempty" default:"25"`
} // @name RenderOptionsMargins

type RenderOptions struct {
	Landscape            bool `json:"landscape,omitempty" default:"false"`
	ExcludeBuiltinStyles bool `json:"excludeBuiltinStyles,omitempty" default:"false"`

	// page size in mm; overrides page format
	PageSize   PageSize `json:"pageSize,omitempty"`
	PageFormat string   `json:"pageFormat,omitempty" default:"A4" enums:"A0,A1,A2,A3,A4,A5,A6,Letter,Legal"`

	// margins in mm; fallback to default if null
	Margins *RenderOptionsMargins `json:"margins,omitempty"`
} // @name RenderOptions
