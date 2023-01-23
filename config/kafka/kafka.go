package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv()  {
	godotenv.Load("misc/.env")
}

func ConnProducer() *kafka.Producer{
	loadEnv()
	kafkaServer := os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func ConnConsumer() *kafka.Consumer{
	loadEnv()
	kafkaServer := os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          "group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	return c

}