import { StackContext, Table, AppSyncApi } from "@serverless-stack/resources";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";

export function MyStack({ stack }: StackContext) {
  //! Create a notes table
  // const newTable = new Table(stack, "TableName", {
  //   fields: {
  //     id: "string",
  //   },
  //   primaryIndex: { partitionKey: "id" },
  // });

  //! Existing table
  const exTable = new Table(stack, "table", {
    cdk: {
      table: dynamodb.Table.fromTableArn(stack, "Table", "arn:aws:dynamodb:eu-west-1:680852762302:table/test-sharebus-appsync-Notes"),
    },
  });

  //! Create the AppSync GraphQL API
  const api = new AppSyncApi(stack, "AppSyncApi", {
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
      myLambda: "functions/main.handler",
    },
    resolvers: {
      //? query/mutation name should be same as defined in schema
      // "Query    queryName": "dataSourceName",
      // "Mutation mutationName": "dataSourceName",
      "Mutation createNote":  "myLambda",
      "Query    listNotes":   "myLambda",
      "Query    getNoteById": "myLambda",
      "Mutation updateNote":  "myLambda",
      "Mutation deleteNote":  "myLambda",
    },
  });

  //! Enable the AppSync API to access the DynamoDB table
  api.attachPermissions([exTable]);

  //! Print some output in console
  stack.addOutputs({ 
    ApiId: api.apiId, 
    ApiKey: String(api.cdk.graphqlApi.apiKey),
    ApiUrl: api.url,
    tableName: exTable.tableName
  })
}
