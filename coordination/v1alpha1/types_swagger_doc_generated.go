/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_LeaseCandidate = map[string]string{
	"":         "LeaseCandidate defines a candidate for a Lease object. Candidates are created such that coordinated leader election will pick the best leader from the list of candidates.",
	"metadata": "More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec contains the specification of the Lease. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (LeaseCandidate) SwaggerDoc() map[string]string {
	return map_LeaseCandidate
}

var map_LeaseCandidateList = map[string]string{
	"":         "LeaseCandidateList is a list of Lease objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is a list of schema objects.",
}

func (LeaseCandidateList) SwaggerDoc() map[string]string {
	return map_LeaseCandidateList
}

var map_LeaseCandidateSpec = map[string]string{
	"":                    "LeaseCandidateSpec is a specification of a Lease.",
	"leaseName":           "LeaseName is the name of the lease for which this candidate is contending. This field is immutable.",
	"pingTime":            "PingTime is the last time that the server has requested the LeaseCandidate to renew. It is only done during leader election to check if any LeaseCandidates have become ineligible. When PingTime is updated, the LeaseCandidate will respond by updating RenewTime.",
	"renewTime":           "RenewTime is the time that the LeaseCandidate was last updated. Any time a Lease needs to do leader election, the PingTime field is updated to signal to the LeaseCandidate that they should update the RenewTime. Old LeaseCandidate objects are also garbage collected if it has been hours since the last renew. The PingTime field is updated regularly to prevent garbage collection for still active LeaseCandidates.",
	"binaryVersion":       "BinaryVersion is the binary version. It must be in a semver format without leading `v`. This field is required when strategy is \"OldestEmulationVersion\"",
	"emulationVersion":    "EmulationVersion is the emulation version. It must be in a semver format without leading `v`. EmulationVersion must be less than or equal to BinaryVersion. This field is required when strategy is \"OldestEmulationVersion\"",
	"preferredStrategies": "PreferredStrategies indicates the list of strategies for picking the leader for coordinated leader election. The list is ordered, and the first strategy supersedes all other strategies. The list is used by coordinated leader election to make a decision about the final election strategy. This follows as - If all clients have strategy X as the first element in this list, strategy X will be used. - If a candidate has [X] and another candidate has [Y, X], Y supersedes X and strategy Y\n  will be used.\n- If a candidate has [X, Y] and another candidate has [Y, X], this is a user error and leader\n  election will not operate the Lease until resolved.\n- In general: [A1, A2, ..., An] > [B1, B2, ..., Bm] if the latter is a sub-sequence of the former, and hence\n  A1 is chosen. For more than two candidates, one must be the maximum of all candidates. Otherwise, this is a user\n  error and leader election will not operate the Lease until resolved.\n(Alpha) Using this field requires the CoordinatedLeaderElection feature gate to be enabled.",
}

func (LeaseCandidateSpec) SwaggerDoc() map[string]string {
	return map_LeaseCandidateSpec
}

// AUTO-GENERATED FUNCTIONS END HERE