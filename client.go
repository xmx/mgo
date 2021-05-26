package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientOperator struct {
	*mongo.Client
}

// New creates a new Client and then initializes it using the Connect method. This is equivalent to calling
// NewClient followed by Client.Connect.
func New(opts ...*options.ClientOptions) (*ClientOperator, error) {
	conn, err := mongo.Connect(nil, opts...)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(nil, nil); err != nil {
		return nil, err
	}
	return &ClientOperator{Client: conn}, nil
}

// Database 创建 DatabaseOperator
func (o *ClientOperator) Database(name string, opts ...*options.DatabaseOptions) *DatabaseOperator {
	return &DatabaseOperator{Database: o.Client.Database(name, opts...)}
}

// Disconnect closes sockets to the topology referenced by this Client. It will
// shut down any monitoring goroutines, close the idle connection pool, and will
// wait until all the in use connections have been returned to the connection
// pool and closed before returning. If the context expires via cancellation,
// deadline, or timeout before the in use connections have returned, the in use
// connections will be closed, resulting in the failure of any in flight read
// or write operations. If this method returns with no errors, all connections
// associated with this Client have been closed.
func (o *ClientOperator) Disconnect(ctx context.Context) {
	_ = o.Client.Disconnect(ctx)
}
