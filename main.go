package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB bağlantı ayarları
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB'ye bağlan
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Bağlantıyı test et
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB'ye başarıyla bağlandı!")

	// ... MongoDB bağlantısı kodu ...

	// Kafka Producer ayarları
	brokerList := []string{"localhost:9092"}
	topic := "my-topic"

	// Kafka'ya bağlan
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokerList,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Mesajı gönder
	err = writer.WriteMessages(context.TODO(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte("Hello, Kafka!"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mesaj başarıyla Kafka'ya gönderildi!")

	// Kafka bağlantısını kapat
	writer.Close()
}
