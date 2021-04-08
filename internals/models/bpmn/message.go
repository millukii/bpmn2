package bpmn

import (
	"sync"
)

type  Message struct {
	id string
	name string
	messageRef MessageFlow

}

var once sync.Once
var message *Message

func NewMessage(id, name string, messageRef MessageFlow) *Message {
    if message == nil {
        once.Do(
            func() {
                message = &Message{
					id, name, messageRef,
				}
            })
    } 
    return message
}

func (m Message) GetId() string{
	return m.id
}

func (m *Message) SetId(id string) {
	m.id = id
}
func (m Message) GetName() string{
	return m.name
}

func (m *Message) SetName(name string) {
	m.name = name
}

func (m Message) GetMessageRef() MessageFlow{
	return m.messageRef
}

func (m *Message) SetMessageRef(messageRef MessageFlow) {
	m.messageRef = messageRef
}