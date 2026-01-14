package datastructures

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Notifier interface {
	Notify(ctx context.Context, msg string) error
}

type EmailNotifier struct {
	From string
}

var (
	ErrEmptyMessage = errors.New("empty message")
	ErrRateLimited  = errors.New("rate limited")
)

func (e *EmailNotifier) Notify(ctx context.Context, msg string) error {
	if msg == "" {
		return ErrEmptyMessage
	}

	// Simula errore "esterno" wrappato
	if len(msg) > 20 {
		return fmt.Errorf("email provider: %w", ErrRateLimited)
	}

	fmt.Println("EMAIL from", e.From, "->", msg)
	return nil
}

var _ Notifier = (*EmailNotifier)(nil)

// 4) Constructor that injects a dependency, so that any struct that implements the interface can be injected.
// useful when you have a service that coordinates the logic, and doesn't matter who or how is the logic executed.
// in this case let's say, the service wants to notify the message. But doesn't care if it's a mail or something else
// so the service depends only on the interface, not on the implementations
type Service struct {
	notifier Notifier
}

func NewService(notifier Notifier) *Service {
	return &Service{notifier: notifier}
}

// 5) using the interface
func (service *Service) SendWelcome(ctx context.Context, user string) error {
	msg := "Welcome " + user
	return service.notifier.Notify(ctx, msg)
}

// 6) Adapter / Decorator (wraps a Notifier and adds a behaviour)
type LoggingNotifier struct {
	next Notifier
}

func NewLoggingNotifier(next Notifier) *LoggingNotifier {
	return &LoggingNotifier{next: next}
}

func (l *LoggingNotifier) Notify(ctx context.Context, msg string) error {
	start := time.Now()
	err := l.next.Notify(ctx, msg)
	d := time.Since(start)

	if err != nil {
		log.Printf("[notify] FAIL decorating error message. msg=%q err=%v took=%s", msg, err, d)
		return err
	}
	log.Printf("[notify] OK   msg=%q took=%s", msg, d)
	return nil
}

var _ Notifier = (*LoggingNotifier)(nil)

// 8) Nil pitfall: interface nil vs concrete nil
// this is important because you have to remember that in the case2 the nil check lies to you.
// beacause the interface is not nil, but the implementation is nil.
// lesson: Never rely on interface == nil unless YOU created the interface value yourself.
func nilPitfallDemo() {
	var n1 Notifier = nil
	fmt.Println("n1 == nil:", n1 == nil) // true

	var e *EmailNotifier = nil
	var n2 Notifier = e
	fmt.Println("n2 == nil:", n2 == nil) // false (type + value)
	//THIS LINE BELOW WOULD PANIC
	//n2.Notify(context.Background(), "provaNotifyNil")
}

//  9. manual Mock for tests (without libraries)
//     (as a struct with function)
type NotifierMock struct {
	NotifyFn func(ctx context.Context, msg string) error
	Calls    []string
}

func (m *NotifierMock) Notify(ctx context.Context, msg string) error {
	m.Calls = append(m.Calls, msg)
	if m.NotifyFn == nil {
		return nil
	}
	return m.NotifyFn(ctx, msg)
}

var _ Notifier = (*NotifierMock)(nil)

// 7) example to check that the wrapping works, but the error returned is still the same
func errorHandlingDemo(n Notifier) {
	ctx := context.Background()

	err := n.Notify(ctx, "i'm the note to notify")

	// error is still the same returned from the notify
	fmt.Println("err:", err)
	if errors.Is(err, ErrRateLimited) {
		fmt.Println("-> recognized: rate limited (errors.Is)")
	}
}

func LaunchConstructorDemonstration() {
	// Costruiamo la catena: EmailNotifier -> LoggingNotifier -> Service
	email := &EmailNotifier{From: "noreply@domna.example"}
	logging := NewLoggingNotifier(email)
	svc := NewService(logging)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_ = svc.SendWelcome(ctx, "Anna")

	errorHandlingDemo(logging)
	nilPitfallDemo()

	// Mock usage (necessary for testing interfaces)
	mock := &NotifierMock{
		NotifyFn: func(ctx context.Context, msg string) error {
			if msg == "Welcome Bob" {
				return nil
			}
			return fmt.Errorf("unexpected msg: %q", msg)
		},
	}
	svc2 := NewService(mock)
	_ = svc2.SendWelcome(context.Background(), "Bob")
	fmt.Println("mock calls:", mock.Calls)
}
