package cloudsearchdomain_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/stretchr/testify/assert"
)

func TestRequireEndpointIfRegionProvided(t *testing.T) {
	svc := cloudsearchdomain.New(&aws.Config{
		Region:                 aws.NewString("mock-region"),
		DisableParamValidation: aws.NewBool(true),
	})
	req, _ := svc.SearchRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointIfNoRegionProvided(t *testing.T) {
	svc := cloudsearchdomain.New(&aws.Config{
		Region:                 aws.NewString(""),
		DisableParamValidation: aws.NewBool(true),
	})
	req, _ := svc.SearchRequest(nil)
	err := req.Build()

	assert.Equal(t, "", svc.Endpoint)
	assert.Error(t, err)
	assert.Equal(t, aws.ErrMissingEndpoint, err)
}

func TestRequireEndpointUsed(t *testing.T) {
	svc := cloudsearchdomain.New(&aws.Config{
		Region:                 aws.NewString("mock-region"),
		DisableParamValidation: aws.NewBool(true),
		Endpoint:               aws.NewString("https://endpoint"),
	})
	req, _ := svc.SearchRequest(nil)
	err := req.Build()

	assert.Equal(t, "https://endpoint", svc.Endpoint)
	assert.NoError(t, err)
}
