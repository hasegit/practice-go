package adapter

func Access(c IClient) string {
	info := c.GetInfo()
	return info
}
