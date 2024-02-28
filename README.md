# Basic Server

Based on John Crickett's Coding Challenges, I've decided to try one of them out - well, actually two of them.

# PART 1: The Server

This ended up being way simpler than I thought it'd be, though that may be becuase I chose Go for this exercise. This is partially due to the fact that I've been exploring Go over the past few months, and also because this is based on a backend server I've been working on for another project.

To start, I made sure we could actually take care of HTTP requests. Go's `net` module has a package for HTTP handling - aptly known as `http` (in fact, some may argue that this makes this project in Go _way too easy_). To start our server project, we need to establish a handler function, as that's how `http` works in Go: establish a handler that takes a request and writes it to some writer (currently, standard I/O). Once that's been done, we have to actually serve the thing, and in Go, we do that by using `http.ListenAndServe()`. I'm using port :8080 for this as it's one of the "common" ports used for things like these (when I was in a coding bootcamp, we often used :3000 and :3001 for full stack projects).

I should note that I did this as a "submodule" and wasn't super sure how to handle it since I couldn't call it in `main` without an explicit return value, so I basically did the ol' C trick - return 0 if it's successful or 1 if something happens. Also, I did a lot of this part in GitHub Codespaces, so there was a bit of wonkiness when trying to start the server - it was able to serve what I had told it to, but couldn't handle any routes for whatever reason. However, it worked fine when I downloiaded the code onto my computer and ran it locally (though the name "localhost" should have made that obvious - especially since CodeSpaces isn't exactly "local").

Once I knew that the server can serve something up, I wanted to serve something slightly more complex: An HTML document. Go's http package has a function that returns a handler: `http.FileServer`. I set the handler to serve a folder I called "public" (borrowing from the world of JS frontend library React), stuck an `index.html` file in there, and then used `http.FileServer` and `http.Dir()` to serve the public folder.

The most important thing to remember here is to go from the ROOT PATH. The actual server definition is "buried" in somme folders, but the app is running from the root (main.go), so I think we need to make sure that the pathing is correct or else we'll end up always getting a 404 not found error.

Also, while running thses in the browser is probably the more "common" way to view what the server is doing, it's not the only way: you can use the `curl` command in a Terminal shell. I note this here because if you try to GET the HTML doc in Terminal, you get the actual HTML with tags and everything; in the browser, you get the HTML translated into a webpage (there's a TON that goes into that whole process, but we're not going to get into that today).

We're not QUITE done here, as there's another important consideration we have to take into account:

Wildcards.

With the simple server, we can actually serve a fair amount of "wildcard" requests. For example, a route such as "localhost:8080/this-is-not-the-droid-that-you-are-looking-for" will actually yield the request route and serve a result. However, we cannot do that with HTML documents.

The generic response I got from the server when trying to find nonextant things in the HTML server was a simple "404 page not found" message, which is, in all honesty, fine. However, I'd like to have a special page for any bad requests like this; IMO, it's a better signal to the end user that they made a bad request.

I made a 404 page, and we're actually able to access it if it's in our public directory by simply requesting `localhost:8080/404.html`. However, what we need is to redirect any "bad" requests to this page; for example, if we tried to get `localhost:8080/random_page_generator` and it didn't exist, we should be rerouted to `404.html`. To do this part, I ended up deep in the research "woods" before realizing that 

Here's where the decision to use Go for this really begins to get interesting: instead of making one server, I tried to write two different servers - one for frontend and one for backend. Obviously, I used different ports for them (:3000 for the HTML pages and :8080 for the backend stuff), but alas, only one can run at a time...or can it?

I know that concurrency is a thing; I used it for a monorepo project WAY back in the day. However, I was, to put it bluntly, kinda stupid back then and probably copied it from somewhere without understanding what was actually going on. However, one of the things Go is most well known for is its concurrency paradigms, so I figured that there MIGHT be a way to have both run at the same time. As it turns out, Goroutines are our friend here.

A Goroutine is, in short, a singular programming thread within Go. Thanks to Goroutines, I was able to get both ports routed to the same server...but that's not quite what I had in mind. However, in order for the two ports to serve different things (i.e. frontend and backend), I did a bit more digging.

<!-- # Part 2: The Load Tester -->

 Originally, I intended to have both a web server and an HTTP Load Tester in this repository. However, it ended up not working out in practice, so I decided to split it into its own project/repository - which you can view here:

# [LOAD TESTER](https://github.com/JamesCalingo/go_loadtester)
