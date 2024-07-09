const { gql } = require("apollo-server");

const typeDefs = gql`
  type Query {
    getSettingsByID (id: ID!): Settings!
    tests: [Test!]! #list of tests, if t1here no tests return null
    settings: [Settings!]! #list of tests, if there no tests return null
    setting(id: ID!): Settings #return 1 test using ID
    projectInfos: [Projectinfos!]!
    projectInfo(id: ID!): Projectinfos
  }

  type Mutation {
    initiateTest: String!
    destroyAll: String!
    setSettings(input: getSettingInput!): Settings!
    setProjectInfo(input: projectInfoInput!): Projectinfos!
    saveTestResult(input: saveTestResultInput!): Test!
    # startTest(input: startTestInput!): String!

    # updateTest(input: updateTestInput!): Test!

    createDockerFile(input: createDockerFileInput!): String!
    cloneGitRepository(input: cloneGitRepInput!): String!
    seleniumPull: String!
    # seleniumStart: String!
    seleniumStart(input: seleniumInput!): String!
    seleniumStop(input: seleniumInput!): String!
    seleniumRemove(input: seleniumInput!): String!

    # projectBuildImage(input: projectBuildImageInput!): String!
    # projectContainerStart(input: projectContainerStartInput!): String!
    # projectContainerStop(input: projectContainerEndInput!): String!
    # projectContainerRemove(input: projectContainerEndInput!): String!
  }

  type Test {
  id: ID!
  username: String! #abdullahsaleem40404@gmail.com
  project: String! #python-project
  type: String! #video/ screenshot/ logs
  urlid: String! #817491248721894782
  testfile: String! #python1
  status: String! #passed/failed/ error
  duration: String! #5s
}

  type Settings {
  id: ID!
  browser: String!
  version: String!
  stepByStepDebugging: Boolean!
  enableLogs: Boolean!
  parallelism: Boolean!
  # numberOfParallelTests: Int!
  }


  type Projectinfos {
  id: ID!
  gitEmail: String!
  gitProjectName: String!
  gitLanguage: String!
  gitNoOfTestCases: Int!
  gitTestCaseFileName: String!
}

  input getSettingInput {
  browser: String!
  version: String!
  stepByStepDebugging: Boolean!
  enableLogs: Boolean!
  parallelism: Boolean!
  # numberOfParallelTests: Int!
}

input projectInfoInput {
  gitEmail: String!
  gitProjectName: String!
  gitLanguage: String!
  gitNoOfTestCases: Int!
  gitTestCaseFileName: String!
}

input saveTestResultInput {
  username: String! #abdullahsaleem40404@gmail.com
  project: String! #python-project
  type: String! #video/ screenshot/ logs
  urlid: String! #817491248721894782
  testfile: String! #python1
  status: String! #passed/failed/ error
  duration: String! #5s
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

  input cloneGitRepInput {
    gitRepPath: String!
  }

  input startTestInput {
    seleniumContainerName: String!
    projectPath: String!
    projectImageName: String!
    projectContainerName: String!
    # //startedby: String!
  }

  input createDockerFileInput {
    dockerCmnds: String!
  }
    # input startTestInput {
    # testPath: String!
    # # //startedby: String!
    # }
  `;

module.exports = { typeDefs };
