---
Title: dealing with mediocrity
Summary: My own take on dealing with his own mediocrity
Author: Peter Paravinja
Published: 9. Aug 2024
---

**How do we define 'mediocre'?**

What would be mediocre? Something that is not awesome, not bad, just in the middle? Would some look at mediocre as fall short of expectations?

I would describe "mediocre" as something that satisfies your current level of knowledge and lack of motivation to push further.

If we compare it with another word that holds similar meaning "average" - mediocre holds a more subjective meaning... we want to express that there was potential for more... which gives a more personal implication.

My work experience has given me opportunity to work with some amazing engineers and a whole-lot of mediocre ones. It would be interesting to see how others define me. We usually see ourselves better than most other people... Since I am a perfectionist with a growth mindset I usually listen to feedback and reflect that feedback with my own subjective thoughts about the issue. Even when I get a praise - I reflect on it... __"Did I really give it my all? Is the praise on point. I know I did a good job - but does this person know that I was capable of much more."__

I sometimes notice that I am slacking off with daily improvement... I really want to get better at software! Yet there are other values in my life that I want to give attention like quality time with partner and friends, sports and practicing guitar.

Since having multiple values I often question myself if these things bar me from achieving something more if I could focus on just one thing. But what is that one thing that would keep my mind at peace?

I figured out soon that I was at least of average intelligence with a strong logical mind. I was raised with poor working habits which I believe have set me back in life... Life took a spin, and I ended up in an operators job where I hated everything - the work,environment, coworkers, corporate structure, working hours... people with Phds that would look down on me for being lower ranked employee. I would say I wasn't even levels of mediocrity.

That was the lowest of low in my life and a moment took place where I decided to take matters into my own hands. Life really isn't easy... You really need to work for it if you want it!

But it was more than want for me! I needed more if I wanted to keep living! I wanted to be stimulated, I wanted challenge, I wanted to build stuff, I needed to feel like I am doing something that challenged me intellectually...

Making my first life self-evaluation high on drugs - I decide that due to loving computers there was only one job for me: Software engineering!

Landing intern role I really took everything very single task very seriously. I didn't know if internship would evolve into a full time job so I just wanted to really use this opportunity and learn as much as I could!

I craved for that moist programming challenge every day! It felt incredible! Every day there was insane progress - I was learning how to ssh to a remote server, JS, PHP, SCSS, git,... insane progress was being made! My first rodeo with classes was mind-blowing! I started seeing real world like classes, with inheritance and abstractions all over the place.

Well... fast-forward 5 years into the future, my knowledge really broadened, and my skills stopped progressing all together on the job. I was placed on a Lead developer role where suddenly I had to learn how to talk to customers, how to lead a team, how to communicate properly, how to guide and mentor junior engineers... It was new, yes. But the thrill... it was gone. I didn't have something that would challenge me anymore. I started to feel stagnant.

Since we're an agency we're dependent on our customers to give us tasks based on their needs. And I can say that 95% of our customer needs are CRUD operations with integrations between systems sprinkled in between.

Let me tell you, CRUD operations with integrations between systems stopped being difficult long time ago. Of course there is business logic surrounding CRUD operations - but even those were usually just spawns of evil created by the customers themselves and with zero proper reasoning on why things had to happen certain way.

I am not downplaying anything or taking away anyone's pride in their job - since I think a lot of us are doing some type of CRUD operations. But usually there's a problem that surrounds the CRUD operation... lets say scale or a smart algorithm to determine how to display the data. These features were sadly not present at with our customers.

Due to repetitive work I found out that I wasn't being challenged daily... I wasn't being challenged even on a monthly basis. I created my own challenges on how to write something really fast but keep the quality intact... Nothing would help. This is where mediocrity starts to sip into my thoughts.

Am I not working daily with enough complex of a problem that would encourage my growth through work? I read a lot of articles that showcase that there are complex problems in the work that you can solve... Am I missing out on my opportunity while working for this company? Should I work on my hobby projects more and see growth there? Should I switch jobs and find something that would give me new environment and new challenges?

How to escape the mediocrity that is breathing behind my neck?

![image](/public/assets/panik.png)

_panik_

Well what worked for me - might not work for you... What I did to encourage new growth was simply to pick up a new programming language and commit hard to it to learn new ways. In my instance I picked up **Go**.

I didn't like Javascript's world with abstractions and blackbox complexity all over the place. I had more issues setting projects up than writing the code... The programmers were encouraging more and more blackbox solutions which would reduce my knowledge of how everything works down to just learning how the library for the service works.

It made me feel even worse. I wanted to puke in my mouth. I really didn't want to have any part in the frontend development world. And I never shy from the challenge - but this all felt terrible.

So Go was an incredibly refreshing change. I started with reading programming books - but for the first time I read a book about programming language. It was extremely pleasant. [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) by Jon Bodner.

Next I tackled [Writing A Compiler In Go](https://compilerbook.com/) by Thorsten Ball which opened my eyes that there really is no magic behind the curtains... Lexers, parsers, AST and REPL - wau...

It was an eye-opener.

![image](/public/assets/zugzug.jpg)

_zug zug_


I gained so much newfound knowledge.

Not only for programming language... I am talking about the entire world of new possibilities.

I would never have thought about creating CLI apps before - everything was website based, then suddenly I got an idea to try to build my own TUI browser which got me thinking what features actually create a browser (lol). 

Next thing I never thought about was how HTTP really works - its just TCP - you have to send status, header and a body and split them with `"\r\n"`... I honestly didn't know about that beforehand.

I always worked with encrypted messages, and I was pleasantly surprised to find out that you can generate your own PEM file with elliptic curve and sign messages with it. It was dark magic that suddenly became very clear (day magic?).

I am slowly peeling the onion that has become today's world of technological abstractions. We really are standing on the shoulders of giants that came before us. It's scary sometimes just thinking how programming looked like 30-40 years ago.

But notice - I became more knowledgeable with 'basics'. I didn't learn a service that would handle everything for me. I didn't go out and learn a Go framework (I actually did tho, just not right from the get go :D) that would just serve requests and handle responses for me. I tried to give focus on getting the grasp on the underlying.

I did the same with tools that I was using.

I like using GoLand and Phpstorm - I like Intellij. This blog is written in GoLand. And I wanted to learn the editor deeper to help me achieve more things faster!

While working with huge JSON files I actually took time, researched tools and learned about [jq](https://jqlang.github.io/jq/).

I wanted to learn more about databases - not just queries but the entire toolset. Postgres provides you with [full text search](https://www.postgresql.org/docs/current/textsearch-intro.html)!

Every obstacle I faced, I took some time off the main project I was developing and learned a new technology - I wanted to broaden everything. Knowledge, tools and especially my way of problem-solving.

When mediocrity inevitably catches up with me again (software is constantly changing plus there are millions of things to learn!), I already have ideas on exploring functional programming, learning Zig and tackling manual memory management. Currently, I don't know what I'll be able to do with these things - and this is where people get stuck most when they're faced with a challenge that does not have obvious implication at first... Why would I learn that? What would be the point?

That is the point! The point is to learn!