package mq

import (
	"LeiliNetdisk/config_example"
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection
var channel *amqp.Channel

// 如果异常关闭，会接收通知
var NotifyClose chan *amqp.Error

func init() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	if !config_example.AsyncTransferEnable {
		return
	}
	if InitChannel() {
		channel.NotifyClose(NotifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-NotifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				InitChannel()
			}
		}
	}()
}

func InitChannel() bool {
	if channel != nil {
		return true
	}

	conn, err := amqp.Dial(config_example.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// Publish : 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !InitChannel() {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条小心
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}
