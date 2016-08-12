// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_service.proto
// DO NOT EDIT!

package google_bigtable_admin_table_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BigtableTableService service

type BigtableTableServiceClient interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(ctx context.Context, in *ColumnFamily, opts ...grpc.CallOption) (*ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type bigtableTableServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableTableServiceClient(cc *grpc.ClientConn) BigtableTableServiceClient {
	return &bigtableTableServiceClient{cc}
}

func (c *bigtableTableServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/ListTables", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) GetTable(ctx context.Context, in *GetTableRequest, opts ...grpc.CallOption) (*Table, error) {
	out := new(Table)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/GetTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteTable(ctx context.Context, in *DeleteTableRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) RenameTable(ctx context.Context, in *RenameTableRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) CreateColumnFamily(ctx context.Context, in *CreateColumnFamilyRequest, opts ...grpc.CallOption) (*ColumnFamily, error) {
	out := new(ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) UpdateColumnFamily(ctx context.Context, in *ColumnFamily, opts ...grpc.CallOption) (*ColumnFamily, error) {
	out := new(ColumnFamily)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) DeleteColumnFamily(ctx context.Context, in *DeleteColumnFamilyRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableTableServiceClient) BulkDeleteRows(ctx context.Context, in *BulkDeleteRowsRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableTableService service

type BigtableTableServiceServer interface {
	// Creates a new table, to be served from a specified cluster.
	// The table can be created with a full set of initial column families,
	// specified in the request.
	CreateTable(context.Context, *CreateTableRequest) (*Table, error)
	// Lists the names of all tables served from a specified cluster.
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
	// Gets the schema of the specified table, including its column families.
	GetTable(context.Context, *GetTableRequest) (*Table, error)
	// Permanently deletes a specified table and all of its data.
	DeleteTable(context.Context, *DeleteTableRequest) (*google_protobuf2.Empty, error)
	// Changes the name of a specified table.
	// Cannot be used to move tables between clusters, zones, or projects.
	RenameTable(context.Context, *RenameTableRequest) (*google_protobuf2.Empty, error)
	// Creates a new column family within a specified table.
	CreateColumnFamily(context.Context, *CreateColumnFamilyRequest) (*ColumnFamily, error)
	// Changes the configuration of a specified column family.
	UpdateColumnFamily(context.Context, *ColumnFamily) (*ColumnFamily, error)
	// Permanently deletes a specified column family and all of its data.
	DeleteColumnFamily(context.Context, *DeleteColumnFamilyRequest) (*google_protobuf2.Empty, error)
	// Delete all rows in a table corresponding to a particular prefix
	BulkDeleteRows(context.Context, *BulkDeleteRowsRequest) (*google_protobuf2.Empty, error)
}

func RegisterBigtableTableServiceServer(s *grpc.Server, srv BigtableTableServiceServer) {
	s.RegisterService(&_BigtableTableService_serviceDesc, srv)
}

func _BigtableTableService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_GetTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).GetTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/GetTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).GetTable(ctx, req.(*GetTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteTable(ctx, req.(*DeleteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_RenameTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/RenameTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).RenameTable(ctx, req.(*RenameTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_CreateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/CreateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).CreateColumnFamily(ctx, req.(*CreateColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_UpdateColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColumnFamily)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/UpdateColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).UpdateColumnFamily(ctx, req.(*ColumnFamily))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_DeleteColumnFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteColumnFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/DeleteColumnFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).DeleteColumnFamily(ctx, req.(*DeleteColumnFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableTableService_BulkDeleteRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.table.v1.BigtableTableService/BulkDeleteRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableTableServiceServer).BulkDeleteRows(ctx, req.(*BulkDeleteRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableTableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.table.v1.BigtableTableService",
	HandlerType: (*BigtableTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _BigtableTableService_CreateTable_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _BigtableTableService_ListTables_Handler,
		},
		{
			MethodName: "GetTable",
			Handler:    _BigtableTableService_GetTable_Handler,
		},
		{
			MethodName: "DeleteTable",
			Handler:    _BigtableTableService_DeleteTable_Handler,
		},
		{
			MethodName: "RenameTable",
			Handler:    _BigtableTableService_RenameTable_Handler,
		},
		{
			MethodName: "CreateColumnFamily",
			Handler:    _BigtableTableService_CreateColumnFamily_Handler,
		},
		{
			MethodName: "UpdateColumnFamily",
			Handler:    _BigtableTableService_UpdateColumnFamily_Handler,
		},
		{
			MethodName: "DeleteColumnFamily",
			Handler:    _BigtableTableService_DeleteColumnFamily_Handler,
		},
		{
			MethodName: "BulkDeleteRows",
			Handler:    _BigtableTableService_BulkDeleteRows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_service.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x95, 0x4f, 0x8b, 0xd4, 0x3e,
	0x18, 0xc7, 0xe9, 0xef, 0xf0, 0x43, 0xaa, 0x78, 0x08, 0xe2, 0x61, 0x0e, 0x1e, 0x0a, 0x5e, 0x06,
	0x69, 0xe8, 0x8c, 0xa2, 0x3b, 0xb0, 0x20, 0x9d, 0xfd, 0x83, 0xae, 0xc2, 0x3a, 0x2a, 0x82, 0x97,
	0x25, 0xed, 0x3c, 0x1b, 0xa3, 0x6d, 0x52, 0x9b, 0x74, 0x64, 0x15, 0x2f, 0xbe, 0x05, 0xbd, 0xea,
	0xc9, 0xa3, 0x27, 0xbd, 0xeb, 0xdd, 0xab, 0x6f, 0xc1, 0x17, 0x62, 0x9a, 0xb4, 0x6e, 0x77, 0x15,
	0xdb, 0xc8, 0x7a, 0x49, 0x43, 0xf3, 0x7c, 0x9f, 0xe7, 0xf3, 0xfc, 0x09, 0xf1, 0x1f, 0x52, 0x21,
	0x68, 0x06, 0x21, 0x15, 0x19, 0xe1, 0x34, 0x14, 0x25, 0xc5, 0x14, 0x78, 0x51, 0x0a, 0x25, 0xb0,
	0x3d, 0x22, 0x05, 0x93, 0x38, 0x61, 0x54, 0x91, 0x24, 0x03, 0x4c, 0x96, 0x39, 0xe3, 0xd8, 0xee,
	0x57, 0xd1, 0xcf, 0xff, 0x7b, 0x76, 0x95, 0x50, 0xae, 0x58, 0x0a, 0xa1, 0xd1, 0xa3, 0x0b, 0x8d,
	0xef, 0xd6, 0x28, 0x34, 0xe2, 0xd0, 0xee, 0x57, 0xd1, 0xe8, 0xc6, 0xb0, 0xd8, 0x7a, 0xc1, 0x8d,
	0xe3, 0x54, 0xf0, 0x7d, 0x46, 0x31, 0xe1, 0x5c, 0x28, 0xa2, 0x98, 0xe0, 0xd2, 0x86, 0x1a, 0x3d,
	0x38, 0xd1, 0x34, 0x96, 0x44, 0x91, 0xc6, 0x71, 0xfa, 0x2f, 0xea, 0xb3, 0x97, 0x83, 0x94, 0x84,
	0x42, 0x4b, 0x3f, 0xa5, 0x4c, 0x3d, 0xaa, 0x92, 0x30, 0x15, 0x39, 0xb6, 0x81, 0xb0, 0x39, 0x48,
	0xaa, 0x7d, 0x5c, 0xa8, 0x83, 0x02, 0x24, 0x86, 0x5c, 0x6f, 0xec, 0x6a, 0x45, 0x93, 0xcf, 0x67,
	0xfc, 0x73, 0x71, 0xe3, 0xfe, 0x5e, 0xbd, 0xdc, 0xb5, 0xce, 0xd1, 0x7b, 0xcf, 0x3f, 0x3d, 0x2f,
	0x81, 0x28, 0xfb, 0x1b, 0x4d, 0xc2, 0x3f, 0xf7, 0x21, 0xec, 0x18, 0x2f, 0xe0, 0x69, 0x05, 0x52,
	0x8d, 0x2e, 0xf6, 0x69, 0x8c, 0x75, 0x30, 0x7b, 0xf5, 0xed, 0xfb, 0xeb, 0xff, 0x2e, 0x07, 0xb8,
	0xce, 0xf5, 0x05, 0x27, 0x39, 0xac, 0x6b, 0xb8, 0xc7, 0x90, 0x2a, 0x89, 0xc7, 0xf8, 0xb9, 0xe0,
	0x50, 0x7f, 0xd3, 0xac, 0x92, 0x0a, 0x4a, 0xbd, 0x7d, 0x69, 0x2b, 0x23, 0x67, 0xde, 0x18, 0x7d,
	0xf0, 0x7c, 0xff, 0x16, 0x93, 0xca, 0x78, 0x92, 0x28, 0xea, 0x8b, 0x78, 0x68, 0xdb, 0x42, 0x4e,
	0x5c, 0x24, 0xb2, 0xd0, 0xd3, 0x02, 0xc1, 0x55, 0x43, 0x1c, 0x21, 0x57, 0x62, 0xf4, 0xd6, 0xf3,
	0x4f, 0x6d, 0x83, 0x75, 0x87, 0x70, 0x5f, 0xe4, 0xd6, 0xd2, 0xb1, 0x9e, 0x6b, 0x86, 0x6e, 0x8a,
	0xa2, 0x81, 0x74, 0x0d, 0x9c, 0xc6, 0x44, 0x6f, 0x74, 0xd7, 0x37, 0x20, 0x83, 0xc1, 0x5d, 0xef,
	0x18, 0xb7, 0x94, 0xe7, 0x5b, 0x4d, 0x3b, 0x7d, 0xe1, 0x66, 0x3d, 0x70, 0x2d, 0xd6, 0xf8, 0x2f,
	0xb0, 0xde, 0x69, 0xac, 0x05, 0xd4, 0x92, 0x81, 0x58, 0x1d, 0xe3, 0x3e, 0xac, 0xb9, 0xc1, 0x5a,
	0x0f, 0xae, 0x39, 0x63, 0xcd, 0x4a, 0x13, 0xa5, 0x1e, 0xc3, 0xaf, 0x9e, 0x8f, 0xec, 0x05, 0x98,
	0x8b, 0xac, 0xca, 0xf9, 0x16, 0xc9, 0x59, 0x76, 0x80, 0xd6, 0x86, 0x5d, 0x9a, 0xae, 0xa6, 0xc5,
	0xbd, 0xd4, 0x2b, 0xed, 0x88, 0x82, 0x1d, 0x93, 0xc4, 0x66, 0x70, 0xdd, 0x39, 0x09, 0x9c, 0x1e,
	0xfa, 0x61, 0xf6, 0x4e, 0x7d, 0xd1, 0xc9, 0xdc, 0x2f, 0x96, 0xc7, 0x93, 0x71, 0x22, 0x72, 0xe4,
	0xbf, 0x6d, 0xf8, 0xb7, 0x47, 0xb1, 0x2b, 0xff, 0x31, 0xfc, 0xba, 0x2b, 0x3a, 0x83, 0x4f, 0x3a,
	0x03, 0x3b, 0x99, 0x6e, 0xed, 0xf8, 0x55, 0xd3, 0x37, 0x3d, 0x37, 0x0d, 0xf8, 0xc6, 0xf8, 0x04,
	0xc0, 0xd1, 0x47, 0xcf, 0x3f, 0x1b, 0x57, 0xd9, 0x13, 0x4b, 0xb1, 0x10, 0xcf, 0x24, 0xba, 0xd2,
	0x47, 0x7c, 0xd4, 0xbe, 0x8f, 0xf6, 0x8e, 0xa1, 0xdd, 0x09, 0xb6, 0x0c, 0xad, 0x7d, 0x4d, 0x9c,
	0x26, 0x3e, 0x39, 0x12, 0x4e, 0x97, 0x3a, 0x8e, 0xfd, 0x40, 0x3f, 0x38, 0x3d, 0x98, 0xf1, 0xe8,
	0x77, 0x6f, 0x8c, 0xdc, 0xad, 0xe9, 0x76, 0xbd, 0xe4, 0x7f, 0x83, 0x39, 0xfd, 0x11, 0x00, 0x00,
	0xff, 0xff, 0x2e, 0xa4, 0x4d, 0x5a, 0x47, 0x08, 0x00, 0x00,
}