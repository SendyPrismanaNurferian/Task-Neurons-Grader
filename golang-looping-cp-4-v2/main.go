package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	// Split email by "@"
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "Invalid email format"
	}

	// Split domain by "."
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "Invalid email format"
	}

	// Get domain and TLD
	domain := domainParts[0]
	tld := strings.Join(domainParts[1:], ".")

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
