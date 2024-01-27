package work

type Team string

const (
	TeamGreen  Team = "Green"
	TeamRed    Team = "Red"
	TeamBlue   Team = "Blue"
	TeamOrange Team = "Orange"
)

var Teams = []Team{TeamGreen, TeamRed, TeamBlue, TeamOrange}

type Status map[Team]float32

func NewStatus() Status {
	s := Status{}

	for _, g := range Teams {
		s[g] = 0
	}

	return s
}

type Gender string

const (
	GenderMale   = "Male"
	GenderFemale = "Female"
)

var Genders = []Gender{GenderMale, GenderFemale}

type GenderStatus map[Gender]float32

func NewGenderStatus() GenderStatus {
	s := GenderStatus{}

	for _, g := range Genders {
		s[g] = 0
	}

	return s
}
