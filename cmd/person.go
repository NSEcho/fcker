package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/lateralusd/fcker/models/fakename"
	"github.com/spf13/cobra"
)

const (
	fakeNameURL = "https://www.fakenamegenerator.com/"
)

var regexes = map[string]*regexp.Regexp{
	"name":     regexp.MustCompile(`class="address">\n.*<h3>(.*)</h3>`),
	"addr":     regexp.MustCompile(`class="adr">\n(.*)<br.*/>(.*)</div>`),
	"ssn":      regexp.MustCompile(`<dt>SSN</dt><dd>(.*).<div`),
	"phone":    regexp.MustCompile(`<dt>Phone</dt>\n.*<dd>(.*)</dd>`),
	"birthday": regexp.MustCompile(`<dt>Birthday</dt>\n.*<dd>(.*)</dd>`),
	"email":    regexp.MustCompile(`<dt>Email Address</dt>\n\n.*<dd>(.*)<div`),
	"username": regexp.MustCompile(`<dt>Username</dt>\n.*<dd>(.*)</dd>`),
	"password": regexp.MustCompile(`<dt>Password</dt>\n.*<dd>(.*)</dd>`),
	"height":   regexp.MustCompile(`<dt>Height</dt>\n.*<dd>(.*)</dd>`),
	"weight":   regexp.MustCompile(`<dt>Weight</dt>\n.*<dd>(.*)</dd>`),
}

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Fetch person from fakenamegenerator.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		us, err := cmd.Flags().GetString("us")
		if err != nil {
			return err
		}

		co, err := cmd.Flags().GetString("co")
		if err != nil {
			return err
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		if err := checkCodes(us, co); err != nil {
			return err
		}

		results := make(map[string]string)

		fetchURL := fmt.Sprintf("%sgen-random-%s-%s.php", fakeNameURL, us, co)

		client := &http.Client{
			Timeout: 3 * time.Second,
		}

		req, err := http.NewRequest("GET", fetchURL, nil)
		if err != nil {
			return err
		}

		req.Header.Add("User-Agent", "curl/7.64.1")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		for key, re := range regexes {
			match := re.FindStringSubmatch(string(body))
			if len(match) > 0 {
				result := match[1]
				if key == "addr" {
					result = strings.TrimLeft(fmt.Sprintf("%s, %s", match[1], match[2]), " ")
				}
				results[key] = result
			}
		}

		user := fakename.Person{
			Name:     results["name"],
			Address:  results["addr"],
			SSN:      results["ssn"],
			Phone:    results["phone"],
			Birthday: results["birthday"],
			Email:    results["email"],
			Username: results["username"],
			Password: results["password"],
			Height:   results["height"],
			Weight:   results["weight"],
		}

		if file != "" {
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			defer f.Close()

			wr := io.MultiWriter(os.Stdout, f)
			user.Write(wr)
		} else {
			user.Write(os.Stdout)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(personCmd)
	personCmd.Flags().StringP("us", "u", "us", "user code")
	personCmd.Flags().StringP("co", "c", "us", "country code")
	personCmd.Flags().StringP("file", "f", "", "save output to file specified")
}
