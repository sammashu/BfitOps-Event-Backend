// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/eventpb/event.proto

package eventpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Eventing struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventId              string   `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventName            string   `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Eventing) Reset()         { *m = Eventing{} }
func (m *Eventing) String() string { return proto.CompactTextString(m) }
func (*Eventing) ProtoMessage()    {}
func (*Eventing) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{0}
}

func (m *Eventing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Eventing.Unmarshal(m, b)
}
func (m *Eventing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Eventing.Marshal(b, m, deterministic)
}
func (m *Eventing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Eventing.Merge(m, src)
}
func (m *Eventing) XXX_Size() int {
	return xxx_messageInfo_Eventing.Size(m)
}
func (m *Eventing) XXX_DiscardUnknown() {
	xxx_messageInfo_Eventing.DiscardUnknown(m)
}

var xxx_messageInfo_Eventing proto.InternalMessageInfo

func (m *Eventing) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Eventing) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Eventing) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Eventing) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type EventRequest struct {
	Eventing             *Eventing `protobuf:"bytes,1,opt,name=eventing,proto3" json:"eventing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{1}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetEventing() *Eventing {
	if m != nil {
		return m.Eventing
	}
	return nil
}

type EventResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{2}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type EventManyTimesRequest struct {
	Eventing             *Eventing `protobuf:"bytes,1,opt,name=eventing,proto3" json:"eventing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventManyTimesRequest) Reset()         { *m = EventManyTimesRequest{} }
func (m *EventManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*EventManyTimesRequest) ProtoMessage()    {}
func (*EventManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{3}
}

func (m *EventManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventManyTimesRequest.Unmarshal(m, b)
}
func (m *EventManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *EventManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventManyTimesRequest.Merge(m, src)
}
func (m *EventManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_EventManyTimesRequest.Size(m)
}
func (m *EventManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventManyTimesRequest proto.InternalMessageInfo

func (m *EventManyTimesRequest) GetEventing() *Eventing {
	if m != nil {
		return m.Eventing
	}
	return nil
}

type EventManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventManyTimesResponse) Reset()         { *m = EventManyTimesResponse{} }
func (m *EventManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*EventManyTimesResponse) ProtoMessage()    {}
func (*EventManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{4}
}

func (m *EventManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventManyTimesResponse.Unmarshal(m, b)
}
func (m *EventManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *EventManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventManyTimesResponse.Merge(m, src)
}
func (m *EventManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_EventManyTimesResponse.Size(m)
}
func (m *EventManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventManyTimesResponse proto.InternalMessageInfo

func (m *EventManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type EventBiRequest struct {
	Eventing             *Eventing `protobuf:"bytes,1,opt,name=eventing,proto3" json:"eventing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventBiRequest) Reset()         { *m = EventBiRequest{} }
func (m *EventBiRequest) String() string { return proto.CompactTextString(m) }
func (*EventBiRequest) ProtoMessage()    {}
func (*EventBiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{5}
}

func (m *EventBiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventBiRequest.Unmarshal(m, b)
}
func (m *EventBiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventBiRequest.Marshal(b, m, deterministic)
}
func (m *EventBiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBiRequest.Merge(m, src)
}
func (m *EventBiRequest) XXX_Size() int {
	return xxx_messageInfo_EventBiRequest.Size(m)
}
func (m *EventBiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventBiRequest proto.InternalMessageInfo

func (m *EventBiRequest) GetEventing() *Eventing {
	if m != nil {
		return m.Eventing
	}
	return nil
}

type EventBiResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventBiResponse) Reset()         { *m = EventBiResponse{} }
func (m *EventBiResponse) String() string { return proto.CompactTextString(m) }
func (*EventBiResponse) ProtoMessage()    {}
func (*EventBiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c88e0126701e74a, []int{6}
}

func (m *EventBiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventBiResponse.Unmarshal(m, b)
}
func (m *EventBiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventBiResponse.Marshal(b, m, deterministic)
}
func (m *EventBiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBiResponse.Merge(m, src)
}
func (m *EventBiResponse) XXX_Size() int {
	return xxx_messageInfo_EventBiResponse.Size(m)
}
func (m *EventBiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventBiResponse proto.InternalMessageInfo

func (m *EventBiResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Eventing)(nil), "event.Eventing")
	proto.RegisterType((*EventRequest)(nil), "event.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "event.EventResponse")
	proto.RegisterType((*EventManyTimesRequest)(nil), "event.EventManyTimesRequest")
	proto.RegisterType((*EventManyTimesResponse)(nil), "event.EventManyTimesResponse")
	proto.RegisterType((*EventBiRequest)(nil), "event.EventBiRequest")
	proto.RegisterType((*EventBiResponse)(nil), "event.EventBiResponse")
}

func init() { proto.RegisterFile("event/eventpb/event.proto", fileDescriptor_7c88e0126701e74a) }

var fileDescriptor_7c88e0126701e74a = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xfd, 0x6d, 0x7f, 0xf6, 0xdf, 0x54, 0x5b, 0x18, 0x6d, 0x49, 0x8b, 0x85, 0x92, 0x8b, 0x15,
	0xa1, 0x86, 0xea, 0x4d, 0xf4, 0x50, 0xf4, 0xe0, 0x41, 0xc1, 0xe8, 0xc9, 0x8b, 0xa4, 0xcd, 0x50,
	0x16, 0xcc, 0x26, 0x66, 0xd3, 0x80, 0x9f, 0xd4, 0xaf, 0x23, 0x4e, 0x36, 0x21, 0x06, 0x54, 0xe8,
	0x25, 0xd9, 0x79, 0xf3, 0xe6, 0xbd, 0xb7, 0xc3, 0xc2, 0x90, 0x52, 0x52, 0xc9, 0x29, 0x7f, 0xa3,
	0x65, 0xf6, 0x9f, 0x45, 0x71, 0x98, 0x84, 0x58, 0xe7, 0xc2, 0x4e, 0xa1, 0x75, 0xf3, 0x75, 0x90,
	0x6a, 0x8d, 0x5d, 0xa8, 0x49, 0xdf, 0x12, 0x13, 0x31, 0x6d, 0xbb, 0x35, 0xe9, 0xe3, 0x10, 0x5a,
	0x4c, 0x7a, 0x91, 0xbe, 0x55, 0x63, 0xb4, 0xc9, 0xf5, 0xad, 0x8f, 0x63, 0x80, 0xac, 0xa5, 0xbc,
	0x80, 0xac, 0xff, 0xdc, 0x6c, 0x33, 0x72, 0xef, 0x05, 0x84, 0x13, 0xe8, 0xf8, 0xa4, 0x57, 0xb1,
	0x8c, 0x12, 0x19, 0x2a, 0x6b, 0x87, 0xfb, 0x65, 0xc8, 0xbe, 0x80, 0x5d, 0xf6, 0x75, 0xe9, 0x6d,
	0x43, 0x3a, 0xc1, 0x13, 0xe3, 0x25, 0xd5, 0x9a, 0x13, 0x74, 0xe6, 0xbd, 0x59, 0x16, 0x37, 0x8f,
	0xe7, 0x16, 0x04, 0xfb, 0x08, 0xf6, 0xcc, 0xb0, 0x8e, 0x42, 0xa5, 0x09, 0x07, 0xd0, 0x88, 0x49,
	0x6f, 0x5e, 0x13, 0x93, 0xde, 0x54, 0xf6, 0x35, 0xf4, 0x99, 0x78, 0xe7, 0xa9, 0xf7, 0x27, 0x19,
	0x90, 0xde, 0xca, 0xce, 0x81, 0x41, 0x55, 0xe5, 0x0f, 0xdf, 0x4b, 0xe8, 0xf2, 0xc4, 0x42, 0x6e,
	0x65, 0x78, 0x0c, 0xbd, 0x62, 0xfc, 0x77, 0xa7, 0xf9, 0x87, 0x30, 0x8b, 0x7c, 0xa4, 0x38, 0x95,
	0x2b, 0xc2, 0x73, 0xa8, 0x73, 0x8d, 0xfb, 0x65, 0x7d, 0x13, 0x63, 0x74, 0xf0, 0x1d, 0xcc, 0xc4,
	0xed, 0x7f, 0xf8, 0x60, 0x02, 0x17, 0x57, 0xc4, 0xc3, 0x32, 0xb3, 0xba, 0xbf, 0xd1, 0xf8, 0x87,
	0x6e, 0x2e, 0xe8, 0x08, 0xbc, 0x82, 0xa6, 0xb9, 0x04, 0xf6, 0xcb, 0xec, 0x62, 0x27, 0xa3, 0x41,
	0x15, 0xce, 0xa7, 0xa7, 0xc2, 0x11, 0x8b, 0xf6, 0x73, 0xd3, 0xbc, 0xdb, 0x65, 0x83, 0x9f, 0xec,
	0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x7e, 0x0b, 0x70, 0xcf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	//Unary
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	EventManyTimes(ctx context.Context, in *EventManyTimesRequest, opts ...grpc.CallOption) (EventService_EventManyTimesClient, error)
	EventBi(ctx context.Context, opts ...grpc.CallOption) (EventService_EventBiClient, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/event.EventService/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) EventManyTimes(ctx context.Context, in *EventManyTimesRequest, opts ...grpc.CallOption) (EventService_EventManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventService_serviceDesc.Streams[0], "/event.EventService/EventManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceEventManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_EventManyTimesClient interface {
	Recv() (*EventManyTimesResponse, error)
	grpc.ClientStream
}

type eventServiceEventManyTimesClient struct {
	grpc.ClientStream
}

func (x *eventServiceEventManyTimesClient) Recv() (*EventManyTimesResponse, error) {
	m := new(EventManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) EventBi(ctx context.Context, opts ...grpc.CallOption) (EventService_EventBiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventService_serviceDesc.Streams[1], "/event.EventService/EventBi", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceEventBiClient{stream}
	return x, nil
}

type EventService_EventBiClient interface {
	Send(*EventBiRequest) error
	Recv() (*EventBiResponse, error)
	grpc.ClientStream
}

type eventServiceEventBiClient struct {
	grpc.ClientStream
}

func (x *eventServiceEventBiClient) Send(m *EventBiRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceEventBiClient) Recv() (*EventBiResponse, error) {
	m := new(EventBiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	//Unary
	Event(context.Context, *EventRequest) (*EventResponse, error)
	EventManyTimes(*EventManyTimesRequest, EventService_EventManyTimesServer) error
	EventBi(EventService_EventBiServer) error
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_EventManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).EventManyTimes(m, &eventServiceEventManyTimesServer{stream})
}

type EventService_EventManyTimesServer interface {
	Send(*EventManyTimesResponse) error
	grpc.ServerStream
}

type eventServiceEventManyTimesServer struct {
	grpc.ServerStream
}

func (x *eventServiceEventManyTimesServer) Send(m *EventManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventService_EventBi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).EventBi(&eventServiceEventBiServer{stream})
}

type EventService_EventBiServer interface {
	Send(*EventBiResponse) error
	Recv() (*EventBiRequest, error)
	grpc.ServerStream
}

type eventServiceEventBiServer struct {
	grpc.ServerStream
}

func (x *eventServiceEventBiServer) Send(m *EventBiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceEventBiServer) Recv() (*EventBiRequest, error) {
	m := new(EventBiRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _EventService_Event_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventManyTimes",
			Handler:       _EventService_EventManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventBi",
			Handler:       _EventService_EventBi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "event/eventpb/event.proto",
}
