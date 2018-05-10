package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	Iter   int
	Time   float64
	Bytes  int
	Allocs int
}

type results []result

func (rs results) MedianTime() float64 {
	t := []float64{}
	for _, r := range rs {
		t = append(t, r.Time)
	}
	sort.Float64s(t)
	return t[len(t)/2]
}

var re = regexp.MustCompile(
	`^Benchmark/([a-zA-Z0-9_-]+)/(.+?)\s+(\d+)\s+([\d.]+) ns/op(?:\s+(\d+) B/op\s+(\d+) allocs/op)?`)

func main() {
	names := map[string]struct{}{}
	report := map[string]map[string]results{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		sm := re.FindStringSubmatch(line)
		if sm == nil {
			continue
		}

		name := sm[1]
		names[name] = struct{}{}
		metric := sm[2]
		m := report[metric]
		if m == nil {
			m = map[string]results{}
			report[metric] = m
		}
		m[name] = append(m[name], result{
			Iter:   atoi(sm[3]),
			Time:   atof(sm[4]),
			Bytes:  atoi(sm[5]),
			Allocs: atoi(sm[6]),
		})
	}

	snames := []string{}
	for name := range names {
		snames = append(snames, name)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(snames)))
	metrics := []string{}
	for metric := range report {
		metrics = append(metrics, metric)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(metrics)))

	fmt.Print("| Metric")
	for _, name := range snames {
		fmt.Printf(" | %s", name)
	}
	fmt.Println(" |")
	fmt.Println(strings.Repeat("|-", len(names)+1) + "|")
	for _, metric := range metrics {
		results := report[metric]
		fmt.Printf("| %s", metric)
		best := bestTime(results)
		for _, name := range snames {
			if r, found := results[name]; found {
				format := " | %s"
				if name == best {
					format = " | &#x1F538; **%s**"
				}
				fmt.Printf(format, strconv.FormatFloat(r.MedianTime(), 'f', -1, 64))
			} else {
				fmt.Print(" | n/a")
			}
		}
		fmt.Println(" |")
	}
}

func bestTime(results map[string]results) string {
	var bestName string
	var bestTime float64 = -1
	for name, result := range results {
		if bestTime == -1 || result.MedianTime() < bestTime {
			bestName = name
			bestTime = result.MedianTime()
		}
	}
	return bestName
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
