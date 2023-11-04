package core_func

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type TestServer struct {
	fu          *FuncRouter
	operationID string
	sessionID   string
	platformID  string
}

const (
	operationID string = "123456"
	sessionID   string = "111"
	platformID  string = "1"
)

func NewTestServer() *TestServer {
	ev := make(chan *EventData, 10)
	return &TestServer{
		operationID: operationID,
		sessionID:   sessionID,
		platformID:  platformID,
		fu:          NewFuncRouter(ev, sessionID),
	}
}

func TestInitSDK(t *testing.T) {
	te := NewTestServer()
	fn := func() bool {
		te.fu.InitSDK(te.operationID, te.platformID)
		msg, err := <-te.fu.respMessage.respMessagesChan

		ret := &EventData{
			OperationID: te.operationID,
			Event:       "InitSDK",
			Data:        "",
		}
		assert.Equal(t, true, err)
		assert.Equal(t, ret, msg)
		return true
	}
	err := quick.Check(fn, nil)
	assert.Nil(t, err)
}
