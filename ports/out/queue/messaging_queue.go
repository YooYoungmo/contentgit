package queue

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type MessagingQueue interface {
	PublishMessage(ctx context.Context, queueName string, message string) error
	ReadMessage(ctx context.Context, queueName string, vt uint) (*MessageEnvelope, error)
	DeleteMessage(ctx context.Context, queueName string, msgId int64) (bool, error)
}

type MessageEnvelope struct {
	MsgId      int64
	ReadCt     int64
	EnqueuedAt time.Time
	// VT is "visibility time". The UTC timestamp at which the message will
	// be available for reading again.
	Vt      time.Time
	Message string
}

type MessagingQueueMock struct {
	mock.Mock
}

func (m *MessagingQueueMock) PublishMessage(ctx context.Context, queueName string, message string) error {
	args := m.Called(ctx, queueName, message)
	return args.Error(0)
}

func (m *MessagingQueueMock) ReadMessage(ctx context.Context, queueName string, vt uint) (*MessageEnvelope, error) {
	args := m.Called(ctx, queueName, vt)
	return args.Get(0).(*MessageEnvelope), args.Error(1)
}

func (m *MessagingQueueMock) DeleteMessage(ctx context.Context, queueName string, msgId int64) (bool, error) {
	args := m.Called(ctx, queueName, msgId)
	return args.Bool(0), args.Error(1)
}
