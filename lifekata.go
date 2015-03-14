package lifekata

type State int
type amount int

const (
	Dead  State = 0
	Alive State = 1
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return "{" + string(c.x) + "," + string(c.y) + "}"
}

type Game struct {
	field      map[coord]State
	born, stay []amount
}

func (g Game) SetCell(c coord, s State) {
	if s == Dead {
		return
	}
	g.field[c] = s
}
func (g Game) GetCell(c coord) State {
	val, _ := g.field[c]
	return val
}

func NewGame() Game {
	g := Game{}
	g.field = make(map[coord]State)
	g.born = []amount{3}
	g.stay = []amount{2, 3}
	return g
}

func (g Game) CountAliveNeighbours(c coord) amount {
	x, y := c.x, c.y
	return g.CountAliveCell(coord{x - 1, y - 1}) + g.CountAliveCell(coord{x - 1, y}) + g.CountAliveCell(coord{x - 1, y + 1}) +
		g.CountAliveCell(coord{x, y - 1}) + g.CountAliveCell(coord{x, y + 1}) +
		g.CountAliveCell(coord{x + 1, y - 1}) + g.CountAliveCell(coord{x + 1, y}) + g.CountAliveCell(coord{x + 1, y + 1})
}

func (g Game) CountAliveCell(c coord) amount {
	if g.GetCell(c) == Alive {
		return 1
	}
	return 0
}

func (g Game) NewGeneration(c coord) State {
	am := g.CountAliveNeighbours(c)
	switch {
	case inAmountList(g.born, am):
		return Alive
	case inAmountList(g.stay, am):
		return g.GetCell(c)
	default:
		return Dead
	}
}

func (g Game) CoordsToGenerate() []coord {
	coords := make([]coord, 0)
	for c, _ := range g.field {
		coords = append(coords, coord{c.x - 1, c.y - 1})
		coords = append(coords, coord{c.x - 1, c.y})
		coords = append(coords, coord{c.x - 1, c.y + 1})
		coords = append(coords, coord{c.x, c.y - 1})
		coords = append(coords, coord{c.x, c.y})
		coords = append(coords, coord{c.x, c.y + 1})
		coords = append(coords, coord{c.x + 1, c.y - 1})
		coords = append(coords, coord{c.x + 1, c.y})
		coords = append(coords, coord{c.x + 1, c.y + 1})
	}
	return coords
}

func (g *Game) GenerateNextGeneration() {
	coords := g.CoordsToGenerate()
	newgame := NewGame()
	for _, c := range coords {
		newgame.SetCell(c, g.NewGeneration(c))
	}
	*g = newgame
}

func NewRulesGame(born, stay []amount) Game {
	g := NewGame()
	g.born = born
	g.stay = stay
	return g
}

func inAmountList(l []amount, value amount) bool {
	for idx := range l {
		if l[idx] == value {
			return true
		}
	}
	return false
}
