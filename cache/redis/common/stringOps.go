package main

func (c *RedisClient) SetString(key string, value string) error {
	return c.Client.Set(c.Ctx, key, value, 0).Err()
}

func (c *RedisClient) GetString(key string) (string, error) {
	return c.Client.Get(c.Ctx, key).Result()
}
