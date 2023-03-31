package modules

// TemplateData holds data from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int32
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRFToken  string // Cross-Site Forgery Token
	Flash     string
	Warning   string
	Error     string
}
