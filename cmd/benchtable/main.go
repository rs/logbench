package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/perf/benchstat"
)

func main() {
	libs := map[string]struct{}{}
	report := map[string]map[string]*benchstat.Row{}

	c := &benchstat.Collection{}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	c.AddConfig("-", data)
	tables := c.Tables()
	for _, row := range tables[0].Rows {
		idx := strings.IndexByte(row.Benchmark[1:], '/')
		if idx == -1 {
			log.Fatalf("invalid benchmark name: %s", row.Benchmark)
		}
		lib := row.Benchmark[1 : idx+1]
		bench := row.Benchmark[idx+2:]
		libs[lib] = struct{}{}
		m := report[bench]
		if m == nil {
			m = map[string]*benchstat.Row{}
			report[bench] = m
		}
		m[lib] = row
	}

	slibs := []string{}
	for lib := range libs {
		slibs = append(slibs, lib)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(slibs)))
	benchs := []string{}
	for bench := range report {
		benchs = append(benchs, bench)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(benchs)))

	fmt.Print("| Name")
	for _, lib := range slibs {
		fmt.Printf(" | %s", lib)
	}
	fmt.Println(" |")
	fmt.Println(strings.Repeat("|-", len(libs)+1) + "|")
	for _, bench := range benchs {
		results := report[bench]
		fmt.Printf("| %s", bench)
		best := bestTime(results)
		for _, lib := range slibs {
			if r, found := results[lib]; found {
				format := " | %s"
				if lib == best {
					format = " | &#x1F538; **%s**"
				}
				fmt.Printf(format, r.Metrics[0].Format(r.Scaler))
			} else {
				fmt.Print(" | n/a")
			}
		}
		fmt.Println(" |")
	}
}

func bestTime(results map[string]*benchstat.Row) string {
	var bestName string
	var bestRow *benchstat.Metrics
	for name, row := range results {
		if bestRow == nil || row.Metrics[0].Mean < bestRow.Mean {
			bestName = name
			bestRow = row.Metrics[0]
		}
	}
	return bestName
}
