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
func TimestampToTime(ts *timestamp.Timestamp) (t *time.Time, err error) {
	if ts != nil {
		t1, _ := ptypes.Timestamp(ts)
		t = &t1
	} else {
		return nil, nil
	}
	return
}
