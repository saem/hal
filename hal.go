package hal

type HalLink struct {
	Href      string `json:"href"`
	Name      string `json:"name,omitempty"`
	Hreflang  string `json:"hreflang,omitempty"`
	Title     string `json:"title,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}

type HalResource struct {
	Href     string `json:"href"`
	Name     string `json:"name,omitempty"`
	Hreflang string `json:"hreflang,omitempty"`
	Title    string `json:"title,omitempty"`
}
