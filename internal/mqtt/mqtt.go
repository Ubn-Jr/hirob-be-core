package mqtt

import (
	"fmt"
	"sync"
	"time"

	mq "github.com/eclipse/paho.mqtt.golang"
)

var (
	client          mq.Client
	clientInitMutex sync.Mutex
)

// messagePubHandler is a variable of type mq.MessageHandler that handles incoming MQTT messages.
// It prints the topic and payload of the message to the console.
var messagePubHandler mq.MessageHandler = func(client mq.Client, msg mq.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())
}

// connectHandler is a callback function that is executed when the MQTT client is connected.
var connectHandler mq.OnConnectHandler = func(client mq.Client) {
	fmt.Println("Mqtt client is Connected Successfully")
}

// connectLostHandler is a callback function that is triggered when the MQTT client loses connection.
var connectLostHandler mq.ConnectionLostHandler = func(client mq.Client, err error) {
	fmt.Printf("Connection lost due to: %+v", err)
}

// createClient creates and returns an MQTT client instance.
// If the client instance already exists, it returns the existing instance.
// The client is configured with the broker IP address, port, username, and password.
// It also sets the default publish handler, connect handler, and connection lost handler.
// The client is then connected to the MQTT broker.
// If the connection fails, it panics with the error message.
// Returns the MQTT client instance.
func createClient() mq.Client {
	clientInitMutex.Lock()
	defer clientInitMutex.Unlock()

	if client != nil {
		return client
	}

	var broker = "51.120.244.98"
	var port = 1883
	opts := mq.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mq_client")
	opts.SetUsername("hirob")
	opts.SetPassword("ubnjr")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client = mq.NewClient(opts)
	print("Client created\n")
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	print("Client connected\n")
	return client
}

// Publish sends the given message to the "/movements" topic using MQTT.
// It creates a client, publishes the message with the specified quality of service (qos),
// and waits until the message is delivered. If there is an error during publishing,
// it prints the error message.
func Publish(message string) {
	var client = createClient()
	token := client.Publish("/sub/movements", 2, false, message)
	//  (retained parameter (false): This behavior is useful for messages that have no lasting significance or are only relevant at the time of publication.)
	token.Wait() // This will block until the message is delivered
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
}

// Subscribe is a function that creates a client and subscribes to a topic in MQTT.
// It listens for new messages on the subscribed topic and prints them to the console.
func Subscribe() {
	var client = createClient()
	// The MQTT client subscribes to the "/movements" topic and when new messages arrive, the function func(client mq.Client, message mq.Message) is called.
	token := client.Subscribe("/pub/movements", 2, func(client mq.Client, message mq.Message) {
		fmt.Printf("Received message on topic %s: %s\n, at %s", message.Topic(), message.Payload(), time.Now().Format("2006-01-02 15:04:05"))
		// Received message on topic /movement: This is the message payload.
	})
	token.Wait() // This will block until the Subscribe is successful

	if token.Error() != nil {
		fmt.Println(token.Error())
	}

}
