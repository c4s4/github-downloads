package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Version as printed with -version option
var Version = "UNKNOWN"

const (
	// Help as printed with -help option
	Help = `github-downloads [-h] [-v] user repo
Github downloads prints downloads information for given Github user repo:
-help     To print this help
-version  To print version
user      The Github user name
repo      The Github repository`
)

// Asset is the structure for an asset
type Asset struct {
	Name      string `json:"name"`
	Downloads int    `json:"download_count"`
}

// Release is the structure for a release
type Release struct {
	Tag    string  `json:"tag_name"`
	Date   string  `json:"published_at"`
	Assets []Asset `json:"assets"`
}

// parseCommandLine parses command line and returns:
// - help: a boolean that tells if we print help
// - version: a boolean that tells if we print version
// - user: Github user name
// - repo: Github repository
func parseCommandLine() (bool, bool, string, string) {
	help := flag.Bool("help", false, "Print help")
	version := flag.Bool("version", false, "Print version")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		println("ERROR: You must pass user and repo on command line")
		println(Help)
		os.Exit(1)
	}
	user := args[0]
	repo := args[1]
	return *help, *version, user, repo
}

// Releases call Github API to get release information
func Releases(user, repo string) ([]Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", user, repo)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var releases []Release
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

// reverse a slice
func reverse[S []E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// PrintReleases prints releases on terminal
func PrintReleases(releases []Release) {
	reverse(releases)
	total := 0
	for _, release := range releases {
		subtotal := 0
		for _, asset := range release.Assets {
			subtotal += asset.Downloads
		}
		total += subtotal
		fmt.Printf("Release %s: %d\n", release.Tag, subtotal)
		if len(release.Assets) > 1 {
			for _, asset := range release.Assets {
				if asset.Downloads > 0 {
					fmt.Printf("- %s: %d\n", asset.Name, asset.Downloads)
				}
			}
		}
	}
	fmt.Printf("Total: %d\n", total)
}

func main() {
	help, version, user, repo := parseCommandLine()
	if help {
		fmt.Println(Help)
		os.Exit(0)
	}
	if version {
		fmt.Println(Version)
		os.Exit(0)
	}
	releases, err := Releases(user, repo)
	if err != nil {
		println(fmt.Sprintf("Error getting release information: %v", err))
	}
	PrintReleases(releases)
}
