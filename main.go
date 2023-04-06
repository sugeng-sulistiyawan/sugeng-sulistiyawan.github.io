package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var simpleIcons = []string{
		"adobe",
		"amazonaws",
		"amazons3",
		"androidstudio",
		"arduino",
		"bower",
		"bootstrap",
		"cplusplus",
		"c",
		"canva",
		"chocolatey",
		"codeigniter",
		"composer",
		"cpanel",
		"css3",
		"dart",
		"delphi",
		"docker",
		"eagle",
		"electron",
		"expo",
		"firebase",
		"figma",
		"flutter",
		"fontawesome",
		"gatsby",
		"git",
		"go",
		"googlecloud",
		"googleanalytics",
		"googlefonts",
		"googleplay",
		"googletagmanager",
		"googlesheets",
		"googlesearchconsole",
		"grafana",
		"graphql",
		"gulp",
		"hugo",
		"html5",
		"icons8",
		"inkscape",
		"ionic",
		"joomla",
		"jquery",
		"jss",
		"javascript",
		"json",
		"kotlin",
		"kubernetes",
		"laragon",
		"laravel",
		"latex",
		"less",
		"letsencrypt",
		"libreoffice",
		"linux",
		"livewire",
		"lumen",
		"mariadb",
		"materialdesignicons",
		"mdx",
		"microsoftoffice",
		"microsoftvisio",
		"mongodb",
		"mysql",
		"nodedotjs",
		"notion",
		"npm",
		"obsstudio",
		"packagist",
		"perl",
		"php",
		"postgresql",
		"postcss",
		"prestashop",
		"prisma",
		"prezi",
		"prometheus",
		"protodotio",
		"pwa",
		"pusher",
		"python",
		"raspberrypi",
		"react",
		"redis",
		"sass",
		"semver",
		"sketchup",
		"socketdotio",
		"stackblitz",
		"swagger",
		"symfony",
		"tailwindcss",
		"travisci",
		"trello",
		"twilio",
		"typescript",
		"vite",
		"visualstudiocode",
		"vuedotjs",
		"webgl",
		"webrtc",
		"wordpress",
		"xampp",
		"yarn",
		"kicad",
		"yiiframework",
	}

	for _, element := range simpleIcons {
		fileUrl := "https://cdn.simpleicons.org/"
		fileSave := "static/images/icons/"
		err := DownloadFile(fileSave+element+".svg", fileUrl+element)
		if err == nil {
			fmt.Println("Downloaded: " + fileUrl + element)
		}
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if resp.StatusCode == 404 {
		return nil
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
