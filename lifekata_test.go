package lifekata

import "testing"
import "reflect"

func TestGame(t *testing.T) {
	_ = NewGame()
}

func TestCoord(t *testing.T) {
	_ = coord{1, 1}
}

func TestSetCell(t *testing.T) {
	g := NewGame()
	g.SetCell(coord{1, 1}, Alive)
	if g.GetCell(coord{1, 1}) != Alive {
		t.Errorf("Expected coord(1,1) to be Alive")
	}
}

func TestDeadCell(t *testing.T) {
	g := NewGame()
	if g.GetCell(coord{1, 1}) != Dead {
		t.Errorf("Expected unset cell to be Dead")
	}
}

func TestCountAliveNeighbours(t *testing.T) {
	g := NewGame()
	if g.CountAliveNeighbours(coord{1, 1}) != 0 {
		t.Errorf("Expected no neighbours to be alive")
	}

	g.SetCell(coord{1, 1}, Alive)
	g.SetCell(coord{1, 2}, Alive)
	if g.CountAliveNeighbours(coord{2, 1}) != 2 {
		t.Errorf("Expected 2 neighbours to be alive")
	}
	if g.CountAliveNeighbours(coord{1, 1}) != 1 {
		t.Errorf("Expected 1 neighbour to be alive")
	}
}

func TestNewGeneration(t *testing.T) {
	g := NewGame()

	g.SetCell(coord{1, 1}, Alive)
	g.SetCell(coord{1, 2}, Alive)
	g.SetCell(coord{1, 3}, Alive)
	if g.NewGeneration(coord{1, 1}) != Dead {
		t.Errorf("Expected coord{1,1} to die")
	}
	if g.NewGeneration(coord{1, 2}) != Alive {
		t.Errorf("Expected coord{1,2} to remain alive")
	}
	if g.NewGeneration(coord{2, 1}) != Dead {
		t.Errorf("Expected coord{2,1} to remain dead")
	}
	if g.NewGeneration(coord{2, 2}) != Alive {
		t.Errorf("Expected coord{2,2} to live")
	}
}

func TestOvercrowded(t *testing.T) {
	g := NewGame()

	g.SetCell(coord{1, 1}, Alive)
	g.SetCell(coord{1, 2}, Alive)
	g.SetCell(coord{1, 3}, Alive)
	g.SetCell(coord{0, 2}, Alive)
	g.SetCell(coord{2, 2}, Alive)

	if g.NewGeneration(coord{1, 2}) != Dead {
		t.Errorf("Expected coord{1,2} to starve")
	}
}

func TestNoCoordsToGenerate(t *testing.T) {
	g := NewGame()

	if !reflect.DeepEqual(g.CoordsToGenerate(), []coord{}) {
		t.Errorf("Expected no coords to generate")
	}
}

func TestCoordsToGenerate(t *testing.T) {
	g := NewGame()
	g.SetCell(coord{1, 1}, Alive)

	crds := g.CoordsToGenerate()
	if !reflect.DeepEqual(crds, []coord{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}) {
		t.Errorf("Expected coords to generate, got %#v", crds)
	}
}

func TestDeadCoordsToGenerate(t *testing.T) {
	g := NewGame()
	g.SetCell(coord{1, 1}, Dead)
	crds := g.CoordsToGenerate()

	if !reflect.DeepEqual(crds, []coord{}) {
		t.Errorf("Expected no coords to generate, got %#v", crds)
	}
}

func TestCompleteGeneration(t *testing.T) {
	g := NewGame()
	g.SetCell(coord{1, 1}, Alive)
	g.SetCell(coord{1, 2}, Alive)
	g.SetCell(coord{1, 3}, Alive)

	g.GenerateNextGeneration()
	if g.GetCell(coord{1, 1}) != Dead {
		t.Errorf("Expected coord{1,1} to die")
	}
	if g.GetCell(coord{1, 2}) != Alive {
		t.Errorf("Expected coord{1,2} to remain alive")
	}
	if g.GetCell(coord{2, 1}) != Dead {
		t.Errorf("Expected coord{2,1} to remain dead")
	}
	if g.GetCell(coord{2, 2}) != Alive {
		t.Errorf("Expected coord{2,2} to live")
	}
}

func TestRules(t *testing.T) {
	g := NewRulesGame([]amount{1}, []amount{})
	g.SetCell(coord{0, 0}, Alive)

	if g.NewGeneration(coord{-1, -1}) != Alive {
		t.Errorf("Expected coord{-1,-1} to be alive")
	}
}
