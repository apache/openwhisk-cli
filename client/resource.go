package client

// resourceService is a generic CRUD mixin for other services to provide basic CRUD functionality.
type resourceService struct {
	resourceName string
	client       *Client
}

func (s resourceService) List() {

}

// TODO :: resource service functions (Create, update, list etc...)
