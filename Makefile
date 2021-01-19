run:
	go run main.go

graph-gen:
	go run github.com/99designs/gqlgen generate

orm-gen:
	sqlboiler --wipe psql
