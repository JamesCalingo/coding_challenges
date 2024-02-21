# Basic Server
A few coding challenge type stuffs

# CHAPTER 1: The Server Itself

## Step 0: The Language

I've chosen Go for this exercise. This is partially due to the fact that I've been exploring Go over the past few months, and also because this is based on a backend server I've been working on for another project.

Here, I'd like to digress a TEENSY bit into file systems. Since I wanted to try and learn how Go projects are 

I also wanted to have multiple things 

## Step 1: Basic HTTP

HTTP is, as should be common knowledge if you're in software, HyperText Transfer Protocol - the standard by which most data is transferred on the internet. 

The most common HTTP request is the GET request. With this request, a client asks some remote server for some data, and said server gives some form of response. There's a LOT that goes into both the request and the response, but we'll get into it a bit later. For now, let's start by having our server serve up a good ol' "Hello, World" when it receives a GET request - because it's not like we have any better messages to send. 

Go's `net` module has a package for HTTP handling - aptly known as `http`. To start our server project, we establish a handler function (as the 

I did this as a "submodule" and wasn't super sure how to handle it since I couldn't call it in `main` without an explicit return value, so I basically did the ol' C trick - return 0 if it's successful or 1 if something happens.

I should note that I did this part in GitHub Codespaces, so there was a bit of wonkiness when trying to start the server - it was able to serve what I had told it to, but couldn't handle any routes for whatever reason.

## Step 2: HTM Losche!

Now that we can serve something up, let's make things a bit more interesting by making said thing a bit more complex: An HTML document.

In case you didn't know, you're reading an HTML document right now! 

## Step 3: Concurrency

Here's where the decision to use Go for this really begins to get interesting.

One of the things Go is most well known for is its concurrency paradigms using things known as channels and Goroutines.

Web servers are not like the checkout line at the local store: they need to be able to handle multiple clients asking them for information all at once.

## Step 4: Wildcards

There's one thing we should shore up for our server before we call it "complete for now": client requests for information we don't have (i.e. the ol' 404 not found).

Go's http module has features that are, to put it bluntly, seemingly pessimistic. For example, the ListenAndServe function we've been using returns an error - but it only returns if said error occurs.

# Chapter 2: Load Testing

We now have a web server that can parse requests from clients and serve them various pieces of data. However, it's important that we're not only able to handle any traffic on our server, but handle it *well*; for example, we wouldn't want to, say, run into massive issues should our server be called upon by [EVERY SINGLE GAMESTOP IN THE UNITED STATES](https://www.polygon.com/2015/4/2/8337499/gamestops-website-down-amiibo-ness).
