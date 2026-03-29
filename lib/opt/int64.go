package libopt

import (
	"flag"
	"strconv"
)

type OptionalInt64 struct {
	value     int64
	something bool
}

var _ flag.Value = &OptionalInt64{}

func NoInt64() OptionalInt64 {
	return OptionalInt64{}
}

func SomeInt64(value int64) OptionalInt64 {
	return OptionalInt64{
		something: true,
		value: value,
	}
}

func (receiver OptionalInt64) Get() (int64, bool) {
	return receiver.value, receiver.something
}

func (receiver OptionalInt64) IsNothing() bool {
	return !receiver.something
}

func (receiver OptionalInt64) IsSomething() bool {
	return receiver.something
}

func (receiver *OptionalInt64) Set(value string) error {
	if nil == receiver {
		return ErrNilReceiver
	}

	i64, err := strconv.ParseInt(value, 10, 64)
	if nil != err {
		return err
	}

	*receiver = SomeInt64(i64)

	return nil
}

func (receiver OptionalInt64) String() string {
	if !receiver.something {
		return ""
	}

	return strconv.FormatInt(receiver.value, 10)
}
