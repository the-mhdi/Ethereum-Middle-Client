package node

import (
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/rpc"
)

type servers struct {
	http     *httpServer
	ws       *httpServer
	httpAuth *httpServer
	wsAuth   *httpServer
	ipc      *ipcServer //for local communication
}

// same as geth's httpServer, but with some modifications

type httpServer struct {
	log      log.Logger
	timeouts rpc.HTTPTimeouts
	mux      http.ServeMux // registered handlers go here

	mu       sync.Mutex
	server   *http.Server
	listener net.Listener // non-nil when server is running

	// HTTP RPC handler things.

	httpConfig  httpConfig
	httpHandler atomic.Pointer[rpcHandler]

	// WebSocket handler things.
	wsConfig  wsConfig
	wsHandler atomic.Pointer[rpcHandler]

	// These are set by setListenAddr.
	endpoint string
	host     string
	port     int

	handlerNames map[string]string
}

type ipcServer struct {
	log      log.Logger
	endpoint string

	mu       sync.Mutex
	listener net.Listener
	srv      *rpc.Server
}

// httpConfig is the JSON-RPC/HTTP configuration.
type httpConfig struct {
	Modules            []string
	CorsAllowedOrigins []string
	Vhosts             []string
	prefix             string // path prefix on which to mount http handler
	rpcEndpointConfig
}

// wsConfig is the JSON-RPC/Websocket configuration
type wsConfig struct {
	Origins []string
	Modules []string
	prefix  string // path prefix on which to mount ws handler
	rpcEndpointConfig
}

type rpcEndpointConfig struct {
	jwtSecret              []byte // optional JWT secret
	batchItemLimit         int
	batchResponseSizeLimit int
	httpBodyLimit          int
}

type rpcHandler struct {
	http.Handler
	prefix string
	server *rpc.Server
}
