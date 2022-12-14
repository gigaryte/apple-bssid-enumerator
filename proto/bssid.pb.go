// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: bssid.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat   *int64 `protobuf:"varint,1,opt,name=lat" json:"lat,omitempty"`
	Lon   *int64 `protobuf:"varint,2,opt,name=lon" json:"lon,omitempty"`
	Unk1  *int64 `protobuf:"varint,3,opt,name=unk1" json:"unk1,omitempty"`
	Unk2  *int64 `protobuf:"varint,4,opt,name=unk2" json:"unk2,omitempty"`
	Unk3  *int64 `protobuf:"varint,5,opt,name=unk3" json:"unk3,omitempty"`
	Unk4  *int64 `protobuf:"varint,6,opt,name=unk4" json:"unk4,omitempty"`
	Unk5  *int64 `protobuf:"varint,7,opt,name=unk5" json:"unk5,omitempty"`
	Unk6  *int64 `protobuf:"varint,8,opt,name=unk6" json:"unk6,omitempty"`
	Unk7  *int64 `protobuf:"varint,9,opt,name=unk7" json:"unk7,omitempty"`
	Unk8  *int64 `protobuf:"varint,10,opt,name=unk8" json:"unk8,omitempty"`
	Unk9  *int64 `protobuf:"varint,11,opt,name=unk9" json:"unk9,omitempty"`
	Unk10 *int64 `protobuf:"varint,12,opt,name=unk10" json:"unk10,omitempty"`
	Unk11 *int64 `protobuf:"varint,21,opt,name=unk11" json:"unk11,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bssid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_bssid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_bssid_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLat() int64 {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return 0
}

func (x *Location) GetLon() int64 {
	if x != nil && x.Lon != nil {
		return *x.Lon
	}
	return 0
}

func (x *Location) GetUnk1() int64 {
	if x != nil && x.Unk1 != nil {
		return *x.Unk1
	}
	return 0
}

func (x *Location) GetUnk2() int64 {
	if x != nil && x.Unk2 != nil {
		return *x.Unk2
	}
	return 0
}

func (x *Location) GetUnk3() int64 {
	if x != nil && x.Unk3 != nil {
		return *x.Unk3
	}
	return 0
}

func (x *Location) GetUnk4() int64 {
	if x != nil && x.Unk4 != nil {
		return *x.Unk4
	}
	return 0
}

func (x *Location) GetUnk5() int64 {
	if x != nil && x.Unk5 != nil {
		return *x.Unk5
	}
	return 0
}

func (x *Location) GetUnk6() int64 {
	if x != nil && x.Unk6 != nil {
		return *x.Unk6
	}
	return 0
}

func (x *Location) GetUnk7() int64 {
	if x != nil && x.Unk7 != nil {
		return *x.Unk7
	}
	return 0
}

func (x *Location) GetUnk8() int64 {
	if x != nil && x.Unk8 != nil {
		return *x.Unk8
	}
	return 0
}

func (x *Location) GetUnk9() int64 {
	if x != nil && x.Unk9 != nil {
		return *x.Unk9
	}
	return 0
}

func (x *Location) GetUnk10() int64 {
	if x != nil && x.Unk10 != nil {
		return *x.Unk10
	}
	return 0
}

func (x *Location) GetUnk11() int64 {
	if x != nil && x.Unk11 != nil {
		return *x.Unk11
	}
	return 0
}

type BSSIDGeo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bssid    *string   `protobuf:"bytes,1,req,name=bssid" json:"bssid,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (x *BSSIDGeo) Reset() {
	*x = BSSIDGeo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bssid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSSIDGeo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSSIDGeo) ProtoMessage() {}

func (x *BSSIDGeo) ProtoReflect() protoreflect.Message {
	mi := &file_bssid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSSIDGeo.ProtoReflect.Descriptor instead.
func (*BSSIDGeo) Descriptor() ([]byte, []int) {
	return file_bssid_proto_rawDescGZIP(), []int{1}
}

func (x *BSSIDGeo) GetBssid() string {
	if x != nil && x.Bssid != nil {
		return *x.Bssid
	}
	return ""
}

func (x *BSSIDGeo) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type WiFiLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk1    *int64      `protobuf:"varint,1,opt,name=unk1" json:"unk1,omitempty"`
	Wifi    []*BSSIDGeo `protobuf:"bytes,2,rep,name=wifi" json:"wifi,omitempty"`
	Unk2    *int32      `protobuf:"varint,3,opt,name=unk2" json:"unk2,omitempty"`
	Single  *int32      `protobuf:"varint,4,opt,name=single" json:"single,omitempty"`
	APIName *string     `protobuf:"bytes,5,opt,name=APIName" json:"APIName,omitempty"`
}

func (x *WiFiLocation) Reset() {
	*x = WiFiLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bssid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WiFiLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WiFiLocation) ProtoMessage() {}

func (x *WiFiLocation) ProtoReflect() protoreflect.Message {
	mi := &file_bssid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WiFiLocation.ProtoReflect.Descriptor instead.
func (*WiFiLocation) Descriptor() ([]byte, []int) {
	return file_bssid_proto_rawDescGZIP(), []int{2}
}

func (x *WiFiLocation) GetUnk1() int64 {
	if x != nil && x.Unk1 != nil {
		return *x.Unk1
	}
	return 0
}

func (x *WiFiLocation) GetWifi() []*BSSIDGeo {
	if x != nil {
		return x.Wifi
	}
	return nil
}

func (x *WiFiLocation) GetUnk2() int32 {
	if x != nil && x.Unk2 != nil {
		return *x.Unk2
	}
	return 0
}

func (x *WiFiLocation) GetSingle() int32 {
	if x != nil && x.Single != nil {
		return *x.Single
	}
	return 0
}

func (x *WiFiLocation) GetAPIName() string {
	if x != nil && x.APIName != nil {
		return *x.APIName
	}
	return ""
}

var File_bssid_proto protoreflect.FileDescriptor

var file_bssid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x73, 0x73, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e,
	0x6b, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x6b, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x34, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x35, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e,
	0x6b, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x36, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x75, 0x6e, 0x6b, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x37, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x37, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x6b, 0x38, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x38, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x39, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e,
	0x6b, 0x39, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x6b, 0x31, 0x30, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x75, 0x6e, 0x6b, 0x31, 0x30, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x6b, 0x31,
	0x31, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x6b, 0x31, 0x31, 0x22, 0x47,
	0x0a, 0x08, 0x42, 0x53, 0x53, 0x49, 0x44, 0x47, 0x65, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x57, 0x69, 0x46, 0x69,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x12, 0x1d, 0x0a, 0x04,
	0x77, 0x69, 0x66, 0x69, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x53, 0x53,
	0x49, 0x44, 0x47, 0x65, 0x6f, 0x52, 0x04, 0x77, 0x69, 0x66, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x6b, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x50, 0x49, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_bssid_proto_rawDescOnce sync.Once
	file_bssid_proto_rawDescData = file_bssid_proto_rawDesc
)

func file_bssid_proto_rawDescGZIP() []byte {
	file_bssid_proto_rawDescOnce.Do(func() {
		file_bssid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bssid_proto_rawDescData)
	})
	return file_bssid_proto_rawDescData
}

var file_bssid_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bssid_proto_goTypes = []interface{}{
	(*Location)(nil),     // 0: Location
	(*BSSIDGeo)(nil),     // 1: BSSIDGeo
	(*WiFiLocation)(nil), // 2: WiFiLocation
}
var file_bssid_proto_depIdxs = []int32{
	0, // 0: BSSIDGeo.location:type_name -> Location
	1, // 1: WiFiLocation.wifi:type_name -> BSSIDGeo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bssid_proto_init() }
func file_bssid_proto_init() {
	if File_bssid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bssid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bssid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSSIDGeo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bssid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WiFiLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bssid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bssid_proto_goTypes,
		DependencyIndexes: file_bssid_proto_depIdxs,
		MessageInfos:      file_bssid_proto_msgTypes,
	}.Build()
	File_bssid_proto = out.File
	file_bssid_proto_rawDesc = nil
	file_bssid_proto_goTypes = nil
	file_bssid_proto_depIdxs = nil
}
