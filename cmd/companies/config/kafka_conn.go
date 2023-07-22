package config

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fakovacic/companies-service/internal/companies/errors"
	"github.com/segmentio/kafka-go"
)

const (
	retryKafkaTimeout time.Duration = 5 * time.Second
	writeTimeout      time.Duration = 10 * time.Second
)

func NewKafkaConn() (*kafka.Conn, error) {
	host := os.Getenv("KAFKA_HOST")
	if host == "" {
		return nil, errors.New("host is empty")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		return nil, errors.New("topic is empty")
	}

	partition := os.Getenv("KAFKA_PARTITION")
	if partition == "" {
		return nil, errors.New("partition is empty")
	}

	partitionValue, err := strconv.Atoi(partition)
	if err != nil {
		return nil, errors.Wrap(err, "partition conversion error")
	}

	conn, err := retryKafkaConn(host, topic, partitionValue)
	if err != nil {
		return nil, errors.Wrap(err, "dial")
	}

	return conn, nil
}

func retryKafkaConn(host, topic string, partitionValue int) (*kafka.Conn, error) {
	for i := 0; i <= 3; i++ {
		conn, err := kafka.DialLeader(context.Background(), "tcp", host, topic, partitionValue)
		if err != nil {
			log.Printf("kafka connection error: %v", err)
			time.Sleep(retryKafkaTimeout)

			continue
		}

		if err == nil {
			return conn, nil
		}

		time.Sleep(retryKafkaTimeout)
	}

	return nil, errors.New("kafka connection retry exceded")
}
