/*
####### loggers (c) 2024 Loïc TROCHET ############################################################## MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package loggers

import (
	"html/template"
	"io"
)

func render(w io.Writer, loggers map[string]Logger) error {
	t, err := template.New("loggers").Parse(
		`<!DOCTYPE html>
<html>
<head>
	<title>loggers v0.0.0</title>
	<style>
		h1 {
			text-align: center;
		}
		p {
			text-align: center;
			font-size: smaller;
		}
		table {
			border-collapse: collapse;
			margin-left: auto;
			margin-right: auto;
		}
		tr {
			border-bottom: 1px solid #ccc;
		}
		th, td {
			text-align: left;
			padding: 5px;
		}
	</style>
</head>
<body>
{{range $id, $l := .}}
	<h1>Loggers</h1>
	<table>
		<thead>
			<tr>
				<th>ID</th>
				<th>Name</th>
				<th>Level</th>
				<th></th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>{{$id}}</td>
				<td>{{$l.Name}}</td>
				<td>{{$l.Level}}</td>
				<td><a href="?id={{$id}}&cmd=inc">+</a></td>
				<td><a href="?id={{$id}}&cmd=dec">-</a></td>
			</tr>
		</tbody>
	</table>
	<p>© 2024 Loïc TROCHET</p>
{{else}} 
	<p><strong>no loggers</strong><p>
{{end}}
</body>
</html>`,
	)
	if err != nil {
		return err
	}

	return t.Execute(w, loggers)
}

/*
####### END ############################################################################################################
*/
