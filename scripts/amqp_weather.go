package scripts

import (
	"basic_framework/configs"
	"basic_framework/core/log"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"pack.ag/amqp"
	"time"
)

// 气象amqp
func AmqpWeatherExecute() {

	if configs.Yaml.Amqp.Weather.Enable {
		w := configs.Yaml.Amqp.Weather
		timestamp := time.Now().UnixMilli()

		userName := fmt.Sprintf("%s|authMode=aksign,signMethod=Hmacsha1,consumerGroupId=%s,authId=%s,iotInstanceId=%s,timestamp=%d|", w.ClientId, w.ConsumerGroupId, w.AccessKey, w.IotInstanceId, timestamp)
		stringToSign := fmt.Sprintf("authId=%s&timestamp=%d", w.AccessKey, timestamp)

		hmacKey := hmac.New(sha1.New, []byte(w.AccessSecret))
		hmacKey.Write([]byte(stringToSign))
		password := base64.StdEncoding.EncodeToString(hmacKey.Sum(nil))

		am := &amqpManager{
			address:  w.Address,
			userName: userName,
			password: password,
		}

		go am.startReceiveMessage(context.Background())
	}

}

type amqpManager struct {
	address  string
	userName string
	password string
	receiver *amqp.Receiver
	session  *amqp.Session
	client   *amqp.Client
}

// 开始监听
func (am *amqpManager) startReceiveMessage(ctx context.Context) {
	childCtx, _ := context.WithCancel(ctx)

	// 开始连接服务
	err := am.generateReceiverWithRetry(childCtx)
	if err != nil {
		log.Error("连接amqp_weather失败，err: ", err)
		return
	}

	defer func() {
		am.receiver.Close(childCtx)
		am.session.Close(childCtx)
		am.client.Close()
	}()

	for {
		// 阻塞接收消息
		message, err := am.receiver.Receive(ctx)

		// 获取成功
		if err == nil {
			// 处理信息
			go am.processMessage(message)

			message.Accept()
		} else {
			log.Error("amqp_weather receive data error: ", err)

			// 主动取消，退出
			select {
			case <-childCtx.Done():
				return
			default:
			}

			// 重新建立连接
			err := am.generateReceiverWithRetry(childCtx)
			if err != nil {
				log.Error("连接amqp_weather失败，err: ", err)
				return
			}

		}

	}

}

// 连接服务
func (am *amqpManager) generateReceiverWithRetry(ctx context.Context) error {
	// 重连时间，10ms依次*2，直到20s
	duration := 10 * time.Microsecond
	maxDuration := 20000 * time.Microsecond
	// 重连次数
	times := 1

	for {
		select {
		case <-ctx.Done():
			return amqp.ErrConnClosed
		default:
		}

		// 开始连接
		err := am.generateReceiver()
		if err != nil {
			time.Sleep(duration)
			if duration < maxDuration {
				duration *= 2
			}
			log.Error("amqp_wather retry, times: ", times, ", duration: ", duration)
			times++
		} else {
			log.Info("amqp_wather init success")
			return nil
		}

	}
}

func (am *amqpManager) generateReceiver() error {
	if am.session != nil {
		receiver, err := am.session.NewReceiver(
			amqp.LinkSourceAddress(configs.Yaml.Amqp.Weather.QueueName),
			amqp.LinkCredit(20),
		)

		// 如果断网等行为发生，Connection会关闭导致Session建立失败，未关闭连接则建立成功。
		if err == nil {
			am.receiver = receiver
			return nil
		}
	}

	// 清理上一个连接
	if am.client != nil {
		am.client.Close()
	}

	client, err := amqp.Dial(am.address, amqp.ConnSASLPlain(am.userName, am.password))
	if err != nil {
		return err
	}
	am.client = client

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	am.session = session

	receiver, err := am.session.NewReceiver(
		amqp.LinkSourceAddress(configs.Yaml.Amqp.Weather.QueueName),
		amqp.LinkCredit(20),
	)
	am.receiver = receiver

	return nil
}

// 处理信息
func (am *amqpManager) processMessage(message *amqp.Message) {
	fmt.Println("这是信息：", string(message.GetData()), "   properties: ", message.ApplicationProperties)
}
