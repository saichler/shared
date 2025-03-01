package interfaces

import (
	"github.com/saichler/shared/go/types"
)
import "github.com/google/uuid"

type IResources interface {
	Registry() IRegistry
	ServicePoints() IServicePoints
	Security() ISecurityProvider
	DataListener() IDatatListener
	Serializer(SerializerMode) ISerializer
	Logger() ILogger
	Config() *types.VNicConfig
	Introspector() IIntrospector
	AddTopic(int32, string)
}

func AddTopic(config *types.VNicConfig, vlan int32, topic string) {
	if config == nil {
		return
	}
	if config.LocalUuid == "" {
		config.LocalUuid = NewUuid()
	}
	if config.Topics == nil {
		config.Topics = &types.Topics{}
		config.Topics.VlanToTopic = make(map[int32]string)
	}
	config.Topics.VlanToTopic[vlan] = topic
}

func NewUuid() string {
	return uuid.New().String()
}
