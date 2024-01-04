package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	pb "github.com/zbaylab/openbytes/api/go"
	impl "github.com/zbaylab/openbytes/pkg/impl/service"
	"google.golang.org/grpc"
)

var (
	port    int
	rootCmd = &cobra.Command{
		Use:   "Launch gRPC && UI server",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			startMuxServer()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	log.SetReportCaller(true)
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8888, "server port")
}

//go:embed all:dist
var ui embed.FS

func startMuxServer() {
	grpcServer := grpc.NewServer()
	pb.RegisterCapturesServer(grpcServer, &impl.CapturesImpl{})

	mux := http.NewServeMux()
	dist, err := fs.Sub(ui, "dist")
	if err != nil {
		log.Fatalf("sub error: %s", err)
		return
	}
	mux.Handle("/", fileServerWithExt(http.FS(dist)))
	log.Infof("listening: %d", port)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(r) {
			wrappedGrpc.ServeHTTP(w, r)
			return
		}
		mux.ServeHTTP(w, r)
	})))
}

// Handle 404 to "/index.html" and gzip compression
func fileServerWithExt(root http.FileSystem) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := root.Open(r.URL.Path)
		if err != nil && os.IsNotExist(err) {
			r.URL.Path = "/"
		}
		if err == nil {
			f.Close()
		}
		http.FileServer(root).ServeHTTP(w, r)
	}))
}
