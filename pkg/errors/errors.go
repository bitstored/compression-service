package errors

import (
	"fmt"
	"strings"
)

type Kind int

// Error represents errors for the service.
type Err struct {
	Msgs []string
	Kind Kind
}

const (
	KindOther                Kind = iota // Unclassified error.
	KindEmptyImage                       // Empty Image
	KindNotPermitedOperation             // Not permited operation
	KindOverflow                         // buffer owerflow
	KindMalformedStructure
)

func NewError(kind Kind, msg string) *Err {
	return &Err{Msgs: []string{msg}, Kind: kind}
}

func NewErrorArr(kind Kind, msgs []string) *Err {
	return &Err{Msgs: msgs, Kind: kind}
}

func (e *Err) AddMessage(msg string) {
	e.Msgs = append(e.Msgs, msg)
}

func (e *Err) AddMessageArr(msgs []string) {
	e.Msgs = append(e.Msgs, msgs...)
}

func (e *Err) Message() string {
	return fmt.Sprintf("%s : %s", e.KindMsg(), strings.Join(e.Msgs, "\n"))
}

func (e *Err) Error() error {
	if e == nil || len(e.Msgs) == 0 {
		return nil
	}
	return fmt.Errorf("%s : %s", e.KindMsg(), strings.Join(e.Msgs, "\n"))
}

func (e *Err) KindMsg() string {
	switch e.Kind {
	case KindOther:
		return "other error"
	case KindEmptyImage:
		return "image is nil or empty"
	case KindNotPermitedOperation:
		return "operation is not permited"
	case KindOverflow:
		return "buffer/matrix owerflow"
	case KindMalformedStructure:
		return "malformed buffer structure"
	default:
		return "something is wrong"
	}
}
