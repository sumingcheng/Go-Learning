package main

func (c *RedisClient) GetSortedSetMembers(key string) ([]string, error) {
	return c.Client.ZRange(c.Ctx, key, 0, -1).Result()
}
