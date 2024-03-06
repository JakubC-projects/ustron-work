package work

const (
	TeamGreen  = "Green"
	TeamRed    = "Red"
	TeamBlue   = "Blue"
	TeamOrange = "Orange"
)

var Teams = []string{TeamGreen, TeamRed, TeamBlue, TeamOrange}

type Status map[string]float32

func NewStatus() Status {
	s := Status{}

	for _, g := range Teams {
		s[g] = 0
	}
	return s
}
