package migrate

import (
	"fmt"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {
	personMap := NamePersonIdMap("../../../persons.json")
	regs := ParseRegistrations("../../../registrations.csv", personMap)

	// fmt.Println(regs)

	sql := RegsToSql(regs)

	f, _ := os.Create("res.sql")

	f.WriteString(sql)

}

func TestLoadPersonMap(t *testing.T) {
	personMap := NamePersonIdMap("../../../persons.json")

	fmt.Println(personMap)
}
