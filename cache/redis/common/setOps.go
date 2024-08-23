package main

func (c *RedisClient) AddToSet(key string, members ...interface{}) error {
	return c.Client.SAdd(c.Ctx, key, members...).Err()
}

func (c *RedisClient) GetSetMembers(key string) ([]string, error) {
	return c.Client.SMembers(c.Ctx, key).Result()
}
