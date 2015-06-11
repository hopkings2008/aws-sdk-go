// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package glacier_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/glacier"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleGlacier_AbortMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.AbortMultipartUploadInput{
		AccountID: aws.StringPtr("string"), // Required
		UploadID:  aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.AbortMultipartUpload(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_AddTagsToVault() {
	svc := glacier.New(nil)

	params := &glacier.AddTagsToVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Tags: map[string]*string{
			"Key": aws.StringPtr("TagValue"), // Required
			// More values...
		},
	}
	resp, err := svc.AddTagsToVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_CompleteMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.CompleteMultipartUploadInput{
		AccountID:   aws.StringPtr("string"), // Required
		UploadID:    aws.StringPtr("string"), // Required
		VaultName:   aws.StringPtr("string"), // Required
		ArchiveSize: aws.StringPtr("string"),
		Checksum:    aws.StringPtr("string"),
	}
	resp, err := svc.CompleteMultipartUpload(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_CreateVault() {
	svc := glacier.New(nil)

	params := &glacier.CreateVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.CreateVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DeleteArchive() {
	svc := glacier.New(nil)

	params := &glacier.DeleteArchiveInput{
		AccountID: aws.StringPtr("string"), // Required
		ArchiveID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DeleteArchive(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DeleteVault() {
	svc := glacier.New(nil)

	params := &glacier.DeleteVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DeleteVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DeleteVaultAccessPolicy() {
	svc := glacier.New(nil)

	params := &glacier.DeleteVaultAccessPolicyInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DeleteVaultAccessPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DeleteVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.DeleteVaultNotificationsInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DeleteVaultNotifications(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DescribeJob() {
	svc := glacier.New(nil)

	params := &glacier.DescribeJobInput{
		AccountID: aws.StringPtr("string"), // Required
		JobID:     aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DescribeJob(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_DescribeVault() {
	svc := glacier.New(nil)

	params := &glacier.DescribeVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.DescribeVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_GetDataRetrievalPolicy() {
	svc := glacier.New(nil)

	params := &glacier.GetDataRetrievalPolicyInput{
		AccountID: aws.StringPtr("string"), // Required
	}
	resp, err := svc.GetDataRetrievalPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_GetJobOutput() {
	svc := glacier.New(nil)

	params := &glacier.GetJobOutputInput{
		AccountID: aws.StringPtr("string"), // Required
		JobID:     aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Range:     aws.StringPtr("string"),
	}
	resp, err := svc.GetJobOutput(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_GetVaultAccessPolicy() {
	svc := glacier.New(nil)

	params := &glacier.GetVaultAccessPolicyInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.GetVaultAccessPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_GetVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.GetVaultNotificationsInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.GetVaultNotifications(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_InitiateJob() {
	svc := glacier.New(nil)

	params := &glacier.InitiateJobInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		JobParameters: &glacier.JobParameters{
			ArchiveID:   aws.StringPtr("string"),
			Description: aws.StringPtr("string"),
			Format:      aws.StringPtr("string"),
			InventoryRetrievalParameters: &glacier.InventoryRetrievalJobInput{
				EndDate:   aws.StringPtr("string"),
				Limit:     aws.StringPtr("string"),
				Marker:    aws.StringPtr("string"),
				StartDate: aws.StringPtr("string"),
			},
			RetrievalByteRange: aws.StringPtr("string"),
			SNSTopic:           aws.StringPtr("string"),
			Type:               aws.StringPtr("string"),
		},
	}
	resp, err := svc.InitiateJob(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_InitiateMultipartUpload() {
	svc := glacier.New(nil)

	params := &glacier.InitiateMultipartUploadInput{
		AccountID:          aws.StringPtr("string"), // Required
		VaultName:          aws.StringPtr("string"), // Required
		ArchiveDescription: aws.StringPtr("string"),
		PartSize:           aws.StringPtr("string"),
	}
	resp, err := svc.InitiateMultipartUpload(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_ListJobs() {
	svc := glacier.New(nil)

	params := &glacier.ListJobsInput{
		AccountID:  aws.StringPtr("string"), // Required
		VaultName:  aws.StringPtr("string"), // Required
		Completed:  aws.StringPtr("string"),
		Limit:      aws.StringPtr("string"),
		Marker:     aws.StringPtr("string"),
		Statuscode: aws.StringPtr("string"),
	}
	resp, err := svc.ListJobs(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_ListMultipartUploads() {
	svc := glacier.New(nil)

	params := &glacier.ListMultipartUploadsInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Limit:     aws.StringPtr("string"),
		Marker:    aws.StringPtr("string"),
	}
	resp, err := svc.ListMultipartUploads(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_ListParts() {
	svc := glacier.New(nil)

	params := &glacier.ListPartsInput{
		AccountID: aws.StringPtr("string"), // Required
		UploadID:  aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Limit:     aws.StringPtr("string"),
		Marker:    aws.StringPtr("string"),
	}
	resp, err := svc.ListParts(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_ListTagsForVault() {
	svc := glacier.New(nil)

	params := &glacier.ListTagsForVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
	}
	resp, err := svc.ListTagsForVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_ListVaults() {
	svc := glacier.New(nil)

	params := &glacier.ListVaultsInput{
		AccountID: aws.StringPtr("string"), // Required
		Limit:     aws.StringPtr("string"),
		Marker:    aws.StringPtr("string"),
	}
	resp, err := svc.ListVaults(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_RemoveTagsFromVault() {
	svc := glacier.New(nil)

	params := &glacier.RemoveTagsFromVaultInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		TagKeys: []*string{
			aws.StringPtr("string"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromVault(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_SetDataRetrievalPolicy() {
	svc := glacier.New(nil)

	params := &glacier.SetDataRetrievalPolicyInput{
		AccountID: aws.StringPtr("string"), // Required
		Policy: &glacier.DataRetrievalPolicy{
			Rules: []*glacier.DataRetrievalRule{
				{ // Required
					BytesPerHour: aws.Int64Ptr(1),
					Strategy:     aws.StringPtr("string"),
				},
				// More values...
			},
		},
	}
	resp, err := svc.SetDataRetrievalPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_SetVaultAccessPolicy() {
	svc := glacier.New(nil)

	params := &glacier.SetVaultAccessPolicyInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Policy: &glacier.VaultAccessPolicy{
			Policy: aws.StringPtr("string"),
		},
	}
	resp, err := svc.SetVaultAccessPolicy(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_SetVaultNotifications() {
	svc := glacier.New(nil)

	params := &glacier.SetVaultNotificationsInput{
		AccountID: aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		VaultNotificationConfig: &glacier.VaultNotificationConfig{
			Events: []*string{
				aws.StringPtr("string"), // Required
				// More values...
			},
			SNSTopic: aws.StringPtr("string"),
		},
	}
	resp, err := svc.SetVaultNotifications(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_UploadArchive() {
	svc := glacier.New(nil)

	params := &glacier.UploadArchiveInput{
		AccountID:          aws.StringPtr("string"), // Required
		VaultName:          aws.StringPtr("string"), // Required
		ArchiveDescription: aws.StringPtr("string"),
		Body:               bytes.NewReader([]byte("PAYLOAD")),
		Checksum:           aws.StringPtr("string"),
	}
	resp, err := svc.UploadArchive(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleGlacier_UploadMultipartPart() {
	svc := glacier.New(nil)

	params := &glacier.UploadMultipartPartInput{
		AccountID: aws.StringPtr("string"), // Required
		UploadID:  aws.StringPtr("string"), // Required
		VaultName: aws.StringPtr("string"), // Required
		Body:      bytes.NewReader([]byte("PAYLOAD")),
		Checksum:  aws.StringPtr("string"),
		Range:     aws.StringPtr("string"),
	}
	resp, err := svc.UploadMultipartPart(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
