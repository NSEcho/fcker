package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	thisPersonURL = "https://thispersondoesnotexist.com/image"
)

var person = &cobra.Command{
	Use:   "person",
	Short: "Fetch image from thispersondoesnotexist.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get(thisPersonURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		f, err := os.Create(name)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, resp.Body); err != nil {
			return err
		}

		fmt.Printf("[*] Saved %s\n", name)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(person)
	person.Flags().StringP("name", "n", "image.jpeg", "name of the fetched image")
}
