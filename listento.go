package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Outline struct {
	Title string `xml:"text,attr"`
	Type  string `xml:"type,attr"`
	Url   string `xml:"xmlUrl,attr"`
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

	// open the opml.file and unmarshal it into our Outline struct
	opml_file, _ := os.Open("export.opml")
	opml_reader := bufio.NewScanner(opml_file)
	opml_text := ""
	for opml_reader.Scan() {
		opml_text += opml_reader.Text() + "\n"
		if strings.Contains(opml_reader.Text(), "<outline ") {
			//        fmt.Println(opml_reader.Text())
			var opml_struct Outline = Outline{}
			xml.Unmarshal([]byte(opml_reader.Text()), &opml_struct)
			fmt.Println(opml_struct.Title)
			fmt.Println(opml_struct.Type)
			fmt.Println(opml_struct.Url)
		}
	}
}
