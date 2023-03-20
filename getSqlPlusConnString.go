package main

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

var sqlConn = "database.username/database.password@database.host:database.port/database.service"

func main() {

	input := os.Args[1]
	if input != "" {
		sDec, _ := b64.StdEncoding.DecodeString(input)
		sqlString := generateSqlSyntax(string(sDec))
		fmt.Println(sqlString)
	} else {
		fmt.Println("Use Kubectl get secret , to get the base64 encoded secret ")
		fmt.Println("Pass it as param to this function")
	}
}

/*
*
Input will be property file with

database.host=oracle-gsd.com
database.port=1521
database.service=server1
database.username=userName
database.password=superPassword

*
*/
func generateSqlSyntax(input string) string {
	p := properties.MustLoadString(input)
	host := p.MustGetString("database.host")
	port := p.MustGetString("database.port")
	service := p.MustGetString("database.service")
	username := p.MustGetString("database.username")
	password := p.MustGetString("database.password")
	sqlString := username + "/" + password + "@" + host + ":" + port + "/" + service
	return sqlString
}
