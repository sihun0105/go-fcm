package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("firebaseconfig.json")
	fmt.Print((opt))
	
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Token: "eyM1nSaQTsmCtD8QpnjUcE:APA91bE6h3TEe1ROAdXaYzA0B-7GbxMJBbw8lyd3MLEwfcN_c5sb5LBxVwHxzYx2KrU5zC5M7r3GCp4oo3bxrvLRx31QJ7Ss_RbUMJHDp6DiKUhG08Xok_0FCHJezUaQu-FmGwwtwreH",
		Notification: &messaging.Notification{
			Title: "FCM Test",
			Body:  "This is a test notification",
		},
	}

	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
	}

	fmt.Println("Successfully sent message:", response)
}
