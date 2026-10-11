# go-dataframe.github.io

The go-dataframe landing page (Hugo), served at <https://go-dataframe.github.io/>.

The layout is the go-* family template (go-fileshare.github.io's, through
`logos/landing-kit`, as go-plot and go-notebook use it): `layouts/index.html`,
the shared stylesheet `layouts/partials/styles.html`, and `[params.brand]` in
`hugo.toml`, read from the logo's gradient.

The example is real: `figure/main.go` builds a frame with
go-dataframe/dataframe, groups it, and writes what it prints
(`static/out/frame.txt`), the HTML table a notebook shows for it
(`static/out/frame.html`) and a bar chart drawn with go-plot
(`static/img/figure.svg`). The page's terminal shows that file, read at build
time, and the `ci` workflow runs it again and fails if a byte differs.

```sh
cd figure && go run . ../static
hugo server
```

| Workflow | What |
| --- | --- |
| `hugo.yml` | builds on pull requests; builds and deploys to Pages on `main` |
| `ci.yml` | the site builds and names every repository; the outputs are what their program writes |
| `links.yml` | every link and image of the built site resolves |

BSD-3-Clause.
