# GoCrawler
Concurrent and parallel GET/POST requests written in GoLang (with consumer-producer design). The Code is well commented and easy to configure for your own needs :)


## Why is it fast?

1. It can initiate multiple GoRoutines (50 in this case for demonstration, but feel free to tune it up). All goRoutines are also concurrent (async). Both parallelism and concurrency are exploited for speed gain.

2. Written in GoLang: GoLang is high-level like Python, and nearly fast as C/C++.


## How does it work?

### basic.go:
It reads URL's from a file (1000 in my case), creates 50 workers/consumers (goRoutines), and 1 producer (main thread). Producer puts the work (URL's) in the unbuffered channel, one by one. And workers get the URL's, one by one. This loop runs, till all of the URL's are processed. Successful outputs are written to a file, in a folder called `200`. This folder is safely created in the beginning of the process

### referer-request.go:
Same as `basic.go`. Only difference is, if the server responds with a successful response, we query another request, but this time with a referer, and process that request along with the original one. If everything goes well, we will have a 2 files for each URL, `url_original` and `url_referer`


## What is it good for?

1. This can be a barebone of a tool (maybe cyber-security), which requires to handle too many GET/POST requests, and must do them fast. The code is very easy to tweak, it's readable and well commented. 
In the second version (referer-request), after the first successful request, we send another request to the same server, but this time with a referer. It's just a demo for showing how easy to configure the code for your own needs.

2. If you are new to GoLang, you can practice with this code. It's short, understandable, and fun to play with. Just put print statements anywhere you like!

