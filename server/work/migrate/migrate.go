package migrate

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jakubc-projects/ustron-work/server/work"
	"github.com/jakubc-projects/ustron-work/server/work/date"
	"github.com/samber/lo"
)

const timestampLayout = "1/2/2006 15:04:05"
const dateLayout = "1/2/2006"

var accentReplacer = strings.NewReplacer(
	"ż", "z",
	"ł", "l",
	"Ł", "L",
	"ę", "e",
	"ś", "s",
	"ń", "n",
)

func ParseRegistrations(filepath string, personMap map[string]int) []work.Registration {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)

	r.Read()

	var regs []work.Registration

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var reg work.Registration

		reg.Uid = uuid.New()

		reg.CreatedAt = lo.Must(time.Parse(timestampLayout, record[0]))

		name := strings.Trim(record[1], " ")

		name = accentReplacer.Replace(name)

		team := mapTeam(record[3])

		reg.Team = team

		key := mapKey(name, team)
		personId := personMap[key]
		if personId == 0 {
			panic("cannot find person: " + key)
		}

		reg.PersonID = personId
		reg.Type = mapType(record[4])

		if reg.Type == work.RegistrationTypeMoney {
			reg.Date = date.DateOf(lo.Must(time.Parse(dateLayout, record[5])))
			reg.PaidSum = float32(lo.Must(strconv.ParseFloat(record[6], 32)))
			reg.Goal = work.RegistrationGoal(record[7])
			reg.Description = record[8]
		} else {
			reg.Date = date.DateOf(lo.Must(time.Parse(dateLayout, record[9])))
			reg.Hours = float32(lo.Must(strconv.ParseFloat(record[10], 32)))
			reg.HourlyWage = lo.Must(strconv.Atoi(record[11]))
			reg.Goal = work.RegistrationGoal(record[12])
			reg.Description = record[13]
		}

		regs = append(regs, reg)
	}

	return regs
}

func RegsToSql(regs []work.Registration) string {
	sql := "INSERT INTO registrations (uid, person_id, team, date, type, hourly_wage, hours, paid_sum, goal, description, created_at) VALUES "

	regValueSql := []string{}

	for _, reg := range regs {
		s := fmt.Sprintf("('%s', %d, '%s', '%s', '%s', %d, %f, %f,'%s','%s', '%s')", reg.Uid, reg.PersonID, reg.Team, reg.Date, reg.Type, reg.HourlyWage, reg.Hours, reg.PaidSum, reg.Goal, reg.Description, reg.CreatedAt.Format(time.RFC3339))
		regValueSql = append(regValueSql, s)
	}

	sql += strings.Join(regValueSql, ", ")
	return sql
}

type person struct {
	PersonId    int       `json:"person_id"`
	DisplayName string    `json:"display_name"`
	Team        work.Team `json:"team"`
}

func NamePersonIdMap(filepath string) map[string]int {

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var persons []person
	err = json.NewDecoder(file).Decode(&persons)
	if err != nil {
		panic(err)
	}

	result := map[string]int{}

	for _, p := range persons {
		p.DisplayName = accentReplacer.Replace(p.DisplayName)

		key := mapKey(p.DisplayName, p.Team)
		if result[key] != 0 {
			fmt.Println("Duplicate", p.DisplayName)
		}
		result[key] = p.PersonId
	}

	return result
}

func mapKey(name string, team work.Team) string {
	return name + " | " + string(team)
}

func mapTeam(dataTeam string) work.Team {
	switch dataTeam {
	case "🔵 BLUE":
		return work.TeamBlue

	case "🟠 ORANGE":
		return work.TeamOrange

	case "🔴 RED":
		return work.TeamRed

	case "🟢 GREEN":
		return work.TeamGreen
	default:
		panic("invalid team: " + dataTeam)
	}
}

func mapType(t string) work.RegistrationType {
	switch t {
	case "Wpłatę":
		return work.RegistrationTypeMoney

	case "Godziny pracy":
		return work.RegistrationTypeWork
	default:
		panic("invalid registration type: " + t)
	}
}
