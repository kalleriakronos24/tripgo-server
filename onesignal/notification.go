package onesignal

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/kalleriakronos24/khaimal-group/config"
)

var configuration = onesignal.NewConfiguration()
var apiClient = onesignal.NewAPIClient(configuration)

func PushNotificationSingleExternalId(deviceId string, msgParam string) {
	appId := config.AppConfig.OneSignalAppID
	restApiKey := config.AppConfig.OneSignalRestApiKey
	osAuthCtx := context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		restApiKey,
	)

	notification := *onesignal.NewNotification(appId)
	notification.IncludeExternalUserIds = []string{deviceId}
	notification.SetIsAndroid(true)
	// notification.SetIsIos(true)
	notification.SetPriority(10)
	message := msgParam
	stringMap := onesignal.StringMap{En: &message}
	notification.Contents = *onesignal.NewNullableStringMap(&stringMap)

	resp, r, err := apiClient.DefaultApi.CreateNotification(osAuthCtx).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Notification ID: %v\n", resp.GetId())
}

// TODO
func PushNotificationMultipleExternalId(deviceId []string) {
	appId := os.Getenv("ONESIGNAL_APP_ID")
	restApiKey := os.Getenv("ONESIGNAL_REST_API_KEY")
	osAuthCtx := context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		restApiKey,
	)

	notification := *onesignal.NewNotification(appId)
	notification.IncludeExternalUserIds = deviceId
	notification.SetIncludedSegments([]string{"Subscribed Users"})
	// notification.SetIsIos(false)
	message := "Go Test Notification"
	stringMap := onesignal.StringMap{En: &message}
	notification.Contents = *onesignal.NewNullableStringMap(&stringMap)

	resp, r, err := apiClient.DefaultApi.CreateNotification(osAuthCtx).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Notification ID: %v\n", resp.GetId())
}
