// Author: Qingshan Luo <edoger@qq.com>
package goqueue

import (
    "github.com/tal-tech/go-queue/dq"
)

/**
GoQueue:
  - Name: default
    Conf:
      Beanstalks:
        - Endpoint: 127.0.0.1:11300
          Tube: tube
        - Endpoint: 127.0.0.1:11301
          Tube: tube
      Redis:
        Host: 127.0.0.1:6379
        Type: node
        Pass: ''
*/
type Configs []Config

type Config struct {
    Name string
    Conf dq.DqConf
}