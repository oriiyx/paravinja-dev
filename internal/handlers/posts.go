package handlers

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	posts2 "github.com/oriiyx/paravinja-dev/views/posts"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/util"
)

// PostsIndex serves homepage
func PostsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		if slug == "" {
			http.Error(w, "Slug missing", http.StatusBadRequest)
			return
		}

		var filePath = "posts/" + slug + ".md"
		postByte, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Post "+filePath+" not found", http.StatusNotFound)
			return
		}

		var md = goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				highlighting.NewHighlighting(
					highlighting.WithWrapperRenderer(customWrapperRenderer),
				),
			),
		)

		var buf bytes.Buffer
		if err := md.Convert(postByte, &buf); err != nil {
			panic(err)
		}

		err = posts2.Index(buf.String(), strings.Replace(slug, "-", " ", -1)).Render(r.Context(), w)
		if err != nil {
			log.Println("Error occurred in rendering homepage: ", err)
			return
		}
	}
}

func customWrapperRenderer(w util.BufWriter, ctx highlighting.CodeBlockContext, entering bool) {
	if entering {
		w.WriteString(`<div class="mockup-code">`)
	} else {
		w.WriteString(`</div>`)
	}
}
