package mqtt

import (
	"fmt"
	"net/http"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MQTTClient mqtt.Client

func connectHandler(client mqtt.Client) {
	fmt.Println("Connected")
}

func connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func Publish(message string) {
	token := MQTTClient.Publish("/movements", 0, false, message)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func messagePubHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Published message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
}

func InitializeMQTTClient() {
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
	MQTTClient = mqtt.NewClient(opts)

	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// messageHandler adlı bir işlev tanımlanır. Bu işlev, bir MQTT mesajını aldığınızda çağrılır ve konu ve mesajı yazdırır,
// ardından sendPayloadToHTTPChannel işlevini çağırarak HTTP kanalına mesajı gönderir.
func messageHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
	// Send the payload message to an HTTP channel
	payload := message.Payload()
	sendPayloadToHTTPChannel(payload)
}

// Bu işlev, bir dizeyi HTTP kanalına POST isteği olarak göndermek için kullanılır. Olası hatalar hata kontrolü ile ele alınır.
func sendPayloadToHTTPChannel(payload []byte) {
	// Replace 'http_channel_url' with the actual URL of your HTTP channel
	url := "http://http_channel_url"
	// Convert payload to string
	payloadStr := string(payload)
	// http.Post işlevi kullanılarak HTTP POST isteği oluşturulur.
	// İsteğin gönderileceği URL, veri türü ("application/json") ve gönderilecek veri içeriği ("strings.NewReader(payloadStr)") belirtilir.
	resp, err := http.Post(url, "application/json", strings.NewReader(payloadStr))
	if err != nil {
		fmt.Println("Error sending payload to HTTP channel:", err)
		return
	}
	// Yanıtın (response) sızdırmadan kapatılması için defer kullanılır.?
	defer resp.Body.Close()

	//HTTP yanıtının durumu, resp.StatusCode ile kontrol edilir.
	//Eğer yanıtın durumu 200 (HTTP OK) değilse, bir hata mesajı yazdırılır. HTTP isteği başarısız olduğunda bu durumda bilgilendirilirsiniz
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed with status code: %d\n", resp.StatusCode)
	}
}
func main() {
	// Initialize the MQTT client
	InitializeMQTTClient()
	// Subscribe to the MQTT client
	//topic string: Bu parametre, abone olunacak MQTT konusunu temsil eder.
	// qos 2: Mesajlar yalnızca bir kez teslim edilir ve herhangi bir tekrarlama veya kaybolma olasılığı yoktur.
	// MQTT istemcisi /movemente abone oluyor ve yeni mesajlar geldiğinde func(client mqtt.Client, message mqtt.Message) işlevi çağrılıyor.
	token := MQTTClient.Subscribe("/movement", byte(2), func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("Received message on topic %s: %s\n", message.Topic(), message.Payload())
		// Received message on topic /movement: This is the message payload.
	})
	token.Wait()
	// Publish a movement command
	Publish("forward")
}
