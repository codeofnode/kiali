package models

// JaegerInfo provides information to access Jaeger UI
type JaegerInfo struct {
	URL               string `json:"url"`
	NamespaceSelector bool   `json:"namespace_selector"`
}
