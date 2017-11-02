---
layout: post
title:  "GoLang Templating Made Easy"
date:   2017-11-01
categories:
---

Golang provides the package [text/template](https://golang.org/pkg/text/template/) for generating textual output from an object.
I will go over a few of the things I've learned while using it.

## Preface ##

The past few months I have been contributing the the Kubernetes community, specifically to [ingress-nginx](https://github.com/kubernetes/ingress-nginx).

While working on this project, I added functionality to alter specific server attributes via annotations. When adding
these features I had to make alterations to several components and create new template functions.

There are 3 Examples below which build on each other. Each contain a `zoo.go` and `zoo.tmpl` file. `zoo.go` contains
all the logic and the `zoo.tmpl` contains the template we are using.

## Creating a Template from an Object ##

An object or range of objects can be used with a template to create a textual output. Pretty much its as simple as a template
recognizing the object being passed in and drawing attributes from the object to create wanted text. It is very useful for
reports or server configurations.

[Here](https://github.com/diazjf/diazjf.github.io/tree/master/resources/sample-code/zoo/example-1) is an example in which I created a ZooKeeper application
which generates a report based off of different zoos and what animals they contain.

## Creating a Template Function ##

Template functions are very useful for processing Objects passed within the template. Suppose the zoos only wanted to
allow animals which are suitable for the climate in the area of the zoo's location. We can create a function to sort
which animals belong in the zoo, so we can generate a template of the acceptable animals.

[Here](https://github.com/diazjf/diazjf.github.io/tree/master/resources/sample-code/zoo/example-2) is an example I have created of how to use a function within a
template. The function run in the template accepts `[]Animal` and a string `Zoo.Climate` and returns a list of Animals which
`Animals[i].Climates` matches `Zoo.Climate`. It can be seen in the FunctionMap with a given name defining it on the template.

## Using a Sub-Template ##

A sub-template would be a template that is loaded within another template. This is useful because, we can move all
common code to another template and reduce duplication of code if the sub-template is used several times in the main template.

[Here](https://github.com/diazjf/diazjf.github.io/tree/master/resources/sample-code/zoo/example-3) is an example I have created of how to use a sub-template.
You can see that instead of writing the for-loops for both acceptable and unacceptable animals separately, we only write it once.

Note: Sub-Templates can be in other `.tmpl` files. You will just need to pass those files into the `ParseFiles()` function
seen in the examples.

## References ##

- [Related Patch I've created for nginx-ingress](https://github.com/kubernetes/ingress-nginx/pull/1123)
- [My Examples](https://github.com/diazjf/diazjf.github.io/tree/master/resources/sample-code/zoo)
