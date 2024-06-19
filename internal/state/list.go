package state

type State struct {
	Name string
}

func (s *State) String() string {
	return s.Name
}