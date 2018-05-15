---
layout: post
title:  "Kubernetes Nginx Mergeable Types"
date:   2018-05-15
categories:
---

Recently with the help of [pleshakov](https://github.com/pleshakov), I added a feature to [nginxinc/kubernetes-ingress](https://github.com/nginxinc/kubernetes-ingress)
to allow for multiple ingress resources with the same host. These multiple ingress
resources "merge" in order to create a single configuration. I will go over the
newly added feature and its significance to the NGINX Kubernetes Ingress Controller.

## Preface ##

In nginxinc/kubernetes-ingress when creating 2 or more Ingress Resource containing
the same host, a Conflict would arise. You would get an warning stating that the
conflicting server name has been ignored.

This would mean that for a particular host you would have to add all annotations
and rules to a single Ingress Resource. This can make configuring and managing a
single host a nightmare, depending on how many different teams are working on the same resource.

Hence Mergeable Types was recently introduced to nginxinc/kubernetes-ingress
to allow the Merging of Ingress rules with the same host into a single configuration.
It provides many benefits in terms of organization.

## Rules and Annotation ##

Mergeable Types are separated into 2 types. There are Masters and there are Minions.

Master: will process all configurations at the host level, which includes the TLS configuration, and any annotations which will be applied for the complete host.

Minion: will be used to append different locations to an ingress resource with the Master value. It also allows for the locations to be processed with certain annotations as well.

In order to enable Mergeable Types, its as simple as adding the `nginx.org/mergeable-ingress-type` annotation with the value `master` to setup a master. Then adding the same annotation with the value `minion` to setup each minion.

Nginx will process the master and its minions, and generate a single configuration file containing all the relevant information.

## Examples ##

Examples can be found in [kubernetes-ingress/examples/mergeable-ingress-types/](https://github.com/nginxinc/kubernetes-ingress/tree/master/examples/mergeable-ingress-types).

## Other ##

This feature is available in [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx),
however it works in a different way.

Annotations apply to each path presented in a host. When another Ingress Resource with the same
host is created, if it contains the same annotations and there are any duplicate paths, then the
annotations are not applied to those paths, but are applied to all other paths. Similarly in host level
annotations, the first annotation applied takes precedence.
