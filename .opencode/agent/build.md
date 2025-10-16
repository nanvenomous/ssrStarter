
---
---

Build new software to the user's specification.
When possible keep with the design patterns and coding standards of the project (unless you intend to improve the standard).
The purpose of this project is to create a ui library around daisyui which primarily solves problems with a server-side-rendered approach using `htmx`.
However, it is still possible to write Typescript in `web/index.ts` when a trip to the server doesn't make sense.
Using `templ` typescript functions can be called like so ```templ @templ.JSFuncCall("funcName", "argument1", "argument2") ```

The project primarily uses
- `golang/go`
- `a-h/templ`
- `tailwindlabs/tailwindcss`
- `saadeghi/daisyui`
- `oven-sh/bun`
- `bigskysoftware/htmx`
the project also uses `go-task/task` for building the project; similar to gnu make you can use `task build` to build the project

Feel free to do some refactoring to keep the code clean and concise.
Also if you deprecate code, then remove that code at the end of the change.
Do not leave unused functions and do not export functions which are not used outside that package.

After implementing the requested changes:

1. Run the build command: `task build`
2. If there are any build errors, analyze them and fix the issues
3. Re-run `task build` until the build passes
5. Provide a summary of changes made and validation results
