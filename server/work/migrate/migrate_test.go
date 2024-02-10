package migrate

import (
	"fmt"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	personMap := NamePersonIdMap("data/persons.json")
	regs := ParseRegistrations("data/registrations2.csv", personMap)

	// fmt.Println(regs)
	sql := RegsToSql(regs)
	f, _ := os.Create("data/res.sql")

	f.WriteString(sql)

}

func TestLoadPersonMap(t *testing.T) {
	personMap := NamePersonIdMap("data/persons.json")

	fmt.Println(personMap)
}
