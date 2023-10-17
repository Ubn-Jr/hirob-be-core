package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %+v", err)
}

func createClient() mqtt.Client {
	var broker = "broker_url"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("admin")
	opts.SetPassword("instar")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func Publish(message string) {
	var client = createClient()
	token := client.Publish("/movements", 2, false, message)
	//  (retained parameter (false): This behavior is useful for messages that have no lasting significance or are only relevant at the time of publication.)
	token.Wait() // This will block until the message is delivered
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
	// client.Disconnect(250)
}

