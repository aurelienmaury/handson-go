package main


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

func (state State) string() string {
	return states[state-1]
}


type Game struct {

}
