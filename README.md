# Basic Server and Load Tester

Based on John Crickett's Coding Challenges, I've decided to try one of them out - well, actually two of them.

# CHAPTER 1: The Server

## Step 0: The Language

I've chosen Go for this exercise. This is partially due to the fact that I've been exploring Go over the past few months, and also because this is based on a backend server I've been working on for another project.

## Step 1: Basic HTTP

HTTP is, as should be common knowledge if you're in software, HyperText Transfer Protocol - the standard by which most data is transferred on the internet.

The most common HTTP request is the GET request. With this request, a client asks some remote server for some data, and said server gives some form of response. There's a LOT that goes into both the request and the response, but we'll get into it a bit later. For now, let's start by having our server serve up a good ol' "Hello, World" when it receives a GET request - because it's not like we have any better messages to send. 

Go's `net` module has a package for HTTP handling - aptly known as `http` (in fact, some may argue that this makes this project in Go _way too easy_). To start our server project, we need to establish a handler function, as that's how `http` works in Go: establish a handler that takes a request and writes it to some writer (currently, standard I/O). Once that's been done, we have to actually serve the thing, and in Go, we do that by using `http.ListenAndServe()`.

I should note that I did this as a "submodule" and wasn't super sure how to handle it since I couldn't call it in `main` without an explicit return value, so I basically did the ol' C trick - return 0 if it's successful or 1 if something happens. Also, I did a lot of this part in GitHub Codespaces, so there was a bit of wonkiness when trying to start the server - it was able to serve what I had told it to, but couldn't handle any routes for whatever reason. However, it worked fine when I downloiaded the code onto my computer and ran it locally (though the name "localhost" should have made that obvious - especially since CodeSpaces isn't exactly "local").

## Step 2: HTM Losche!

Now that we can serve something up, let's make things a bit more interesting by making said thing a bit more complex: An HTML document.

In case you didn't know, you're reading an HTML document right now! The internet is basically a bunch of these HTML documents, and the http requests we make are for our computer to find and fetch these documents from various servers.

As for how we can create a way to serve said HTML, Go's http package has a function that returns a handler: `FileServer`.

The most important thing to remember here is to go from the ROOT PATH. The actual server definition is "buried" in somme folders, but the app is running from the root (main.go), so I think we need to make sure that the pathing is correct or else we'll end up always getting a 404 not found error.

Also, while running thses in the browser is probably the more "common" way to view what the server is doing, it's not the only way: you can use the `curl` command in a Terminal shell. I note this here because if you try to GET the HTML doc in Terminal, you get the actual HTML with tags and everything; in the browser, you get the HTML translated into a webpage (there's a TON that goes into that whole process, but we're not going to get into that today).

We're not QUITE done here, as there's another important consideration we have to take into account:

## Step 4: Wildcards
(Yes I know that this is out of order. I'll explain later.)

There's one thing we should shore up for our server: client requests for information we don't have (i.e. the ol' 404 not found).

In step one, we were able to serve up whatever thanks to our URL path. However, we changed a few things in step 2 to serve an HTML document. 

The generic response I got from the server when trying to find nonextant things was a simple "404 page not found", which is fine for now. However, I personally consider that a bit weak especially for a website; I'd like to have a special page for any bad requests like this.

I made a 404 page, and we're actually able to access it if it's in our public directory by simply requesting `localhost:8080/404.html` (I used 8080 for this). However, what we need is to redirect any "bad" requests to this page; for example, if we tried to get `localhost:8080/random_page_generator`, we should be rerouted to `404.html`. To do this part, Go

## Step 3: Concurrency

You may notice that this section is technically BEFORE the section on wildcard routing in the coding challenges website. However, I swapped them as I ended up doing the wildcard routing first with HTML.

__\*COMING SOON*__

<!-- Here's where the decision to use Go for this really begins to get interesting.

One of the things Go is most well known for is its concurrency paradigms using things known as channels and Goroutines.

Web servers are not like the checkout line at the local store: they need to be able to handle multiple clients asking them for information all at once. -->

<!-- # Chapter 2: Load Testing -->

<!-- We now have a web server that can parse requests from clients and serve them various pieces of data. However, it's important that we're not only able to handle any traffic on our server, but handle it *well*; for example, we wouldn't want to, say, run into massive issues should our server be called upon by [EVERY SINGLE GAMESTOP IN THE UNITED STATES](https://www.polygon.com/2015/4/2/8337499/gamestops-website-down-amiibo-ness).

## Step 0: The language

Let's not get weird here and stick with Go. 

## Step 1

I have no clue what this is, but  -->