package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var firstInstall string = "No"
var secondInstall string = "No"
var thirdInstall string = "No"

func main() {
	colleges := []string{
		"Technical Computer Engineering",
		"Medical Lab Technology",
		"Civil Engineering",
		"Media",
		"Law",
		"English Language",
		"Business Administration",
		"Accounting",
		"Arabic Language",
		"Pharmacy",
	}
	var resutl [][]string
	var record []string
	ffile, err := ioutil.ReadFile("files/first.csv")
	if err != nil {
		log.Fatalln(err)
	}
	sfile, err := ioutil.ReadFile("files/second.csv")
	if err != nil {
		log.Fatalln(err)
	}
	bfile, err := ioutil.ReadFile("files/third.csv")
	if err != nil {
		log.Fatalln(err)
	}
	sliceBytes := bytes.Split(bfile, []byte("\n"))
	count := 0
	for _, c := range colleges {
		resutl = append(resutl, []string{"التسلسل", "الاسم", "البريد", "الكلية", "المرحلة", "القسط الاول", "القسط الثاني", "القسط الثالث"})
		for _, v := range sliceBytes[1:] {
			bslice := bytes.Split(v, []byte(","))
			year := string(bslice[6])
			if year != "First Year" {
				ok := paidStatus(bslice)
				if ok {
					id := string(bslice[0])
					email := string(bslice[1])
					name := string(bslice[4])
					college := string(bslice[7])
					if year == "Fouth Year" {
						year = "Fourth Year"
					}
					if c == college {
						count++
						first := checkOtherInstallments(id, ffile)
						second := checkOtherInstallments(id, sfile)
						if first {
							firstInstall = "Yes"
						}
						if second {
							secondInstall = "Yes"
						}
						record = []string{strconv.Itoa(count), email, name, college, year, firstInstall, secondInstall, thirdInstall}
						resutl = append(resutl, record)
						firstInstall = "No"
						secondInstall = "No"
					}
				}
			}
		}
		fname := fmt.Sprintf("%s.%s", c, "csv")
		file, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		w := csv.NewWriter(file)
		w.WriteAll(resutl)
		if err := w.Error(); err != nil {
			log.Fatalln(err)
		}
		resutl = [][]string{}
		record = []string{}
		count = 0
	}
}

func paidStatus(slice [][]byte) bool {
	mstatus := string(slice[8])
	ystatus := string(slice[9])

	email := string(slice[1])
	if mstatus == "" && ystatus == "" {
		if strings.Contains(email, "mpu.university") {
			return false
		}
		return true
	}
	return false
}

func checkOtherInstallments(id string, slice []byte) bool {
	sliceBytes := bytes.Split(slice, []byte("\n"))
	for _, v := range sliceBytes[1:] {
		bslice := bytes.Split(v, []byte(","))
		stdid := string(bslice[0])
		if stdid == id {
			status := string(bslice[6])
			if status == "paid" {
				return true
			}
			return false
		}
	}
	return false
}
