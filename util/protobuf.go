package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (handler *Util) ConvertStringToTimestampPb(s string) (*timestamppb.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, err
	}
	pb := timestamppb.New(t)
	return pb, pb.CheckValid()
}
