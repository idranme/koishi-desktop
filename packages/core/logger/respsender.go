package logger

import (
	"github.com/samber/do"
	"gopkg.ilharper.com/koi/core/god/proto"
	"gopkg.ilharper.com/koi/core/koicmd"
	"gopkg.ilharper.com/x/rpl"
	"sync"
)

type ResponseSender struct {
	c chan rpl.Log
}

func NewResponseSender(i *do.Injector) (*ResponseSender, error) {
	wg := do.MustInvoke[*sync.WaitGroup](i)

	r := &ResponseSender{
		c: make(chan rpl.Log),
	}
	ch := do.MustInvokeNamed[chan<- *proto.Response](i, koicmd.ServiceKoiCmdResponseChan)

	go func(r1 *ResponseSender, ch1 chan<- *proto.Response) {
		for {
			log := <-r1.c

			wg.Add(1)

			ch1 <- proto.NewLog(log)

			wg.Wait()
		}
	}(r, ch)

	return r, nil
}

func (responseSender *ResponseSender) Writer() chan<- rpl.Log {
	return responseSender.c
}