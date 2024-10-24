package synchronization

import (
	"testing"
)

// TestStatusUnmarshal tests that unmarshaling from a string specification
// succeeeds for Status.
func TestStatusUnmarshal(t *testing.T) {
	// Set up test cases.
	testCases := []struct {
		text          string
		expected      Status
		expectFailure bool
	}{
		{"", Status_Disconnected, true},
		{"asdf", Status_Disconnected, true},
		{"disconnected", Status_Disconnected, false},
		{"halted-on-root-emptied", Status_HaltedOnRootEmptied, false},
		{"halted-on-root-deletion", Status_HaltedOnRootDeletion, false},
		{"halted-on-root-type-change", Status_HaltedOnRootTypeChange, false},
		{"connecting-alpha", Status_ConnectingAlpha, false},
		{"connecting-beta", Status_ConnectingBeta, false},
		{"watching", Status_Watching, false},
		{"scanning", Status_Scanning, false},
		{"waiting-for-rescan", Status_WaitingForRescan, false},
		{"reconciling", Status_Reconciling, false},
		{"staging-alpha", Status_StagingAlpha, false},
		{"staging-beta", Status_StagingBeta, false},
		{"transitioning", Status_Transitioning, false},
		{"saving", Status_Saving, false},
	}

	// Process test cases.
	for _, testCase := range testCases {
		var status Status
		if err := status.UnmarshalText([]byte(testCase.text)); err != nil {
			if !testCase.expectFailure {
				t.Errorf("unable to unmarshal text (%s): %s", testCase.text, err)
			}
		} else if testCase.expectFailure {
			t.Error("unmarshaling succeeded unexpectedly for text:", testCase.text)
		} else if status != testCase.expected {
			t.Errorf(
				"unmarshaled status (%s) does not match expected (%s)",
				status,
				testCase.expected,
			)
		}
	}
}
