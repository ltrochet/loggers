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
	<title>loggers</title>
	<style>
		body {
			background-color: #314051;
			color: #32caa9;
		}
		a {
			text-decoration: none;
			padding: 0 10px;
			border: 1px solid #0d8e75;
			color: #0d8e75;
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
			padding: 10px 5px;
		}
		th {
			text-align: left;
			padding: 5px 5px;
		}
	</style>
</head>
<body>
{{range $id, $l := .}}
	<h1>Logger(s)</h1>
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
				<td><a href="?id={{$id}}&cmd=dec" title="Decrement log level">-</a></td>
				<td><a href="?id={{$id}}&cmd=inc" title="Increment log level">+</a></td>
			</tr>
		</tbody>
	</table>
	<p>v0.0.0</p>
	<p>© 2024 Loïc TROCHET</p>
{{else}} 
	<h1>No logger</h1>
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
