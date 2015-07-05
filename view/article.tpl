<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
	<title>{{ article.Title }}</title>
	<link rel="stylesheet" href="/assets/css/style.css" media="all">
  </head>
  <body>
	<div id="content">
	  <h1 class="title">{{ article.Title }}</h1>
	  <div class="body">
		{{ article.Text }}
	  </div>
	  <span class="date">Updated at: {{ article.UpdatedAt | date: "2006/01/02 03:04:05" }}</span>
	</div>
  </body>
</html>
