Hopefully this README will be more fleshed out eventually.

## enVC 

is a project for exploring various ideas that have been in my mind for a while. It aims to be a writing app that promotes calm and organized workflows with ways of handling text more loosely than pages or files tend to allow.

I've tried a few things in this direction before which haven't materialized, and have decided it will be more fun and interesting to try building and learning in public. This approach adds some pressure on me to either create an app that others can use or at least to fail in a more useful way.

## Status

I started using a simple structure to store all my notes in MariaDB earlier this year. The initial goal for the app is to be able to do this inside of a terminal on the front end of the app, which won't require handling all SQL functions yet.  Once this is in place I would like to start running background tasks on this data and playing around with presentation.

## Tech

Eventually I hope to do things which might require handling many DOM elements. To prepare for this, I'm taking this project as an opportunity to really work with Tailwind and SolidJS, both of which I've sampled and been impressed by in the past.  I believe they will perform better than the CSS-in-JS + React stack I've used before.

In a previous iteration, I tried creating a backend using Rust, but found it difficult to prototype in, at least as a newcomer to the language. I decided to try Go for this version, which has been great for quickly iterating on ideas. Go's strong support for concurrency has influenced me to move more fully towards a message passing architecture. This shift has in turn shaped the user interface, encouraging me to start thinking about how to have background results meaningfully available to the user without interrupting their workflow.