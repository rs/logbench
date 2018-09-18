package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/perf/benchstat"
)

type Chart struct {
	Labels   []string  `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
}

func main() {
	libs := map[string]struct{}{}
	// group -> test -> lib
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
		libs[lib] = struct{}{}
		name := row.Benchmark[idx+2:]
		n := report[name]
		if n == nil {
			n = map[string]*benchstat.Row{}
			report[name] = n
		}
		n[lib] = row
	}

	chart := Chart{}
	for name := range report {
		chart.Labels = append(chart.Labels, name)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(chart.Labels)))
	slibs := []string{}
	for lib := range libs {
		slibs = append(slibs, lib)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(slibs)))
	for _, lib := range slibs {
		chart.Datasets = append(chart.Datasets, Dataset{
			Label: lib,
			Data:  make([]float64, len(chart.Labels)),
		})
	}
	for i, name := range chart.Labels {
		results := report[name]
		for j, ds := range chart.Datasets {
			if r, found := results[ds.Label]; found {
				chart.Datasets[j].Data[i] = r.Metrics[0].Mean
			}
		}
	}

	b, err := json.Marshal(chart)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
