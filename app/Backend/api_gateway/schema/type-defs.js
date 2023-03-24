const { gql } = require("apollo-server");

const typeDefs = gql`
  type Settings {
  id: ID!
  browser: String!
  version: String!
  stepByStepDebugging: Boolean!
  enableLogs: Boolean!
  parallelism: Boolean!
  numberOfParallelTests: Int!
  }

  input getSettingInput {
  browser: String!
  version: String!
  stepByStepDebugging: Boolean!
  enableLogs: Boolean!
  parallelism: Boolean!
  numberOfParallelTests: Int!
  }

  type Test {
  id: ID!
  status: String!
  # //startedby: String!
  testPath: String!
  }

  input startTestInput {
  testPath: String!
  # //startedby: String!
  }

  input seleniumInput {
  containerName: String!
}

input projectBuildImageInput {
  imgName: String!
  projectPath: String! #contextPath: pythonProject
}

input projectContainerStartInput {
  imageName: String!
  containerName: String!
  # seleniumContainerName: String!
}
input projectContainerEndInput {
  containerName: String!
}

input updateTestInput {
  id: ID!
  status: String!
  testPath: String!
}

  type Query{
    getSettingsByID (id: ID!): Settings!
    tests: [Test!]! #list of tests, if there no tests return null
  }

  type Mutation {
    setSettings(input: getSettingInput!): Settings!
    startTest(input: startTestInput!): Test!

    updateTest(input: updateTestInput!): Test!

    seleniumPull: String!
    # seleniumStart: String!
    seleniumStart(input: seleniumInput!): String!
    seleniumStop(input: seleniumInput!): String!
    seleniumRemove(input: seleniumInput!): String!

    projectBuildImage(input: projectBuildImageInput!): String!
    projectContainerStart(input: projectContainerStartInput!): String!
    projectContainerStop(input: projectContainerEndInput!): String!
    projectContainerRemove(input: projectContainerEndInput!): String!
  }
`;

module.exports = { typeDefs };
