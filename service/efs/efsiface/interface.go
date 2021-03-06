// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package efsiface provides an interface for the Amazon Elastic File System.
package efsiface

import (
	"github.com/aws/aws-sdk-go/service/efs"
)

// EFSAPI is the interface type for efs.EFS.
type EFSAPI interface {
	CreateFileSystem(*efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)

	CreateMountTarget(*efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)

	CreateTags(*efs.CreateTagsInput) (*efs.CreateTagsOutput, error)

	DeleteFileSystem(*efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)

	DeleteMountTarget(*efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)

	DeleteTags(*efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)

	DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)

	DescribeMountTargetSecurityGroups(*efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)

	DescribeMountTargets(*efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)

	DescribeTags(*efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)

	ModifyMountTargetSecurityGroups(*efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
}
