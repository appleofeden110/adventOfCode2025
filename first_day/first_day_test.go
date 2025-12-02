package firstday

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFindZero(t *testing.T) {
	p := Password{ActualPassword: 0}

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatalf("Err reading input file: %v\n", err)
	}
	scanner := bufio.NewScanner(f)

	start := DialStart
	for scanner.Scan() {
		r := GetRotation(scanner.Text())

		res := FindZero(start, *r)
		start = res
		p.LastRotation = *r

		if res == 0 {
			if r.NumChange == 0 {
				log.Println(scanner.Text())
			}
			p.ActualPassword++
		}
	}
	fmt.Printf("Actual password: %d\n", p.ActualPassword)
}

func TestFindZeroMethod0x434C49434B(t *testing.T) {
	p := Password{ActualPassword: 0}

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatalf("Err reading input file: %v\n", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	start := DialStart
	for scanner.Scan() {
		r := GetRotation(scanner.Text())

		res, atZero := FindZeroMethod0x434C49434B(start, *r)
		start = res
		p.LastRotation = *r

		p.ActualPassword += atZero

		if res == 0 {
			if r.NumChange == 0 {
				log.Println(scanner.Text())
			}
		}

	}
	fmt.Printf("Actual password by method 0x434C49434B: %d\n", p.ActualPassword)
}
