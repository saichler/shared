package infra

import (
	"errors"
	"github.com/saichler/shared/go/share/interfaces"
	"github.com/saichler/shared/go/types"
	"google.golang.org/protobuf/proto"
)

var Log interfaces.ILogger

type TestServicePointHandler struct {
	Name         string
	PostNumber   int
	PutNumber    int
	PatchNumber  int
	DeleteNumber int
	GetNumber    int
	FailedNumber int
	Tr           bool
	ErrorMode    bool
}

const (
	TEST_TOPIC = "TestProto"
)

func NewTestServicePointHandler(name string) *TestServicePointHandler {
	tsp := &TestServicePointHandler{}
	tsp.Name = name
	return tsp
}

func (tsp *TestServicePointHandler) Post(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Post -", tsp.Name, "- Test callback")
	tsp.PostNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Post - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) Put(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Put -", tsp.Name, "- Test callback")
	tsp.PutNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Put - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) Patch(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Patch -", tsp.Name, "- Test callback")
	tsp.PatchNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Patch - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) Delete(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Delete -", tsp.Name, "- Test callback")
	tsp.DeleteNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Delete - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) GetCopy(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Get -", tsp.Name, "- Test callback")
	tsp.GetNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("GetCopy - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) Get(pb proto.Message, resourcs interfaces.IResources) (proto.Message, error) {
	Log.Debug("Get -", tsp.Name, "- Test callback")
	tsp.GetNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Get - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) Failed(pb proto.Message, resourcs interfaces.IResources, info *types.Message) (proto.Message, error) {
	dest := "n/a"
	msg := "n/a"
	if info != nil {
		dest = info.SourceUuid
		msg = info.FailMsg
	}
	Log.Debug("Failed -", tsp.Name, " to ", dest, "- Test callback")
	Log.Debug("Failed Reason is ", msg)
	tsp.FailedNumber++
	var err error
	if tsp.ErrorMode {
		err = errors.New("Failed - TestServicePointHandler Error")
	}
	return pb, err
}
func (tsp *TestServicePointHandler) EndPoint() string {
	return "/Tests"
}
func (tsp *TestServicePointHandler) Topic() string {
	return TEST_TOPIC
}
func (tsp *TestServicePointHandler) Transactional() bool {
	return tsp.Tr
}
