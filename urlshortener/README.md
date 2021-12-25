# Url Shortener

A basic url shortener, running on http server.

To create a new url, you must specify the url you need to shorten & optional expiring time in seconds. If the time of expiration is not specified, the url will not be deleted.

For example, if you want to add a new short url for youtube.com that expires in 1 hour, you need to call a method by URL:

<i>localhost:4000/create?url=https://youtube.com&expires=3600</i>

The response body of the method is the short url for the given long url. After shortening, you can just go by the given url, and the service will redirect you to the long url, or reply status 404, if the link is expired or does not exist.
