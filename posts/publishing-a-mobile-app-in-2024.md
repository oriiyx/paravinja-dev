---
Title: publishing a mobile app in 2024
Summary: First hand experience trying to publish a mobile application in 2024 / 2025.
Author: Peter Paravinja
---

Let me start by acknowledging that I completely understand that Google and Apple want to create a nice marketplace where
it's not over-saturated with shit apps that developers spent 10-20 hours on and just want to release them quick without
any thought if they're actually useful for anyone.

Shout-out to all the calculator & todo tutorials out there.

But seriously... you as a developer can't expect someone to actually host that shit on their platform...

For more context let me describe my own application that I'm currently in process of publishing. For clarity, I think
Apple or Google are not suggesting I didn't put in the effort to make an application that users can have some sort of
benefit. I just want to plug in some of my own work tbh.

## Booktime

Booktime is a React Native application that is very similar to Strava - but for book readers and book clubs if you can
imagine. You track your readings, you can interact with other people socially. Ultimately I hope that users will have
additional motivation to read by interacting with their own circle of followers and friends.

Stack is:

- React Native for mobile app
- Golang for backend
- Postgres for database
- Redis for caching
- Remix for frontend
- AWS for email service

Ultimate goal of this application is for me learn from start to finish - how it is to launch a mobile application and
create some type of a product.

## Expectations

While planning this thing - my ultimate goal was to get some deeper software skills and some good practices - get
familiar with a new backend language Golang - get familiar with Remix - get familiar with Redis - get familiar with AWS.
I thought those would be the 'hard parts'...
It somehow managed to slip my mind that the software part would be more natural to me and closer to what I wanted.
Making the machine do what I wanted was somehow very easy compared to dealing with Google and Apple.

Now that I understand the process a little better I know that I underestimated the Google and Apple part by a mile... I
thought that there would be a process where I would have to somehow compile the react native application to some binary
that Google and Apple need in order to run the application on the hardware. Completely understandable.

That wasn't the issue...

## Google

After starting the research on how to publish an application on the Google Store I found their official documentation,
First I discovered play.google.com where I had to register a developer account. It cost me like 30€ ($$ for American
friends). This was normal - one time payment - something to show that you're at least a little dedicated and prepared to
invest into the program. Everything goes on pretty smoothly must say.

I create follow their official post where they guide you how to create a new Application.
A lot of information.
Some questions are better understood than others... Hopefully I didn't mess anything up.

Then comes the part where they need ME TO GET 12 TESTERS TO TEST THE APPLICATION FOR 14 DAYS...

Ok - lets get this straight - I have friends. I have family. I could ask 20 of them...

![image](/public/assets/forever-alone.png)

_I have friends okay!_

But I could never ask someone to help me out... and then proceed to nag him for 14 days if he used the application once
per day...

I thought about going to reddit... I thought to myself - asking for help - that's cool, I don't mind it. I have to get
some people that are interested in the application - so I went to specific subreddits and read their Rules... Every
single subreddit had very restrictive rules regarding asking for help... I understand I guess... but it's impossible to
get or interact with anyone and show them anything you're doing... I always have a feeling like I am being intrusive if
I want to talk about some things I am working on... even asking for help... :(

So I went on fiverr and gave € ($$ for American friends) to a nice fellow from Afghanistan. He was actually really
helpful and guided me all the way through. With everything I needed to do... really cool dude.
But he's not like a QA - he just has a few friends that they click every app they sold testing for 5 mins a day - good
side gig if you ask me.

The problem for me is that they expect you to do testing - they'll monitor the app user interactions - but that's all...
I checked logs and I didn't get any useful information from them... I had to do all the real testing myself... And I
don't have access to a tablet, so I don't have any clue what's going on there... I used Simulators but... I guess I hope
that they simulate the real thing perfectly... if not... I am fucked...

I am currently in process of applying for production.

## Apple

I am stuck in "Enroll into developer program" for 1.5months now... I wrote 6 different support mails asking for help and
everybody is just saying 'Be careful with information you input in the form'... bitch please - I have a feeling like
they're telling me I don't know how to fucking submit a Customer Cart form on any ecommerce site... wtf... and its
100€ ($$ for American friends) that you need to pay YEARLY!

I'll update the post - but the process for Apple is really, really annoying and fucked tbh... why do I need to beg for
them to enroll me in their developer program... the whole thing is really bizarre!

## Conclusion

- I didn't publish an application yet
- It's pretty difficult process
- Development is not the main protagonist as you might think
- If you're a small fry like me its hard mentally
- Feeling that nobody cares