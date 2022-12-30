// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/PdfComposeService.proto

package pdfcompose

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

// PdfComposeServiceClient is the client API for PdfComposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PdfComposeServiceClient interface {
	MergePdf(ctx context.Context, opts ...grpc.CallOption) (PdfComposeService_MergePdfClient, error)
}

type pdfComposeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPdfComposeServiceClient(cc grpc.ClientConnInterface) PdfComposeServiceClient {
	return &pdfComposeServiceClient{cc}
}

func (c *pdfComposeServiceClient) MergePdf(ctx context.Context, opts ...grpc.CallOption) (PdfComposeService_MergePdfClient, error) {
	stream, err := c.cc.NewStream(ctx, &PdfComposeService_ServiceDesc.Streams[0], "/pdfcompose.PdfComposeService/MergePdf", opts...)
	if err != nil {
		return nil, err
	}
	x := &pdfComposeServiceMergePdfClient{stream}
	return x, nil
}

type PdfComposeService_MergePdfClient interface {
	Send(*File) error
	Recv() (*MergedPdf, error)
	grpc.ClientStream
}

type pdfComposeServiceMergePdfClient struct {
	grpc.ClientStream
}

func (x *pdfComposeServiceMergePdfClient) Send(m *File) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pdfComposeServiceMergePdfClient) Recv() (*MergedPdf, error) {
	m := new(MergedPdf)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfComposeServiceServer is the server API for PdfComposeService service.
// All implementations must embed UnimplementedPdfComposeServiceServer
// for forward compatibility
type PdfComposeServiceServer interface {
	MergePdf(PdfComposeService_MergePdfServer) error
	mustEmbedUnimplementedPdfComposeServiceServer()
}

// UnimplementedPdfComposeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPdfComposeServiceServer struct {
}

func (UnimplementedPdfComposeServiceServer) MergePdf(PdfComposeService_MergePdfServer) error {
	return status.Errorf(codes.Unimplemented, "method MergePdf not implemented")
}
func (UnimplementedPdfComposeServiceServer) mustEmbedUnimplementedPdfComposeServiceServer() {}

// UnsafePdfComposeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PdfComposeServiceServer will
// result in compilation errors.
type UnsafePdfComposeServiceServer interface {
	mustEmbedUnimplementedPdfComposeServiceServer()
}

func RegisterPdfComposeServiceServer(s grpc.ServiceRegistrar, srv PdfComposeServiceServer) {
	s.RegisterService(&PdfComposeService_ServiceDesc, srv)
}

func _PdfComposeService_MergePdf_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PdfComposeServiceServer).MergePdf(&pdfComposeServiceMergePdfServer{stream})
}

type PdfComposeService_MergePdfServer interface {
	Send(*MergedPdf) error
	Recv() (*File, error)
	grpc.ServerStream
}

type pdfComposeServiceMergePdfServer struct {
	grpc.ServerStream
}

func (x *pdfComposeServiceMergePdfServer) Send(m *MergedPdf) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pdfComposeServiceMergePdfServer) Recv() (*File, error) {
	m := new(File)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfComposeService_ServiceDesc is the grpc.ServiceDesc for PdfComposeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PdfComposeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfcompose.PdfComposeService",
	HandlerType: (*PdfComposeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MergePdf",
			Handler:       _PdfComposeService_MergePdf_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/PdfComposeService.proto",
}
