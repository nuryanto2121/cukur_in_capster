package fcmgetway

import (
	"context"
	"fmt"
	"log"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type SendFCM struct {
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	DeviceToken []string `json:"device_token"`
	JumlahNotif int      `json:"jumlah_notif"`
}

// func (s *SendFCM) SendPushNotification(ctx context.Context) error {
func (s *SendFCM) SendPushNotification() error {
	// [START send_multicast]
	// Create a list containing up to 100 registration tokens.
	// This registration tokens come from the client FCM SDKs.

	opt := option.WithCredentialsFile("FCM-Key.json")
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Messaging(context.Background()) //app.Messaging(ctx)
	if err != nil {
		return err
	}
	// fmt.Printf("%v", client)

	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"jumlah_notif": strconv.Itoa(s.JumlahNotif),
			// "time":  "2:45",
		},
		Tokens: s.DeviceToken,
		Notification: &messaging.Notification{
			Title: s.Title,
			Body:  s.Body,
		},
	}

	br, err := client.SendEachForMulticast(context.Background(), message)
	// br, err := client.SendMulticast(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	// See the BatchResponse reference documentation
	// for the contents of response.
	if br.FailureCount > 0 {
		// fmt.Println("%s", br.Responses)
		for _, k := range br.Responses {
			fmt.Printf("messages were failed : %s \n", k.Error.Error())
		}
	} else {
		fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
	}
	// [END send_multicast]

	return nil
}
