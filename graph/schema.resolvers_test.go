package graph

import (
	"context"
	"testing"

	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	"S3_FriendManagement_Graphql/modelss/api-models"
	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/require"
)

func TestMutationResolver_CreateUser(t *testing.T) {
	type mockIsUserExisted struct {
		input  string
		result bool
		err    error
	}

	type mockCreateUserService struct {
		input  *api_models.UserServiceInput
		result int
		err    error
	}

	testCases := []struct {
		name                  string
		input                 graphqlmodels.NewUser
		mockIsUserExisted     mockIsUserExisted
		mockCreateUserService mockCreateUserService
		expectedResult        *graphqlmodels.User
		expectedError         error
	}{
		{
			name: "Failed",
			input: graphqlmodels.NewUser{
				Email: "gmail.com",
			},
			expectedResult: nil,
			expectedError:  errors.New("\"email\"'s format is not valid. (ex: \"andy@abc.xyz\")"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			ctx := context.Background()

			mockUserService := new(mockUserService)

			mockUserService.On("IsExistedUser", testCase.mockIsUserExisted.input).
				Return(testCase.mockIsUserExisted.result, testCase.mockIsUserExisted.err)
			mockUserService.On("CreateUser", testCase.mockCreateUserService.input).
				Return(testCase.mockCreateUserService.result, testCase.mockCreateUserService.err)

			r := Resolver{
				IUserService: mockUserService,
			}
			mut := r.Mutation()

			//When
			result, err := mut.CreateUser(ctx, testCase.input)

			//Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}
