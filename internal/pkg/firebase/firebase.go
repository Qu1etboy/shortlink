package firebase

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Firebase() (*firebase.App, *firestore.Client, context.Context) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("firebase_credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return app, client, ctx
}
