package whisk

// NOTE :: deprecated
type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Annotations []KeyValue

// type Annotations map[string]interface{}

type Parameters []KeyValue

// type Parameters map[string]interface{}

type Limits struct {
	Timeout int `json:"timeout,omitempty"`
	Memory  int `json:"memory,omitempty"`
}
