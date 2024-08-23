package main

func (c *RedisClient) PushToList(key string, values ...string) error {
	return c.Client.RPush(c.Ctx, key, values).Err()
}

func (c *RedisClient) GetList(key string, start, stop int64) ([]string, error) {
	return c.Client.LRange(c.Ctx, key, start, stop).Result()
}
