package client

type ActivationService struct {
	client *Client
}

type Activation struct {
	Namespace    string `json:"namespace,omitempty"`
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	Publish      bool   `json:"publish,omitempty"`
	Subject      string `json:"subject,omitempty"`
	ActivationID string `json:"activationId,omitempty"`
	Start        string `json:"start,omitempty"`
	End          string `json:"end,omitempty"`
	Result       `json:"result,omitempty"`
	Logs         string `json:"logs,omitempty"`
}

type Result struct {
	Status string                 `json:"status,omitempty"`
	Value  map[string]interface{} `json:"value,omitempty"`
}

type ActivationRequest struct {
}
