package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ErrUserNotFound    = errors.New("codes: user code not found")
	ErrCountryNotFound = errors.New("codes: country code not found")
)

var userCodes = map[string]string{
	"ar":     "Arabic",
	"au":     "Australian",
	"br":     "Brazil",
	"celat":  "Chechen (Latin)",
	"ch":     "Chinese",
	"zhtw":   "Chinese (Traditional)",
	"hr":     "Croatian",
	"cs":     "Czech",
	"dk":     "Danish",
	"nl":     "Dutch",
	"en":     "England/Wales",
	"er":     "Eritrean",
	"fi":     "Finnish",
	"fr":     "French",
	"gr":     "German",
	"gl":     "Greenland",
	"sp":     "Hispanic",
	"hobbit": "Hobbit",
	"hu":     "Hungarian",
	"is":     "Icelandic",
	"ig":     "Igbo",
	"it":     "Italian",
	"jpja":   "Japanese",
	"jp":     "Japanese (Anglicized)",
	"tlh":    "Klingon",
	"ninja":  "Ninja",
	"no":     "Norwegian",
	"fa":     "Persian",
	"pl":     "Polish",
	"ru":     "Russian",
	"rucyr":  "Russian (Cyrillic)",
	"gd":     "Scottish",
	"sl":     "Slovenian",
	"sw":     "Swedish",
	"th":     "Thai",
	"us":     "United States",
	"vn":     "Vietnamese",
}

var countryCodes = map[string]string{
	"au":   "Australia",
	"as":   "Austria",
	"bg":   "Belgium",
	"br":   "Brazil",
	"ca":   "Canada",
	"cyen": "Cyprus (Anglicized)",
	"cygk": "Cyprus (Greek)",
	"cz":   "Czech Republic",
	"dk":   "Denmark",
	"ee":   "Estonia",
	"fi":   "Finland",
	"fr":   "France",
	"gr":   "Germany",
	"gl":   "Greenland",
	"hu":   "Hungary",
	"is":   "Iceland",
	"it":   "Italy",
	"nl":   "Netherlands",
	"nz":   "New Zealand",
	"no":   "Norway",
	"pl":   "Poland",
	"pt":   "Portugal",
	"sl":   "Slovenia",
	"za":   "South Africa",
	"sp":   "Spain",
	"sw":   "Sweden",
	"sz":   "Switzerland",
	"tn":   "Tunisia",
	"uk":   "United Kingdom",
	"us":   "United States",
	"uy":   "Uruguay",
}

var codesCmd = &cobra.Command{
	Use:   "codes",
	Short: "List user and country codes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("\tUSER SET\n")
			printCodes(userCodes)
			fmt.Println("\tCOUNTRY CODES")
			printCodes(countryCodes)
			return
		}

		switch args[0] {
		case "u":
			fmt.Printf("\tUSER SET\n")
			printCodes(userCodes)
		case "c":
			fmt.Println("\tCOUNTRY CODES")
			printCodes(countryCodes)
		default:
			fmt.Fprintf(os.Stderr, "Unknown code, possible values are 'u' and 'c'!\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(codesCmd)
}

func printCodes(set map[string]string) {
	fmt.Printf("%-10s | %-10s\n", "Code", "Country")
	fmt.Println("============================")
	for k, v := range set {
		fmt.Printf("%-10s | %-10s\n", k, v)
	}
	fmt.Printf("\n\n")
}

func checkCodes(us, co string) error {
	if _, ok := userCodes[us]; !ok {
		return ErrUserNotFound
	}
	if _, ok := countryCodes[co]; !ok {
		return ErrCountryNotFound
	}
	return nil
}
