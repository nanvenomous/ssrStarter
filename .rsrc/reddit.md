After using htmx for a while I'm convinced it does solve a majority of needs around the web with speed and simplicity.
However, I admit I've struggled a bit in terms of riding on the forefront of industry adoption.
In some groups I feel like the only person evangelizing the technology.
One thing js frameworks do well is libraries which are part of the endless abstraction problem, **but** also make things easier sometimes.

Last week I collected some of the most common components from some of my server side rendered web projects.
https://github.com/nanvenomous/ssrStarter

I had considered making a library or framework of sort (Yes I'm aware htmx is somewhat anti-framework).
Decided I don't have enough time to build a general component library right now by myself.

I have been very inspired by the following open source projects, which compose my favorite tech stack:
- [go](https://github.com/golang/go)
- [templ](https://github.com/a-h/templ)
- [htmx](https://github.com/bigskysoftware/htmx)
- [tailwindcss](https://github.com/tailwindlabs/tailwindcss)
- [daisyui](https://github.com/saadeghi/daisyui)
- [bun](https://github.com/oven-sh/bun)
- [air](https://github.com/air-verse/air)
- [task](https://github.com/go-task/task)
- [phosphor icons](https://github.com/phosphor-icons/homepage)

My focus is the unison point between htmx <-> go/templ <-> tailwindcss/daisyui/bun (css, ts)
I chose daisyui instead of a js framework since it works out of the box with htmx. 

I find the issue to be not with the javascript itself, but with initialization/deinitialization of javascript after/before htmx swaps.
Few to no js frameworks or js component libraries were designed with htmx swapping in mind.
Please tell me if I'm wrong here. I don't have time to try them all.

So, **if you try it out**, let me know what you think of the starter project.
https://github.com/nanvenomous/ssrStarter
I'm sure it could use improvement.
And also let me know your thoughts on building frameworks for hypertext driven server-side-rendered applications.
Or let me know if you are simply interested in the same problem space, have a similar project, or have had some of the same frustrations.
