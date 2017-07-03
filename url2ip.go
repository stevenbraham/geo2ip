package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	url := os.Args[1]
	domain := getDomainFromUrl(url)

	//execute ping request
	out, _ := exec.Command("ping", "-n", "-c1", domain).Output()

	//parse ping output
	regex := regexp.MustCompile(`(\d{2,3}\.\d{2,3}\.\d{2,3}\.\d{2,3})`)

	fmt.Println(regex.FindString((string(out))))
}

func getDomainFromUrl(url string) string {
	regex := regexp.MustCompile(`\:\/\/(.+\.\w+)`)
	domain := regex.FindString(url)
	domain = strings.Replace(domain, "://", "", 2)
	return domain
}
