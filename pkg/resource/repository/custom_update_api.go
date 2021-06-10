// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ecr"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

var (
	defaultImageScanningConfig = svcsdk.ImageScanningConfiguration{
		ScanOnPush: aws.Bool(false),
	}
	defaultImageTagMutability = svcsdk.ImageTagMutabilityMutable
)

// customUpdateRepository implements specialized logic for handling Repository
// resource updates. The ECR API has 4 separate API calls to update a
// Repository, depending on the Repository attribute that has changed:
//
// * PutImageScanningConfiguration for when the
//   Repository.imageScanningConfiguration struct changed
// * PutImageTagMutability for when the Repository.imageTagMutability attribute
//   changed
// * PutLifecyclePolicy for when the Repository.lifecyclePolicy changed
// * SetRepositoryPolicy for when the Repository.policy changed (yes, it uses
//   "Set" and not "Put"... no idea why this is inconsistent)
func (rm *resourceManager) customUpdateRepository(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	var err error
	var updated *resource
	updated = desired
	fmt.Println("called custom")
	if delta.DifferentAt("ScanConfig") {
		fmt.Println("FOUND DIFF")
		updated, err = rm.updateImageScanningConfiguration(ctx, updated)
		if err != nil {
			return nil, err
		}
	}
	if delta.DifferentAt("ImageTagMutability") {
		updated, err = rm.updateImageTagMutability(ctx, updated)
		if err != nil {
			return nil, err
		}
	}
	return updated, nil
}

// updateImageScanningConfiguration calls the PutImageScanningConfiguration ECR
// API call for a specific repository
func (rm *resourceManager) updateImageScanningConfiguration(
	ctx context.Context,
	desired *resource,
) (*resource, error) {
	fmt.Println("UPDATING REPO")
	dspec := desired.ko.Spec
	input := &svcsdk.PutImageScanningConfigurationInput{
		RepositoryName: aws.String(*dspec.Name),
	}
	if dspec.ScanConfig == nil {
		// There isn't any "reset" behaviour and the image scanning
		// configuration field should always be set...
		input.SetImageScanningConfiguration(&defaultImageScanningConfig)
	} else {
		isc := svcsdk.ImageScanningConfiguration{
			ScanOnPush: dspec.ScanConfig.ScanOnPush,
		}
		input.SetImageScanningConfiguration(&isc)
	}
	fmt.Println("UPDATING REPO with", input.ImageScanningConfiguration.ScanOnPush)
	_, err := rm.sdkapi.PutImageScanningConfigurationWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return desired, nil
}

// updateImageTagMutability calls the PutImageTagMutability ECR API call for a
// specific repository
func (rm *resourceManager) updateImageTagMutability(
	ctx context.Context,
	desired *resource,
) (*resource, error) {
	dspec := desired.ko.Spec
	input := &svcsdk.PutImageTagMutabilityInput{
		RepositoryName: aws.String(*dspec.Name),
	}
	if dspec.ImageTagMutability == nil {
		// There isn't any "reset" behaviour and the image scanning
		// configuration field should always be set...
		input.SetImageTagMutability(defaultImageTagMutability)
	} else {
		input.SetImageTagMutability(*dspec.ImageTagMutability)
	}
	_, err := rm.sdkapi.PutImageTagMutabilityWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return desired, nil
}
