# pursuit of software simplicity

Simplicity.

There is something beautiful to my eyes when I see something designed with simplistic approach. 

Simplicity and minimalism are usually intertwined with each other. Simple life is minimal. Minimal design contains elements of simplest forms.

I enjoy minimalistic approach in a lot of aspects of my life. Minimalist design or approach to a design to me is something which has the fewest and simplest elements used to create the effect you're pursuing.

But how did I get this appreciation for something simple?

When I first started out with programming I figured that 10x engineers combat complex problems and create these astounding systems that are resilient and provide good user/developer experience. I wanted to work on complex projects.

Well... as Doctor Jagger once said: 

> you can't always get what you want... but if you try sometimes, you might find you get what you need

![image](/public/assets/aww-yeah.png)

_aww yeah_

Thus, I was reassigned to a big project inside the company that was having very difficult times. Reassigned is a loosey-goosey word... I was thrown into a tsunami with a message "have fun". 

It shocked me how big and intertwined the projects can become.

Even more of a shocker was how big the projects can become without having any proper tooling around them... People were deploying via FTP.

You had to run the project in production environment otherwise it would not start.

You could not run almost ANY of the commands straight out of the box since you could make changes on production servers.

Git? Rofl lmao

Staging environment? Okay, relax there buddy.

Developers responsible for this mess? Wanted to work on other projects and are 'busy' with other customers.

Customer happiness? Shouting at me on meetings level and threatening with lawsuits. Very nice.


The codebase? Absolute disaster. Bugs everywhere - weekly sprints were mostly filled with bug reports and 1 small new feature.

Technology used? Already old at the point that I joined with 0 documentation provided.

This type of system was complex. Very complex. You had to think and somehow test 10 different things when you made one change. You had 0 tests to help you with testing. You could not run some tests due to possibility of running production services.

We had to create a plan of steps that would reduce the complexity of the entire project. Reformat the codebase was easier said than done since somehow things broke even when reformatting.

Debugging was hard - abstracting method names and calling abstract $set and $get methods without prior checking if method exists on an object was nightmarish.

All those experiences shaped my mind that complex is not something that you really want as a programmer. You really want as simple as you can get and plan for it before you introduce complexity. Because once you introduced complexity, it's really hard to restore back simplicity. I often read that people rewrite the entire system rather than to refactor the existing system. Same thing happened when an internal team was formed, and they successfully sold (to client) that rewrite would solve their and our problems.

Fast-forward a few years, and we're currently inside the "javascript boom". Suddenly I was surrounded with new developers that had a completely different mindset. We're currently an agency that develops for our customers custom solutions based on their needs.

But this new javascript wave of developers actually flipped the experience priority. They pushed for all the new fancy frameworks that emerged and ultimately won - due to being able to sell some fancy keywords like "headless".

I don't argue. Headless has its place, and it truly can be wonderful.

But forcing it down everyone's throat can be a real pain in the ass - especially if you're not on the same bandwagon. Headless introduced a full stack of problems for us. First let's decide on how we want to transfer the data from backend system to frontend? We decided on graphql - cool cool cool. We even found a nice library that would enable a GUI for creating queries and mutations based on our classes.

Everything was fine and dandy - until you get a special client edge case.

And since we decided on graphql and library approach we had to create an entire solution around that edge case. Dear god... Instead of using a REST API approach and just modifying the logic of fetching the data we had to create listeners that would listen to the request and alter the data being sent.

Since we couldn't use the traditional CMS features in our system we had to mold and recreate the entire CMS system just so that we could fit in inside this headless structure.

The data structure was nuts! I still can't fully understand everything - I have to check and understand it first...

We created something incredibly complex - which in practice should be a simple CRUD app... Traffic is really not even a debate.

The main reason we reached for complexity was not the demand of our customer. We pushed for it. And in the end we created a mess out of it. Editor experience is awful... but developer experience was nice for the frontend folk. They had autocomplete and types generated via graphql in the correct format.

![image](/public/assets/this-is-fine.png)

_this is fine_


Should we strive for something so complex without any real business need for it?
Should we do something manually without codegen if it means that we don't need to introduce another system/library just to make it work?

I understand that people want to have great developer experiences.

But would you sacrifice another experience for it?
You're unfortunately not writing a library for developers - you're writing software for an office worker.

Are we creating artificial complexity and using super 'edge' technology that is changing faster than anyone can keep up just so we can feel any happiness from our day-to-day job?

This blog runs on 0 javascript and without a database.

Was there a real need for it? No.
Will there be in the future? Possibly? Idk... 

I am actually super happy I managed to make it run like I did. I feel like if I made it with Next.js I would have to battle caching systems, routing - serving static shit with islands etc.

I'm just not down to do that sort of stuff for a 3 page website with a blog.

I like it simple.  I like it neat.
