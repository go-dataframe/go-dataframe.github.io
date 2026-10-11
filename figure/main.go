// Command figure writes what the landing page shows beside its terminal:
//
//	cd figure && go run . ../static
//
// writes static/out/frame.txt (what the program prints), static/out/frame.html
// (the table a notebook shows for the same frame, its Display text/html) and
// static/img/figure.svg (the frame drawn by dfplot with go-plot). The page's
// terminal shows this file from func main on, read at build time, and the CI
// runs it again and fails if a byte differs.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-dataframe/dataframe"
	"github.com/go-dataframe/dataframe/dfplot"
)

func main() {
	// Yearly rainfall (mm) and rainy days at three stations; made-up numbers.
	df := dataframe.Must(dataframe.New(
		dataframe.Col("city", []string{"Lyon", "Paris", "Nice", "Lyon", "Paris", "Nice"}),
		dataframe.Col("year", []int{2023, 2023, 2023, 2024, 2024, 2024}),
		dataframe.Col("rain", []float64{830.5, 640.2, 702.3, 810.0, 598.7, 760.1}),
		dataframe.Col("days", []int{112, 118, 64, 107, 111, 70}),
	))

	// df.groupby("city")[["rain", "days"]].mean(), the wettest first.
	g := dataframe.Must(df.GroupBy("city"))
	g = dataframe.Must(g.Select("rain", "days"))
	means := dataframe.Must(g.Mean())
	means = dataframe.Must(means.SortValues([]string{"rain"}, false))
	fmt.Println(means)

	// What a notebook cell shows, and a bar chart.
	out := os.Args[1]
	write(filepath.Join(out, "out", "frame.txt"), means.Display()["text/plain"])
	write(filepath.Join(out, "out", "frame.html"), means.Display()["text/html"])
	fig := dataframe.Must(dfplot.BarOpts(means, []string{"rain"},
		[]dfplot.Option{dfplot.WithTitle("Mean yearly rainfall, 2023–2024 (mm)")}))
	f := dataframe.Must(os.Create(filepath.Join(out, "img", "figure.svg")))
	defer f.Close()
	if err := fig.SaveSVG(f); err != nil {
		panic(err)
	}
}

func write(path string, b []byte) {
	if err := os.WriteFile(path, append(b, '\n'), 0o644); err != nil {
		panic(err)
	}
}
