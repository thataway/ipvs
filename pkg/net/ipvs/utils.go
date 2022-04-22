package ipvs

import (
	"github.com/thataway/protos/pkg/api/ipvs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	//String2NetworkTransport String to NetworkTransport
	String2NetworkTransport = make(map[string]ipvs.NetworkTransport)
	//NetworkTransport2String NetworkTransport to String
	NetworkTransport2String = make(map[ipvs.NetworkTransport]string)

	//String2ScheduleMethod String to ScheduleMethod
	String2ScheduleMethod = make(map[string]ipvs.ScheduleMethod)
	//ScheduleMethod2String ScheduleMethod to String
	ScheduleMethod2String = make(map[ipvs.ScheduleMethod]string)

	//PacketFwdMethod2String PacketFwdMethod to String
	PacketFwdMethod2String = make(map[ipvs.PacketFwdMethod]string)
	//String2PacketFwdMethod String to PacketFwdMethod
	String2PacketFwdMethod = make(map[string]ipvs.PacketFwdMethod)
)

type protoEnumWalker struct {
	protoreflect.EnumDescriptor
}

func (walker protoEnumWalker) traverseValues(extInfo *protoimpl.ExtensionInfo,
	f func(value protoreflect.EnumNumber, valueExt interface{})) {

	vals := walker.EnumDescriptor.Values()
	for i := 0; i < vals.Len(); i++ {
		v := vals.Get(i)
		var valueExt interface{}
		if extInfo != nil {
			switch o := v.Options().(type) {
			case *descriptorpb.EnumValueOptions:
				valueExt = proto.GetExtension(o, extInfo)
			}
		}
		f(v.Number(), valueExt)
	}
}

func init() {
	protoEnumWalker{EnumDescriptor: ipvs.NetworkTransport(0).Descriptor()}.
		traverseValues(ipvs.E_Transport,
			func(value protoreflect.EnumNumber, valueExt interface{}) {
				switch s := valueExt.(type) {
				case string:
					String2NetworkTransport[s] = ipvs.NetworkTransport(value)
					NetworkTransport2String[ipvs.NetworkTransport(value)] = s
				}
			},
		)

	protoEnumWalker{EnumDescriptor: ipvs.ScheduleMethod(0).Descriptor()}.
		traverseValues(ipvs.E_ScheduleAlg,
			func(value protoreflect.EnumNumber, valueExt interface{}) {
				switch s := valueExt.(type) {
				case string:
					ScheduleMethod2String[ipvs.ScheduleMethod(value)] = s
					String2ScheduleMethod[s] = ipvs.ScheduleMethod(value)
				}
			},
		)

	protoEnumWalker{EnumDescriptor: ipvs.PacketFwdMethod(0).Descriptor()}.
		traverseValues(ipvs.E_FwdAlg,
			func(value protoreflect.EnumNumber, valueExt interface{}) {
				switch s := valueExt.(type) {
				case string:
					PacketFwdMethod2String[ipvs.PacketFwdMethod(value)] = s
					String2PacketFwdMethod[s] = ipvs.PacketFwdMethod(value)
				}
			},
		)
}
