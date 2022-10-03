import { App } from "@serverless-stack/resources";
import { MyStack } from "./MyStack";

export default function (app: App) {
  app.setDefaultFunctionProps({
    runtime: "go1.x",
    srcPath: "services",
  });
  app.stack(MyStack);
}
