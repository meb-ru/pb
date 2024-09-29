// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chests/v1/chests.proto

package chestsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CipherChestsService_GetChests_FullMethodName       = "/chests.v1.CipherChestsService/GetChests"
	CipherChestsService_GetChestTrymap_FullMethodName  = "/chests.v1.CipherChestsService/GetChestTrymap"
	CipherChestsService_SubmitGuess_FullMethodName     = "/chests.v1.CipherChestsService/SubmitGuess"
	CipherChestsService_SetNotification_FullMethodName = "/chests.v1.CipherChestsService/SetNotification"
)

// CipherChestsServiceClient is the client API for CipherChestsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CipherChestsServiceClient interface {
	// Get all chests visible to a user
	GetChests(ctx context.Context, in *GetChestsRequest, opts ...grpc.CallOption) (*GetChestsResponse, error)
	GetChestTrymap(ctx context.Context, in *GetChestTrymapRequest, opts ...grpc.CallOption) (*GetChestTrymapResponse, error)
	SubmitGuess(ctx context.Context, in *SubmitGuessRequest, opts ...grpc.CallOption) (*SubmitGuessResponse, error)
	SetNotification(ctx context.Context, in *SetNotificationRequest, opts ...grpc.CallOption) (*SetNotificationResponse, error)
}

type cipherChestsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCipherChestsServiceClient(cc grpc.ClientConnInterface) CipherChestsServiceClient {
	return &cipherChestsServiceClient{cc}
}

func (c *cipherChestsServiceClient) GetChests(ctx context.Context, in *GetChestsRequest, opts ...grpc.CallOption) (*GetChestsResponse, error) {
	out := new(GetChestsResponse)
	err := c.cc.Invoke(ctx, CipherChestsService_GetChests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipherChestsServiceClient) GetChestTrymap(ctx context.Context, in *GetChestTrymapRequest, opts ...grpc.CallOption) (*GetChestTrymapResponse, error) {
	out := new(GetChestTrymapResponse)
	err := c.cc.Invoke(ctx, CipherChestsService_GetChestTrymap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipherChestsServiceClient) SubmitGuess(ctx context.Context, in *SubmitGuessRequest, opts ...grpc.CallOption) (*SubmitGuessResponse, error) {
	out := new(SubmitGuessResponse)
	err := c.cc.Invoke(ctx, CipherChestsService_SubmitGuess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipherChestsServiceClient) SetNotification(ctx context.Context, in *SetNotificationRequest, opts ...grpc.CallOption) (*SetNotificationResponse, error) {
	out := new(SetNotificationResponse)
	err := c.cc.Invoke(ctx, CipherChestsService_SetNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CipherChestsServiceServer is the server API for CipherChestsService service.
// All implementations must embed UnimplementedCipherChestsServiceServer
// for forward compatibility
type CipherChestsServiceServer interface {
	// Get all chests visible to a user
	GetChests(context.Context, *GetChestsRequest) (*GetChestsResponse, error)
	GetChestTrymap(context.Context, *GetChestTrymapRequest) (*GetChestTrymapResponse, error)
	SubmitGuess(context.Context, *SubmitGuessRequest) (*SubmitGuessResponse, error)
	SetNotification(context.Context, *SetNotificationRequest) (*SetNotificationResponse, error)
	mustEmbedUnimplementedCipherChestsServiceServer()
}

// UnimplementedCipherChestsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCipherChestsServiceServer struct {
}

func (UnimplementedCipherChestsServiceServer) GetChests(context.Context, *GetChestsRequest) (*GetChestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChests not implemented")
}
func (UnimplementedCipherChestsServiceServer) GetChestTrymap(context.Context, *GetChestTrymapRequest) (*GetChestTrymapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChestTrymap not implemented")
}
func (UnimplementedCipherChestsServiceServer) SubmitGuess(context.Context, *SubmitGuessRequest) (*SubmitGuessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGuess not implemented")
}
func (UnimplementedCipherChestsServiceServer) SetNotification(context.Context, *SetNotificationRequest) (*SetNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotification not implemented")
}
func (UnimplementedCipherChestsServiceServer) mustEmbedUnimplementedCipherChestsServiceServer() {}

// UnsafeCipherChestsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CipherChestsServiceServer will
// result in compilation errors.
type UnsafeCipherChestsServiceServer interface {
	mustEmbedUnimplementedCipherChestsServiceServer()
}

func RegisterCipherChestsServiceServer(s grpc.ServiceRegistrar, srv CipherChestsServiceServer) {
	s.RegisterService(&CipherChestsService_ServiceDesc, srv)
}

func _CipherChestsService_GetChests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherChestsServiceServer).GetChests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CipherChestsService_GetChests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherChestsServiceServer).GetChests(ctx, req.(*GetChestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CipherChestsService_GetChestTrymap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChestTrymapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherChestsServiceServer).GetChestTrymap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CipherChestsService_GetChestTrymap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherChestsServiceServer).GetChestTrymap(ctx, req.(*GetChestTrymapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CipherChestsService_SubmitGuess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitGuessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherChestsServiceServer).SubmitGuess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CipherChestsService_SubmitGuess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherChestsServiceServer).SubmitGuess(ctx, req.(*SubmitGuessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CipherChestsService_SetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherChestsServiceServer).SetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CipherChestsService_SetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherChestsServiceServer).SetNotification(ctx, req.(*SetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CipherChestsService_ServiceDesc is the grpc.ServiceDesc for CipherChestsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CipherChestsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chests.v1.CipherChestsService",
	HandlerType: (*CipherChestsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChests",
			Handler:    _CipherChestsService_GetChests_Handler,
		},
		{
			MethodName: "GetChestTrymap",
			Handler:    _CipherChestsService_GetChestTrymap_Handler,
		},
		{
			MethodName: "SubmitGuess",
			Handler:    _CipherChestsService_SubmitGuess_Handler,
		},
		{
			MethodName: "SetNotification",
			Handler:    _CipherChestsService_SetNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chests/v1/chests.proto",
}