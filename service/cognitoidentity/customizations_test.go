package cognitoidentity_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/stretchr/testify/assert"
)

var svc = cognitoidentity.New(&aws.Config{
	Region: aws.NewString("mock-region"),
})

func TestUnsignedRequest_GetID(t *testing.T) {
	req, _ := svc.GetIDRequest(&cognitoidentity.GetIDInput{
		IdentityPoolID: aws.StringPtr("IdentityPoolId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetOpenIDToken(t *testing.T) {
	req, _ := svc.GetOpenIDTokenRequest(&cognitoidentity.GetOpenIDTokenInput{
		IdentityID: aws.StringPtr("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_GetCredentialsForIdentity(t *testing.T) {
	req, _ := svc.GetCredentialsForIdentityRequest(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityID: aws.StringPtr("IdentityId"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
