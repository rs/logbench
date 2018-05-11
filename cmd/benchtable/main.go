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
	report := map[string]map[string]map[string]*benchstat.Row{}

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
		libs[lib] = struct{}{}
		gidx := strings.LastIndexByte(row.Benchmark, '/')
		group := row.Benchmark[idx+2 : gidx]
		name := row.Benchmark[gidx+1:]
		g := report[group]
		if g == nil {
			g = map[string]map[string]*benchstat.Row{}
			report[group] = g
		}
		n := g[name]
		if n == nil {
			n = map[string]*benchstat.Row{}
			g[name] = n
		}
		n[lib] = row
	}

	slibs := []string{}
	for lib := range libs {
		slibs = append(slibs, lib)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(slibs)))
	groups := []string{}
	for group := range report {
		groups = append(groups, group)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(groups)))

	fmt.Print("<table><tr><td><b>Name</b></td>")
	for _, lib := range slibs {
		fmt.Printf("<td><b>%s</b></td>", lib)
	}
	fmt.Println("</tr>")
	for _, group := range groups {
		fmt.Printf("<tr><td colspan=%d><b>%s</b></td></tr>\n", len(libs)+1, group)
		for name, results := range report[group] {
			fmt.Printf("<tr><td>%s</td>", name)
			best := bestTime(results)
			for _, lib := range slibs {
				if r, found := results[lib]; found {
					format := "<td>%s</td>"
					if lib == best {
						format = "<td>&#x1F538; <b>%s</b></td>"
					}
					fmt.Printf(format, r.Metrics[0].Format(r.Scaler))
				} else {
					fmt.Print("<td>n/a</td>")
				}
			}
			fmt.Println("</tr>")
		}
	}
	fmt.Println("</table>")
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
