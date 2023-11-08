// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0--rc1
// source: proto/library_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LibraryClient is the client API for Library service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryClient interface {
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error)
	GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error)
	GetAll(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*BookList, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	SearchBookByName(ctx context.Context, in *BookName, opts ...grpc.CallOption) (*BookList, error)
}

type libraryClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryClient(cc grpc.ClientConnInterface) LibraryClient {
	return &libraryClient{cc}
}

func (c *libraryClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookID, error) {
	out := new(BookID)
	err := c.cc.Invoke(ctx, "/Library/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/Library/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) GetAll(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/Library/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/Library/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/Library/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryClient) SearchBookByName(ctx context.Context, in *BookName, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/Library/SearchBookByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServer is the server API for Library service.
// All implementations must embed UnimplementedLibraryServer
// for forward compatibility
type LibraryServer interface {
	AddBook(context.Context, *Book) (*BookID, error)
	GetBook(context.Context, *BookID) (*Book, error)
	GetAll(context.Context, *wrapperspb.StringValue) (*BookList, error)
	UpdateBook(context.Context, *Book) (*wrapperspb.StringValue, error)
	DeleteBook(context.Context, *BookID) (*wrapperspb.StringValue, error)
	SearchBookByName(context.Context, *BookName) (*BookList, error)
	mustEmbedUnimplementedLibraryServer()
}

// UnimplementedLibraryServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServer struct {
}

func (UnimplementedLibraryServer) AddBook(context.Context, *Book) (*BookID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedLibraryServer) GetBook(context.Context, *BookID) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedLibraryServer) GetAll(context.Context, *wrapperspb.StringValue) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedLibraryServer) UpdateBook(context.Context, *Book) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedLibraryServer) DeleteBook(context.Context, *BookID) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedLibraryServer) SearchBookByName(context.Context, *BookName) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBookByName not implemented")
}
func (UnimplementedLibraryServer) mustEmbedUnimplementedLibraryServer() {}

// UnsafeLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServer will
// result in compilation errors.
type UnsafeLibraryServer interface {
	mustEmbedUnimplementedLibraryServer()
}

func RegisterLibraryServer(s grpc.ServiceRegistrar, srv LibraryServer) {
	s.RegisterService(&Library_ServiceDesc, srv)
}

func _Library_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).GetAll(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).DeleteBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Library_SearchBookByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServer).SearchBookByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Library/SearchBookByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServer).SearchBookByName(ctx, req.(*BookName))
	}
	return interceptor(ctx, in, info, handler)
}

// Library_ServiceDesc is the grpc.ServiceDesc for Library service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Library_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Library",
	HandlerType: (*LibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _Library_AddBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Library_GetBook_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Library_GetAll_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Library_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Library_DeleteBook_Handler,
		},
		{
			MethodName: "SearchBookByName",
			Handler:    _Library_SearchBookByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/library_service.proto",
}
