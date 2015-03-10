wadereddi
=========

Reddit-like demo, written with the Wade.Go framework

How to run this:
1. Install Wade's Template Compiler (`go get github.com/phaikawl/wade`).
2. Go to this package's directory and run `wade` to generate Go files for the templates.
3. Install this forked version of `fresh` the dev server runner `go get github.com/phaikawl/fresh`.
4. Inside this package's directory, invoke `./run_fresh` for the server side, then make a new terminal and invoke `./run_gopherjs` for the client side.
If all goes well, you can browse the page at http://localhost:3000.
The client-side functional (still native Go) test is in `client`, you can run it.
