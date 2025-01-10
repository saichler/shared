package type_registry

import (
	"github.com/saichler/shared/go/share/interfaces"
	"reflect"
)

/*
Type info contains only two attributes, the golang type of, usually, a struct.
And a Serializer, which indicates a which serializer should be used to transalte
this type into bytes. Serializer can be nil...
*/
type TypeInfo struct {
	/* The golang reflect type */
	typ reflect.Type
	/* The serializers */
	serializers map[interfaces.SerializerMode]interfaces.Serializer
}

// NewTypeInfo /* Constructs a new type info with the given attributes */
func NewTypeInfo(typ reflect.Type) *TypeInfo {
	if typ == nil {
		interfaces.Error("Cannot register a nil type")
		return nil
	}
	return &TypeInfo{typ: typ, serializers: make(map[interfaces.SerializerMode]interfaces.Serializer)}
}

// Type /* Return the reflect type of this TypeInfo */
func (info *TypeInfo) Type() reflect.Type {
	return info.typ
}

// Name /* Return the name of the type/struct */
func (info *TypeInfo) Name() string {
	return info.typ.Name()
}

// Serializer /* Return the serializer to be used with this type. Can be nil... */
func (info *TypeInfo) Serializer(mode interfaces.SerializerMode) interfaces.Serializer {
	return info.serializers[mode]
}

func (info *TypeInfo) AddSerializer(serializer interfaces.Serializer) {
	info.serializers[serializer.Mode()] = serializer
}

// NewInstance /* Construct a new instance via reflect using the type */
func (info *TypeInfo) NewInstance() (interface{}, error) {
	n := reflect.New(info.typ)
	if !n.IsValid() {
		return nil, interfaces.Error("was not able to create new instance of type ", info.typ.Name())
	}
	return n.Interface(), nil
}
