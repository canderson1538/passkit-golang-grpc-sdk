// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: io/core/a_rpc_images.proto

package io

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImagesClient is the client API for Images service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImagesClient interface {
	SetProfileImage(ctx context.Context, in *ProfileImageInput, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetProfileImage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Url, error)
	CreateImages(ctx context.Context, in *CreateImageInput, opts ...grpc.CallOption) (*ImageIds, error)
	UpdateImage(ctx context.Context, in *UpdateImageInput, opts ...grpc.CallOption) (*ImageRecord, error)
	GetImageURL(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Url, error)
	GetStampImageURL(ctx context.Context, in *StampImageRequest, opts ...grpc.CallOption) (*Url, error)
	GetLocalizedImageURL(ctx context.Context, in *LocalizedImageInput, opts ...grpc.CallOption) (*Url, error)
	GetProfileImageById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Url, error)
	// returns a zip bundle containing all images for that ID
	GetImageBundle(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ImageBundle, error)
	GetImageData(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ImageRecord, error)
	DeleteImage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteLocalizedImage(ctx context.Context, in *LocalizedImageInput, opts ...grpc.CallOption) (*ImageRecord, error)
	ListImagesForUserDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (Images_ListImagesForUserDeprecatedClient, error)
	ListImagesForUser(ctx context.Context, in *Filters, opts ...grpc.CallOption) (Images_ListImagesForUserClient, error)
	ListImagesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (Images_ListImagesDeprecatedClient, error)
	ListImages(ctx context.Context, in *Filters, opts ...grpc.CallOption) (Images_ListImagesClient, error)
	CountImagesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error)
	CountImages(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error)
	CountImagesForUserDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error)
	CountImagesForUser(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error)
}

type imagesClient struct {
	cc grpc.ClientConnInterface
}

func NewImagesClient(cc grpc.ClientConnInterface) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) SetProfileImage(ctx context.Context, in *ProfileImageInput, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.Images/setProfileImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetProfileImage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/io.Images/getProfileImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) CreateImages(ctx context.Context, in *CreateImageInput, opts ...grpc.CallOption) (*ImageIds, error) {
	out := new(ImageIds)
	err := c.cc.Invoke(ctx, "/io.Images/createImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) UpdateImage(ctx context.Context, in *UpdateImageInput, opts ...grpc.CallOption) (*ImageRecord, error) {
	out := new(ImageRecord)
	err := c.cc.Invoke(ctx, "/io.Images/updateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetImageURL(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/io.Images/getImageURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetStampImageURL(ctx context.Context, in *StampImageRequest, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/io.Images/getStampImageURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetLocalizedImageURL(ctx context.Context, in *LocalizedImageInput, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/io.Images/getLocalizedImageURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetProfileImageById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/io.Images/getProfileImageById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetImageBundle(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ImageBundle, error) {
	out := new(ImageBundle)
	err := c.cc.Invoke(ctx, "/io.Images/getImageBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) GetImageData(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ImageRecord, error) {
	out := new(ImageRecord)
	err := c.cc.Invoke(ctx, "/io.Images/getImageData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) DeleteImage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.Images/deleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) DeleteLocalizedImage(ctx context.Context, in *LocalizedImageInput, opts ...grpc.CallOption) (*ImageRecord, error) {
	out := new(ImageRecord)
	err := c.cc.Invoke(ctx, "/io.Images/deleteLocalizedImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) ListImagesForUserDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (Images_ListImagesForUserDeprecatedClient, error) {
	stream, err := c.cc.NewStream(ctx, &Images_ServiceDesc.Streams[0], "/io.Images/listImagesForUserDeprecated", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagesListImagesForUserDeprecatedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Images_ListImagesForUserDeprecatedClient interface {
	Recv() (*ImageRecord, error)
	grpc.ClientStream
}

type imagesListImagesForUserDeprecatedClient struct {
	grpc.ClientStream
}

func (x *imagesListImagesForUserDeprecatedClient) Recv() (*ImageRecord, error) {
	m := new(ImageRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imagesClient) ListImagesForUser(ctx context.Context, in *Filters, opts ...grpc.CallOption) (Images_ListImagesForUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Images_ServiceDesc.Streams[1], "/io.Images/listImagesForUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagesListImagesForUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Images_ListImagesForUserClient interface {
	Recv() (*ImageRecord, error)
	grpc.ClientStream
}

type imagesListImagesForUserClient struct {
	grpc.ClientStream
}

func (x *imagesListImagesForUserClient) Recv() (*ImageRecord, error) {
	m := new(ImageRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imagesClient) ListImagesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (Images_ListImagesDeprecatedClient, error) {
	stream, err := c.cc.NewStream(ctx, &Images_ServiceDesc.Streams[2], "/io.Images/listImagesDeprecated", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagesListImagesDeprecatedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Images_ListImagesDeprecatedClient interface {
	Recv() (*ImageRecord, error)
	grpc.ClientStream
}

type imagesListImagesDeprecatedClient struct {
	grpc.ClientStream
}

func (x *imagesListImagesDeprecatedClient) Recv() (*ImageRecord, error) {
	m := new(ImageRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imagesClient) ListImages(ctx context.Context, in *Filters, opts ...grpc.CallOption) (Images_ListImagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Images_ServiceDesc.Streams[3], "/io.Images/listImages", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagesListImagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Images_ListImagesClient interface {
	Recv() (*ImageRecord, error)
	grpc.ClientStream
}

type imagesListImagesClient struct {
	grpc.ClientStream
}

func (x *imagesListImagesClient) Recv() (*ImageRecord, error) {
	m := new(ImageRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imagesClient) CountImagesDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/io.Images/countImagesDeprecated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) CountImages(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/io.Images/countImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) CountImagesForUserDeprecated(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/io.Images/countImagesForUserDeprecated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) CountImagesForUser(ctx context.Context, in *Filters, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/io.Images/countImagesForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImagesServer is the server API for Images service.
// All implementations should embed UnimplementedImagesServer
// for forward compatibility
type ImagesServer interface {
	SetProfileImage(context.Context, *ProfileImageInput) (*emptypb.Empty, error)
	GetProfileImage(context.Context, *emptypb.Empty) (*Url, error)
	CreateImages(context.Context, *CreateImageInput) (*ImageIds, error)
	UpdateImage(context.Context, *UpdateImageInput) (*ImageRecord, error)
	GetImageURL(context.Context, *Id) (*Url, error)
	GetStampImageURL(context.Context, *StampImageRequest) (*Url, error)
	GetLocalizedImageURL(context.Context, *LocalizedImageInput) (*Url, error)
	GetProfileImageById(context.Context, *Id) (*Url, error)
	// returns a zip bundle containing all images for that ID
	GetImageBundle(context.Context, *Id) (*ImageBundle, error)
	GetImageData(context.Context, *Id) (*ImageRecord, error)
	DeleteImage(context.Context, *Id) (*emptypb.Empty, error)
	DeleteLocalizedImage(context.Context, *LocalizedImageInput) (*ImageRecord, error)
	ListImagesForUserDeprecated(*Pagination, Images_ListImagesForUserDeprecatedServer) error
	ListImagesForUser(*Filters, Images_ListImagesForUserServer) error
	ListImagesDeprecated(*Pagination, Images_ListImagesDeprecatedServer) error
	ListImages(*Filters, Images_ListImagesServer) error
	CountImagesDeprecated(context.Context, *Pagination) (*Count, error)
	CountImages(context.Context, *Filters) (*Count, error)
	CountImagesForUserDeprecated(context.Context, *Pagination) (*Count, error)
	CountImagesForUser(context.Context, *Filters) (*Count, error)
}

// UnimplementedImagesServer should be embedded to have forward compatible implementations.
type UnimplementedImagesServer struct {
}

func (UnimplementedImagesServer) SetProfileImage(context.Context, *ProfileImageInput) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfileImage not implemented")
}
func (UnimplementedImagesServer) GetProfileImage(context.Context, *emptypb.Empty) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileImage not implemented")
}
func (UnimplementedImagesServer) CreateImages(context.Context, *CreateImageInput) (*ImageIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImages not implemented")
}
func (UnimplementedImagesServer) UpdateImage(context.Context, *UpdateImageInput) (*ImageRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImage not implemented")
}
func (UnimplementedImagesServer) GetImageURL(context.Context, *Id) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageURL not implemented")
}
func (UnimplementedImagesServer) GetStampImageURL(context.Context, *StampImageRequest) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStampImageURL not implemented")
}
func (UnimplementedImagesServer) GetLocalizedImageURL(context.Context, *LocalizedImageInput) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalizedImageURL not implemented")
}
func (UnimplementedImagesServer) GetProfileImageById(context.Context, *Id) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileImageById not implemented")
}
func (UnimplementedImagesServer) GetImageBundle(context.Context, *Id) (*ImageBundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageBundle not implemented")
}
func (UnimplementedImagesServer) GetImageData(context.Context, *Id) (*ImageRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageData not implemented")
}
func (UnimplementedImagesServer) DeleteImage(context.Context, *Id) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImagesServer) DeleteLocalizedImage(context.Context, *LocalizedImageInput) (*ImageRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocalizedImage not implemented")
}
func (UnimplementedImagesServer) ListImagesForUserDeprecated(*Pagination, Images_ListImagesForUserDeprecatedServer) error {
	return status.Errorf(codes.Unimplemented, "method ListImagesForUserDeprecated not implemented")
}
func (UnimplementedImagesServer) ListImagesForUser(*Filters, Images_ListImagesForUserServer) error {
	return status.Errorf(codes.Unimplemented, "method ListImagesForUser not implemented")
}
func (UnimplementedImagesServer) ListImagesDeprecated(*Pagination, Images_ListImagesDeprecatedServer) error {
	return status.Errorf(codes.Unimplemented, "method ListImagesDeprecated not implemented")
}
func (UnimplementedImagesServer) ListImages(*Filters, Images_ListImagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedImagesServer) CountImagesDeprecated(context.Context, *Pagination) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountImagesDeprecated not implemented")
}
func (UnimplementedImagesServer) CountImages(context.Context, *Filters) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountImages not implemented")
}
func (UnimplementedImagesServer) CountImagesForUserDeprecated(context.Context, *Pagination) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountImagesForUserDeprecated not implemented")
}
func (UnimplementedImagesServer) CountImagesForUser(context.Context, *Filters) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountImagesForUser not implemented")
}

// UnsafeImagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImagesServer will
// result in compilation errors.
type UnsafeImagesServer interface {
	mustEmbedUnimplementedImagesServer()
}

func RegisterImagesServer(s grpc.ServiceRegistrar, srv ImagesServer) {
	s.RegisterService(&Images_ServiceDesc, srv)
}

func _Images_SetProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).SetProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/setProfileImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).SetProfileImage(ctx, req.(*ProfileImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getProfileImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetProfileImage(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_CreateImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CreateImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/createImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CreateImages(ctx, req.(*CreateImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_UpdateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).UpdateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/updateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).UpdateImage(ctx, req.(*UpdateImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetImageURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetImageURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getImageURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetImageURL(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetStampImageURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StampImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetStampImageURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getStampImageURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetStampImageURL(ctx, req.(*StampImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetLocalizedImageURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalizedImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetLocalizedImageURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getLocalizedImageURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetLocalizedImageURL(ctx, req.(*LocalizedImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetProfileImageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetProfileImageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getProfileImageById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetProfileImageById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetImageBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetImageBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getImageBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetImageBundle(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_GetImageData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).GetImageData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/getImageData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).GetImageData(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/deleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).DeleteImage(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_DeleteLocalizedImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalizedImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).DeleteLocalizedImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/deleteLocalizedImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).DeleteLocalizedImage(ctx, req.(*LocalizedImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_ListImagesForUserDeprecated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImagesServer).ListImagesForUserDeprecated(m, &imagesListImagesForUserDeprecatedServer{stream})
}

type Images_ListImagesForUserDeprecatedServer interface {
	Send(*ImageRecord) error
	grpc.ServerStream
}

type imagesListImagesForUserDeprecatedServer struct {
	grpc.ServerStream
}

func (x *imagesListImagesForUserDeprecatedServer) Send(m *ImageRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Images_ListImagesForUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filters)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImagesServer).ListImagesForUser(m, &imagesListImagesForUserServer{stream})
}

type Images_ListImagesForUserServer interface {
	Send(*ImageRecord) error
	grpc.ServerStream
}

type imagesListImagesForUserServer struct {
	grpc.ServerStream
}

func (x *imagesListImagesForUserServer) Send(m *ImageRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Images_ListImagesDeprecated_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImagesServer).ListImagesDeprecated(m, &imagesListImagesDeprecatedServer{stream})
}

type Images_ListImagesDeprecatedServer interface {
	Send(*ImageRecord) error
	grpc.ServerStream
}

type imagesListImagesDeprecatedServer struct {
	grpc.ServerStream
}

func (x *imagesListImagesDeprecatedServer) Send(m *ImageRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Images_ListImages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filters)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImagesServer).ListImages(m, &imagesListImagesServer{stream})
}

type Images_ListImagesServer interface {
	Send(*ImageRecord) error
	grpc.ServerStream
}

type imagesListImagesServer struct {
	grpc.ServerStream
}

func (x *imagesListImagesServer) Send(m *ImageRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _Images_CountImagesDeprecated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CountImagesDeprecated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/countImagesDeprecated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CountImagesDeprecated(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_CountImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CountImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/countImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CountImages(ctx, req.(*Filters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_CountImagesForUserDeprecated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CountImagesForUserDeprecated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/countImagesForUserDeprecated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CountImagesForUserDeprecated(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_CountImagesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CountImagesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Images/countImagesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CountImagesForUser(ctx, req.(*Filters))
	}
	return interceptor(ctx, in, info, handler)
}

// Images_ServiceDesc is the grpc.ServiceDesc for Images service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Images_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setProfileImage",
			Handler:    _Images_SetProfileImage_Handler,
		},
		{
			MethodName: "getProfileImage",
			Handler:    _Images_GetProfileImage_Handler,
		},
		{
			MethodName: "createImages",
			Handler:    _Images_CreateImages_Handler,
		},
		{
			MethodName: "updateImage",
			Handler:    _Images_UpdateImage_Handler,
		},
		{
			MethodName: "getImageURL",
			Handler:    _Images_GetImageURL_Handler,
		},
		{
			MethodName: "getStampImageURL",
			Handler:    _Images_GetStampImageURL_Handler,
		},
		{
			MethodName: "getLocalizedImageURL",
			Handler:    _Images_GetLocalizedImageURL_Handler,
		},
		{
			MethodName: "getProfileImageById",
			Handler:    _Images_GetProfileImageById_Handler,
		},
		{
			MethodName: "getImageBundle",
			Handler:    _Images_GetImageBundle_Handler,
		},
		{
			MethodName: "getImageData",
			Handler:    _Images_GetImageData_Handler,
		},
		{
			MethodName: "deleteImage",
			Handler:    _Images_DeleteImage_Handler,
		},
		{
			MethodName: "deleteLocalizedImage",
			Handler:    _Images_DeleteLocalizedImage_Handler,
		},
		{
			MethodName: "countImagesDeprecated",
			Handler:    _Images_CountImagesDeprecated_Handler,
		},
		{
			MethodName: "countImages",
			Handler:    _Images_CountImages_Handler,
		},
		{
			MethodName: "countImagesForUserDeprecated",
			Handler:    _Images_CountImagesForUserDeprecated_Handler,
		},
		{
			MethodName: "countImagesForUser",
			Handler:    _Images_CountImagesForUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "listImagesForUserDeprecated",
			Handler:       _Images_ListImagesForUserDeprecated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listImagesForUser",
			Handler:       _Images_ListImagesForUser_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listImagesDeprecated",
			Handler:       _Images_ListImagesDeprecated_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "listImages",
			Handler:       _Images_ListImages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "io/core/a_rpc_images.proto",
}
