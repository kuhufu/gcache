package gcache

import (
	"fmt"
	"strconv"
)

var _ = Result(memResult{})

type memResult struct {
	reply interface{}
	err   error
}

func (r memResult) Value() (interface{}, error) {
	return r.reply, r.err
}

func (r memResult) Bool() (reply bool, err error) {
	if r.err != nil {
		return false, r.err
	}

	switch reply := r.reply.(type) {
	case bool:
		return reply, nil
	case int:
		return reply != 0, nil
	case string:
		return strconv.ParseBool(reply)
	case []byte:
		return strconv.ParseBool(string(reply))
	default:
		return false, fmt.Errorf("unexpected type for bool, got type %T", reply)
	}
}

func (r memResult) Int() (reply int, err error) {
	if r.err != nil {
		return 0, r.err
	}

	switch reply := r.reply.(type) {
	case int:
		return reply, nil
	case int64:
		return int(reply), nil
	case string:
		return strconv.Atoi(reply)
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 0)
		return int(n), err
	default:
		return 0, fmt.Errorf("unexpected type for int, got type %T", reply)
	}
}

func (r memResult) Int64() (reply int64, err error) {
	if r.err != nil {
		return 0, r.err
	}

	switch reply := r.reply.(type) {
	case int64:
		return reply, nil
	case int:
		return int64(reply), nil
	case int32:
		return int64(reply), nil
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 64)
		return n, err
	default:
		return 0, fmt.Errorf("unexpected type for int64, got type %T", reply)
	}
}

func (r memResult) Float64() (reply float64, err error) {
	if r.err != nil {
		return 0, r.err
	}

	switch reply := r.reply.(type) {
	case float64:
		return reply, nil
	case float32:
		return float64(reply), nil
	case int32:
		return float64(reply), nil
	default:
		return 0, fmt.Errorf("unexpected type for float64, got type %T", reply)
	}
}

func (r memResult) String() (reply string, err error) {
	if r.err != nil {
		return "", r.err
	}

	switch reply := r.reply.(type) {
	case string:
		return reply, nil
	case []byte:
		return string(reply), nil
	case int64:
		return strconv.Itoa(int(reply)), nil
	case int:
		return strconv.Itoa(reply), nil
	case int32:
		return strconv.Itoa(int(reply)), nil
	default:
		return "", fmt.Errorf("unexpected type for string, got type %T", reply)
	}
}

func (r memResult) Bytes() (reply []byte, err error) {
	if r.err != nil {
		return nil, r.err
	}
	switch reply := r.reply.(type) {
	case []byte:
		return reply, nil
	case string:
		return []byte(reply), nil
	default:
		return nil, fmt.Errorf("unexpected type for []byte, got type %T", reply)
	}
}

func (r memResult) Reply() (reply interface{}) {
	return r.reply
}

func (r memResult) Error() (err error) {
	return r.err
}
