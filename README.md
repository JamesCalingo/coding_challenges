# Basic Server and Load Tester

Based on John Crickett's Coding Challenges, I've decided to try one of them out - well, actually two of them.

The first is a basic web server. Technically, this isn't COMPLETELY new as I've seen these before, 

# CHAPTER 1: The Server

## Step 0: The Language

I've chosen Go for this exercise. This is partially due to the fact that I've been exploring Go over the past few months, and also because this is based on a backend server I've been working on for another project. As it turns out

Here, I'd like to digress a TEENSY bit into file systems. Since I wanted to try and learn how Go projects are 

I also wanted to have multiple things 

## Step 1: Basic HTTP

HTTP is, as should be common knowledge if you're in software, HyperText Transfer Protocol - the standard by which most data is transferred on the internet. 

The most common HTTP request is the GET request. With this request, a client asks some remote server for some data, and said server gives some form of response. There's a LOT that goes into both the request and the response, but we'll get into it a bit later. For now, let's start by having our server serve up a good ol' "Hello, World" when it receives a GET request - because it's not like we have any better messages to send. 

Go's `net` module has a package for HTTP handling - aptly known as `http` (in fact, some may argue that this makes this project in Go _way too easy_). To start our server project, we establish a handler function (as the 

I did this as a "submodule" and wasn't super sure how to handle it since I couldn't call it in `main` without an explicit return value, so I basically did the ol' C trick - return 0 if it's successful or 1 if something happens.

I should note that I did a lot of this part in GitHub Codespaces, so there was a bit of wonkiness when trying to start the server - it was able to serve what I had told it to, but couldn't handle any routes for whatever reason. However, it worked fine when I downloiaded the code onto my computer and ran it locally (though the name "localhost" should have made that obvious - especially since CodeSpaces isn't exactly "local").

## Step 2: HTM Losche!

Now that we can serve something up, let's make things a bit more interesting by making said thing a bit more complex: An HTML document.

In case you didn't know, you're reading an HTML document right now! 

As for how we can create a

The most important thing to remember here is to go from the ROOT PATH. The actual server definition is "buried" in somme folders, but the app is running from the root (main.go), so I think we need to make sure

## Step 3: Concurrency

Here's where the decision to use Go for this really begins to get interesting.

One of the things Go is most well known for is its concurrency paradigms using things known as channels and Goroutines.

Web servers are not like the checkout line at the local store: they need to be able to handle multiple clients asking them for information all at once.

## Step 4: Wildcards

There's one thing we should shore up for our server before we call it "complete for now": client requests for information we don't have (i.e. the ol' 404 not found).

Go's http module has features that are, to put it bluntly, seemingly pessimistic. For example, the ListenAndServe function we've been using returns an error - but it only returns if said error occurs.

# Chapter 2: Load Testing

We now have a web server that can parse requests from clients and serve them various pieces of data. However, it's important that we're not only able to handle any traffic on our server, but handle it *well*; for example, we wouldn't want to, say, run into massive issues should our server be called upon by [EVERY SINGLE GAMESTOP IN THE UNITED STATES](https://www.polygon.com/2015/4/2/8337499/gamestops-website-down-amiibo-ness).

## Step 0: The language

Let's not get weird here and stick with Go. 
