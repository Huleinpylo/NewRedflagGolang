package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"time"

	
)

func banner() {
	fmt.Println("NewRedflag --- by 0xlildoudou -- fork XCOM go ctf\n")
}

func downloadAll(link string, output string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	filename := "redflag_all.txt"
	if output != "" {
		filename = output
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}

	return filename
}

func downloadLatest(link string, output string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02")
	filename := "redflag_" + date + ".txt"
	if output != "" {
		filename = output
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}

	return filename
}

func main() {
	latest := flag.Bool("latest", false, "latest list")
	all := flag.Bool("all", false, "All list")
	output := flag.String("output", "", "Output file")
	verbose := flag.Bool("verbose", false, "Verbose output")
	day := flag.Int("day", 0, "Day of list")

	flag.Parse()

	banner()

	var outputFile string

	if *all {
		link := "https://dl.red.flag.domains/red.flag.domains.txt"
		if *verbose {
			fmt.Println("[Download]:", link)
		}
		outputFile = downloadAll(link, *output)
	} else if *latest {
		link := "https://dl.red.flag.domains/daily/" + time.Now().Format("2006-01-02") + ".txt"
		if *verbose {
			fmt.Println("[Download]:", link)
		}
		outputFile = downloadLatest(link, *output)
	} else if *day != 0 {
		date := time.Now().AddDate(0, 0, -*day)
		link := "https://dl.red.flag.domains/daily/" + date.Format("2006-01-02") + ".txt"
		if *verbose {
			fmt.Println("[Download]:", link)
		}
		outputFile = downloadLatest(link, *output)
	}

	fmt.Println("--->", outputFile)
}
