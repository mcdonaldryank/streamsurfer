// Templates for webpages
package main

const ReportMainPageTemplate =`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>HLS Probe II Reports</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
</head>
<body><h1>Reports are:</h1>
 <ul>
  <li><a href="rprt/3hours">Errors for all streams for last 3 hours</a></li>
  <li><a href="rprt/last">Last stream errors</a></li>
 </ul>
</body>
</html>
`

const Report3HoursTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Errors per 3 hours for all groups :: HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
</head>
<body><h1>Errors per 3 hours for all groups</h1>
<table class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th rowspan="2">Group</th>
            <th rowspan="2">Name</th>
            <th colspan="7">Last hour</th>
            <th colspan="7">Two hours ago</th>
            <th colspan="7">Three hours ago</th>
          </tr>
          <tr>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="List is empty">LE</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="HLS parsing problem">HLS</th>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="List is empty">LE</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="HLS parsing problem">HLS</th>
            <th rel="tooltip" title="Slow response">SR</th>
            <th rel="tooltip" title="Bad status">BS</th>
            <th rel="tooltip" title="Bad URI">BU</th>
            <th rel="tooltip" title="List is empty">LE</th>
            <th rel="tooltip" title="Timeout on read">RT</th>
            <th rel="tooltip" title="Timeout on connect">CT</th>
            <th rel="tooltip" title="HLS parsing problem">HLS</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr>
            <td>{{group}}</td>
            <td><a href="{{uri}}">{{name}}</a></td>
            <td{{#sw-severity}} class="{{sw-severity}}"{{/sw-severity}}>{{sw}}</td>
            <td{{#bs-severity}} class="{{bs-severity}}"{{/bs-severity}}>{{bs}}</td>
            <td{{#bu-severity}} class="{{bu-severity}}"{{/bu-severity}}>{{bu}}</td>
            <td{{#le-severity}} class="{{le-severity}}"{{/le-severity}}>{{le}}</td>
            <td{{#rt-severity}} class="{{rt-severity}}"{{/rt-severity}}>{{rt}}</td>
            <td{{#ct-severity}} class="{{ct-severity}}"{{/ct-severity}}>{{ct}}</td>
            <td{{#hls-severity}} class="{{hls-severity}}"{{/hls-severity}}>{{hls}}</td>
            <td>{{sw2}}</td>
            <td>{{bs2}}</td>
            <td>{{bs2}}</td>
            <td>{{le2}}</td>
            <td>{{rt2}}</td>
            <td>{{ct2}}</td>
            <td>{{hls2}}</td>
            <td>{{sw3}}</td>
            <td>{{bs3}}</td>
            <td>{{bu3}}</td>
            <td>{{le3}}</td>
            <td>{{rt3}}</td>
            <td>{{ct3}}</td>
            <td>{{hls3}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`

const ReportLastTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>All groups last errors :: HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
</head>
<body><h1>Last errors</h1>
<table class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th>Group</th>
            <th>Name</th>
            <th>Status</th>
            <th>Length</th>
            <th>Request Duration</th>
            <th>Last Checked</th>
            <th>Error</th>
            <th>Last Hour</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr class="{{severity}}">
            <td>{{group}}</td>
            <td><a href="{{uri}}">{{name}}</a></td>
            <td>{{status}}</td>
            <td>{{contentlength}}</td>
            <td>{{elapsed}}</td>
            <td>{{started}}</td>
            <td>{{error}}</td>
            <td>{{totalerrs}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`

const ReportGroupLastTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{#Vars}}{{group}} :: {{/Vars}}HLS Probe II</title>
<meta name="description" content="Creating a table with Twitter Bootstrap. Learn how to use Twitter Bootstrap toolkit to create Tables with examples.">
<link href="/css/bootstrap.css" rel="stylesheet">
<link href="/css/custom.css" rel="stylesheet">
</head>
<body>{{#Vars}}<h1>{{group}} last errors</h1>{{/Vars}}
<table class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Length</th>
            <th>Request Duration</th>
            <th>Last Checked</th>
            <th>Error</th>
            <th>Last Hour</th>
          </tr>
        </thead>
        <tbody>
        {{#TableData}}
          <tr class="{{severity}}">
            <td><a href="{{uri}}">{{name}}</a></td>
            <td>{{status}}</td
            <td>{{contentlength}}</td>
            <td>{{elapsed}}</td>
            <td>{{started}}</td>
            <td>{{error}}</td>
            <td>{{totalerrs}}</td>
          </tr>
        {{/TableData}}
        </tbody>
</table>
Generated by <a href="http://github.com/grafov/hlsprobe2">HLS Probe II</a>
</body>
</html>
`
