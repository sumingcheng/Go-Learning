package main

func (c *RedisClient) SetHash(key string, values map[string]interface{}) error {
	return c.Client.HSet(c.Ctx, key, values).Err()
}

func (c *RedisClient) GetHash(key string) (map[string]string, error) {
	return c.Client.HGetAll(c.Ctx, key).Result()
}
