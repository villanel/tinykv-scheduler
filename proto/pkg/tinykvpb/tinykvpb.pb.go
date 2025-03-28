// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tinykvpb.proto

package tinykvpb

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	coprocessor "github.com/villanel/tinykv-scheduler/proto/pkg/coprocessor"

	kvrpcpb "github.com/villanel/tinykv-scheduler/proto/pkg/kvrpcpb"

	raft_serverpb "github.com/villanel/tinykv-scheduler/proto/pkg/raft_serverpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TinyKv service

type TinyKvClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(ctx context.Context, opts ...grpc.CallOption) (TinyKv_RaftClient, error)
	Snapshot(ctx context.Context, opts ...grpc.CallOption) (TinyKv_SnapshotClient, error)
	// Coprocessor
	Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error)
	SAdd(ctx context.Context, in *kvrpcpb.SetAddRequest, opts ...grpc.CallOption) (*kvrpcpb.SetAddResponse, error)
	SMembers(ctx context.Context, in *kvrpcpb.SetMembersRequest, opts ...grpc.CallOption) (*kvrpcpb.SetMembersResponse, error)
	SRem(ctx context.Context, in *kvrpcpb.SetRemoveRequest, opts ...grpc.CallOption) (*kvrpcpb.SetRemoveResponse, error)
	SIsMember(ctx context.Context, in *kvrpcpb.SetIsMemberRequest, opts ...grpc.CallOption) (*kvrpcpb.SetIsMemberResponse, error)
	SCard(ctx context.Context, in *kvrpcpb.SetCardRequest, opts ...grpc.CallOption) (*kvrpcpb.SetCardResponse, error)
	SUnion(ctx context.Context, in *kvrpcpb.SetUnionRequest, opts ...grpc.CallOption) (*kvrpcpb.SetUnionResponse, error)
}

type tinyKvClient struct {
	cc *grpc.ClientConn
}

func NewTinyKvClient(cc *grpc.ClientConn) TinyKvClient {
	return &tinyKvClient{cc}
}

func (c *tinyKvClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvPrewrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvCheckTxnStatus(ctx context.Context, in *kvrpcpb.CheckTxnStatusRequest, opts ...grpc.CallOption) (*kvrpcpb.CheckTxnStatusResponse, error) {
	out := new(kvrpcpb.CheckTxnStatusResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvCheckTxnStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvBatchRollback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/KvResolveLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) RawScan(ctx context.Context, in *kvrpcpb.RawScanRequest, opts ...grpc.CallOption) (*kvrpcpb.RawScanResponse, error) {
	out := new(kvrpcpb.RawScanResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/RawScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) Raft(ctx context.Context, opts ...grpc.CallOption) (TinyKv_RaftClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TinyKv_serviceDesc.Streams[0], "/tinykvpb.TinyKv/Raft", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyKvRaftClient{stream}
	return x, nil
}

type TinyKv_RaftClient interface {
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tinyKvRaftClient struct {
	grpc.ClientStream
}

func (x *tinyKvRaftClient) Send(m *raft_serverpb.RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tinyKvRaftClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tinyKvClient) Snapshot(ctx context.Context, opts ...grpc.CallOption) (TinyKv_SnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TinyKv_serviceDesc.Streams[1], "/tinykvpb.TinyKv/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyKvSnapshotClient{stream}
	return x, nil
}

type TinyKv_SnapshotClient interface {
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
	grpc.ClientStream
}

type tinyKvSnapshotClient struct {
	grpc.ClientStream
}

func (x *tinyKvSnapshotClient) Send(m *raft_serverpb.SnapshotChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tinyKvSnapshotClient) CloseAndRecv() (*raft_serverpb.Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(raft_serverpb.Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tinyKvClient) Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error) {
	out := new(coprocessor.Response)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/Coprocessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SAdd(ctx context.Context, in *kvrpcpb.SetAddRequest, opts ...grpc.CallOption) (*kvrpcpb.SetAddResponse, error) {
	out := new(kvrpcpb.SetAddResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SMembers(ctx context.Context, in *kvrpcpb.SetMembersRequest, opts ...grpc.CallOption) (*kvrpcpb.SetMembersResponse, error) {
	out := new(kvrpcpb.SetMembersResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SRem(ctx context.Context, in *kvrpcpb.SetRemoveRequest, opts ...grpc.CallOption) (*kvrpcpb.SetRemoveResponse, error) {
	out := new(kvrpcpb.SetRemoveResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SRem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SIsMember(ctx context.Context, in *kvrpcpb.SetIsMemberRequest, opts ...grpc.CallOption) (*kvrpcpb.SetIsMemberResponse, error) {
	out := new(kvrpcpb.SetIsMemberResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SIsMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SCard(ctx context.Context, in *kvrpcpb.SetCardRequest, opts ...grpc.CallOption) (*kvrpcpb.SetCardResponse, error) {
	out := new(kvrpcpb.SetCardResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyKvClient) SUnion(ctx context.Context, in *kvrpcpb.SetUnionRequest, opts ...grpc.CallOption) (*kvrpcpb.SetUnionResponse, error) {
	out := new(kvrpcpb.SetUnionResponse)
	err := c.cc.Invoke(ctx, "/tinykvpb.TinyKv/SUnion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TinyKv service

type TinyKvServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvCheckTxnStatus(context.Context, *kvrpcpb.CheckTxnStatusRequest) (*kvrpcpb.CheckTxnStatusResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	RawScan(context.Context, *kvrpcpb.RawScanRequest) (*kvrpcpb.RawScanResponse, error)
	// Raft commands (tinykv <-> tinykv).
	Raft(TinyKv_RaftServer) error
	Snapshot(TinyKv_SnapshotServer) error
	// Coprocessor
	Coprocessor(context.Context, *coprocessor.Request) (*coprocessor.Response, error)
	SAdd(context.Context, *kvrpcpb.SetAddRequest) (*kvrpcpb.SetAddResponse, error)
	SMembers(context.Context, *kvrpcpb.SetMembersRequest) (*kvrpcpb.SetMembersResponse, error)
	SRem(context.Context, *kvrpcpb.SetRemoveRequest) (*kvrpcpb.SetRemoveResponse, error)
	SIsMember(context.Context, *kvrpcpb.SetIsMemberRequest) (*kvrpcpb.SetIsMemberResponse, error)
	SCard(context.Context, *kvrpcpb.SetCardRequest) (*kvrpcpb.SetCardResponse, error)
	SUnion(context.Context, *kvrpcpb.SetUnionRequest) (*kvrpcpb.SetUnionResponse, error)
}

func RegisterTinyKvServer(s *grpc.Server, srv TinyKvServer) {
	s.RegisterService(&_TinyKv_serviceDesc, srv)
}

func _TinyKv_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvCheckTxnStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CheckTxnStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvCheckTxnStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvCheckTxnStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvCheckTxnStatus(ctx, req.(*kvrpcpb.CheckTxnStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_RawScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).RawScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/RawScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).RawScan(ctx, req.(*kvrpcpb.RawScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_Raft_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TinyKvServer).Raft(&tinyKvRaftServer{stream})
}

type TinyKv_RaftServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.RaftMessage, error)
	grpc.ServerStream
}

type tinyKvRaftServer struct {
	grpc.ServerStream
}

func (x *tinyKvRaftServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tinyKvRaftServer) Recv() (*raft_serverpb.RaftMessage, error) {
	m := new(raft_serverpb.RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TinyKv_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TinyKvServer).Snapshot(&tinyKvSnapshotServer{stream})
}

type TinyKv_SnapshotServer interface {
	SendAndClose(*raft_serverpb.Done) error
	Recv() (*raft_serverpb.SnapshotChunk, error)
	grpc.ServerStream
}

type tinyKvSnapshotServer struct {
	grpc.ServerStream
}

func (x *tinyKvSnapshotServer) SendAndClose(m *raft_serverpb.Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tinyKvSnapshotServer) Recv() (*raft_serverpb.SnapshotChunk, error) {
	m := new(raft_serverpb.SnapshotChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TinyKv_Coprocessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(coprocessor.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).Coprocessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/Coprocessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).Coprocessor(ctx, req.(*coprocessor.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SAdd(ctx, req.(*kvrpcpb.SetAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SMembers(ctx, req.(*kvrpcpb.SetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SRem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SRem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SRem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SRem(ctx, req.(*kvrpcpb.SetRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SIsMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetIsMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SIsMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SIsMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SIsMember(ctx, req.(*kvrpcpb.SetIsMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SCard(ctx, req.(*kvrpcpb.SetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyKv_SUnion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.SetUnionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyKvServer).SUnion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinykvpb.TinyKv/SUnion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyKvServer).SUnion(ctx, req.(*kvrpcpb.SetUnionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TinyKv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tinykvpb.TinyKv",
	HandlerType: (*TinyKvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _TinyKv_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _TinyKv_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _TinyKv_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _TinyKv_KvCommit_Handler,
		},
		{
			MethodName: "KvCheckTxnStatus",
			Handler:    _TinyKv_KvCheckTxnStatus_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _TinyKv_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _TinyKv_KvResolveLock_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _TinyKv_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _TinyKv_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _TinyKv_RawDelete_Handler,
		},
		{
			MethodName: "RawScan",
			Handler:    _TinyKv_RawScan_Handler,
		},
		{
			MethodName: "Coprocessor",
			Handler:    _TinyKv_Coprocessor_Handler,
		},
		{
			MethodName: "SAdd",
			Handler:    _TinyKv_SAdd_Handler,
		},
		{
			MethodName: "SMembers",
			Handler:    _TinyKv_SMembers_Handler,
		},
		{
			MethodName: "SRem",
			Handler:    _TinyKv_SRem_Handler,
		},
		{
			MethodName: "SIsMember",
			Handler:    _TinyKv_SIsMember_Handler,
		},
		{
			MethodName: "SCard",
			Handler:    _TinyKv_SCard_Handler,
		},
		{
			MethodName: "SUnion",
			Handler:    _TinyKv_SUnion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Raft",
			Handler:       _TinyKv_Raft_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Snapshot",
			Handler:       _TinyKv_Snapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tinykvpb.proto",
}

func init() { proto.RegisterFile("tinykvpb.proto", fileDescriptor_tinykvpb_e8b743394c9c28e7) }

var fileDescriptor_tinykvpb_e8b743394c9c28e7 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x5b, 0xa9, 0x2b, 0x9d, 0xd1, 0x60, 0xb8, 0x05, 0xb6, 0x6c, 0x04, 0x69, 0x57, 0x5c,
	0x15, 0x09, 0x90, 0x90, 0x80, 0x21, 0x75, 0xa9, 0x98, 0x50, 0x36, 0xa9, 0x8a, 0xb7, 0x6b, 0xe4,
	0xa6, 0x67, 0x6d, 0x95, 0x36, 0x0e, 0xb6, 0xeb, 0xb2, 0x37, 0xe1, 0x39, 0x78, 0x0a, 0x2e, 0x79,
	0x04, 0x54, 0x5e, 0x04, 0xa5, 0xa9, 0x9d, 0xb8, 0x49, 0xef, 0x92, 0xef, 0x3f, 0xff, 0x6f, 0xd9,
	0x3e, 0xc7, 0xe8, 0x91, 0x9c, 0xc6, 0xf7, 0x91, 0x4a, 0x86, 0xdd, 0x84, 0x33, 0xc9, 0x70, 0x4b,
	0xff, 0x3b, 0x07, 0x91, 0xe2, 0x49, 0xa8, 0x05, 0xa7, 0xcd, 0xe9, 0x9d, 0xfc, 0x26, 0x80, 0x2b,
	0xe0, 0x06, 0x3e, 0x09, 0x59, 0xc2, 0x59, 0x08, 0x42, 0x30, 0xbe, 0x41, 0x9d, 0x31, 0x1b, 0xb3,
	0xf5, 0xe7, 0xeb, 0xf4, 0x2b, 0xa3, 0x6f, 0x7e, 0x21, 0xd4, 0xbc, 0x99, 0xc6, 0xf7, 0xbe, 0xc2,
	0xef, 0xd0, 0x9e, 0xaf, 0x2e, 0x41, 0xe2, 0x76, 0x57, 0xaf, 0x70, 0x09, 0x32, 0x80, 0xef, 0x0b,
	0x10, 0xd2, 0xe9, 0xd8, 0x50, 0x24, 0x2c, 0x16, 0x70, 0x56, 0xc3, 0xef, 0x51, 0xd3, 0x57, 0x24,
	0xa4, 0x31, 0xce, 0x2b, 0xd2, 0x5f, 0xed, 0x7b, 0xba, 0x45, 0x8d, 0xd1, 0x43, 0xc8, 0x57, 0x03,
	0x0e, 0x4b, 0x3e, 0x95, 0x80, 0x8f, 0x4c, 0x99, 0x46, 0x3a, 0xe0, 0xb8, 0x42, 0x31, 0x21, 0xe7,
	0xa8, 0xe5, 0x2b, 0x8f, 0xcd, 0xe7, 0x53, 0x89, 0x9f, 0x99, 0xc2, 0x0c, 0xe8, 0x80, 0xe7, 0x25,
	0x6e, 0xec, 0xb7, 0xe8, 0xd0, 0x57, 0xde, 0x04, 0xc2, 0xe8, 0xe6, 0x47, 0x4c, 0x24, 0x95, 0x0b,
	0x81, 0xdd, 0xbc, 0xdc, 0x12, 0x74, 0xdc, 0xcb, 0x9d, 0xba, 0x89, 0x0d, 0xd0, 0x63, 0x5f, 0x5d,
	0x50, 0x19, 0x4e, 0x02, 0x36, 0x9b, 0x0d, 0x69, 0x18, 0xe1, 0x17, 0xc6, 0x65, 0x71, 0x1d, 0xea,
	0xee, 0x92, 0x4d, 0xe6, 0x15, 0x3a, 0xf0, 0x55, 0x00, 0x82, 0xcd, 0x14, 0x5c, 0xb1, 0x30, 0xc2,
	0x27, 0xc6, 0x52, 0xa0, 0x3a, 0xef, 0xb4, 0x5a, 0x34, 0x69, 0x1f, 0x51, 0x33, 0xa0, 0xcb, 0xf4,
	0xb2, 0xf3, 0x53, 0xcb, 0x40, 0xf9, 0xd4, 0x34, 0xdf, 0x32, 0x0f, 0x16, 0x5b, 0xe6, 0xc1, 0xa2,
	0xda, 0xbc, 0xe6, 0xc6, 0xdc, 0x47, 0xfb, 0x01, 0x5d, 0xf6, 0x61, 0x06, 0x12, 0xf0, 0x71, 0xb1,
	0x2e, 0x63, 0x3a, 0xc2, 0xa9, 0x92, 0x4c, 0xca, 0x67, 0xf4, 0x20, 0xa0, 0xcb, 0x75, 0xdb, 0x59,
	0x6b, 0x15, 0x3b, 0xef, 0xa8, 0x2c, 0x14, 0xb6, 0xd0, 0x08, 0xe8, 0x9d, 0xc4, 0x4e, 0xd7, 0x9e,
	0x9e, 0x14, 0x5e, 0x83, 0x10, 0x74, 0x0c, 0x4e, 0x7b, 0x4b, 0xeb, 0xb3, 0x18, 0xce, 0x6a, 0xaf,
	0xea, 0xb8, 0x87, 0x5a, 0x24, 0xa6, 0x89, 0x98, 0x30, 0x89, 0x4f, 0xb7, 0x8a, 0xb4, 0xe0, 0x4d,
	0x16, 0x71, 0xb4, 0x3b, 0xe2, 0x13, 0x7a, 0xe8, 0xe5, 0x13, 0x8a, 0x3b, 0xdd, 0xe2, 0xbc, 0xe6,
	0xa3, 0x63, 0xd3, 0xc2, 0xcc, 0x35, 0x48, 0x6f, 0x34, 0x2a, 0x1c, 0x3f, 0x01, 0xd9, 0x1b, 0x8d,
	0xca, 0xc7, 0xaf, 0x79, 0x66, 0xc5, 0x1e, 0x6a, 0x91, 0x6b, 0x98, 0x0f, 0x81, 0x0b, 0xec, 0x14,
	0x8b, 0x36, 0x50, 0x07, 0x9c, 0x54, 0x6a, 0x9b, 0x90, 0x73, 0xd4, 0x20, 0x01, 0xcc, 0x0b, 0x97,
	0x47, 0xd2, 0xf6, 0x98, 0x33, 0x55, 0x71, 0x79, 0x05, 0x69, 0x63, 0xff, 0x82, 0xf6, 0xc9, 0x57,
	0x91, 0x85, 0x62, 0x6b, 0x21, 0x4d, 0xcb, 0x4d, 0x6c, 0x89, 0x9b, 0x9c, 0x0f, 0x68, 0x8f, 0x78,
	0x94, 0x8f, 0xb0, 0xb5, 0xdb, 0x94, 0x94, 0x1b, 0xc0, 0x08, 0x66, 0x0b, 0x4d, 0x72, 0x1b, 0x4f,
	0x59, 0x8c, 0xad, 0x9a, 0x35, 0x2a, 0xbf, 0x3b, 0xb9, 0x92, 0xd9, 0x2f, 0x0e, 0x7f, 0xaf, 0xdc,
	0xfa, 0x9f, 0x95, 0x5b, 0xff, 0xbb, 0x72, 0xeb, 0x3f, 0xff, 0xb9, 0xb5, 0x61, 0x73, 0xfd, 0x9a,
	0xbe, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x58, 0x9f, 0xac, 0x17, 0xb6, 0x05, 0x00, 0x00,
}
