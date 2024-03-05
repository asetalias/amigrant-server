package v1

import (
	"testing"

	"github.com/asetalias/amigrant-server/internal/dto"
	"github.com/stretchr/testify/require"
)

// This file should contain functions to start a Test and various functions to be repetited in the testing process
// We will be using the table driven testing approach to test the various test cases
func TestPingPong(t *testing.T) {
	testCases := []struct {
		name          string
		checkResponse func(t *testing.T, resp *string, err *dto.ErrorResponse)
	}{
		{
			name: "Test Public Hello",
			checkResponse: func(t *testing.T, resp *string, err *dto.ErrorResponse) {
				require.NotNil(t, resp)
				require.Nil(t, err)
				require.Equal(t, "pong", *resp)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			routeHandler := NewRouteHandler()
			context, _ := createTestContext()

			resp, err := routeHandler.Ping(context)

			tc.checkResponse(t, resp, err)
		})
	}
}
