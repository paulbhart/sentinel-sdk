package rpc

import (
	"context"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/proto/go"
	"google.golang.org/grpc"
)

// ImportPlugin is the goplugin.Plugin implementation to serve sdk.Import.
type ImportPlugin struct {
	goplugin.NetRPCUnsupportedPlugin

	F func() sdk.Import
}

func (p *ImportPlugin) GRPCServer(_ *goplugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterImportServer(s, &ImportGRPCServer{F: p.F})
	return nil
}

func (p *ImportPlugin) GRPCClient(_ context.Context, _ *goplugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ImportGRPCClient{Client: proto.NewImportClient(c)}, nil
}
