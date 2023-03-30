package modules

// TemplateData holds data from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int32
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRFToken  string
	Flash     string
	Warning   string
	Error     string
}
