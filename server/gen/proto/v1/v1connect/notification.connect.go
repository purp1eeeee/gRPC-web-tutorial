// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/v1/notification.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	v1 "server/gen/proto/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// NotificationServiceName is the fully-qualified name of the NotificationService service.
	NotificationServiceName = "proto.v1.NotificationService"
)

// NotificationServiceClient is a client for the proto.v1.NotificationService service.
type NotificationServiceClient interface {
	SendNotification(context.Context, *connect_go.Request[v1.SendNotificationRequest]) (*connect_go.Response[emptypb.Empty], error)
	NotificationList(context.Context, *connect_go.Request[v1.NotificationListRequest]) (*connect_go.Response[v1.NotificationListResponse], error)
}

// NewNotificationServiceClient constructs a client for the proto.v1.NotificationService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNotificationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NotificationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &notificationServiceClient{
		sendNotification: connect_go.NewClient[v1.SendNotificationRequest, emptypb.Empty](
			httpClient,
			baseURL+"/proto.v1.NotificationService/SendNotification",
			opts...,
		),
		notificationList: connect_go.NewClient[v1.NotificationListRequest, v1.NotificationListResponse](
			httpClient,
			baseURL+"/proto.v1.NotificationService/NotificationList",
			opts...,
		),
	}
}

// notificationServiceClient implements NotificationServiceClient.
type notificationServiceClient struct {
	sendNotification *connect_go.Client[v1.SendNotificationRequest, emptypb.Empty]
	notificationList *connect_go.Client[v1.NotificationListRequest, v1.NotificationListResponse]
}

// SendNotification calls proto.v1.NotificationService.SendNotification.
func (c *notificationServiceClient) SendNotification(ctx context.Context, req *connect_go.Request[v1.SendNotificationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.sendNotification.CallUnary(ctx, req)
}

// NotificationList calls proto.v1.NotificationService.NotificationList.
func (c *notificationServiceClient) NotificationList(ctx context.Context, req *connect_go.Request[v1.NotificationListRequest]) (*connect_go.Response[v1.NotificationListResponse], error) {
	return c.notificationList.CallUnary(ctx, req)
}

// NotificationServiceHandler is an implementation of the proto.v1.NotificationService service.
type NotificationServiceHandler interface {
	SendNotification(context.Context, *connect_go.Request[v1.SendNotificationRequest]) (*connect_go.Response[emptypb.Empty], error)
	NotificationList(context.Context, *connect_go.Request[v1.NotificationListRequest]) (*connect_go.Response[v1.NotificationListResponse], error)
}

// NewNotificationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNotificationServiceHandler(svc NotificationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/proto.v1.NotificationService/SendNotification", connect_go.NewUnaryHandler(
		"/proto.v1.NotificationService/SendNotification",
		svc.SendNotification,
		opts...,
	))
	mux.Handle("/proto.v1.NotificationService/NotificationList", connect_go.NewUnaryHandler(
		"/proto.v1.NotificationService/NotificationList",
		svc.NotificationList,
		opts...,
	))
	return "/proto.v1.NotificationService/", mux
}

// UnimplementedNotificationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNotificationServiceHandler struct{}

func (UnimplementedNotificationServiceHandler) SendNotification(context.Context, *connect_go.Request[v1.SendNotificationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.v1.NotificationService.SendNotification is not implemented"))
}

func (UnimplementedNotificationServiceHandler) NotificationList(context.Context, *connect_go.Request[v1.NotificationListRequest]) (*connect_go.Response[v1.NotificationListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.v1.NotificationService.NotificationList is not implemented"))
}
