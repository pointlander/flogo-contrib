package log

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/ext/activity"
	"github.com/TIBCOSoftware/flogo-lib/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("log")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &LogActivity{metadata: md}

	tc := test.NewTestActivityContext()
	//setup attrs

	act.Eval(tc)

	//check result attr
}