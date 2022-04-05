package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fullstorydev/grpcui/standalone"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cc, err := grpc.DialContext(ctx, "dns:///localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot DialContext: %v", err)
	}
	h, err := standalone.HandlerViaReflection(ctx, cc, "dns:///my-grpc-server:8080")
	if err != nil {
		log.Fatalf("cannot get handler: %v", err)
	}

	http.ListenAndServe(":8081", h)
}
