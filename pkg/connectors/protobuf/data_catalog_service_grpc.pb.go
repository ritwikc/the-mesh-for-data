// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// DataCatalogServiceClient is the client API for DataCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataCatalogServiceClient interface {
	GetDatasetInfo(ctx context.Context, in *CatalogDatasetRequest, opts ...grpc.CallOption) (*CatalogDatasetInfo, error)
	RegisterDatasetInfo(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error)
}

type dataCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCatalogServiceClient(cc grpc.ClientConnInterface) DataCatalogServiceClient {
	return &dataCatalogServiceClient{cc}
}

func (c *dataCatalogServiceClient) GetDatasetInfo(ctx context.Context, in *CatalogDatasetRequest, opts ...grpc.CallOption) (*CatalogDatasetInfo, error) {
	out := new(CatalogDatasetInfo)
	err := c.cc.Invoke(ctx, "/connectors.DataCatalogService/GetDatasetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCatalogServiceClient) RegisterDatasetInfo(ctx context.Context, in *RegisterAssetRequest, opts ...grpc.CallOption) (*RegisterAssetResponse, error) {
	out := new(RegisterAssetResponse)
	err := c.cc.Invoke(ctx, "/connectors.DataCatalogService/RegisterDatasetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataCatalogServiceServer is the server API for DataCatalogService service.
// All implementations must embed UnimplementedDataCatalogServiceServer
// for forward compatibility
type DataCatalogServiceServer interface {
	GetDatasetInfo(context.Context, *CatalogDatasetRequest) (*CatalogDatasetInfo, error)
	RegisterDatasetInfo(context.Context, *RegisterAssetRequest) (*RegisterAssetResponse, error)
	mustEmbedUnimplementedDataCatalogServiceServer()
}

// UnimplementedDataCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataCatalogServiceServer struct {
}

func (UnimplementedDataCatalogServiceServer) GetDatasetInfo(context.Context, *CatalogDatasetRequest) (*CatalogDatasetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetInfo not implemented")
}
func (UnimplementedDataCatalogServiceServer) RegisterDatasetInfo(context.Context, *RegisterAssetRequest) (*RegisterAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDatasetInfo not implemented")
}
func (UnimplementedDataCatalogServiceServer) mustEmbedUnimplementedDataCatalogServiceServer() {}

// UnsafeDataCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataCatalogServiceServer will
// result in compilation errors.
type UnsafeDataCatalogServiceServer interface {
	mustEmbedUnimplementedDataCatalogServiceServer()
}

func RegisterDataCatalogServiceServer(s grpc.ServiceRegistrar, srv DataCatalogServiceServer) {
	s.RegisterService(&DataCatalogService_ServiceDesc, srv)
}

func _DataCatalogService_GetDatasetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatalogDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServiceServer).GetDatasetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.DataCatalogService/GetDatasetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServiceServer).GetDatasetInfo(ctx, req.(*CatalogDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCatalogService_RegisterDatasetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServiceServer).RegisterDatasetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.DataCatalogService/RegisterDatasetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServiceServer).RegisterDatasetInfo(ctx, req.(*RegisterAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataCatalogService_ServiceDesc is the grpc.ServiceDesc for DataCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.DataCatalogService",
	HandlerType: (*DataCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatasetInfo",
			Handler:    _DataCatalogService_GetDatasetInfo_Handler,
		},
		{
			MethodName: "RegisterDatasetInfo",
			Handler:    _DataCatalogService_RegisterDatasetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_catalog_service.proto",
}
