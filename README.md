## Background Story

Deal with low resource to render complex web UI it's a painful process, especially in React.js server side, since nodeJS single thread by nature. It has many limitation on performance comparing with another languange which support multithreading and multiprocessing.

So what we do to deal with it, we need to route some traffic where need high resource usage to something more efficience service, such as golang service.

But how to handle the the View and/or data fetching, where currently running super well in React.js. Or should we create View and data fetching in 2 services now ?.

Based on the story, we create some solution on how we can still do coding UI in mature ecosystem which is React.js but processing for bot request using golang. With no coding View in Golang. We create solution, call `Template as a Service`. Where React.js will become source of truth for View and golang will do the rest when traffic coming from bot using template provided by React.js to render the View. So we still has consistent UI between bot and user.

## How It Works

Contract is required to be created between template service and the consumer. Developer require to create mock data fetching contract which will render when `consumer` request the template. Then the consumer will save the latest template locally and use it to handle real bot traffic.

## Development

Open 3 terminals to running this project:

1. Enter to directory `next-app`, and running nextJS app with `yarn && yarn dev`
2. After port `3000` up, enter to directory `renderer`, and running golang app with `go mod tidy && go run main.go`
3. From working directory, executre `make run` to running development environment in docker.
4. Open `localhost` with normal user agent or test open with custom bot user agent to get response from different service.

## References

- [Golang Template](https://www.practical-go-lessons.com/chap-32-templates)
