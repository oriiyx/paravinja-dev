# Dealing with my own mediocrity

**How do we define 'mediocre'?**

What would be mediocre? Something that is not awesome, not bad, just in the middle? Would some look at mediocre as fall short of expectations?

I would say that mediocre would pass as something that satiates your current state of knowledge and low willpower to continue progressing further. If we compare it with another word that holds similar meaning "average" - mediocre holds a more subjective meaning... we want to express that there was potential for more... which gives a more personal implication.

My work experience has given me opportunity to work with some amazing engineers and a whole-lot of mediocre ones. It would be interesting to see how others define me. We usually see ourselves better than most other people - yet I hold a high regard for my harsh self evaluation which usually focuses on my weaknesses.

I often notice that I am slacking off... I want to get better at writing software! Yet there are other values in my life that I want to give attention like sports, quality time with partner and guitar.

Since having multiple values I often question myself if these things bar me from achieving something more if I could focus on just one thing.

I figured out soon that I was at least average smart with good logical mind. I was unfortunately not raised to have strong working habits which set me back a few years... Life took a spin, and I was placed on an operators job where I hated everything - environment, coworkers, corporate structure, working hours... people with Phds would look down on me for being lower ranked employee. I would say I was not even hitting mediocrity.

That is where the moment came and I had to take matters into my own hands.
It was more than want! I needed more! I wanted to be stimulated, I wanted challenge, I wanted to build stuff, I needed to feel like I am doing something that challenged me...

Making my first life self-evaluation high on drugs - I decide that due to loving computers there was only one job for me: Software engineering baby!

Landing an intern role I swiftly took everything very seriously. I craved for that moist challenge every day! It felt incredible! Every day there was insane progress - I was learning SSH, JS, PHP, SCSS, git,... insane progress was being made! My first rodeo with classes was mind-blowing! I started seeing real world like classes, with inheritance and abstractions all over the place.

Well... fast-forward 5 years into the future, my knowledge has broadened, my skills have capped at my company.

Since we're an agency we're dependent on our customers to give us tasks based on their needs. And I can say that 95% of our customer needs are CRUD operations with integrations between systems.

Let me tell you, CRUD operations with integrations between systems is not difficult.

Of course, a task like that would be difficult for a person with 0 knowledge. You need to have programming language down, the framework you're working with, authentication methods that are required, databases you're working with, design patters that you're implementing and the dockerized environment you're running your code in. I am not downplaying anything.

I am talking about the need to be challenged daily. This is where mediocrity starts to sip into my thoughts.

Am I not working daily with enough complex of a problem that would encourage my growth through work? Should I work on my hobby projects more and see growth there? Should I switch jobs and find something that would give me new environment and new challenges?

How to escape the mediocrity that is breathing behind my neck?

What I did to encourage new growth was to pick up a new programming language. **Go**.

I didn't like JS world with abstractions and blackbox complexity all over the place. I had more issues setting projects up than writing the code... The programmers were encouraging more and more blackbox solutions which would reduce my knowledge of how everything works down to a small margin.

It made me feel even worse.

So Go was an incredibly refreshing change. I started to read programming books - but for the first time I read a book about programming language. It was extremely pleasant. [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) by Jon Bodner.

It was an eye-opener.

![image](/public/assets/zugzug.jpg)

_zug zug_


I gained so much newfound knowledge.

Not only for programming language... I am talking about the entire world of possibilities.

I would never have thought about creating CLI apps before - everything was website based. I never thought how HTTP really works - its just TCP - `conn.Write([]byte(fmt.Sprintf("%s\r\n%s\r\n%s", status, header, body)))` you send some status, some header and a body... I honestly didn't know about that beforehand.

I always worked with encrypted messages, and I was pleasantly surprised to find out that you can generate your own PEM file with `elliptic.P256()` and sign messages with it. It was dark magic that suddenly became very clear (day magic?).

I was becoming more and more powerful in the knowledge department.

But notice - I became more knowledgeable with 'basics'. I didn't learn a service that would handle everything for me. I didn't go out and learn a Go framework (I actually did tho, just not right from the get go :D) that would just serve requests and handle responses for me. I tried to give focus on getting the grasp on the underlying.

I did the same with tools that I was using.

I like using GoLand and Phpstorm - I like Intellij products. This blog is written in GoLand. And I wanted to learn the editor deeper to help me achieve even more!

I wanted to learn more about databases - not just queries but the entire toolset. Postgres provides you with [full text search](https://www.postgresql.org/docs/current/textsearch-intro.html) go out there and check it out! You could try something other than `LIKE` and `ILIKE` (but don't introduce unneeded complexity to your system).

Every day I move and learn something new I'll be a step ahead of my own mediocrity.

When mediocrity inevitably catches up with me again (software is constantly changing plus there are millions of things to learn!), I already have ideas on exploring functional programming. Learning Zig and tackling memory safety and even lower level language. Currently, I don't know what I'll be able to do with these things - and this is where people get stuck most I think... Why would I learn that? What would be the point?

That is the point! The point is to learn!