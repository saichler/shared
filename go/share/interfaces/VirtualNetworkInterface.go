package interfaces

import (
	"github.com/saichler/shared/go/types"
)

type IVirtualNetworkInterface interface {
	Start()
	Shutdown()
	Name() string
	SendMessage([]byte) error
	Unicast(types.Action, string, interface{}) error
	Multicast(CastMode, types.Action, int32, string, interface{}) error
	Request(CastMode, types.Action, int32, string, interface{}) error
	API(int32) API
	Resources() IResources
}

type API interface {
	Post(interface{}) (interface{}, error)
	Put(interface{}) (interface{}, error)
	Patch(interface{}) (interface{}, error)
	Delete(interface{}) (interface{}, error)
	Get(string) (interface{}, error)
}

type CastMode int

const (
	All    CastMode = 0
	Single CastMode = 1
	Leader CastMode = 2
)

type IDatatListener interface {
	ShutdownVNic(IVirtualNetworkInterface)
	HandleData([]byte, IVirtualNetworkInterface)
	Failed([]byte, IVirtualNetworkInterface, string)
}

func NewVNicConfig(maxDataSize uint64, txQueueSize, rxQueueSize uint64, vNetPort uint32) *types.VNicConfig {
	mc := &types.VNicConfig{
		MaxDataSize: maxDataSize,
		TxQueueSize: txQueueSize,
		RxQueueSize: rxQueueSize,
		VnetPort:    vNetPort,
	}
	return mc
}
