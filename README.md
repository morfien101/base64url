# base64url
A simple Golang based tool that will encode to base64 URL format from a shell

## What does it do?
It runs the base64.URLEncoding.EncodeToString function in go's encoding/base64 package in the command line argument that you give it.

## Why did I make it?
I could not find an easy way in Linux to do base64 URL encoding that wouldn't want me to write a shell script of some kind that I didn't understand. I was using it with curl.

## How do you use this?
Download the binary for either Windows or Linux (Mac can build from source if you want.). Put it in a folder that is in your path.

Then use it with:

```
base64url "This message to be encoded."
```
