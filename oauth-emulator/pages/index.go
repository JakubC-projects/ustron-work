package pages

import (
	"io"

	"github.com/jakubc-projects/ustron-work/oauth-emulator/users"
)

type Login struct {
	AvailableUsers []users.User
}

func RenderLogin(w io.Writer, d Login) error {
	return templates.ExecuteTemplate(w, "login.html", d)
}
