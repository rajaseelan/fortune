package pkg

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// LoadFortunes reads fortunes from the supplied file name
// the file name is the set name
func LoadFortunes(fileName string) *Fortune {

	fortune := Fortune{
		Cookies: []*string{},
		SetName: &fileName,
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Unable to locate fortune file - %s - %s", fileName, err)
	}
	defer file.Close()

	cookie := ""

	delimiterRegex := regexp.MustCompile(`^%$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// do stuff

		line := scanner.Text()

		if !delimiterRegex.MatchString(line) {
			cookie += line + "\n"
		}

		if delimiterRegex.MatchString(line) {

			cookieText := cookie
			// cookies := fortune.Cookies
			fortune.Cookies = append(fortune.Cookies, &cookieText)
			// fortune.Cookies = cookies
			cookie = ""

			continue
		}
	}

	// add remaining cookie
	fortune.Cookies = append(fortune.Cookies, &cookie)

	return &fortune
}
