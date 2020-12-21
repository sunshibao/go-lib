// Author: Qingshan Luo <edoger@qq.com>
package goqueue

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/tal-tech/go-queue/dq"

	"go-lib/log"
)

const DefaultName = "default"

type GoQueue struct {
	cs map[string]dq.Consumer
	ps map[string]dq.Producer
}

/*
q := NewGoQueue(configs)
q.PublishJSON(obj, time.Second*30)
q.PublishJSON(obj, time.Second*30, "producer-name")

type MyObject struct { ... }

q.ConsumeJSON((*MyObject)(nil), func (arg interface{}) error {
	obj := arg.(*MyObject)
	... bala bala ...
})
*/
func NewGoQueue(configs Configs) *GoQueue {
	r := &GoQueue{
		cs: make(map[string]dq.Consumer, len(configs)),
		ps: make(map[string]dq.Producer, len(configs)),
	}
	for i, j := 0, len(configs); i < j; i++ {
		r.cs[configs[i].Name] = NewConsumer(configs[i].Conf)
		r.ps[configs[i].Name] = NewProducer(configs[i].Conf)
	}
	return r
}

func (r *GoQueue) Consumer(name ...string) dq.Consumer {
	if len(name) == 0 {
		return r.cs[DefaultName]
	}
	return r.cs[name[0]]
}

func (r *GoQueue) Producer(name ...string) dq.Producer {
	if len(name) == 0 {
		return r.ps[DefaultName]
	}
	return r.ps[name[0]]
}

func (r *GoQueue) Consume(f func([]byte) error, name ...string) {
	go func() {
		r.Consumer(name...).Consume(func(body []byte) {
			defer func() {
				if v := recover(); v != nil {
					log.Errorf("GoQueue.Consume: panic %v, with message %s", v, string(body))
				}
			}()
			err := f(body)
			if err != nil {
				log.Errorf("GoQueue.Consume: error %s, with message %s", err, string(body))
			}
		})
	}()
}

func (r *GoQueue) ConsumeJSON(arg interface{}, f func(interface{}) error, name ...string) {
	rt := reflect.TypeOf(arg)
	if kind := rt.Kind(); kind != reflect.Ptr || !reflect.New(rt.Elem()).CanInterface() {
		panic("GoQueue.ConsumeJSON: invalid arguments pointer")
	}
	r.Consume(func(body []byte) error {
		obj := reflect.New(rt.Elem()).Interface()
		if err := json.Unmarshal(body, obj); err != nil {
			return err
		}
		return f(obj)
	}, name...)
}

func (r *GoQueue) Publish(msg []byte, delay time.Duration, name ...string) error {
	_, err := r.Producer(name...).Delay(msg, delay)
	return err
}

func (r *GoQueue) PublishJSON(msg interface{}, delay time.Duration, name ...string) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return r.Publish(data, delay, name...)
}
