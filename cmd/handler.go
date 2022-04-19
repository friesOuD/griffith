package cmd

import (
	"bufio"
	"fmt"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"os"
)

type docsForm struct {
	id    string
	name  string
	text  string
	date  string
	owner string
}

func InputDoc() {

	var inc docsForm
	fmt.Println("Введите ID:")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	inc.id = myscanner.Text()
	fmt.Println("Введите название инцидента:")
	myscanner.Scan()
	inc.name = myscanner.Text()
	fmt.Println("Введите решение инцидента:")
	myscanner.Scan()
	inc.text = myscanner.Text()
	fmt.Println("Введите даты открытия/зарытия инцидента:")
	myscanner.Scan()
	inc.date = myscanner.Text()
	fmt.Println("Введите инициатора инцидента:")
	myscanner.Scan()
	inc.owner = myscanner.Text()
	fmt.Println(inc)

}

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey("cd91451c5faf3424e59ed9df3cfec9678ab158c0fb01d654d2531e6bc8f91907")
	if err != nil {
		panic(err)
	}
}

func Test() {

	doc, _ := document.OpenTemplate("template.docx")
	defer doc.Close()

	doc.SaveToFile("simple.docx")
}

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}
