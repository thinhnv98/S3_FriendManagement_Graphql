// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphqlmodels

type Friends struct {
	Friends []string `json:"friends"`
}

type IsFriend struct {
	Success bool `json:"success"`
}

type NewUser struct {
	Email string `json:"email"`
}

type Success struct {
	Status string `json:"status"`
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}