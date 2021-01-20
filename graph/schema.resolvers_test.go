package graph

import (
	"testing"

	"S3_FriendManagement_Graphql/graph/generated"
	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	"S3_FriendManagement_Graphql/modelss/api-models"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/handler"
	"github.com/friendsofgo/errors"
	"github.com/mitchellh/mapstructure"
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
		query                 string
		mockIsUserExisted     mockIsUserExisted
		mockCreateUserService mockCreateUserService
		expectedResult        *graphqlmodels.User
		expectedErr           error
	}{
		{
			name: "Format failed",
			query: `mutation {
					  createUser(input: {
						  email: "abcgmail.com"
					  }){
						  id,
						  email
					  }
					}`,
			expectedResult: &graphqlmodels.User{
				ID:    1,
				Email: "abc@gmail.com",
			},
			expectedErr: errors.New("[{\"message\":\"\\\"email\\\"'s format is not valid. (ex: \\\"andy@abc.xyz\\\")\",\"path\":[\"createUser\"]}]"),
		},
		{
			name: "Existed user",
			query: `mutation {
					  createUser(input: {
						  email: "abc@mail.com"
					  }){
						  id,
						  email
					  }
					}`,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@mail.com",
				result: true,
				err:    nil,
			},
			expectedResult: nil,
			expectedErr:    errors.New("[{\"message\":\"this email address existed\",\"path\":[\"createUser\"]}]"),
		},
		{
			name: "Success",
			query: `mutation {
					  createUser(input: {
						  email: "abc@gmail.com"
					  }){
						  id,
						  email
					  }
					}`,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@gmail.com",
				result: false,
				err:    nil,
			},
			mockCreateUserService: mockCreateUserService{
				input: &api_models.UserServiceInput{
					Email: "abc@gmail.com",
				},
				result: 1,
				err:    nil,
			},
			expectedResult: &graphqlmodels.User{
				ID:    1,
				Email: "abc@gmail.com",
			},
			expectedErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserService := new(mockUserService)
			mockFriendService := new(mockFriendService)
			mockSubscriptionService := new(mockSubscriptionService)
			mockBlockingService := new(mockBlockingService)

			mockUserService.On("IsExistedUser", testCase.mockIsUserExisted.input).
				Return(testCase.mockIsUserExisted.result, testCase.mockIsUserExisted.err)
			mockUserService.On("CreateUser", testCase.mockCreateUserService.input).
				Return(testCase.mockCreateUserService.result, testCase.mockCreateUserService.err)

			graphqlServer := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
				IUserService:         mockUserService,
				IFriendService:       mockFriendService,
				ISubscriptionService: mockSubscriptionService,
				IBlockingService:     mockBlockingService,
			}}))

			testClient := client.New(graphqlServer)
			result := &graphqlmodels.User{}
			if testCase.expectedErr != nil {
				err := testClient.Post(testCase.query, &result)
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				resp, err := testClient.RawPost(testCase.query)
				if err != nil {
					t.Error(err)
				}
				if err := mapstructure.Decode(resp.Data.(map[string]interface{})["createUser"], &result); err != nil {
					t.Error(err)
				}
				require.EqualValues(t, testCase.expectedResult, result)
			}
		})
	}
}
