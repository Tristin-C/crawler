package utils

import (
	"github.com/bwmarrin/snowflake"
)

type IdGenService struct {
	node *snowflake.Node
}

func NewIdGenService() *IdGenService {
	instance := &IdGenService{}
	snowflake.Epoch = 1652330356000 //2022-05-13
	var err error

	snowflake.NodeBits = 6
	snowflake.StepBits = 14 //其他的不重要的id,预留多一点
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	instance.node = node
	return instance
}

func (s *IdGenService) GetId() int64 {
	return s.node.Generate().Int64()
}
