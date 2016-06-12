package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// A struct to hold the OPML file info
type OPML struct {
	Player   string    `xml:"head>title"`
	Podcasts []Outline `xml:"body>outline>outline"`
}

// String() is called everytime we use fmt.Print().
// We want to print the info in a meaningful way.
func (this OPML) String() string {
	return this.Player + "\n\n" + fmt.Sprintln(this.Podcasts)
}

// This holds one outline field from the OPML file
type Outline struct {
	Title string `xml:"text,attr"`
	Type  string `xml:"type,attr"`
	Url   string `xml:"xmlUrl,attr"`
}

// Print an Outline variable in a meaningful way.
func (this Outline) String() string {
	return fmt.Sprintf("Title: %v\nURL: %v", this.Title, this.Url)
}

// This holds information about a podcast that we get from its RSS feed.
type Podcast struct {
	Title       string   `xml:"channel>title"`
	Url         string   `xml:"channel>link"`
	Description string   `xml:"channel>description"`
	Copyright   string   `xml:"channel>copyright"`
	Author      string   `xml:"channel>author"`
	Categroy    Category `xml:"channel>category"`
	Image       Image    `xml:"channel>image"`
}

// Print a Podcast variable in a meaningful way.
func (this Podcast) String() string {
	return fmt.Sprintf("Title: %v\nAuthor: %v\nDescription: %v\nCopyright: %v", this.Title, this.Author, this.Description, this.Copyright)
}

// We need this to get the category attribute from the RSS feed.
type Category struct {
	Name string `xml:"text,attr"`
}

// We need this to get the image URL attribute from the RSS feed.
type Image struct {
	URL string `xml:"href,attr"`
}

func new_user() string {
	fmt.Println("What is your Name?")
	name_scanner := bufio.NewScanner(os.Stdin)
	name_scanner.Scan()
	name := name_scanner.Text()
	return name
}

func main() {

	//Lets greet our new friends
	fmt.Println("Welcome to ListenTo, the tool that helps you share your podcast information!")
	name := "Jay"
	if name == "" {
		name = new_user()
	}
	fmt.Printf("Hello %v\n", name)

	//We need some information from them... Like their OPML 💩
	fmt.Println("We need to get your OPML file where is it?")

	// Parse the entire feed into an OPML
	var opmlStruct OPML = OPML{}
	fileBytes, err := ioutil.ReadFile("export.opml")
	if err == nil {
		xml.Unmarshal(fileBytes, &opmlStruct)
		fmt.Println(opmlStruct)
	}

	// Create a slice (an array like thing) to hold all the podcast info we collect using this OPML file.
	Podcasts := make([]Podcast, len(opmlStruct.Podcasts))
	for i, val := range opmlStruct.Podcasts {
		// For each RSS url we get from the OPML file, we get it and extract data from the RSS feed.
		resp, err := http.Get(val.Url)
		if err == nil {
			fileBytes, err = ioutil.ReadAll(resp.Body)
			xml.Unmarshal(fileBytes, &Podcasts[i])
			resp.Body.Close()
		}
	}

	// Print the podcast information.
	for _, val := range Podcasts {
		fmt.Println(val)
	}
}
