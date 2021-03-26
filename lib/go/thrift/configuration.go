package thrift

import "fmt"

const (
	DEFAULT_MAX_MESSAGE_SIZE = 100 * 1024 * 1024
)

func checkSizeForProtocol(size int32) error {
	if size < 0 {
		return NewTProtocolExceptionWithType(
			NEGATIVE_SIZE,
			fmt.Errorf("negative size: %d", size),
		)
	}
	if size > DEFAULT_MAX_MESSAGE_SIZE {
		return NewTProtocolExceptionWithType(
			SIZE_LIMIT,
			fmt.Errorf("size exceeded max allowed: %d", size),
		)
	}
	return nil
}
