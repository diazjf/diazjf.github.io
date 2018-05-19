---
layout: post
title:  "Make Ingress-Nginx Work for you, and the Community"
date:   2018-05-19
categories:
---

I recently presented at KubeCon Europe 2018 in beautiful Copenhagen on developing for Kubernetes Ingress-Nginx. You can checkout the YouTube video ["Make Ingress-Nginx Work for you, and the Community"](https://youtu.be/GDm-7BlmPPg).

I wanted to talk a little about the session, and describe some of the steps in writing.

## Preface ##

In the session we went over how the ingress-controller works, from all the internals,
to the template. We added a simple annotation, built and deployed the ingress controller.

In this post, I will go over so of the steps I went over in the video linked above. You can
follow the steps below and examine the video to see the actual code.

## Adding a Simple Annotation ##

There are a few changes involves in creating an annotation. I will go over how
to create an annotation to add a "nickname" to a group of servers.

1. In `internal/ingress/types.go` we add a Nickname type as part of a Server.
We will be scoping a nickname to a server, meaning that all servers in a particular
ingress resource will have the same nickname.

2. Create a folder `internal/ingress/annotations/nickname` and add `main.go`.
`main.go` contains all the logic involving parsing and processing the annotation.
Since we are directly setting the value, there is really only parsing involved.
We add the `Nickname` to the annotation extractor seen `internal/ingress/annotations/annotations.go`

3. In `internal/ingress/controller/controller.go` we add the Nickname from the
annotation when each server is created/edited.

4. In the `rootfs/etc/nginx/template/nginx.tmpl`, we add a conditional, which states
that if the server has a nickname, we can generate a comment in `nginx.conf` giving
the nickname of the server.

## Adding a ConfigMap Change ##

ConfigMaps apply to a complete configuration. We will be adding a value, that
when applied as "true", will add a `/hello` location to the default server.

All we have to do is:

1. Add the config-map variable in `internal/ingress/controller/config/config.go`.
This allows the Ingress Controller to recognize a key in the config-map.

2. Add the conditional to `rootfs/etc/nginx/template/nginx.tmpl`. This uses the value
in the config-map if a particular key is found.

## References ##

- [List of Commands for Building, Testing, Etc. ](https://github.com/diazjf/diazjf.github.io/tree/master/resources/sample-code/kubecon-europe-2018)
