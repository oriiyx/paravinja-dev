---
Title: pursuit of software simplicity
Summary: Stung by over-complex systems, my own take on how important it is to pursuit the simplicity in design. 
Author: Peter Paravinja
Published: 14. Aug 2024
---

Simplicity.

There is something beautiful to my eyes when I see something designed with simplistic approach.

Simplicity and minimalism are usually intertwined with each other. Simple life is minimal. Minimal design contains
elements of simplest forms.

I enjoy minimalistic approach in a lot of aspects of my life. Minimalist design or approach to a design to me is
something which has the fewest and simplest elements used to create the effect you're pursuing.

But how did I get this appreciation for something simple?

When I first started out with programming I figured that 10x engineers combat complex problems and create these
astounding systems that are resilient and provide good user/developer experience. I wanted to work on complex projects.

Well... as Doctor Jagger once said:

> you can't always get what you want... but if you try sometimes, you might find you get what you need

![image](/public/assets/aww-yeah.png)

_aww yeah_

Thus, I was reassigned to a big project inside the company that was having very difficult times. Reassigned is a
loosey-goosey word... I was thrown into a tsunami with a memo "have fun".

It shocked me how big and intertwined projects can become.

Even more of a shocker was how big the projects can become without having any proper tooling around them... People were
deploying via FTP, but locally you had git instance where you had to track everything - but production wasn't synced
with git - it was just for local development? What?

You had to run the project in production environment otherwise it would not start.

You could not run almost ANY of the commands straight out of the box since you could make changes on production servers
with the commands.

Staging environment? Okay, relax there buddy.

Developers responsible for this mess? Working on other projects, side eyeing while being 'busy' with other customers.

Customer happiness? Shouting at me on meetings and threatening with lawsuits level. Very nice.

The codebase? Absolute disaster. Bugs everywhere - weekly sprints mostly filled with bug fixes.

Technologies used? Badly researched inside the company with bad coding practices and non-idiomatic approach all over the
place. But overall it was a "good" choice for the project.

This type of system was complex. Very complex. You had to think and somehow test 10 different things when you made one
change. You had 0 programmatic tests to help you with testing. You couldn't run some tests due to possibility of running
production
services.

We had to create a plan of steps that would reduce the complexity of the entire project.

Debugging was hard - abstracting method names and calling abstract $set and $get methods without prior checking if
method exists on an object was nightmarish while function jumping while functions just had another function call without
1 param(why tho? Lazy?).

I implemented a lot of additional logging to at least understand where we're shooting ourselves in the foot in the
production.

All those experiences shaped my mind that complex is not something that you really want as a programmer. You really want
as simple as you can get and plan for it before you introduce complexity. Because once you introduced complexity, it's
really hard to restore back simplicity. I often read that people rewrite the entire system rather than to refactor the
existing system. Same thing happened when an internal team was formed, and they successfully sold (to client) that
rewrite would solve their and our problems.

Fast-forward a few years, and we're currently inside the "javascript boom". Suddenly I was surrounded with new
developers that had a completely different mindset. I am currently working in an agency that develops custom solutions
for our customers based on their needs. And suddenly developer experience is the most important thing in the world. And
I get it. But the business needs are also important. And sometimes you have to make a compromise and embrace the "
sucks".

Maybe this is because my position requires me to talk to the customer and understand their needs. I usually have to
break down their issues down to simple things in order for both parties to understand the "real" issues we're trying to
solve.

![image](/public/assets/raptor-comparison.png)

_raptor engines comparison_

I saw the image that Spacex posted that depicted how they simplified Raptors engine design.
What my eyes see is ability to innovated so darn much that they've managed to simplify something as complex as a rocket
engine. Amazing! I love this timeline!

Now lets face it... most of us are not lucky enough to work on something as cool as a rocket engine. But I think that
the innovation and simplification can be applied to almost anything. Programming in general has already tackled and
achieved a lot of simplifications.

But I think the simplification has swayed in the wrong direction.

If I take Next.js as an example. It's a framework that abstracts huge amounts of complexity and gives you a simple
interface and APIs to work with.

I created a simple dashboard with CRUD operations on top of my database in a few hours. Vercel deployment was amazingly
simple.

But... I had a moment when I realized I have 0 clue what is actually happening under the hood. I had a button that
created a new entity - I added all the data in the form and submitted. Yet when I redirected after success, data was not
there.
Next.js was caching some data and I had to invalidate the cache manually.

Of course, it's skill issues - you have to read the manual. But it got me wondering. How much abstraction is too much?
Are you simplifying things for me or are you creating something more complex for me?

My code does not indicate that something is being cached.
If I onboard a new team member, they would actually have to read the manual to understand what is happening.

Does only version 13 Next.js do this? What about version 12? What about version 14? Do I need version specific
knowledge? What the hell.

Simplicity in the software for me is the ability to understand what is happening when I read the code.

If I look at Next.js:

- There is javascript factor where javascript has some magic to it built into the language
- There is typescript factor
- There is react factor
- There is database factor
- There is prisma factor to for abstracting the database
- There is Next.js factor
- There is Vercel factor

It's a lot of factors to consider when you're trying to understand.

I have a feeling that we're using library on top of library on top of library and writing 1% of the actual logic code in
the end... it's not even mid maxing at this point.

I understand that we want great UI and UX, and we want to build things fast.
I read a lot about HTMX and how it's a game changer in terms that you can do 90% of reactivity on the frontend with 5%
of complexity in comparison of introducing a framework like react.
Think about what you're building and what you're trying to achieve. Do you really need a framework for that? Do you
realize how many thing you're introducing with that framework?
HTMX is definitely next on my list to try out.

Ok. We get it. But...

![image](/public/assets/this-is-fine.png)

_this is fine_

So what is the point of this post?

I think that we should strive for simplicity in our software.

My pursuit lead me to C + Go programming language. While C is very low level and you have a lot of shoot yourself in the
foot features, it's very simple.
You start to understand how things work under the hood.

Go is a bit higher level than C, but it's still very simple - Rob and the Go team really did an awesome job with the
language design. The standard library is extremely wonderful and provides you with so many tools to craft the software
you're building. If you compare it with javascript where you "need" the ecosystem to match these features - and that
ecosystem is broken into many pieces and each piece has its own way of doing things... it's almost like comparing jungle
with a garden.

Go is incredibly fast to build things with. The standard library supports 3/4 of the use cases you will need to build
things with. I am not talking about use case where something will cache things for you without you knowing it use case.
I am talking about building things you have control over use case.
This blog platform was written in a day and deployed in an hour with GitHub, Coolify and Hetzner - really simple and I
have good control over it - only Coolify is a blackbox for me - and I reduced the "blackbox"-ines with learning a little
about how Caddy proxy works inside Coolify.

I think the simplicity is something that you can understand with minimal layers of abstractions and minimal complexity.
You add another system? You're adding complexity and removing simplicity in my mind. You have to understand how that
system works and how it interacts with other systems. When you have a net of systems communicating with each other, you
have a complex system, and you need a team of experts to maintain it.

Now there are situations where you need to face this complexity. There are situations where you just need the fastest
things to prototype something and release it fast to test the market.

It all "depends" in the end... but I think that we generally abandoned simplicity in the software for the sake of
using "cool" things and frameworks without understanding the implications or reasons why we've created them in the first
place.

