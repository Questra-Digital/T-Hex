import React, { useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";

const SettingConfiguration = () => {
  const containerStyle = { width: "85%" };
  const [isClickable1, setIsClickable1] = React.useState(true);
  const [inputBrowser, setInputBrowser] = useState("Google Chrome");
  const [inputVersion, setInputVersion] = useState("latest");
  const [inputDebugging, setInputDebugging] = useState(false);
  const [inputLogs, setInputLogs] = useState(false);
  const [inputParallelism, setInputParallelism] = useState(false);

  // function submitForm(event) {
  //   event.preventDefault();
  //   console.log("submitForm");
  const handleSetSettings = async (e) => {
    

    console.log(
      "Browser " +
      inputBrowser +
      " Version " +
      inputVersion +
      "deb " +
      inputDebugging +
      "Logs " +
      inputLogs +
      "parallelism" +
      inputParallelism
    );

    // setIsClickable1(false);
    // make request to set the settings of test environment in the database
    const { data } = await client.mutate({
      
      mutation: gql`
          mutation SetSettings($input: getSettingInput!) {
            setSettings(input: $input) {
              id
              browser
              version
              stepByStepDebugging
              enableLogs
              parallelism
            }
          }
        `,
      variables: {
        input: {
          browser: inputBrowser,
          version: inputVersion,
          stepByStepDebugging: inputDebugging,
          parallelism: inputParallelism,
          enableLogs: inputLogs,
        },
      },
    });

   // once added, empty the input box
    // setInputBrowser("Google Chrome");
    // setInputVersion("Latest");
    // setInputDebugging(false);
    // setInputParallelism(false);
    // setInputLogs(false);
  };
  //   setIsClickable1(false);
  // }

  return (
    <>
      <div class="container" style={containerStyle}>
        <p>Test Configuration</p>
        <hr />

        {/* <form onSubmit="submitForm(event)"> */}
          <label for="browser">Browser:</label>

          <select id="browser" name="browser" onChange={(e) => setInputBrowser(e.target.value)} value={inputBrowser}>
            <option value="chrome">Google Chrome</option>
            <option value="firefox">Mozilla Firefox</option>
            <option value="safari">Apple Safari</option>
            <option value="edge">Microsoft Edge</option>
            <option value="opera">Opera</option>
            {/* onChange={(e) => setInputBrowser(e.target.value)}
            value={inputBrowser} */}
          </select>
          <br />
          <label for="version">Version:</label>
          <select id="version" name="version" onChange={(e) => setInputVersion(e.target.value)} value={inputVersion}>
            <option value="latest">Latest</option>
            <option value="older">Older</option>
            {/* onChange={(e) => setInputVersion(e.target.value)}
            value={inputVersion} */}
          </select>

          <div>
            <label>Step by Step Debugging:</label>
            <input
              type="checkbox"
              id="debugging"
              name="debugging"
              onChange={(e) => setInputDebugging(e.target.checked)}
              checked={inputDebugging}
            />
          </div>
          <div>
            <label>Enable Logs:</label>
            <input
              type="checkbox"
              id="logs"
              name="logs"
              onChange={(e) => setInputLogs(e.target.checked)}
              checked={inputLogs}
            />
          </div>
          <div>
            <label>Parallelism:</label>
            <input
              type="checkbox"
              id="Parallelism"
              name="Parallelism"
              onChange={(e) => setInputParallelism(e.target.checked)}
              checked={inputParallelism}
            />
          </div>
          <br />
          {/* <label>fghjsm:</label> */}
          <input 
            type="submit"
            value="Submit"
            onClick={handleSetSettings} 
            // disabled={!isClickable1}
          // onClick={handleSetSettings}
          />
          {/* <button
            className="bg-emerald-500 text-white active:bg-emerald-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
            type="button"
            onClick={handleSetSettings}
          ></button> */}
        {/* </form> */}
      </div>
    </>
  );
};

export default SettingConfiguration;
