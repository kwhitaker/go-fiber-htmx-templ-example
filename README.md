# go-fiber-htmx-templ-example

This is a to-do app I made to play with the following goodies:
- [Go Fiber](https://gofiber.io/)
- [HTMX](https://htmx.org/)
- [Templ](https://templ.guide)
- [AlpineJS](https://alpinejs.dev/)
- [Bun](https://bun.sh)

### Why?
Because I wanted to 🤷. After a decade of doing things with React and Typescript, I wanted to poke around with Go and HTMX.
I feel like 90% of the web-apps out there could be done with that stack, so I decided to experiment.

Also, I really wanted to checkout Templ, and I had to poke around a bit to figure out how to get it to work with Fiber, so I thought
I could save someone else a bit of digging later.

### But why did you use X library/package/whatever?
Because I wanted to 🤷. Seriously, its not that deep. Go touch grass, nerd.

### Are there tests?
After-a-fashion. Go doesn't seem to have much of an ecosystem around E2E testing, which is what I wanted to do for this thing. 
I ended up picking [playwright-go](https://github.com/playwright-community/playwright-go).

### Why Tailwind?
Because its 2024 (or was 2024, if you are in the FUTURE), and I don't have time for any CSS nonsense. Also, [DaisyUI](https://daisyui.com/) makes it even easier to bootstrap a quick project like this.

## Install And Run
To install dependencies:

```bash
go mod install
bun install
```

To run
1. Open a terminal and execute `air`
2. Open a second terminal and execute `bun watch`

### Run Tests
I couldn't find a good, Go-standard way to spawn the main App from within the tests, and then use that instance to run the tests,
so you have to do a bit of jank to run them.

1. Open a terminal and execute `MODE=test air`
2. Open a second terminal and execute `go test`