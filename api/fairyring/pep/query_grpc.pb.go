// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: fairyring/pep/query.proto

package pep

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
	Query_Params_FullMethodName                   = "/fairyring.pep.Query/Params"
	Query_EncryptedTx_FullMethodName              = "/fairyring.pep.Query/EncryptedTx"
	Query_EncryptedTxAll_FullMethodName           = "/fairyring.pep.Query/EncryptedTxAll"
	Query_EncryptedTxAllFromHeight_FullMethodName = "/fairyring.pep.Query/EncryptedTxAllFromHeight"
	Query_LatestHeight_FullMethodName             = "/fairyring.pep.Query/LatestHeight"
	Query_PepNonce_FullMethodName                 = "/fairyring.pep.Query/PepNonce"
	Query_PepNonceAll_FullMethodName              = "/fairyring.pep.Query/PepNonceAll"
	Query_PubKey_FullMethodName                   = "/fairyring.pep.Query/PubKey"
	Query_KeyshareReq_FullMethodName              = "/fairyring.pep.Query/KeyshareReq"
	Query_KeyshareReqAll_FullMethodName           = "/fairyring.pep.Query/KeyshareReqAll"
	Query_PrivateKeyshareReq_FullMethodName       = "/fairyring.pep.Query/PrivateKeyshareReq"
	Query_DecryptData_FullMethodName              = "/fairyring.pep.Query/DecryptData"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a EncryptedTx by index.
	EncryptedTx(ctx context.Context, in *QueryEncryptedTxRequest, opts ...grpc.CallOption) (*QueryEncryptedTxResponse, error)
	// Queries a list of EncryptedTx items.
	EncryptedTxAll(ctx context.Context, in *QueryEncryptedTxAllRequest, opts ...grpc.CallOption) (*QueryEncryptedTxAllResponse, error)
	// Queries a list of EncryptedTx items.
	EncryptedTxAllFromHeight(ctx context.Context, in *QueryEncryptedTxAllFromHeightRequest, opts ...grpc.CallOption) (*QueryEncryptedTxAllFromHeightResponse, error)
	// Queries a list of LatestHeight items.
	LatestHeight(ctx context.Context, in *QueryLatestHeightRequest, opts ...grpc.CallOption) (*QueryLatestHeightResponse, error)
	// Queries a PepNonce by index.
	PepNonce(ctx context.Context, in *QueryPepNonceRequest, opts ...grpc.CallOption) (*QueryPepNonceResponse, error)
	// Queries a list of PepNonce items.
	PepNonceAll(ctx context.Context, in *QueryPepNonceAllRequest, opts ...grpc.CallOption) (*QueryPepNonceAllResponse, error)
	// Queries the public keys
	PubKey(ctx context.Context, in *QueryPubKeyRequest, opts ...grpc.CallOption) (*QueryPubKeyResponse, error)
	// Queries a Keyshare request by identity
	KeyshareReq(ctx context.Context, in *QueryKeyshareReqRequest, opts ...grpc.CallOption) (*QueryKeyshareReqResponse, error)
	// queries a list of keyshare requests
	KeyshareReqAll(ctx context.Context, in *QueryKeyshareReqAllRequest, opts ...grpc.CallOption) (*QueryKeyshareReqAllResponse, error)
	// Queries a list of ShowPrivateKeyshareReq items.
	PrivateKeyshareReq(ctx context.Context, in *QueryPrivateKeyshareReqRequest, opts ...grpc.CallOption) (*QueryPrivateKeyshareReqResponse, error)
	// Queries a list of DecryptData items.
	DecryptData(ctx context.Context, in *QueryDecryptDataRequest, opts ...grpc.CallOption) (*QueryDecryptDataResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncryptedTx(ctx context.Context, in *QueryEncryptedTxRequest, opts ...grpc.CallOption) (*QueryEncryptedTxResponse, error) {
	out := new(QueryEncryptedTxResponse)
	err := c.cc.Invoke(ctx, Query_EncryptedTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncryptedTxAll(ctx context.Context, in *QueryEncryptedTxAllRequest, opts ...grpc.CallOption) (*QueryEncryptedTxAllResponse, error) {
	out := new(QueryEncryptedTxAllResponse)
	err := c.cc.Invoke(ctx, Query_EncryptedTxAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EncryptedTxAllFromHeight(ctx context.Context, in *QueryEncryptedTxAllFromHeightRequest, opts ...grpc.CallOption) (*QueryEncryptedTxAllFromHeightResponse, error) {
	out := new(QueryEncryptedTxAllFromHeightResponse)
	err := c.cc.Invoke(ctx, Query_EncryptedTxAllFromHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LatestHeight(ctx context.Context, in *QueryLatestHeightRequest, opts ...grpc.CallOption) (*QueryLatestHeightResponse, error) {
	out := new(QueryLatestHeightResponse)
	err := c.cc.Invoke(ctx, Query_LatestHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PepNonce(ctx context.Context, in *QueryPepNonceRequest, opts ...grpc.CallOption) (*QueryPepNonceResponse, error) {
	out := new(QueryPepNonceResponse)
	err := c.cc.Invoke(ctx, Query_PepNonce_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PepNonceAll(ctx context.Context, in *QueryPepNonceAllRequest, opts ...grpc.CallOption) (*QueryPepNonceAllResponse, error) {
	out := new(QueryPepNonceAllResponse)
	err := c.cc.Invoke(ctx, Query_PepNonceAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PubKey(ctx context.Context, in *QueryPubKeyRequest, opts ...grpc.CallOption) (*QueryPubKeyResponse, error) {
	out := new(QueryPubKeyResponse)
	err := c.cc.Invoke(ctx, Query_PubKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeyshareReq(ctx context.Context, in *QueryKeyshareReqRequest, opts ...grpc.CallOption) (*QueryKeyshareReqResponse, error) {
	out := new(QueryKeyshareReqResponse)
	err := c.cc.Invoke(ctx, Query_KeyshareReq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeyshareReqAll(ctx context.Context, in *QueryKeyshareReqAllRequest, opts ...grpc.CallOption) (*QueryKeyshareReqAllResponse, error) {
	out := new(QueryKeyshareReqAllResponse)
	err := c.cc.Invoke(ctx, Query_KeyshareReqAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PrivateKeyshareReq(ctx context.Context, in *QueryPrivateKeyshareReqRequest, opts ...grpc.CallOption) (*QueryPrivateKeyshareReqResponse, error) {
	out := new(QueryPrivateKeyshareReqResponse)
	err := c.cc.Invoke(ctx, Query_PrivateKeyshareReq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DecryptData(ctx context.Context, in *QueryDecryptDataRequest, opts ...grpc.CallOption) (*QueryDecryptDataResponse, error) {
	out := new(QueryDecryptDataResponse)
	err := c.cc.Invoke(ctx, Query_DecryptData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a EncryptedTx by index.
	EncryptedTx(context.Context, *QueryEncryptedTxRequest) (*QueryEncryptedTxResponse, error)
	// Queries a list of EncryptedTx items.
	EncryptedTxAll(context.Context, *QueryEncryptedTxAllRequest) (*QueryEncryptedTxAllResponse, error)
	// Queries a list of EncryptedTx items.
	EncryptedTxAllFromHeight(context.Context, *QueryEncryptedTxAllFromHeightRequest) (*QueryEncryptedTxAllFromHeightResponse, error)
	// Queries a list of LatestHeight items.
	LatestHeight(context.Context, *QueryLatestHeightRequest) (*QueryLatestHeightResponse, error)
	// Queries a PepNonce by index.
	PepNonce(context.Context, *QueryPepNonceRequest) (*QueryPepNonceResponse, error)
	// Queries a list of PepNonce items.
	PepNonceAll(context.Context, *QueryPepNonceAllRequest) (*QueryPepNonceAllResponse, error)
	// Queries the public keys
	PubKey(context.Context, *QueryPubKeyRequest) (*QueryPubKeyResponse, error)
	// Queries a Keyshare request by identity
	KeyshareReq(context.Context, *QueryKeyshareReqRequest) (*QueryKeyshareReqResponse, error)
	// queries a list of keyshare requests
	KeyshareReqAll(context.Context, *QueryKeyshareReqAllRequest) (*QueryKeyshareReqAllResponse, error)
	// Queries a list of ShowPrivateKeyshareReq items.
	PrivateKeyshareReq(context.Context, *QueryPrivateKeyshareReqRequest) (*QueryPrivateKeyshareReqResponse, error)
	// Queries a list of DecryptData items.
	DecryptData(context.Context, *QueryDecryptDataRequest) (*QueryDecryptDataResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) EncryptedTx(context.Context, *QueryEncryptedTxRequest) (*QueryEncryptedTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptedTx not implemented")
}
func (UnimplementedQueryServer) EncryptedTxAll(context.Context, *QueryEncryptedTxAllRequest) (*QueryEncryptedTxAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptedTxAll not implemented")
}
func (UnimplementedQueryServer) EncryptedTxAllFromHeight(context.Context, *QueryEncryptedTxAllFromHeightRequest) (*QueryEncryptedTxAllFromHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EncryptedTxAllFromHeight not implemented")
}
func (UnimplementedQueryServer) LatestHeight(context.Context, *QueryLatestHeightRequest) (*QueryLatestHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestHeight not implemented")
}
func (UnimplementedQueryServer) PepNonce(context.Context, *QueryPepNonceRequest) (*QueryPepNonceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PepNonce not implemented")
}
func (UnimplementedQueryServer) PepNonceAll(context.Context, *QueryPepNonceAllRequest) (*QueryPepNonceAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PepNonceAll not implemented")
}
func (UnimplementedQueryServer) PubKey(context.Context, *QueryPubKeyRequest) (*QueryPubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubKey not implemented")
}
func (UnimplementedQueryServer) KeyshareReq(context.Context, *QueryKeyshareReqRequest) (*QueryKeyshareReqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyshareReq not implemented")
}
func (UnimplementedQueryServer) KeyshareReqAll(context.Context, *QueryKeyshareReqAllRequest) (*QueryKeyshareReqAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyshareReqAll not implemented")
}
func (UnimplementedQueryServer) PrivateKeyshareReq(context.Context, *QueryPrivateKeyshareReqRequest) (*QueryPrivateKeyshareReqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrivateKeyshareReq not implemented")
}
func (UnimplementedQueryServer) DecryptData(context.Context, *QueryDecryptDataRequest) (*QueryDecryptDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecryptData not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncryptedTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncryptedTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncryptedTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EncryptedTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncryptedTx(ctx, req.(*QueryEncryptedTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncryptedTxAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncryptedTxAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncryptedTxAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EncryptedTxAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncryptedTxAll(ctx, req.(*QueryEncryptedTxAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EncryptedTxAllFromHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEncryptedTxAllFromHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EncryptedTxAllFromHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_EncryptedTxAllFromHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EncryptedTxAllFromHeight(ctx, req.(*QueryEncryptedTxAllFromHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LatestHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLatestHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LatestHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_LatestHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LatestHeight(ctx, req.(*QueryLatestHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PepNonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPepNonceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PepNonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PepNonce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PepNonce(ctx, req.(*QueryPepNonceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PepNonceAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPepNonceAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PepNonceAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PepNonceAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PepNonceAll(ctx, req.(*QueryPepNonceAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PubKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PubKey(ctx, req.(*QueryPubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeyshareReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeyshareReqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeyshareReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_KeyshareReq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeyshareReq(ctx, req.(*QueryKeyshareReqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeyshareReqAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeyshareReqAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeyshareReqAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_KeyshareReqAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeyshareReqAll(ctx, req.(*QueryKeyshareReqAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PrivateKeyshareReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPrivateKeyshareReqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PrivateKeyshareReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PrivateKeyshareReq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PrivateKeyshareReq(ctx, req.(*QueryPrivateKeyshareReqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DecryptData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDecryptDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DecryptData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_DecryptData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DecryptData(ctx, req.(*QueryDecryptDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fairyring.pep.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EncryptedTx",
			Handler:    _Query_EncryptedTx_Handler,
		},
		{
			MethodName: "EncryptedTxAll",
			Handler:    _Query_EncryptedTxAll_Handler,
		},
		{
			MethodName: "EncryptedTxAllFromHeight",
			Handler:    _Query_EncryptedTxAllFromHeight_Handler,
		},
		{
			MethodName: "LatestHeight",
			Handler:    _Query_LatestHeight_Handler,
		},
		{
			MethodName: "PepNonce",
			Handler:    _Query_PepNonce_Handler,
		},
		{
			MethodName: "PepNonceAll",
			Handler:    _Query_PepNonceAll_Handler,
		},
		{
			MethodName: "PubKey",
			Handler:    _Query_PubKey_Handler,
		},
		{
			MethodName: "KeyshareReq",
			Handler:    _Query_KeyshareReq_Handler,
		},
		{
			MethodName: "KeyshareReqAll",
			Handler:    _Query_KeyshareReqAll_Handler,
		},
		{
			MethodName: "PrivateKeyshareReq",
			Handler:    _Query_PrivateKeyshareReq_Handler,
		},
		{
			MethodName: "DecryptData",
			Handler:    _Query_DecryptData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fairyring/pep/query.proto",
}
