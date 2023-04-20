package redis

type Client struct {
	db *redis.Client
}

func GetClient() Client {
	return Client{db}
}
