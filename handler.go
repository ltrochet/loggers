/*
####### loggers (c) 2024 Loïc TROCHET ############################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"bytes"
	"net/http"
)

type (
	Logger interface {
		ID() string
		Name() string
		Level() string
		LessLog()
		MoreLog()
	}
)

func Handler(loggers map[string]Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()

			if id := q.Get("id"); id != "" {
				if l, ok := loggers[id]; ok {
					cmd := q.Get("cmd")

					switch cmd {
					case "less":
						l.LessLog()
					case "more":
						l.MoreLog()
					}
				}
			}

			var buf bytes.Buffer

			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			if err := render(&buf, loggers); err == nil {
				_, _ = w.Write(buf.Bytes())
			} else {
				_, _ = w.Write([]byte(err.Error()))
			}
		},
	)
}

/*
####### END ############################################################################################################
*/
