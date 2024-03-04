package main

import (
	"context"
	"fmt"
	"log"

	"gofcm/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func main()  {
	firebaseConfig, err := config.ReadFirebaseConfig("firebaseconfig.json")
	if err != nil {
		log.Fatalf("error reading firebase config: %v\n", err)
	}
	app := initializeAppWithServiceAccount(firebaseConfig)
	sendToToken(app)
}

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	registrationToken := "YOUR_REGISTRATION_TOKEN"

	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}

func initializeAppWithServiceAccount(opt *config.FirebaseConfig) *firebase.App {
	
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}