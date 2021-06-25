// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/driver"
)

type Session struct {
	InfoStub        func() view.SessionInfo
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns     struct {
		result1 view.SessionInfo
	}
	infoReturnsOnCall map[int]struct {
		result1 view.SessionInfo
	}
	SendStub        func(payload []byte) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		payload []byte
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	SendErrorStub        func(payload []byte) error
	sendErrorMutex       sync.RWMutex
	sendErrorArgsForCall []struct {
		payload []byte
	}
	sendErrorReturns struct {
		result1 error
	}
	sendErrorReturnsOnCall map[int]struct {
		result1 error
	}
	ReceiveStub        func() <-chan *view.Message
	receiveMutex       sync.RWMutex
	receiveArgsForCall []struct{}
	receiveReturns     struct {
		result1 <-chan *view.Message
	}
	receiveReturnsOnCall map[int]struct {
		result1 <-chan *view.Message
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Session) Info() view.SessionInfo {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.recordInvocation("Info", []interface{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.infoReturns.result1
}

func (fake *Session) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *Session) InfoReturns(result1 view.SessionInfo) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 view.SessionInfo
	}{result1}
}

func (fake *Session) InfoReturnsOnCall(i int, result1 view.SessionInfo) {
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 view.SessionInfo
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 view.SessionInfo
	}{result1}
}

func (fake *Session) Send(payload []byte) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		payload []byte
	}{payloadCopy})
	fake.recordInvocation("Send", []interface{}{payloadCopy})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(payload)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendReturns.result1
}

func (fake *Session) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *Session) SendArgsForCall(i int) []byte {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return fake.sendArgsForCall[i].payload
}

func (fake *Session) SendReturns(result1 error) {
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *Session) SendReturnsOnCall(i int, result1 error) {
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Session) SendError(payload []byte) error {
	var payloadCopy []byte
	if payload != nil {
		payloadCopy = make([]byte, len(payload))
		copy(payloadCopy, payload)
	}
	fake.sendErrorMutex.Lock()
	ret, specificReturn := fake.sendErrorReturnsOnCall[len(fake.sendErrorArgsForCall)]
	fake.sendErrorArgsForCall = append(fake.sendErrorArgsForCall, struct {
		payload []byte
	}{payloadCopy})
	fake.recordInvocation("SendError", []interface{}{payloadCopy})
	fake.sendErrorMutex.Unlock()
	if fake.SendErrorStub != nil {
		return fake.SendErrorStub(payload)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendErrorReturns.result1
}

func (fake *Session) SendErrorCallCount() int {
	fake.sendErrorMutex.RLock()
	defer fake.sendErrorMutex.RUnlock()
	return len(fake.sendErrorArgsForCall)
}

func (fake *Session) SendErrorArgsForCall(i int) []byte {
	fake.sendErrorMutex.RLock()
	defer fake.sendErrorMutex.RUnlock()
	return fake.sendErrorArgsForCall[i].payload
}

func (fake *Session) SendErrorReturns(result1 error) {
	fake.SendErrorStub = nil
	fake.sendErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *Session) SendErrorReturnsOnCall(i int, result1 error) {
	fake.SendErrorStub = nil
	if fake.sendErrorReturnsOnCall == nil {
		fake.sendErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Session) Receive() <-chan *view.Message {
	fake.receiveMutex.Lock()
	ret, specificReturn := fake.receiveReturnsOnCall[len(fake.receiveArgsForCall)]
	fake.receiveArgsForCall = append(fake.receiveArgsForCall, struct{}{})
	fake.recordInvocation("Receive", []interface{}{})
	fake.receiveMutex.Unlock()
	if fake.ReceiveStub != nil {
		return fake.ReceiveStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.receiveReturns.result1
}

func (fake *Session) ReceiveCallCount() int {
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	return len(fake.receiveArgsForCall)
}

func (fake *Session) ReceiveReturns(result1 <-chan *view.Message) {
	fake.ReceiveStub = nil
	fake.receiveReturns = struct {
		result1 <-chan *view.Message
	}{result1}
}

func (fake *Session) ReceiveReturnsOnCall(i int, result1 <-chan *view.Message) {
	fake.ReceiveStub = nil
	if fake.receiveReturnsOnCall == nil {
		fake.receiveReturnsOnCall = make(map[int]struct {
			result1 <-chan *view.Message
		})
	}
	fake.receiveReturnsOnCall[i] = struct {
		result1 <-chan *view.Message
	}{result1}
}

func (fake *Session) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *Session) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Session) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.sendErrorMutex.RLock()
	defer fake.sendErrorMutex.RUnlock()
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Session) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ driver.Session = new(Session)