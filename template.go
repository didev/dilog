package main

import (
	"fmt"
	"strings"
)

const headHTML = `<!DOCTYPE html><head><title>dilog</title>
	<meta charset="utf-8">
	<style>
	@charset "utf-8";
	html, body {
		margin:0px;
		padding:0px;
		background:#000000;
		font-family: Courier, 'Courier New', Courier_New;
	}
	a:link {text-decoration:none; color:#5AAEFE;}
	a:visited {text-decoration:none; color:#5AAEFE;}
	a:active {text-decoration:none; color:#5AAEFE;}
	a:hover {text-decoration:none; color:#5AAEFE; cursor:pointer;}
	div {
		display: inline-block;
		line-height: 6px;
	}
	.searchform{
		width:450px;
		padding: 8px;
		margin: 20px auto;
		overflow: hidden;
		border-width: 1px;
		border-style: solid;
		border-color: #C7C7C7;
		font-family: Courier, 'Courier New', Courier_New;
		color: #FFF;
		background-color: #000000;
		vertical-align: bottom;
		font-size:15px;
		font-weight: bold;
		-webkit-appearance: none;
	}

	.button {
		width:100px;
		padding: 7px;
		margin: 20px auto;
		text-align:center;
		color:#FFFFFF;
		border-width: 1px;
		border-style: solid;
		border-color: #C7C7C7;
		background-color: #222;
		font-family: Courier, 'Courier New', Courier_New;
		font-size:15px;
		font-weight: bold;
		letter-spacing: 0.5px;
		vertical-align: bottom;
	}
	.info {
		text-align: left;
		margin-top: 10px;
		margin-left: 10px;
		font-size: 14px;
		color: #B6C751;
		font-weight : bold;
	}

	.title {
		margin-left: 10px;
		margin-right: 10px;
		font-size: 16px;
		color: #B6C751;
		font-weight : bold;
		font-family: Courier, 'Courier New', Courier_New;
	}

	.log {
		margin-left: 10px;
		margin-right: 10px;
		font-size: 16px;
		color: #FFFFFF;
		font-family: Courier, 'Courier New', Courier_New;
		display:table-cell;
		font-weight: bold;
		white-space: nowrap;
	}

	.noresult {
		font-size: 14px;
		color: #FF0000;
		font-weight : bold;
	}

	</style>
	<link rel="icon" type="image/png" href="http://10.0.90.193:8080/template/icon/dilog.png">
	</head><body>`

func infoHTML(tool, project, slug string) string {
	result := ""
	if tool != "" {
		result = "Info : " + tool
	}
	if project != "" {
		result = result +" > "+ project
	}
	if slug != "" {
		result = result +" > "+ slug
	}
	if result == "" {
		return `<div class="info">&nbsp;</div><br>`
	} else {
		return `<div class="info">` + result + "</div><br>"
	}
}

func searchboxHTML(searchword string) string {
	return fmt.Sprintf(`
	<center><form action="/" method="POST" name="checkboxs" class="editform">
	<input class="searchform" id="search" placeholder="Search word.." type="text" maxlength="50" name="searchword" autofocus="autofocus" value="%s">
	<button class="button" id="button">Search</button><br>
	</form></center>`, searchword)
}

func flength(text string, length int) string {
	if length >= len(text) {
		spacenum := length - len(text)
		return strings.Repeat("&nbsp;", spacenum)
	} else {
		return ""
	}
}


func title() string {
	return fmt.Sprintf(`<br><div class="title">%s%s
						%s%s
						%s%s
						%s%s
						%s%s
						%s%s
						%s%s
						%s</div><br><br>`,
			"Time",flength("Time",26),
			"Keep",flength("Keep",5),
			"IP",flength("IP",14),
			"User",flength("User",19),
			"Tool",flength("Tool",11),
			"Project",flength("Project",11),
			"Slug",flength("Slug",15),
			"Log")
}

func linktool(tool string) string {
	if tool != "" {
		return fmt.Sprintf(`<a href="/log/%s">%s</a>`, tool, tool)
	} else {
		return ""
	}
}

func linkproject(tool, project string) string {
	if tool != "" && project != "" {
		return fmt.Sprintf(`<a href="/log/%s/%s">%s</a>`, tool, project, project)
	} else {
		return project
	}
}

func linkslug(tool, project, slug string) string {
	if tool != "" && project != "" && slug != "" {
		return fmt.Sprintf(`<a href="/log/%s/%s/%s">%s</a>`, tool, project, slug, slug)
	} else {
		return slug
	}
}

func linklog(log string) string {
	var rstring string
	rstring = ""
	loglist := strings.Split(log, " ")
	for _, i := range loglist {
		if strings.Contains(i, "/show") || strings.Contains(i, "/lustre") {
			rstring = rstring + fmt.Sprintf(`<a href="dilink://%s">%s</a>`, i, i) + " "
		} else {
			rstring = rstring + i + " "
		}
	}
	return rstring
}

func log2tag(log Log) string {
	return fmt.Sprintf(`<div class="log">&nbsp;%s&nbsp;
						%s%s&nbsp;
						%s%s&nbsp;
						%s%s&nbsp;
						%s%s&nbsp;
						%s%s&nbsp;
						%s%s&nbsp;
						%s</div><br>`,
			log.Time,
			log.Keep, flength(log.Keep,4),
			log.Cip, flength(log.Cip,13),
			log.User, flength(log.User,18),
			linktool(log.Tool), flength(log.Tool,10),
			linkproject(log.Tool, log.Project), flength(log.Project,10),
			linkslug(log.Tool, log.Project, log.Slug), flength(log.Slug,14),
			linklog(log.Log))
}

func logHTML(logs []Log) string {
	var result string = ""
	if len(logs) == 0 {
		result = `<br><center><div class="noresult">No Result.</div></center>`
	} else {
		result = result + title()
		for _, i := range logs {
			result = result + log2tag(i)
		}
	}
	return result
}
