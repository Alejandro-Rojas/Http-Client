package gohttp

type httpClient struct{}

fun new() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	GET()
	Post()
	Put()
	patch()
	delete()
}

func (c *httpClient) GET() {
	response, err := c.Do("GET")
}


func (c *httpClient) Post() {}

func (c *httpClient) Put() {}

func (c *httpClient) patch() {}

func (c *httpClient) delete() {}
