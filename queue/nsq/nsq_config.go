package nsq

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// ----------------------------------------
//  NSQ 队列配置项集合
// ----------------------------------------
type NSQConfigs struct {
	Items map[string]*NSQConfig `toml:"nsq"`
}

// 绑定数据到当前 Redis 配置清单
func (configs *NSQConfigs) Bind(bs []byte) error {
	if err := toml.Unmarshal(bs, configs); err != nil {
		return fmt.Errorf("the nsq configs bind failed: %s", err)
	}
	return nil
}

// ----------------------------------------
//  NSQ 队列配置项（TOML 支持, Viper 支持）
// ----------------------------------------
type NSQConfig struct {
	Topic              string `toml:"topic" mapstructure:"topic"`                             // Topic 名称，消费者和生产者必须
	Channel            string `toml:"channel" mapstructure:"channel"`                         // Channel 名称，消费者必须
	ConsumerCount      int    `toml:"consumer_count" mapstructure:"consumer_count"`           // 同时开启的消费者客户端数量（默认 1）
	ConsumerConcurrent int    `toml:"consumer_concurrent" mapstructure:"consumer_concurrent"` // 每一个消费者的消息处理并发数（默认 1）
	Address            string `toml:"address" mapstructure:"address"`                         // NSQD 或者 Loockup 地址（多个地址使用逗号连接）
	Lookup             bool   `toml:"lookup" mapstructure:"lookup"`                           // Address 是否是 Loockup 地址
	MsgTimeout         string `toml:"msg_timeout" mapstructure:"msg_timeout"`                 // msg 超时时间
}

// 获取一个 NSQ 队列配置项的完整副本
func (config *NSQConfig) Copy() *NSQConfig {
	return &NSQConfig{
		Topic:              config.Topic,
		Channel:            config.Channel,
		ConsumerCount:      config.ConsumerCount,
		ConsumerConcurrent: config.ConsumerConcurrent,
		Address:            config.Address,
		Lookup:             config.Lookup,
		MsgTimeout:         config.MsgTimeout,
	}
}
