// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.2
// source: api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskManager_GetTask_FullMethodName       = "/proto.TaskManager/GetTask"
	TaskManager_GetTaskId_FullMethodName     = "/proto.TaskManager/GetTaskId"
	TaskManager_CreateNewTask_FullMethodName = "/proto.TaskManager/CreateNewTask"
	TaskManager_UpdateTask_FullMethodName    = "/proto.TaskManager/UpdateTask"
	TaskManager_DeleteTask_FullMethodName    = "/proto.TaskManager/DeleteTask"
)

// TaskManagerClient is the client API for TaskManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagerClient interface {
	GetTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTaskResponse, error)
	GetTaskId(ctx context.Context, in *GetTaskIdRequest, opts ...grpc.CallOption) (*GetTaskIdResponse, error)
	CreateNewTask(ctx context.Context, in *CreateNewTaskRequest, opts ...grpc.CallOption) (*CreateNewTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
}

type taskManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagerClient(cc grpc.ClientConnInterface) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) GetTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, TaskManager_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) GetTaskId(ctx context.Context, in *GetTaskIdRequest, opts ...grpc.CallOption) (*GetTaskIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskIdResponse)
	err := c.cc.Invoke(ctx, TaskManager_GetTaskId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) CreateNewTask(ctx context.Context, in *CreateNewTaskRequest, opts ...grpc.CallOption) (*CreateNewTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewTaskResponse)
	err := c.cc.Invoke(ctx, TaskManager_CreateNewTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TaskManager_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, TaskManager_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagerServer is the server API for TaskManager service.
// All implementations must embed UnimplementedTaskManagerServer
// for forward compatibility.
type TaskManagerServer interface {
	GetTask(context.Context, *emptypb.Empty) (*GetTaskResponse, error)
	GetTaskId(context.Context, *GetTaskIdRequest) (*GetTaskIdResponse, error)
	CreateNewTask(context.Context, *CreateNewTaskRequest) (*CreateNewTaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	mustEmbedUnimplementedTaskManagerServer()
}

// UnimplementedTaskManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskManagerServer struct{}

func (UnimplementedTaskManagerServer) GetTask(context.Context, *emptypb.Empty) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskManagerServer) GetTaskId(context.Context, *GetTaskIdRequest) (*GetTaskIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskId not implemented")
}
func (UnimplementedTaskManagerServer) CreateNewTask(context.Context, *CreateNewTaskRequest) (*CreateNewTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTask not implemented")
}
func (UnimplementedTaskManagerServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskManagerServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTaskManagerServer) mustEmbedUnimplementedTaskManagerServer() {}
func (UnimplementedTaskManagerServer) testEmbeddedByValue()                     {}

// UnsafeTaskManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagerServer will
// result in compilation errors.
type UnsafeTaskManagerServer interface {
	mustEmbedUnimplementedTaskManagerServer()
}

func RegisterTaskManagerServer(s grpc.ServiceRegistrar, srv TaskManagerServer) {
	// If the following call pancis, it indicates UnimplementedTaskManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskManager_ServiceDesc, srv)
}

func _TaskManager_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_GetTaskId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetTaskId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_GetTaskId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetTaskId(ctx, req.(*GetTaskIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_CreateNewTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateNewTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_CreateNewTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateNewTask(ctx, req.(*CreateNewTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskManager_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManager_ServiceDesc is the grpc.ServiceDesc for TaskManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _TaskManager_GetTask_Handler,
		},
		{
			MethodName: "GetTaskId",
			Handler:    _TaskManager_GetTaskId_Handler,
		},
		{
			MethodName: "CreateNewTask",
			Handler:    _TaskManager_CreateNewTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskManager_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskManager_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
