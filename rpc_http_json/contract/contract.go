package contract

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}
