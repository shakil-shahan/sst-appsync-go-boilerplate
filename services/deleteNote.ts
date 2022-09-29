import { DynamoDB } from "aws-sdk";

const dynamoDb = new DynamoDB.DocumentClient();

export default async function deleteNote(noteId: string): Promise<string> {
  const params = {
    Key: { id: noteId },
    TableName: process.env.TABLE_NAME as string,
  };

  await dynamoDb.delete(params).promise();

  return noteId;
}
