package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func TimeToTimestamp(t *time.Time) (ts *timestamp.Timestamp, err error) {
	if t != nil {
		ts, err = ptypes.TimestampProto(*t)
	} else {
		return nil, nil
	}
	return
}
