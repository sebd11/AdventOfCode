package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filename = "input.txt"

const birthYear = "byr"
const issueYear = "iyr"
const expirationYear = "eyr"
const height = "hgt"
const hairColor = "hcl"
const eyeColor = "ecl"
const passportID = "pid"
const countryID = "cid"

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	passports := extractInputs()
	firstPart(passports)
	secondPart(passports)
}

func extractInputs() []Passport {
	passports := []Passport{}
	passport := Passport{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, passport)
			passport = Passport{}
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				splitField := strings.Split(field, ":")
				switch splitField[0] {
				case birthYear:
					passport.byr = splitField[1]
				case issueYear:
					passport.iyr = splitField[1]
				case expirationYear:
					passport.eyr = splitField[1]
				case height:
					passport.hgt = splitField[1]
				case hairColor:
					passport.hcl = splitField[1]
				case eyeColor:
					passport.ecl = splitField[1]
				case passportID:
					passport.pid = splitField[1]
				case countryID:
					passport.cid = splitField[1]
				}
			}
		}
	}
	passports = append(passports, passport)
	return passports
}

func firstPart(passports []Passport) {
	validPassports := 0
	for _, passport := range passports {
		isValid := passport.byr != "" && passport.iyr != "" && passport.eyr != "" &&
			passport.hgt != "" && passport.hcl != "" && passport.ecl != "" && passport.pid != ""

		if isValid {
			validPassports += 1
		}
	}
	fmt.Printf("Part 1 - Number of valid passports:  %d\n", validPassports)
}

func validatePassport(passport Passport) bool {
	if passport.byr != "" {
		if !validateYear(passport.byr, 1920, 2002) {
			return false
		}
	} else {
		return false
	}
	if passport.iyr != "" {
		if !validateYear(passport.iyr, 2010, 2020) {
			return false
		}
	} else {
		return false
	}
	if passport.eyr != "" {
		if !validateYear(passport.eyr, 2020, 2030) {
			return false
		}
	} else {
		return false
	}
	if passport.hgt != "" {
		if strings.Contains(passport.hgt, "cm") {
			num, err := strconv.Atoi(strings.Split(passport.hgt, "cm")[0])
			if !(err == nil && num >= 150 && num <= 193) {
				return false
			}
		} else if strings.Contains(passport.hgt, "in") {
			num, err := strconv.Atoi(strings.Split(passport.hgt, "in")[0])
			if !(err == nil && num >= 59 && num <= 76) {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	if passport.hcl != "" {
		re := regexp.MustCompile("^#(?:[0-9a-fA-F]{3}){1,2}$")
		hexCode := re.FindString(passport.hcl)
		if hexCode == "" {
			return false
		}
	} else {
		return false
	}
	if passport.ecl != "" {
		eyeColorIsValid := false
		for _, v := range eyeColors {
			if passport.ecl == v {
				eyeColorIsValid = true
				break
			}
		}
		if !eyeColorIsValid {
			return false
		}
	} else {
		return false
	}
	if passport.pid != "" {
		re := regexp.MustCompile("^[0-9]{9}$")
		digits := re.FindString(passport.pid)
		if digits == "" {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateYear(value string, min int, max int) bool {
	re := regexp.MustCompile("^[0-9]{4}$")
	digits := re.FindString(value)
	num, err := strconv.Atoi(digits)
	if !(err == nil && num >= min && num <= max) {
		return false
	}
	return true
}

func secondPart(passports []Passport) {
	validPassports := 0
	for _, passport := range passports {

		if validatePassport(passport) {
			validPassports += 1
		}
	}
	fmt.Printf("Part 2 - Number of valid passports:  %d", validPassports)
}
