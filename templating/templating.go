package templating

import (
	"fmt"
	"os"
	"text/template"
)

func GenTemplateFiles(paths []string, data map[string]string) error {
	for _, p := range paths {
		t := template.Must(template.ParseFiles(p))
		f, err := os.Create(fmt.Sprintf("%s", p))
		if err != nil {
			return err
		}
		err = t.Execute(f, data)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
