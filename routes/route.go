package routes

import (
	"database/sql"

	"S3_FriendManagement_Graphql/graph"
	"S3_FriendManagement_Graphql/graph/generated"
	"S3_FriendManagement_Graphql/repositories"
	"S3_FriendManagement_Graphql/services"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
)

type Routes struct {
	Db *sql.DB
}

func (_self Routes) Register() *chi.Mux {
	chiServer := chi.NewRouter()
	//GraphQL chiServer

	graphqlServer := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		IUserService: services.UserService{
			IUserRepo: repositories.UserRepo{
				Db: _self.Db,
			},
		},
		IFriendService: services.FriendService{
			IFriendRepo: repositories.FriendRepo{
				Db: _self.Db,
			},
			IUserRepo: repositories.UserRepo{
				Db: _self.Db,
			},
		},
	}}))

	chiServer.Handle("/", playground.Handler("GraphQL playground", "/query"))
	chiServer.Handle("/query", graphqlServer)

	return chiServer
}
