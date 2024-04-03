package skeletons

import (
	"testing"
)

func TestNew(t *testing.T) {
	skeletons := New()

	if skeletons.V1.API != V1API {
		t.Errorf("expected field to be %v, but got %v", V1API, skeletons.V1.API)
	}
	if skeletons.V1.CLI != V1CLI {
		t.Errorf("expected field to be %v, but got %v", V1CLI, skeletons.V1.CLI)
	}
	if skeletons.V1.GRPC != V1GRPC {
		t.Errorf("expected field to be %v, but got %v", V1GRPC, skeletons.V1.GRPC)
	}
	if skeletons.V1.ProjectFiles != V1ProjectFiles {
		t.Errorf("expected field to be %v, but got %v", V1ProjectFiles, skeletons.V1.ProjectFiles)
	}
}
