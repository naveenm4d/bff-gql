package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	graph "github.com/naveenm4d/bff-gql/graph/model"
	resolvers "github.com/naveenm4d/bff-gql/graph/resolvers"
	"github.com/naveenm4d/bff-gql/internal/app/services/query"
	"github.com/naveenm4d/bff-gql/internal/config"
	blogpb "github.com/naveenm4d/blog-svc/proto"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Cancelling", sig)

		cancel()
	}()

	fmt.Println("Initializing blog service client...")

	blogSvcConn, err := grpc.NewClient(*config.Config.BlogSvcURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer blogSvcConn.Close()

	blogSvclient := blogpb.NewBlogSvcClient(blogSvcConn)

	fmt.Println("Initializing services...")

	queryService := query.NewQueryService(blogSvclient)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		QueryService: queryService,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *config.Config.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+*config.Config.HTTPPort, nil))
}
