/*
####### loggers (c) 2024 Loïc TROCHET ############################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"html/template"
	"io"
	"slices"
	"strings"
)

func render(w io.Writer, loggers map[string]Logger) error {
	t, err := template.New("loggers").Parse(
		`<!DOCTYPE html>
<html>
<head>
	<title>loggers</title>
	<style>
		body {
			background-color: #314051;
			color: #32caa9;
		}
		a {
			text-decoration: none;
			padding: 0 10px;
			border: 1px solid #32caa9;
			color: #32caa9;
		}
		a:hover {
			color: #ebf0f1;
			border: 1px solid #ebf0f1;
		}
		h1 {
			text-align: center;
			color: #0d8e75;
		}
		p {
			margin: 0;
			text-align: center;
			font-size: smaller;
			color: #475b6e;
		}
		table {
			border-collapse: collapse;
			margin-left: auto;
			margin-right: auto;
		}
		tr {
			border-bottom: 1px solid #475b6e;
		}
		td {
			text-align: left;
			padding: 10px 10px;
		}
		td.level {
			color: #ff7f50;
		}
		th {
			text-align: left;
			padding: 5px 10px;
		}
	</style>
</head>
<body>
	<h1>Logger(s)</h1>
	<table>
		<thead>
			<tr>
				<th></th>
				<th>ID</th>
				<th>Name</th>
				<th>Level</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
		{{range $n, $l := .}}
			<tr>
				<td>{{$n}}</td>
				<td>{{$l.ID}}</td>
				<td>{{$l.Name}}</td>
				<td class="level">{{$l.Level}}</td>
				<td>
					<a href="?id={{$l.ID}}&cmd=dec" title="Decrement log level">-</a>
					<a href="?id={{$l.ID}}&cmd=inc" title="Increment log level">+</a>
				</td>
			</tr>
		{{end}}
		</tbody>
	</table>
	<p>v0.0.0</p>
	<p>© 2024 Loïc TROCHET</p>
</body>
</html>`,
	)
	if err != nil {
		return err
	}

	list := make([]Logger, 0, len(loggers))

	for _, l := range loggers {
		list = append(list, l)
	}

	slices.SortFunc(
		list,
		func(a, b Logger) int {
			return strings.Compare(a.Name(), b.Name())
		},
	)

	return t.Execute(w, list)
}

/*
####### END ############################################################################################################
*/
