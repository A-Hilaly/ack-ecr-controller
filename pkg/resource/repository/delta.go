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

// Code generated by ack-generate. DO NOT EDIT.

package repository

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration) {
		delta.Add("Spec.EncryptionConfiguration", a.ko.Spec.EncryptionConfiguration, b.ko.Spec.EncryptionConfiguration)
	} else if a.ko.Spec.EncryptionConfiguration != nil && b.ko.Spec.EncryptionConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType) {
			delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
		} else if a.ko.Spec.EncryptionConfiguration.EncryptionType != nil && b.ko.Spec.EncryptionConfiguration.EncryptionType != nil {
			if *a.ko.Spec.EncryptionConfiguration.EncryptionType != *b.ko.Spec.EncryptionConfiguration.EncryptionType {
				delta.Add("Spec.EncryptionConfiguration.EncryptionType", a.ko.Spec.EncryptionConfiguration.EncryptionType, b.ko.Spec.EncryptionConfiguration.EncryptionType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey) {
			delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
		} else if a.ko.Spec.EncryptionConfiguration.KMSKey != nil && b.ko.Spec.EncryptionConfiguration.KMSKey != nil {
			if *a.ko.Spec.EncryptionConfiguration.KMSKey != *b.ko.Spec.EncryptionConfiguration.KMSKey {
				delta.Add("Spec.EncryptionConfiguration.KMSKey", a.ko.Spec.EncryptionConfiguration.KMSKey, b.ko.Spec.EncryptionConfiguration.KMSKey)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability) {
		delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
	} else if a.ko.Spec.ImageTagMutability != nil && b.ko.Spec.ImageTagMutability != nil {
		if *a.ko.Spec.ImageTagMutability != *b.ko.Spec.ImageTagMutability {
			delta.Add("Spec.ImageTagMutability", a.ko.Spec.ImageTagMutability, b.ko.Spec.ImageTagMutability)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ScanConfig, b.ko.Spec.ScanConfig) {
		delta.Add("Spec.ScanConfig", a.ko.Spec.ScanConfig, b.ko.Spec.ScanConfig)
	} else if a.ko.Spec.ScanConfig != nil && b.ko.Spec.ScanConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ScanConfig.ScanOnPush, b.ko.Spec.ScanConfig.ScanOnPush) {
			delta.Add("Spec.ScanConfig.ScanOnPush", a.ko.Spec.ScanConfig.ScanOnPush, b.ko.Spec.ScanConfig.ScanOnPush)
		} else if a.ko.Spec.ScanConfig.ScanOnPush != nil && b.ko.Spec.ScanConfig.ScanOnPush != nil {
			if *a.ko.Spec.ScanConfig.ScanOnPush != *b.ko.Spec.ScanConfig.ScanOnPush {
				delta.Add("Spec.ScanConfig.ScanOnPush", a.ko.Spec.ScanConfig.ScanOnPush, b.ko.Spec.ScanConfig.ScanOnPush)
			}
		}
	}

	return delta
}
