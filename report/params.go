package report

type Filter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GenerateReportParams struct {
	ForUserID     string `json:"-"`
	Type          string `json:"type" validate:"required"`
	Filter        Filter `json:"filter" validate:"required"`
	Format        string `json:"format"`
	Currency      string `json:"currency"`
	ReportVersion string `json:"report_version"`
}

type GetReportParams struct {
	ForUserID string `json:"-"`
	ReportID  string `json:"report_id" validate:"required"`
}
