package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/veqryn/protobuf/ptypes"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // Allow gzip encoding of grpc responses
	"google.golang.org/grpc/reflection"
)

type TestAPI struct {
	UnimplementedTestAPIServer
}

func (*TestAPI) Timestamp(ctx context.Context, req *TimestampReq) (*TimestampResp, error) {
	log.Println(req)
	if req.MyTime != nil {
		return &TimestampResp{MyTime: req.MyTime}, nil
	}
	myTime, err := ptypes.StringTimestamp(time.RFC3339Nano, myTime)
	return &TimestampResp{MyTime: myTime}, err
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var (
	myTime            = "2019-08-14T08:56:15.985618600Z"
	grpcServer        *grpc.Server
	httpRestServer    *http.Server
	gatewayCancelFunc context.CancelFunc
)

func setup() {
	// API function handler
	api := &TestAPI{}

	// Create and run GRPC Server
	grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_validator.StreamServerInterceptor())),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_validator.UnaryServerInterceptor())),
	)
	RegisterTestAPIServer(grpcServer, api)
	reflection.Register(grpcServer)

	go func() {
		listener, listenErr := net.Listen("tcp", ":50051")
		if listenErr != nil {
			log.Fatal(listenErr)
		}
		if serveErr := grpcServer.Serve(listener); serveErr != nil {
			log.Fatal(serveErr)
		}
	}()

	// Create GRPC Gateway Router/Mux
	gwMux := gwruntime.NewServeMux(gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{OrigName: true, EmitDefaults: true}))

	// Register GRPC Gateway with GRPC Server
	var gatewayCtx context.Context
	gatewayCtx, gatewayCancelFunc = context.WithCancel(context.Background())
	if err := RegisterTestAPIHandlerFromEndpoint(gatewayCtx, gwMux, ":50051", []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatal()
	}

	// Create and run HTTP REST Server
	httpRestServer = &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		Handler:      h2c.NewHandler(gwMux, &http2.Server{}), // Wrap to support http/2 cleartext
	}
	go func() {
		if serveErr := httpRestServer.ListenAndServe(); serveErr != nil && serveErr != http.ErrServerClosed {
			log.Fatal(serveErr)
		}
	}()

}
func shutdown() {
	grpcServer.GracefulStop()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer shutdownCancel()
	if shutdownErr := httpRestServer.Shutdown(shutdownCtx); shutdownErr != nil {
		log.Println(shutdownErr)
	}
	gatewayCancelFunc()
}

func TestGatewayMarshalResponse(t *testing.T) {
	t.Parallel()

	resp, err := http.Get("http://localhost:8080/timestamp")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectation := fmt.Sprintf(`{"my_time":"%s"}`, myTime)
	if string(body) != expectation {
		t.Error("Expected: ", expectation, "; Got: ", string(body))
	}
}

func TestGatewayMarshalRequest(t *testing.T) {
	t.Parallel()

	// Get a timestamp with at least one nanosecond on it
	newTime := time.Now().Truncate(10 * time.Nanosecond).Add(time.Nanosecond).UTC().Format(time.RFC3339Nano)
	resp, err := http.Get("http://localhost:8080/timestamp?my_time=" + newTime)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectation := fmt.Sprintf(`{"my_time":"%s"}`, newTime)
	if string(body) != expectation {
		t.Error("Expected: ", expectation, "; Got: ", string(body))
	}
}
