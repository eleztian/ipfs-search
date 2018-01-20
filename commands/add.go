package commands

import (
	"github.com/ipfs-search/ipfs-search/queue"
)

// AddHash queues a single IPFS hash for indexing
func AddHash(hash string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	conn, err := queue.NewConnection(config.AMQPURL)
	if err != nil {
		return err
	}

	ch, err := conn.NewChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.NewQueue("hashes")
	if err != nil {
		return err
	}

	err = queue.Publish(map[string]interface{}{
		"hash": hash,
	})

	return err
}