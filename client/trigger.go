package client

type TriggerService struct {
	client *Client
}

type Trigger struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	ID          string `json:"id"`
	Annotations `json:"annotations"`
	Parameters  `json:"parameters"`
	Limits      `json:"limits"`
}

type TriggerListOptions struct {
	Limit string `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}
