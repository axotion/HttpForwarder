# HttpForwarder
_Super-duper splitter for your critial endpoints_

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/gluten-free.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
![alt text](https://media.giphy.com/media/XoM2ufTqkS10YdBUKr/giphy.gif)

# Features!

  - Split one http request to many without effort!
  - (Think of something else)

I created this tool for better control requests from companies I work with. 504 (and another critial status code) or lost data beacause of _something_ are not dangerous now! Splitter will try to deliver your requests no matter what.

### Installation

HttoForward requires x86/x64/arm to complie and at least GO 1.6.

 ```sh
$ go get github.com/axotion/HttpForwarder
```

### Usage

Move sites.json to your project folder and edit it

```json
[
    {
        "_client" : "TEST2",
        "identificator" : "7ac2970651534830874ba712a30de940", 
        "forward" : [
            {
                "address" : "http:/url.com/187q5lg1",
                "method" : "POST",
                "auth" : "",
                "username" : "",
                "password" : "",
                "retry" : 10,
                "expected_status" : 200
            },
            
            {
                "address" : "https://www.dobretrojany.pl/fdsfdsfdsfsfs",
                "method" : "POST",
                "auth" : "basic",
                "username" : "username",
                "password" : "password",
                "retry" : 30,
                "expected_status" : 200
            }
        ]
    },
    {
        "identificator" : "TEST1", 
        "forward" : [
            {
                "address" : "http://test1.com/id",
                "method" : "POST",
                "auth" : "",
                "username" : "",
                "password" : "",
                "retry" : 10,
                "expected_status" : 200
            },
            {
                "address" : "http://bing1.com/id",
                "method" : "POST",
                "auth" : "basic",
                "username" : "",
                "password" : "",
                "retry" : 30,
                "expected_status" : 200
            }
        ]
    }
]
```

Then in your main section (or whatever you want!) put 
```GO
httpForwarder := httpforwarder.New()
httpForwarder.Run(host, port)
```
**This will run small server for incoming requests**

You can invoke splitter by request on specific URL
```
0.0.0.0:9000/forward/identificator with POST method
```

These headers will be append
```
"X-Real-IP"
"X-Forwarder-For"
"X-Forwarded-Host"
```


And that's it! Nothing more!

### Todos

 - Write MORE Tests
 - Write ANY Tests
 - More auth methods

License
----

MIT


**Free Software, Hell Yeah!**