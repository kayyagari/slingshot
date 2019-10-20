package main

import (
	"slingshot/ui"
)

func main() {
	//dirPath, err := jnlp.ParseAndDownload("http://localhost:8080/webstart.jnlp")
	//if err != nil {
	//	os.Exit(1)
	//}
	//fmt.Println("saved jar files to ", dirPath)
	cd := ui.NewConnectionDialog()
	cd.Show()
}
