package utils

import (
	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
)

func Init(sf *sonyflake.Sonyflake) *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(
		sonyflake.Settings{
			MachineID: awsutil.AmazonEC2MachineID,
		})
}

func Sonyflake(sf *sonyflake.Sonyflake) (uint64, error) {
	return sf.NextID()
}
