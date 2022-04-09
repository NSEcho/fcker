package fakename

import (
	"io"
	"text/template"
)

// Person represents person fetched from the fakenamegenerator.com
type Person struct {
	Name     string
	Address  string
	SSN      string
	Phone    string
	Birthday string
	Email    string
	Username string
	Password string
	Height   string
	Weight   string
}

var tpl = `Name: {{ .Name }}
Address: {{ .Address }}

SSN: {{ .SSN }}

Phone: {{ .Phone }}
Birthday: {{ .Birthday }}
Email: {{ .Email }}
Username: {{ .Username }}
Password: {{ .Password }}

Height: {{ .Height }}
Weight: {{ .Weight }}
`

// Write executes the template on the writer specified by the user
func (p *Person) Write(w io.Writer) error {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return err
	}
	return t.Execute(w, p)
}
