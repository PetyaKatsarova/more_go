package main

import "fmt"

/*
PS C:\Users\petya.katsarova\OneDrive - CGI\Desktop\github_folder\pr8_justforfunc\37consts> go build
PS C:\Users\petya.katsarova\OneDrive - CGI\Desktop\github_folder\pr8_justforfunc\37consts> .\37consts.exe
go build -tags prod saying this is production

--------------
go build
./36consts
output: developement
go build -tags prod
./36consts
output: production
*/

func main() {
	fmt.Printf("running version %s\n", version)
}