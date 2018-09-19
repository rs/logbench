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
	Label string     `json:"label"`
	Data  []*float64 `json:"data"`
}

// group -> test -> lib
type Report map[string]map[string]*benchstat.Metrics

func main() {
	libs := map[string]struct{}{}
	reports := map[string]Report{}

	c := &benchstat.Collection{}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	c.AddConfig("-", data)
	c.Tables() // trigger stats update
	for key, metrics := range c.Metrics {
		report := reports[key.Unit]
		if report == nil {
			report = Report{}
			reports[key.Unit] = report
		}
		idx := strings.IndexByte(key.Benchmark[1:], '/')
		if idx == -1 {
			log.Fatalf("invalid benchmark name: %s", key.Benchmark)
		}
		lib := key.Benchmark[1 : idx+1]
		libs[lib] = struct{}{}
		name := key.Benchmark[idx+2:]
		n := report[name]
		if n == nil {
			n = map[string]*benchstat.Metrics{}
			report[name] = n
		}
		n[lib] = metrics
	}

	slibs := []string{}
	for lib := range libs {
		slibs = append(slibs, lib)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(slibs)))

	charts := map[string]Chart{}
	for unit, report := range reports {
		charts[unit] = buildChart(report, slibs)
	}
	b, err := json.Marshal(charts)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func buildChart(report Report, libs []string) Chart {
	chart := Chart{}
	for name := range report {
		chart.Labels = append(chart.Labels, name)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(chart.Labels)))
	for _, lib := range libs {
		chart.Datasets = append(chart.Datasets, Dataset{
			Label: lib,
			Data:  make([]*float64, len(chart.Labels)),
		})
	}
	for i, name := range chart.Labels {
		results := report[name]
		for j, ds := range chart.Datasets {
			if m, found := results[ds.Label]; found {
				chart.Datasets[j].Data[i] = &m.Mean
			}
		}
	}
	return chart
}
