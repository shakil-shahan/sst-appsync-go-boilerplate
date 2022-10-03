import { AppSyncApi, StackContext, Table } from "@serverless-stack/resources";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";

export function MyStack({ stack }: StackContext) {
  /**
  //? This whole comment-out block is for templating.
  //! Create a notes table
  const newTable = new Table(stack, "TableName", {
    fields: {
      id: "string",
    },
    primaryIndex: { partitionKey: "id" },
  });

  //! Existing table
  const exTable = new Table(stack, "table", {
    cdk: {
      table: dynamodb.Table.fromTableArn(stack, "Table", "arn:aws:dynamodb:eu-west-1:680852762302:table/test-sharebus-appsync-Notes")
    },
  });

  //! Create the AppSync GraphQL API
  const appSyncApi = new AppSyncApi(stack, "AppSyncApi", {
    schema: "services/graphql/schema.graphql",
    defaults: {
      function: {
        // Pass the table name to the function as env variable
        // You can access this TABLE_NAME as a process.env value in all functions
        // You can pass any env variable from here
        environment: {
          TABLE_NAME: exTable.tableName,
        },
      },
    },
    dataSources: {
      myLambda: "functions/appsync-lambda.go",
    },
    resolvers: {
      //? query/mutation name should be same as defined in schema
      // "Query    queryName": "dataSourceName",
      // "Mutation mutationName": "dataSourceName",
    },
  });

  //! Enable the AppSync API to access the DynamoDB table
  appSyncApi.attachPermissions([exTable]);


  //! Create the HTTP API
  const httpApi = new Api(stack, "Api", {
    routes: {
      "GET /": "services/functions/functionName.go",
    },
  });

  //! Print some output in console (HTTP)
  stack.addOutputs({
    //! (AppSync)
    // ApiId: appSyncApi.apiId,
    // ApiKey: String(appSyncApi.cdk.graphqlApi.apiKey),
    // ApiUrl: appSyncApi.url,
    // tableName: exTable.tableName
    //! (HTTP)
    // ApiEndpoint: api.url,
  });
  */

  //? As we are not using SST's stack in our prod,
  //? So we are not using any DB things here
  const appSyncApi = new AppSyncApi(stack, "AppSyncApi", {
    schema: "services/graphql/schema.graphql",
    dataSources: {
      myLambda: "functions/appsync-lambda.go",
    },
    resolvers: {
      //? query/mutation name should be same as defined in schema
      // "Query    queryName": "dataSourceName",
      // "Mutation mutationName": "dataSourceName",
      "Query    getList": "myLambda",
      "Query    getOne": "myLambda",
      "Mutation create": "myLambda",
      "Mutation update": "myLambda",
      "Mutation delete": "myLambda",
    },
  });

  //! As we are using an existing table, so connected it with arn
  const exTable = new Table(stack, "table", {
    cdk: {
      table: dynamodb.Table.fromTableArn(stack, "Table", "arn:aws:dynamodb:eu-west-1:XXXXXXXXXX:table/table-name")
    },
  });

  //! Enable the AppSync API to access the DynamoDB table
  appSyncApi.attachPermissions([exTable]);

  //! Print some output in console (AppSync)
  stack.addOutputs({
    ApiId: appSyncApi.apiId,
    ApiKey: String(appSyncApi.cdk.graphqlApi.apiKey),
    ApiUrl: appSyncApi.url
  })
}
