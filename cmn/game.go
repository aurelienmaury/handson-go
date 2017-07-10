package Cmn

type State int

const (
	WAIT State = 1 + iota
	RUN
	LOST
	WIN
)

var states = [...]string{
	"WAIT",
	"RUN",
	"LOST",
	"WIN",
}

func (state State) String() string {
	return states[state - 1]
}

type GameState struct {
	Partie State `json:"partie"`
	Turn   string `json:"turn"`
	Grid  [3]map[string]string `json:"grid_lines"`
}


/*

{
  “partie”: WAIT | RUN | LOST | WIN,
  “turn”: None | {joueur-uuid}
  “grid_lines”: [
     { “0”: - | (joueur-uuid), “1”: - | (joueur-uuid), “2”: - | (joueur-uuid) },
     { “0”: - | (joueur-uuid), “1”: - | (joueur-uuid), “2”: - | (joueur-uuid) },
     { “0”: - | (joueur-uuid), “1”: - | (joueur-uuid), “2”: - | (joueur-uuid) },
   ]
}
 */