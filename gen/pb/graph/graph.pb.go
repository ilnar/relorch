// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: graph/graph.proto

package graph

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionResult int32

const (
	ActionResult_INVALID ActionResult = 0
	ActionResult_ERROR   ActionResult = 1
	ActionResult_SUCCESS ActionResult = 2
)

// Enum value maps for ActionResult.
var (
	ActionResult_name = map[int32]string{
		0: "INVALID",
		1: "ERROR",
		2: "SUCCESS",
	}
	ActionResult_value = map[string]int32{
		"INVALID": 0,
		"ERROR":   1,
		"SUCCESS": 2,
	}
)

func (x ActionResult) Enum() *ActionResult {
	p := new(ActionResult)
	*p = x
	return p
}

func (x ActionResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionResult) Descriptor() protoreflect.EnumDescriptor {
	return file_graph_graph_proto_enumTypes[0].Descriptor()
}

func (ActionResult) Type() protoreflect.EnumType {
	return &file_graph_graph_proto_enumTypes[0]
}

func (x ActionResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionResult.Descriptor instead.
func (ActionResult) EnumDescriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{0}
}

type TransitionCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ActionResult `protobuf:"varint,1,opt,name=result,proto3,enum=graph.ActionResult,oneof" json:"result,omitempty"`
}

func (x *TransitionCondition) Reset() {
	*x = TransitionCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionCondition) ProtoMessage() {}

func (x *TransitionCondition) ProtoReflect() protoreflect.Message {
	mi := &file_graph_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionCondition.ProtoReflect.Descriptor instead.
func (*TransitionCondition) Descriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{0}
}

func (x *TransitionCondition) GetResult() ActionResult {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return ActionResult_INVALID
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr    string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Name    string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_graph_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{1}
}

func (x *Action) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Action *Action `protobuf:"bytes,3,opt,name=action,proto3,oneof" json:"action,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_graph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_graph_graph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromId    string               `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToId      string               `protobuf:"bytes,2,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	Condition *TransitionCondition `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_graph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_graph_graph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{3}
}

func (x *Edge) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *Edge) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *Edge) GetCondition() *TransitionCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

type Graph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string  `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Edges []*Edge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Graph) Reset() {
	*x = Graph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_graph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Graph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Graph) ProtoMessage() {}

func (x *Graph) ProtoReflect() protoreflect.Message {
	mi := &file_graph_graph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Graph.ProtoReflect.Descriptor instead.
func (*Graph) Descriptor() ([]byte, []int) {
	return file_graph_graph_proto_rawDescGZIP(), []int{4}
}

func (x *Graph) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Graph) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Graph) GetEdges() []*Edge {
	if x != nil {
		return x.Edges
	}
	return nil
}

var File_graph_graph_proto protoreflect.FileDescriptor

var file_graph_graph_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x65,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x61, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x63, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2a, 0x33, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6c, 0x6e, 0x61, 0x72, 0x2f, 0x77, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graph_graph_proto_rawDescOnce sync.Once
	file_graph_graph_proto_rawDescData = file_graph_graph_proto_rawDesc
)

func file_graph_graph_proto_rawDescGZIP() []byte {
	file_graph_graph_proto_rawDescOnce.Do(func() {
		file_graph_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_graph_graph_proto_rawDescData)
	})
	return file_graph_graph_proto_rawDescData
}

var file_graph_graph_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_graph_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_graph_graph_proto_goTypes = []interface{}{
	(ActionResult)(0),           // 0: graph.ActionResult
	(*TransitionCondition)(nil), // 1: graph.TransitionCondition
	(*Action)(nil),              // 2: graph.Action
	(*Node)(nil),                // 3: graph.Node
	(*Edge)(nil),                // 4: graph.Edge
	(*Graph)(nil),               // 5: graph.Graph
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
}
var file_graph_graph_proto_depIdxs = []int32{
	0, // 0: graph.TransitionCondition.result:type_name -> graph.ActionResult
	6, // 1: graph.Action.timeout:type_name -> google.protobuf.Duration
	2, // 2: graph.Node.action:type_name -> graph.Action
	1, // 3: graph.Edge.condition:type_name -> graph.TransitionCondition
	3, // 4: graph.Graph.nodes:type_name -> graph.Node
	4, // 5: graph.Graph.edges:type_name -> graph.Edge
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_graph_graph_proto_init() }
func file_graph_graph_proto_init() {
	if File_graph_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graph_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionCondition); i {
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
		file_graph_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_graph_graph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_graph_graph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
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
		file_graph_graph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Graph); i {
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
	file_graph_graph_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_graph_graph_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_graph_graph_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_graph_graph_proto_goTypes,
		DependencyIndexes: file_graph_graph_proto_depIdxs,
		EnumInfos:         file_graph_graph_proto_enumTypes,
		MessageInfos:      file_graph_graph_proto_msgTypes,
	}.Build()
	File_graph_graph_proto = out.File
	file_graph_graph_proto_rawDesc = nil
	file_graph_graph_proto_goTypes = nil
	file_graph_graph_proto_depIdxs = nil
}