# Basic Server
A few coding challenge type stuffs

# Step 0: The Language

I've chosen Go for this exercise. This is partially due to the fact that I've been exploring Go over the past few months, and also because this is based on a backend server I've been working on for another project.

Here, I'd like to digress a TEENSY bit into file systems. Since I wanted 

# Step 1: Basic HTTP

HTTP is, as should be common knowledge if you're in software, HyperText Transfer Protocol - the standard by which most data is transferred on the internet. 

The most common HTTP request is the GET request. With this request, a client asks some remote server for some data, and said server gives some form of response. There's a LOT that goes into both the request and the response, but we'll get into it a bit later. For now, let's start by having our server serve up a good ol' "Hello, World" when it receives a GET request - because it's not like we have any better messages to send. 