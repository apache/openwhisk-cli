package whisk

type SdkService struct {
	client *Client
}

// Structure for SDK request responses
type Sdk struct {
	// TODO :: Add SDK fields
}

type SdkRequest struct {
	// TODO :: Add SDK
}

// Install artifact {component = docker || swift}
func (s *SdkService) Install(component string) {

}
