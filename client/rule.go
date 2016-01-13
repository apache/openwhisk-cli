package client

type RuleService struct {
	client *Client
}

type Rule struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	Status  string `json:"status"`
	Trigger string `json:"trigger"`
	Action  string `json:"action"`
}

type RuleListOptions struct {
	Limit string `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}
