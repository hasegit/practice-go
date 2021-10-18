package infra

type Client struct {
	api string
}

func (c Client) GetInfo() string {
	return c.api
}

func GetClient(api string) Client {
	c := Client{api}
	return c
}
