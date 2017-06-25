# vulnerable

Lab scenario to test some website common attacks, such as
[CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery) (Cross-Site
Request Forgery) and [XSS](https://en.wikipedia.org/wiki/Cross-site_scripting)
(Cross-site scripting). To test the CSRF attack we created 2 websites, the good
guy, that is a simple registration, login and update, and the bad guy website,
that will try to change the information while the session is still valid.

This lab was created for the [LACTLD](http://lactld.org/) 2017 technical
workshop at Costa Rica.