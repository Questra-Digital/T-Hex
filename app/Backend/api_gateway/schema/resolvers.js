const { GraphQLClient, gql } = require("graphql-request");
// const workspaceClient = new GraphQLClient("http://localhost:4001/query");
// const projectClient = new GraphQLClient("http://localhost:4002/query");

const settingClient = new GraphQLClient("http://localhost:8080/query");
const testExecutionClient = new GraphQLClient("http://localhost:8081/query");

const resolvers = {
  Query: {
    getSettingsByID: async (parent, args) => {
      const query = gql`
        query ($id: ID!) {
          getSettingsByID(id: $id) {
          browser
          version
          stepByStepDebugging
          enableLogs
          parallelism
          numberOfParallelTests
          }
        }
      `;
      const id = args.id;
      const data = await settingClient.request(query, { id: id });
      return data.getSettingsByID;
    },
    tests: async () => {
      const query = gql`
        query {
          tests {
            id
            status
            testPath
          }
        }
      `;
      const data = await testExecutionClient.request(query);
      return data.tests;
    }
  },
  Mutation: {
    setSettings: async (parent, args) => {
      const workspace = args.input;
      console.log(workspace);
      const query = gql`
        mutation SetSettings($input: getSettingInput!) {
          setSettings(input: $input) {
          id
          browser
          version
          stepByStepDebugging
          enableLogs
          parallelism
          numberOfParallelTests
          }
        }
      `;
      const data = await settingClient.request(query, { input: workspace });
      console.log(query);
      console.log(data);

      return data.setSettings;
    },
    startTest: async (parent, args) => {
      const testInput = args.input;
      console.log(testInput);
      const query = gql`
        mutation StartTest($input: startTestInput!) {
          startTest(input: $input) {
          id
          browser
          version
          stepByStepDebugging
          enableLogs
          parallelism
          numberOfParallelTests
          }
        }
      `;
      const data = await testExecutionClient.request(query, { input: testInput });
      return data.startTest;
    },
    seleniumPull: async (parent, args) => {
      const selInput = args.input;
      console.log(selInput);
      const query = gql`
        mutation SeleniumPull($input: String!) {
          seleniumPull(input: $input) {
          }
        }
      `;
      const data = await testExecutionClient.request(query, { input: selInput });
      return data.seleniumPull;
    },
    seleniumStart: async (parent, args) => {
      const selInput = args.input;
      console.log(selInput);
      const query = gql`
        mutation SeleniumStart($input: seleniumInput!) {
          seleniumStart(input: $input) {
          }
        }
      `;
      const data = await testExecutionClient.request(query, { input: selInput });
      return data.seleniumStart;
    },
    seleniumStop: async (parent, args) => {
      const selInput = args.input;
      console.log(selInput);
      const query = gql`
        mutation seleniumStop($input: seleniumInput!) {
          seleniumStop(input: $input) {
          }
        }
      `;
      const data = await testExecutionClient.request(query, { input: selInput });
      return data.seleniumStop;
    }
  }
};

module.exports = { resolvers };
