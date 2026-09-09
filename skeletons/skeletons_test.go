package skeletons

import (
	"testing"
)

func TestNew(t *testing.T) {
	skeletons := New()

	if skeletons.V1.GoAPI != V1GoAPI {
		t.Errorf("expected field to be %v, but got %v", V1GoAPI, skeletons.V1.GoAPI)
	}
	if skeletons.V1.GoEmbed != V1GoEmbed {
		t.Errorf("expected field to be %v, but got %v", V1GoEmbed, skeletons.V1.GoEmbed)
	}
	if skeletons.V1.GoCLI != V1GoCLI {
		t.Errorf("expected field to be %v, but got %v", V1GoCLI, skeletons.V1.GoCLI)
	}
	if skeletons.V1.GoGRPC != V1GoGRPC {
		t.Errorf("expected field to be %v, but got %v", V1GoGRPC, skeletons.V1.GoGRPC)
	}
	if skeletons.V1.NodeJS != V1NodeJS {
		t.Errorf("expected field to be %v, but got %v", V1NodeJS, skeletons.V1.NodeJS)
	}
	if skeletons.V1.Python != V1Python {
		t.Errorf("expected field to be %v, but got %v", V1Python, skeletons.V1.Python)
	}
	if skeletons.V1.ProjectFiles != V1ProjectFiles {
		t.Errorf("expected field to be %v, but got %v", V1ProjectFiles, skeletons.V1.ProjectFiles)
	}
}
