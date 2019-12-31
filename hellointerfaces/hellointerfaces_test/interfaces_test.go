package hellointerfaces_test

import (
	"testing"

	"github.com/EmilianoFraga/helloGo/hellointerfaces"
)

func assertGeometryArea(t *testing.T, g hellointerfaces.Geometry, expectedArea float64) {
	if actualArea := g.Area(); actualArea != expectedArea {
		t.Errorf("Wrong geometry area %f, expected %f", actualArea, expectedArea)
	}
}

func assertGeometryPerim(t *testing.T, g hellointerfaces.Geometry, expectedPerim float64) {
	if actualPerim := g.Perim(); actualPerim != expectedPerim {
		t.Errorf("Wrong geometry perimeter %f, expected %f", actualPerim, expectedPerim)
	}
}
