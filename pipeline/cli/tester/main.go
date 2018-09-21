package main

import (
	"os"

	cmd "github.com/naveego/api/pipeline/cli/pub"
	_ "github.com/naveego/pipeline-publishers"
	_ "github.com/naveego/pipeline-subscribers"
	"github.com/sirupsen/logrus"
)

func main() {

	cmd.TypeName = os.Args[1]
	cmd.RootCmd.SetArgs([]string{"publish", "pub-1", "-a=http://localhost:8084", "-t=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ0ZXN0IiwiZXhwIjoxNDgxODkwMTI3LCJpc3MiOiJuYXZlZWdvIiwibmFtZSI6InB1Yi0xIiwibmJmIjoxNDUxNjQ5NjAwLCJyb2xlcyI6WyJwdWJsaXNoZXIiXSwic3ViIjoicHVibGlzaGVyL3B1Yi0xIn0.PueQqWzstI9VyyLHK62iW9aw2RbQmbSHKtI9VWHsRKZwUhf5cwwka5Q4Umrr3U6K58EaDkaaTijUbzPX9aPLOjc8QnlhzUwiYcGEeVvDXCJ71BuPJztO8Z0nd2mKuR9IqKMVV57PAMIph01cSClIWdqVxSPTLYAwBjkmCAqpwtg", "-v"})

	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
