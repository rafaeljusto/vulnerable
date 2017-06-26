# vulnerable

Lab scenario to test some website common attacks, such as
[CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery) (Cross-Site
Request Forgery) and [XSS](https://en.wikipedia.org/wiki/Cross-site_scripting)
(Cross-site scripting). To test the CSRF attack we created 2 websites, the good
guy, that is a simple registration, login and update, and the bad guy website,
that will try to change the information while the session is still valid.

This lab was created for the [LACTLD](http://lactld.org/) 2017 technical
workshop at Costa Rica.

## Install

Download the project with the following command:

```
go get -u github.com/rafaeljusto/vunerable/...
```

Make sure that your `$GOPATH/bin` is in your `$PATH` environment. The just run
the commands:

```
goodguy -port 8080
badguy -port 8081 -attack-server localhost:8080
```

## CSRF

1. Go to the Good Guy website (`http://localhost:8080`), register a new account
   and login into the system.

2. Check registered name in Good Guy's search page
   (`http://localhost:8080/search`) to see the currently name of the registered
   user.

3. With your session still valid in the Good Guy website, access the Bad Guy
   website (`http://localhost:8081`) to get your account hijacked.

4. Go back to the Good Guy's search page (`http://localhost:8080/search`) to
   check that your account name was changed. Also, if you logout and try to
   login again it will failed, because the password was changed.