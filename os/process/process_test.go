package gxprocess

import (
	"os"
	"testing"
)

import (
	"github.com/alexstocks/goext/log"
)

func TestFindProcess(t *testing.T) {
	p, err := FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if p == nil {
		t.Fatal("should have process")
	}
	t.Logf("current pid:%d, linux process:%#v", os.Getpid(), p)

	if p.Pid() != os.Getpid() {
		t.Fatalf("bad: %#v", p.Pid())
	}
}

func TestProcesses(t *testing.T) {
	// This test works because there will always be SOME processes
	// running.
	p, err := Processes()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	found := false
	for _, p1 := range p {
		if p1.Executable() == "go" || p1.Executable() == "go.exe" {
			found = true
			break
		}
	}
	t.Logf("/proc processes:%#v", gxlog.PrettyString(p))

	if !found {
		t.Fatal("should have Go")
	}
}
