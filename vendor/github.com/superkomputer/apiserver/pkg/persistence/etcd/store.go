package etcd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

const (
	//EtcdLocal is entrpiont of local etcd instance
	EtcdLocal = "http://127.0.0.1:2379"
)

// Store is a store implementation backed by etcd
type Store struct {
	Endpoints []string
	Client    *client.Client
}

// NewStore creates a Store instance
func NewStore(endpoints []string) *Store {
	s := &Store{
		Endpoints: endpoints,
	}

	config := client.Config{
		Endpoints:               s.Endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(config)
	if err != nil {
		log.Fatal(err)
	}
	s.Client = &c

	return s
}

// Create key-value in etcd instance
func (s Store) Create(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	api := client.NewKeysAPI(*s.Client)
	_, err := api.Create(ctx, key, value)
	if err != nil {
		log.Fatal(err)
	}
}

// Get value by key in etcd instance
func (s Store) Get(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	api := client.NewKeysAPI(*s.Client)
	resp, err := api.Get(ctx, key, nil)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Node.Value
}

// Update key-value in etcd instance
func (s Store) Update(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	api := client.NewKeysAPI(*s.Client)
	_, err := api.Set(ctx, key, value, nil)
	if err != nil {
		if err == context.DeadlineExceeded {
			// TODO request took longer than 5s
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
}

// Delete value by key in etcd instance
func (s Store) Delete(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	api := client.NewKeysAPI(*s.Client)
	_, err := api.Delete(ctx, key, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type value struct {
	data string
	ttl  time.Duration
}

func (v *value) TTL() time.Duration {
	return v.ttl
}

func (v *value) Unmarshal(out interface{}) error {
	return json.Unmarshal([]byte(v.data), out)
}

func ttl2Dur(ttl int64, err error) time.Duration {
	if err != nil || ttl < 0 {
		return time.Duration(-1)
	}
	return time.Duration(ttl) * time.Millisecond
}

func dur2TTL(dur time.Duration) int64 {
	return int64(dur / time.Millisecond)
}
