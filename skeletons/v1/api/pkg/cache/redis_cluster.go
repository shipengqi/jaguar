package cache

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"github.com/shipengqi/log"
)

// Client cache implement
type Client struct {
	cli    redis.UniversalClient
	locker *redislock.Client
}

// NewPool create a redis cluster pool.
func NewPool(opts *Options) *Client {
	// redisSingletonMu is locked and we know the singleton is nil
	log.Debug("Creating new Redis connection pool")

	// poolSize applies per cluster node and not for the whole cluster.
	poolSize := 500
	if opts.MaxActive > 0 {
		poolSize = opts.MaxActive
	}

	timeout := 5 * time.Second

	if opts.Timeout > 0 {
		timeout = time.Duration(opts.Timeout) * time.Second
	}

	var tlsConfig *tls.Config

	if opts.UseTLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: opts.SSLInsecureSkipVerify,
		}
	}

	var client redis.UniversalClient
	cfg := &Config{
		Addrs:           getaddrs(opts),
		MasterName:      opts.MasterName,
		Password:        opts.Password,
		DB:              opts.Database,
		DialTimeout:     timeout,
		ReadTimeout:     timeout,
		WriteTimeout:    timeout,
		ConnMaxIdleTime: 240 * timeout,
		PoolSize:        poolSize,
		TLSConfig:       tlsConfig,
	}

	if opts.MasterName != "" {
		log.Info("==> [REDIS] Creating sentinel-backed failover client")
		client = redis.NewFailoverClient(cfg.failover())
	} else if opts.EnableCluster {
		log.Info("==> [REDIS] Creating cluster client")
		client = redis.NewClusterClient(cfg.cluster())
	} else {
		log.Info("==> [REDIS] Creating single-node client")
		client = redis.NewClient(cfg.simple())
	}

	return &Client{cli: client}
}

func (c *Client) UniversalClient() redis.UniversalClient {
	return c.cli
}

func (c *Client) Ping() error {
	var err error
	_, err = c.cli.Ping(context.TODO()).Result()
	return err
}

func (c *Client) Close() error {
	return c.cli.Close()
}

// Get from key
func (c *Client) Get(key string) (string, error) {
	return c.cli.Get(context.TODO(), key).Result()
}

// Set value with key and expire time
func (c *Client) Set(key string, val interface{}, expire int) error {
	return c.cli.Set(context.TODO(), key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (c *Client) Del(key string) error {
	return c.cli.Del(context.TODO(), key).Err()
}

// HashGet from key
func (c *Client) HashGet(hk, key string) (string, error) {
	return c.cli.HGet(context.TODO(), hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (c *Client) HashDel(hk, key string) error {
	return c.cli.HDel(context.TODO(), hk, key).Err()
}

// Increase key
func (c *Client) Increase(key string) error {
	return c.cli.Incr(context.TODO(), key).Err()
}

// Decrease key
func (c *Client) Decrease(key string) error {
	return c.cli.Decr(context.TODO(), key).Err()
}

// Expire set expire time for the given key
func (c *Client) Expire(key string, dur time.Duration) error {
	return c.cli.Expire(context.TODO(), key, dur).Err()
}

// Lock tries to obtain a new lock using a key with the given TTL.
// Don't forget to release.
// See https://github.com/bsm/redislock
func (c *Client) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	if c.locker == nil {
		c.locker = redislock.New(c.cli)
	}
	return c.locker.Obtain(context.TODO(), key, time.Duration(ttl)*time.Second, options)
}
