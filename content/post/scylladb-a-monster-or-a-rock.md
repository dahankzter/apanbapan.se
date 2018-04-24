---
title: "Scylladb a Monster or a Rock"
date: 2018-04-24T08:36:11+02:00
draft: true
categories: ["Databases", "NoSQL", "Backend"]
tags: ["scylladb", "cassandra", "nosql"]
---

__This is a cross post from the [Eniro developer blog](http://developer.eniro.com/blog/post/scylladb-a-monster-or-a-rock/) that I wrote some time ago__

# ScyllaDB - A journey into the unknown...

Pretentious as that might sound it is not far from the truth.
We, me and my coworkers, have traditionally been using databases like [PostgreSQL](https://www.postgresql.org/) and [MongoDB](https://www.mongodb.com/) for most of our needs and we have had great success doing so. We have had no dramatic problems either with scalability or reliability. Neither are either storages particularly difficult to work with.
So why endeavor to try another database? A brand new one with an unproven track record to boot. Well, honestly, there was a fair bit of curiosity involved, at least initially. It was also motivated by the apparent performance benefit we potentially could gain.
I am personally a bit of a performance freak and I spend what is probably an inordinate amount of time on squeezing that last bit of latency or throughput out of a system. Premature optimization? Perhaps but I prefer to call it good engineering and high quality. We also have had issues scaling [PostgreSQL](https://www.postgresql.org/) laterally across a cluster of machines. There has always been the master to track and a little hands-on maintenance is almost always required. Don't get me wrong, [PostgreSQL](https://www.postgresql.org/) is a marvel of engineering and my absolute favorite when it comes to relational databases. It is stable, safe and performs very well in even very demanding scenarios.

## ScyllaDB - A Cassandra clone

A clone you say? Why yes it is indeed a clone in some sense. It is however not just a rewrite in Java but it has taken that step closer to the metal that is often needed to get the most performance out of a system that you can. There are no guarantees but at least the door to performance is opened and you can get to work on the algorithms more than catering to the whims of the JVM. Sure, it is a bit of flame-bait but there is some truth to it so lets stick to it. [Cassandra](http://cassandra.apache.org/) however has long seen heavy use in many a high profile site and is definitely one of those systems that inspire awe and respect in most developers minds. To set out with the express purpose of beating it on its home turf, performance and scalability, is nothing short of awesome! Time will tell whether or not the developer community will switch or if [Cassandra](http://cassandra.apache.org/) will accept the challenge and push on. Either way, competition is healthy and I for one welcome it.

## Our use case

Our use case was initially that we needed to relate a large set of transient identifiers A to another quite large set of other identifiers B. It may seem like a trivial thing to do and we thought so as well and started using [Redis](http://redis.io/) for this. [Redis](http://redis.io/) is known for its blazing performance and the use case seemed cut and dried. It soon became apparent though that we needed to extract all the identifiers A corresponding to given B which also seemed straightforward with [Redis](http://redis.io/) SCAN method. Happily we were satisfied and it ran fast and stable for some time. However, after a while we began seeing high latency in the system and started to investigate. It turned out that we had severely underestimated the number of transient identifiers and
we were struck rather bluntly by [Redis](http://redis.io/)'s single threaded nature. Agreed, this is by design and we were aware of how it worked. We just forgot some due diligence regarding our target flow of data. That being somewhat embarrassing we set out to correct the oversight. We tried to be clever with [Redis](http://redis.io/)'s SCAN trying to do it in steps but it quickly proved untenable since we seemed to hit too large latencies too quickly. Perhaps we could have gotten a bigger server but that felt like admitting defeat although it really is not. Sometimes it can be the right solution. We then started to think a little, occasionally it has been known to happen so we thought we would try it.
What about our old and faithful [PostgreSQL](https://www.postgresql.org/) and [MongoDB](https://www.mongodb.com/), could they be the answer? We concluded that they could possibly both be the answer but we were not entirely convinced of which to choose of the two or if it was worth the trouble to try. There would be some coding and setup, tuning indexes and and queries. Nothing too hard but not done very quickly either.
I had followed [ScyllaDB](http://www.scylladb.com/) since I first read about it in a pre 1.0 blog post some time ago and thought "What if we were to try [ScyllaDB](http://www.scylladb.com/)...?". It has the characteristics we want, low latency reads and writes as well as sharding and simple querying. After surprisingly little discussion we decided to try it.

## So, how are things going?

The bottom line is that we are very pleased with [ScyllaDB](http://www.scylladb.com/). The problem of listing the identifier mappings is completely gone. It is in fact a use case eminently suitable for [ScyllaDB](http://www.scylladb.com/). Normal simple lookups are also blazing fast and while [Redis](http://redis.io/) may be faster in the un-contended case a scaled scenario can be very different...? Benchmarking is fun but hard, perhaps there is time in the future for a test, who knows. Basically the only hitch we have hit so far is that during the installation we realized that our RHEL 7 setup had some as of yet not identified policies in place that got in the way. We found workarounds but so far no clear path to entirely friction free installs. The [ScyllaDB](http://www.scylladb.com/) repositories however worked nicely and the actual install of the packages was a breeze. We also bumped into a couple of already reported bugs in a couple of [ScyllaDB](http://www.scylladb.com/)'s setup scripts that afaik have been fixed at the time of this writing.
Getting our minds around the all so familiar SQL query language into the realm of CQL has been, not hard but gave us pause to think from time to time. There are just some limitations that you have to keep in mind when transitioning from a traditional relational query pattern to that of CQL.

### ScyllaDB in production

The database have been kicking nicely and we have expanded its usage to more than the initial scope with great ease. We are using the eminent [Gocql](http://gocql.github.io/) Go library and have had little to no issues so far and awesome performance. We have also upgraded [ScyllaDB](http://www.scylladb.com/) twice to version 1.2.1 which has been progressively easy and the second time it took very little time and without any problems. This is always a point of fear for me when it comes to databases. You always have to take a database dump and after upgrade reload it and hope that it works. Most of the times it works fine but I never feel really comfortable.
In all fairness [ScyllaDB](http://www.scylladb.com/) upgrade instructions advice taking a snapshot which we did but there was no explicit need to restore it after the upgrade which is very nice.

## Show me the money or "how fast is it"?

Well, it is fast. We have started to get used to seeing latencies in the microseconds as opposed to milliseconds that has usually been the standard. And this is not just the latencies reported by __nodetool__ but also from our API handlers after JSON serialization but before writing to the network.

|Percentile|SSTables|Writes (μs)|Reads (μs)|Partition (bytes)|
|----|----|----|----|----|
|50%|0.00|12.00|0.00|770|
|75%|0.00|17.00|0.00|770|
|95%|0.00|29.00|0.00|770|
|98%|0.00|35.00|0.00|924|
|99%|0.00|35.00|0.00|924|
|Min|0.00|2.00|0.00|104|
|Max|0.00|51012.00|0.00|1331|

As you can see there seems to be no immediate cause for concern although we probably should check that max value of ~50 ms.

Our API's are really fast as well and no caching has so far been needed. Both our standard REST endpoints and our [Nats](http://nats.io/) enabled microservices work really nice!

This brings an almost perverse satisfaction in the entire stack, [Go](https://golang.org/), [Echo](https://github.com/labstack/echo), [Nats](http://nats.io/) and of course [ScyllaDB](http://www.scylladb.com/) that makes me doubt a large chunk of my previous experience using other stacks. It should be noted also that no extra tuning and tweaking has been done and so far we see no reason to do it other than possibly for fun!

## So what about the rock or monster thing?##
It was my perhaps farfetched attempt at being witty. Recalling, somewhat vaguely, the greek mythology of [Scylla and Charybdis](https://en.wikipedia.org/wiki/Between_Scylla_and_Charybdis) we can now get at a meaning. On the one hand we have the rock, on the other hand the monster and we have to navigate in the middle. This is where [ScyllaDB](http://www.scylladb.com/) comes in. It's both the monster of a database that you envision as well as the big rock to lean on when it is windy! :)

*Happy hacking!*

*Henrik*
