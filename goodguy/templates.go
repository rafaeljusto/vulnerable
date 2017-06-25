package main

var headerTmpl = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Good Guy</title>

    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">
    <style>
      body {
        padding-top: 7rem;
      }
    </style>
  </head>
  <body>
    <nav class="navbar navbar-toggleable-md navbar-inverse bg-primary fixed-top">
      <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <a class="navbar-brand" href="#">GoodGuy</a>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
          {{if eq .LoggedUsername ""}}
          <li class="nav-item {{if eq .Menu "login"}}active{{end}}">
            <a class="nav-link" href="/">Login</a>
          </li>
          {{else}}
          <li class="nav-item {{if eq .Menu "update"}}active{{end}}">
            <a class="nav-link" href="/">Update</a>
          </li>
          {{end}}
          <li class="nav-item {{if eq .Menu "register"}}active{{end}}">
            <a class="nav-link" href="/register">Register</a>
          </li>
          <li class="nav-item {{if eq .Menu "search"}}active{{end}}">
            <a class="nav-link" href="/search">Search</a>
          </li>
          {{if ne .LoggedUsername ""}}
          <li class="nav-item">
            <a class="nav-link" href="/logout"><strong>Logout</strong></a>
          </li>
          {{end}}
        </ul>
      </div>

      {{if ne .LoggedUsername ""}}
      <span class="navbar-text">
        [{{.LoggedUsername}}]
      </span>
      {{end}}
    </nav>
    <div class="container">
      {{if gt (len .SuccessMessage) 0}}
      <div class="alert alert-success" role="alert">
        {{.SuccessMessage}}
      </div>
      {{end}}
      {{if gt (len .ErrorMessage) 0}}
      <div class="alert alert-danger" role="alert">
        {{.ErrorMessage}}
      </div>
      {{end}}
`

var footerTmpl = `
    </div>
    <script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" integrity="sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" integrity="sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn" crossorigin="anonymous"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.6.4/angular.min.js" crossorigin="anonymous"></script>
  </body>
</html>
`

var loginTmpl = headerTmpl + `
      <form action="/" method="post">
        <fieldset class="form-group">
          <legend>Login</legend>
          <p>
            <label for="username">Username</label>
            <input type="text" name="username" id="username" value="{{.Username}}" class="form-control" />
          </p>
          <p>
            <label for="password">Password</label>
            <input type="password" name="password" id="password" class="form-control" />
          </p>
          <p>
            <button type="submit" class="btn btn-primary">Login</button>
          </p>
        </fieldset>
      </form>` + footerTmpl

var registerTmpl = headerTmpl + `
    <form action="/register" method="post">
      <fieldset class="form-group">
        <legend>Registration</legend>
        <p>
          <label for="name">Name</label>
          <input type="text" name="name" id="name" value="{{.Name}}" class="form-control" />
        </p>
        <p>
          <label for="username">Username</label>
          <input type="text" name="username" id="username" value="{{.Username}}" class="form-control" />
        </p>
        <p>
          <label for="password">Password</label>
          <input type="password" name="password" id="password" class="form-control" />
        </p>
        <p>
          <button type="submit" class="btn btn-primary">Register</button>
        </p>
      </fieldset>
    </form>` + footerTmpl

var updateTmpl = headerTmpl + `
    <form action="/update" method="post">
      <fieldset class="form-group">
        <legend>Update {{.Username}}</legend>
        <p>
          <label for="name">Name</label>
          <input type="text" name="name" id="name" value="{{.Name}}" class="form-control" />
        </p>
        <p>
          <label for="password">Password</label>
          <input type="password" name="password" id="password" class="form-control" />
        </p>
        <p>
          <button type="submit" class="btn btn-primary">Update</button>
        </p>
      </fieldset>
    </form>` + footerTmpl

var searchTmpl = headerTmpl + `
    <form action="/search" method="post">
      <fieldset class="form-group">
        <legend>Search</legend>
        <p>
          <label for="username">Username</label>
          <input type="text" name="username" id="username" value="{{.Username}}" class="form-control" />
        </p>
        <p>
          <button type="submit" class="btn btn-primary">Search</button>
        </p>
      </fieldset>
    </form>
    {{if ne .Result.Username ""}}
    <div data-ng-app>
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
          <h3>
            {{toHTML .Result.Username}}
            <small class="text-muted">result</small>
          </h3>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-4 col-md-4 col-sm-12 col-xs-12">
          <label><strong>Name</strong></label>
        </div>
        <div class="col-lg-8 col-md-8 col-sm-12 col-xs-12">
          {{toHTML .Result.Name}}
        </div>
      </div>
    </div>
    {{end}}` + footerTmpl

var faviconData = `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
fDbrl2ICAgICAgICAgICAg//2Q==`
