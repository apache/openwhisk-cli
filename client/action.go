package client

type ActionService struct {
	client *Client
}

type Action struct {
	Namespace   string `json:"namespace,omitempty"`
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Publish     bool   `json:"publish,omitempty"`
	Exec        `json:"exec,omitempty"`
	Annotations `json:"annotations,omitempty"`
	Parameters  `json:"parameters,omitempty"`
	Limits      `json:"limits,omitempty"`
}

type Exec struct {
	Code  string `json:"code,omitempty"`
	Image string `json:"image,omitempty"`
	Init  string `json:"init,omitempty"`
}

type ActionRequest struct {
}

////////////////////
// Action Methods //
////////////////////

func (s *ActionService) Create() {

}
