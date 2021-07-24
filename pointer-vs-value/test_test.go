package pointer_vs_value

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/joshcarp/benchmarks/pointer-vs-value/example"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	port, ln, err := example.ServeRand(example.ServerImplValue{})
	defer ln.Close()
	require.NoError(b, err)
	conn, err := setup(true, fmt.Sprintf("localhost:%d", port))
	require.NoError(b, err)
	client := example.NewExampleServiceClient(conn)
	for i := 0; i < 100000; i ++ {
		_, _ = client.GetExample(context.Background(), &example.Example{Whatever: "", Name: ""})
	}
}


func BenchmarkPointer(b *testing.B) {
	port, ln, err := example.ServeRand(&example.ServerImplPointer{})
	defer ln.Close()
	require.NoError(b, err)
	conn, err := setup(true, fmt.Sprintf("localhost:%d", port))
	require.NoError(b, err)
	client := example.NewExampleServiceClient(conn)
	for i := 0; i < 100000; i ++ {
		_, _ = client.GetExample(context.Background(), &example.Example{Whatever: "", Name: ""})
	}
}



func setup(plaintext bool, targetURL string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}
	if !plaintext {
		cp, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		opts = []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(cp, "")),
		}
	}
	ctx := context.Background()
	cc, err := grpc.DialContext(ctx, targetURL, opts...)
	if err != nil {
		return nil, fmt.Errorf("%v: failed to connect to server", err)
	}
	return cc, nil
}