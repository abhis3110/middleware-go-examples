# middleware-go-examples
Real world example:
This middleware example has two goals:
    1. Write a middleware that makes sure request has Header "Content-Type" application/json
    2. Write a middleware that adds current server time to the reponse cookie





Chaining of Middleware
----------------------
This will return the handler and run the middlewares in the order:

    1. filterContentType
    2. setTimeCookie
    3. ... and then our handler
