# GO Lambda function local development
### Build GraphQL API with AWS AppSync and Lambda with Go by using SST

<img src="https://d33wubrfki0l68.cloudfront.net/cc1e61abe8db8f1be699c2cf5adde992db3ab776/00968/img/logo.svg" alt="Go" height="25" />
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="Go" height="25" />
<img src="https://miro.medium.com/max/1400/1*q_DYf7YPkJ7G-m9gC06h5w.jpeg" alt="Go" height="25" />

An appsync-boilerplate with Go, created with SST.

## Getting Started

[**Read the tutorial**](https://sst.dev/examples/how-to-create-a-serverless-graphql-api-with-aws-appsync.html)

Install the example.

```bash
npm intall
```

Start the Live Lambda Development environment.

```bash
npm sst start
```

Change the default stage and region

```bash
# Start
npx sst start --stage alpha --region us-west-1

# Build
npx sst build --stage alpha --region us-west-1

# Deploy
npx sst deploy --stage alpha --region us-west-1

# Remove
npx sst remove --stage alpha --region us-west-1
```

Once your local environment is ready, [head over to the AppSync console](https://console.aws.amazon.com/appsync).

Here you can run queries and mutations and make changes locally to test your Lambda resolvers.

```graphql
mutation create {
  create(input: { id: "001", content: "My note" }) {
    id
    content
  }
}

query getOne {
  getOne(input: {id: "001"}) {
    id
    content
  }
}

query list {
  list {
    id
    content
  }
}

mutation update {
  update(input: { id: "001", content: "My updated note" }) {
    id
    content
  }
}

mutation delete {
  delete(input: {id: "001"})
}
```

## Commands

### `npm run start`

Starts the local Lambda development environment.

### `npm run build`

Build your app and synthesize your stacks.

Generates a `.build/` directory with the compiled files and a `.build/cdk.out/` directory with the synthesized CloudFormation stacks.

### `npm run deploy [stack]`

Deploy all your stacks to AWS. Or optionally deploy a specific stack.

### `npm run remove [stack]`

Remove all your stacks and all of their resources from AWS. Or optionally remove a specific stack.

### `npm run test`

Runs your tests using Jest. Takes all the [Jest CLI options](https://jestjs.io/docs/en/cli).

## Documentation

Learn more about SST.

- [Docs](https://docs.sst.dev)
- [@serverless-stack/cli](https://docs.sst.dev/packages/cli)
- [@serverless-stack/resources](https://docs.sst.dev/packages/resources)

## Community

[Follow them on Twitter](https://twitter.com/sst_dev) or [post on their forums](https://discourse.sst.dev).
