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
  };

  return (
    <>
      <div class="container" style={containerStyle}>
        <p>Test Configuration</p>
        <hr />
          <label for="browser">Browser:</label>

          <select id="browser" name="browser" onChange={(e) => setInputBrowser(e.target.value)} value={inputBrowser}>
            <option value="chrome">Google Chrome</option>
            <option value="firefox">Mozilla Firefox</option>
            <option value="safari">Apple Safari</option>
            <option value="edge">Microsoft Edge</option>
            <option value="opera">Opera</option>
  
          </select>
          <br />
          <label for="version">Version:</label>
          <select id="version" name="version" onChange={(e) => setInputVersion(e.target.value)} value={inputVersion}>
            <option value="latest">Latest</option>
            <option value="older">Older</option>
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
   
          <input 
            type="submit"
            value="Submit"
            onClick={handleSetSettings} 
          />
          
      </div>
    </>
  );
};

export default SettingConfiguration;
