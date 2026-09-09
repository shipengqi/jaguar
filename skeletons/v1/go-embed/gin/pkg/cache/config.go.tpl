package cache

import "github.com/redis/go-redis/v9"

// Config is the overridden type of redis.UniversalOptions. simple(), failover() and cluster() functions are not public in redis
// library.
// Therefore, they are redefined in here to use in creation of new redis cluster logic.
// We don't want to use redis.NewUniversalClient() logic.
type Config redis.UniversalOptions

func (c *Config) cluster() *redis.ClusterOptions {
	if len(c.Addrs) == 0 {
		c.Addrs = []string{"127.0.0.1:6379"}
	}

	return &redis.ClusterOptions{
		Addrs:     c.Addrs,
		OnConnect: c.OnConnect,

		Password: c.Password,

		MaxRedirects:   c.MaxRedirects,
		ReadOnly:       c.ReadOnly,
		RouteByLatency: c.RouteByLatency,
		RouteRandomly:  c.RouteRandomly,

		MaxRetries:      c.MaxRetries,
		MinRetryBackoff: c.MinRetryBackoff,
		MaxRetryBackoff: c.MaxRetryBackoff,

		DialTimeout:     c.DialTimeout,
		ReadTimeout:     c.ReadTimeout,
		WriteTimeout:    c.WriteTimeout,
		PoolSize:        c.PoolSize,
		MinIdleConns:    c.MinIdleConns,
		MaxIdleConns:    c.MaxIdleConns,
		PoolTimeout:     c.PoolTimeout,
		ConnMaxIdleTime: c.ConnMaxIdleTime,
		ConnMaxLifetime: c.ConnMaxLifetime,

		TLSConfig: c.TLSConfig,
	}
}

func (c *Config) simple() *redis.Options {
	addr := "127.0.0.1:6379"
	if len(c.Addrs) > 0 {
		addr = c.Addrs[0]
	}

	return &redis.Options{
		Addr:      addr,
		OnConnect: c.OnConnect,

		DB:       c.DB,
		Password: c.Password,

		MaxRetries:      c.MaxRetries,
		MinRetryBackoff: c.MinRetryBackoff,
		MaxRetryBackoff: c.MaxRetryBackoff,

		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,

		PoolSize:        c.PoolSize,
		MinIdleConns:    c.MinIdleConns,
		MaxIdleConns:    c.MaxIdleConns,
		PoolTimeout:     c.PoolTimeout,
		ConnMaxIdleTime: c.ConnMaxIdleTime,
		ConnMaxLifetime: c.ConnMaxLifetime,

		TLSConfig: c.TLSConfig,
	}
}

func (c *Config) failover() *redis.FailoverOptions {
	if len(c.Addrs) == 0 {
		c.Addrs = []string{"127.0.0.1:26379"}
	}

	return &redis.FailoverOptions{
		SentinelAddrs: c.Addrs,
		MasterName:    c.MasterName,
		OnConnect:     c.OnConnect,

		DB:       c.DB,
		Password: c.Password,

		MaxRetries:      c.MaxRetries,
		MinRetryBackoff: c.MinRetryBackoff,
		MaxRetryBackoff: c.MaxRetryBackoff,

		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,

		PoolSize:        c.PoolSize,
		MinIdleConns:    c.MinIdleConns,
		MaxIdleConns:    c.MaxIdleConns,
		PoolTimeout:     c.PoolTimeout,
		ConnMaxIdleTime: c.ConnMaxIdleTime,
		ConnMaxLifetime: c.ConnMaxLifetime,

		TLSConfig: c.TLSConfig,
	}
}
