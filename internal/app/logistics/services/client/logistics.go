package client

import (
    "context"

    api "github.com/coopnorge/interview-backend/internal/app/logistics/api/v1"
    "github.com/google/wire"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

// ServiceSetForClient providers
var ServiceSetForClient = wire.NewSet(NewLogisticsClient)

// APILogisticsClient to send requests about cargo unit movements
type APILogisticsClient struct {
    api  api.CoopLogisticsEngineAPIClient
    conn *grpc.ClientConn
}

// NewLogisticsClient instance
func NewLogisticsClient() *APILogisticsClient {
    return &APILogisticsClient{}
}

// Connect to gRPC API
func (lc *APILogisticsClient) Connect(serverAddr string, ctx context.Context) error {
    conn, dialErr := grpc.DialContext(
        ctx,
        serverAddr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(),
    )
    if dialErr != nil {
        return dialErr
    }

    lc.conn = conn
    lc.api = api.NewCoopLogisticsEngineAPIClient(lc.conn)

    return nil
}

// Disconnect from gRPC API
func (lc *APILogisticsClient) Disconnect() error {
    return lc.conn.Close()
}

// MoveUnit to new location
func (lc *APILogisticsClient) MoveUnit(ctx context.Context, req *api.MoveUnitRequest) error {
    _, moveRespErr := lc.api.MoveUnit(ctx, req)
    if moveRespErr != nil {
        return moveRespErr
    }

    return nil
}

// UnitReachedWarehouse report that reach warehouse
func (lc *APILogisticsClient) UnitReachedWarehouse(ctx context.Context, req *api.UnitReachedWarehouseRequest) error {
    _, moveRespErr := lc.api.UnitReachedWarehouse(ctx, req)
    if moveRespErr != nil {
        return moveRespErr
    }

    return nil
}
