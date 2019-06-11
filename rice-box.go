package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "dilog.html",
		FileModTime: time.Unix(1560217133, 0),

		Content: string("{{define \"dilog\" }}\n<!DOCTYPE html>\n<head>\n        <title>dilog</title>\n        <meta charset=\"utf-8\">\n        <link rel=\"stylesheet\" href=\"/assets/bootstrap-4/css/bootstrap.min.css\">\n        <link rel=\"stylesheet\" href=\"/assets/css/dilog.css\">\n        <link rel=\"icon\" type=\"image/png\" href=\"/assets/img/dilog.png\">\n</head>\n<body>\n<!--info-->\n<div class=\"info pl-5 pb-2 pt-1 text-secondary\">\n        Info:\n        {{if .Tool}}\n                <span class=\"text-primary\">{{.Tool}}</span>\n                {{if .Project}}\n                > <span class=\"text-success\">{{.Project}}</span>\n                {{if .Slug}}\n                > <span class=\"text-info\">{{.Slug}}</span>\n                {{end}}\n                {{end}}\n        {{end}}\n</div>\n<!--searchbox-->\n<div class=\"container p-3\">\n        <div class=\"row justify-content-center align-items-center p-3\">\n                <form action=\"/search\" method=\"POST\" name=\"checkboxs\" class=\"editform text-center\">\n                <div class=\"input-group\">\n                        <input class=\"form-control bg-dark text-light\" id=\"search\" placeholder=\"Search word..\" type=\"text\" maxlength=\"50\" size=50 name=\"searchword\" autofocus=\"autofocus\" value=\"{{.Searchword}}\">\n                        <div class=\"input-group-append\">\n                        <button class=\"btn btn-dark border-light\" id=\"button\">Search</button>\n                        </div>\n                </div>\n                </form>\n        </div>\n</div>\n\n<!--print log-->\n<div class=\"p-5\">\n        {{if .Error}}\n        <div class=\"text-danger text-center\">\n                {{.Error}}\n        </div>\n        {{else if .Logs}}\n                <div class=\"row text-warning mb-3\">\n                        <div class=\"col-lg-2\">Time / ID</div>\n                        <div class=\"col-lg-1\">Keep</div>\n                        <div class=\"col-lg-1\">User / IP</div>\n                        <div class=\"col-lg-1\">Tool</div>\n                        <div class=\"col-lg-1\">Project</div>\n                        <div class=\"col-lg-1\">Slug</div>\n                        <div class=\"col-lg-5\">Logs</div>\n                </div>\n                {{range .Logs}}\n                <div class=\"row text-secondary mb-1\">\n                        <div class=\"col-lg-2\">{{.Time}}<br>{{.ID}}</div>\n                        <div class=\"col-lg-1\">{{.Keep}}</div>\n                        <div class=\"col-lg-1\">{{.User}}<br>{{.Cip}}</div>\n                        <div class=\"col-lg-1\"><a href=\"/search?tool={{.Tool}}&page=1\">{{.Tool}}</a></div>\n                        <div class=\"col-lg-1\"><a href=\"/search?tool={{.Tool}}&project={{.Project}}&page=1\" class=\"text-success\">{{.Project}}</a></div>\n                        <div class=\"col-lg-1\"><a href=\"/search?tool={{.Tool}}&project={{.Project}}&slug={{.Slug}}&page=1\" class=\"text-info\">{{.Slug}}</a></div>\n                        <div class=\"col-lg-5 text-white\">{{addLink .Log}}</div>\n                </div>\n                <hr/>\n                {{end}}\n                \n        \n        {{else}}\n                <div class=\"text-warning text-center\">\n                No Result.\n                </div>\n        {{end}}\n        \n</div>\n\n<!--print page-->\n<div class=\"pages text-center m-5\">\n        {{range .TotalPagenum}}\n                <a href=\"/search?tool={{$.Tool}}&project={{$.Project}}&searchword={{$.Searchword}}&slug={{$.Slug}}&page={{.}}\" class=\"btn btn-dark btn-sm\">{{.}}</a>\n        {{end}}\n</div>\n        \n<div class=\"footer text-center text-secondary p-3\">\n        © 2019 <a href=\"https://lazypic.org\" class=\"text-secondary\">lazypic</a> & <a href=\"http://www.digitalidea.co.kr\" class=\"text-secondary\">digitalidea</a>\n</div>\n</body>\n<script src=\"/assets/bootstrap-4/js/bootstrap.min.js\"></script>\n</html>\n{{end}}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1560208725, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "dilog.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`assets/template`, &embedded.EmbeddedBox{
		Name: `assets/template`,
		Time: time.Unix(1560208725, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"dilog.html": file2,
		},
	})
}
