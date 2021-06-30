package mqttclient

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"main/dao"
	"main/model"
)

var broker = "tcp://localhost:1883"

func ConnectMqtt()(bool, mqtt.Client){
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return false, client
	}
	return true, client
}

func subCallBackFunc(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	var jsonInfo = string(msg.Payload())
	var info = model.DeviceInfo{}
	err := json.Unmarshal([]byte(jsonInfo), &info)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	dao.InsertDeviceInfo(info)
}

func Subcribe()  {
	err, client := ConnectMqtt()
	if !err{
		fmt.Println("Connect to mqtt failed!")
		return
	}
	client.Subscribe("testapp",0x00, subCallBackFunc)
}
