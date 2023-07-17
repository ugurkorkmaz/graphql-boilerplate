[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# Golang Graphql Boilerplate

This boilerplate is built using Golang, Chi and Gqlgen, providing a scalable and maintainable architecture to build your GraphQL API with ease.

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`APP_NAME`,`APP_PORT`,`APP_SECRET_KEY`,`APP_SECRET_HEADER_NAME`,`APP_DEBUG`,`GQL_COMPLEXITY_LIMIT`

## Run Locally

Clone the project

```bash
  git clone https://github.com/ugurkorkmaz/graphql-boilerplate 
```

Go to the project directory

```bash
  cd graphql-boilerplate
```
Run commands

```bash
  go run ./bin all
```

| Command | Description |
| --- | --- |
| build | Compiles the binary. |
| install | Installs dependencies for the go.mod file. |
| update | Updates project dependencies. |
| gqlgen | Generates code from the GraphQL schema. |
| frontend | Builds the frontend templates. |
| run | Runs the binary. Prints the server URL and allows stopping it with Ctrl+C. |
| test | Runs project tests and lints. |
| clean | Removes the binary. |
| all | Runs all the commands in order: gqlgen, update, build, and run. |
| help | Shows available commands and their descriptions. |

## API Reference

| Path | Description |
| --- | --- |
| `/graphql` | Handles GraphQL requests and serves the GraphQL API. |
| `/websocket` | Handles WebSocket connections for real-time communication with the GraphQL API. |
| `/` | Serves the Playground UI for interacting with the GraphQL API. |




## Tech Stack
**Server:** Go, Gqlgen, Chi


## Authors

- [@ugurkorkmaz](https://www.github.com/ugurkorkmaz)


## License

[MIT](https://github.com/ugurkorkmaz/graphql-boilerplate/blob/main/LICENSE)


## Contributing

Contributions are always welcome!
