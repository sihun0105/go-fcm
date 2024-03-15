package main

import (
	"context"
	"fmt"
	"log"

	"gofcm/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	firebaseConfig, err := config.ReadFirebaseConfig("firebaseconfig.json")
	if err != nil {
		log.Fatalf("error reading firebase config: %v\n", err)
	}
	opt := option.WithCredentialsJSON([]byte(fmt.Sprintf(`{
		"type": "%s",
		"project_id": "%s",
		"private_key_id": "%s",
		"private_key": "%s",
		"client_email": "%s",
		"client_id": "%s",
		"auth_uri": "%s",
		"token_uri": "%s",
		"auth_provider_x509_cert_url": "%s",
		"client_x509_cert_url": "%s",
	}`, firebaseConfig.Type, firebaseConfig.ProjectId, firebaseConfig.PrivateKeyId, firebaseConfig.PrivateKey, firebaseConfig.ClientEmail, firebaseConfig.ClientId, firebaseConfig.AuthUri, firebaseConfig.TokenUri, firebaseConfig.AuthProviderX509CertUrl, firebaseConfig.ClientX509CertUrl)))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Token: "ftnnBejlzk1HlvfZIW9GK0:APA91bFWmtw4wZPanV8iNtDZtQAYJFRs36sQsVKNO9lgAEz_OAiz5XCJVKv7u3yR1FDd9Uu9k8hzABcCHAu--VeQI_NXz1ZtORTK_Wmdy852MN-WfdlxLcvgRmyIMxmilCbQs49sompY",
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
