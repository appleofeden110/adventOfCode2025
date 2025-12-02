package firstday_test

import (
	firstday "advent2026/first_day"
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFindZero(t *testing.T) {
	p := firstday.Password{ActualPassword: 0}

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatalf("Err reading input file: %v\n", err)
	}
	scanner := bufio.NewScanner(f)

	start := firstday.DialStart
	for scanner.Scan() {
		r := firstday.GetRotation(scanner.Text())

		res := firstday.FindZero(start, *r)
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
	p := firstday.Password{ActualPassword: 0}

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0777)
	if err != nil {
		t.Fatalf("Err reading input file: %v\n", err)
	}
	scanner := bufio.NewScanner(f)

	start := firstday.DialStart
	for scanner.Scan() {
		r := firstday.GetRotation(scanner.Text())

		res, atZero := firstday.FindZeroMethod0x434C49434B(start, *r)
		start = res
		p.LastRotation = *r

		p.ActualPassword += atZero

		if res == 0 {
			if r.NumChange == 0 {
				log.Println(scanner.Text())
			}
			p.ActualPassword++
		}

	}
	fmt.Printf("Actual password by method 0x434C49434B: %d\n", p.ActualPassword)
}
