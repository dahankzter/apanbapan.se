+++
title = "Scylladb - a New Gig"
date = 2018-06-17T10:12:32+02:00
description = "I have a new gig and it's great!"
draft = true
toc = false
categories = ["work"]
tags = ["scylladb", "social"]
+++

A couple of months ago I joined [ScyllaDB](https://www.scylladb.com/).

The idea came to me after seeing a tweet from [Scyllas official Twitter](https://twitter.com/ScyllaDB) 
stating that backend engineers were sought after. Having done little else for almost a couple of decades
and also really loving the database and it's core philosophies I thought, __Why not?__. At first though,
I was very hesitant... Not that I really doubted my abilities, I have been around too long to not know
what I can and can not do. Nevertheless, this is a company not just standing on the shoulders of giants
but some of the them are those actual giants that the saying refers to. For example,
[Avi Kiviti](https://twitter.com/AviKivity), our CTO, created [KVM](https://www.linux-kvm.org/page/Main_Page)
together with the [RedHat](https://www.redhat.com/en/topics/virtualization/what-is-KVM) team led by our
CEO [Dor Laor](https://twitter.com/DorLaor) which is afaik the worlds most used virtualization ever since
[Amazon decided to use KVM](https://www.theregister.co.uk/2017/11/07/aws_writes_new_kvm_based_hypervisor_to_make_its_cloud_go_faster/).
I have no stats to support this of course but KVM is clearly software that is quite heavily in use to say the least...

I had met some of the team before on [Scylla Summit 2016](https://www.scylladb.com/2016/08/09/scylla-summit-highlights/)
in San Jose and it seemed to be a very mellow and cool set of people from several parts of the world.
I thus knew that most likely a company spearheaded by these people would be a basically decent place with
fun people. This has definitely turned out to be true. ScyllaDB is a company whose spirit is truly formed
in the shape of it's people. You might wonder how this is any different from other companies and you are
probably right that this is the case everywhere and I just happen to like this particular spirit...
There is anyway something special here and without having analyzed it deeper I think it has to do with
the level of genuine trust that management places in people combined with great autonomy and responsibility.
Many companies are filled with good people but there are always pretty comprehensive control mechanisms
in place that makes you feel like slogans such as _we trust each other_ feels a little hollow. That's not
the case with Scylla, I truly feel that people trust eachother and what enables it is the inherent culture
of the company, a culture that I am sure comes, at least in part, from the open source background of the founders.
Another aspect is that most people in most roles seems to have great technical knowledge. This is perhaps
not so surprising in a very high tech company but still it is very refreshing to me. 
Product owner [Tzach Livyatan](https://twitter.com/TzachL) for example can freely move from advanced tech
discussions to UX onto marketing without breaking stride which is truly not common. My closest boss
[Shlomi Livne](https://twitter.com/ShlomiLivne) also possesses super high tech skills and I don't have to
explain anything really when we discuss something. Our Principal Architect [Glauber Costa](https://twitter.com/glcst)
is one such other very colorful person who has this scary and often, by developers, ridiculed enterpricy title
but I am pretty sure that the person who once coined it had Glauber in his mind. He is for sure the most
hands down and capable programmer there is who can still take in and help steer the big picture.
Just peek at this post about [how Scylla handles compactions](https://www.scylladb.com/2018/06/12/scylla-leverages-control-theory/),
it's a quite amazing piece and you will knowwhat I from now will expect when someone has _architect_
in their title. It has math and data to back it up just as it should in cases like the case but are so often missing.

It is basically the way it should be in any company where there is complex work to be done. Your administrative
 superiors should be at least at the same or comparable technical level. The alternative to this is for managers
to only do simple administrative formal things like vacations or working hours and what not. 
This inevitably results in a command and control structure and calling the bits and pieces all the names from
the agile playbook is not going to change that. This topic deserves a whole post in and of itself so I will
leave it at that for now.

The perhaps biggest change daily has been the fact that I am now working remote full time for the first time
in my career. It has been a change to which I am not yet fully acclimatized. The plan has been to work from
home mostly with a day here and there in a cafe or a one of the startup incubator floors where a lot of
usually very young and unspoilt enginners and hackers sit and work ungodly hours... We will see where I 
land but I will probably take the advice of my closest work colleague [Michal Matczuk](https://twitter.com/michalmatczuk)
and take a couple of mornings for exercise since even the basic _run to the subway_ routine is decimated.

So what will I actually be doing? I will, at least initially, work on the [Scylla Manager](http://docs.scylladb.com/operating-scylla/manager/)
together with Michal. It is basically a server written in [Go](https://golang.org/) that manages your
Scylla clusters and performs various routine tasks at regular intervals. It has a [Swagger](https://swagger.io/)
generated REST API and there is a command line client that wraps the most common operations for you.
It's scope is planned to grow but we will grow it organically from customer feedback as well as other informed
sources not to mention all our own Scylla developers who live and breathe Scylla daily to end users of our
open source version that developers are at an increasing reate picking up. Your input is valuable and will
serve as a guide for what we implement.

Anyway, this became a post not simply about me starting a new gig but also a bit about a lot of things I will
need to dig into later... As usual, writing begets writing!