package client

type KeyValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Annotations []KeyValue

type Parameters []KeyValue

type Limits struct {
	Timeout int `json:"timeout,omitempty"`
	Memory  int `json:"memory,omitempty"`
}
