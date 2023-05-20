import React from "react";
const SettingConfiguration = () => {
  const containerStyle = { width: "85%" };
  const [isClickable1, setIsClickable1] = React.useState(true);

  function submitForm(event) {
    event.preventDefault();

    setIsClickable1(false);
    
  }
  
  
  return (
    <>
      <div class="container" style={containerStyle}>
        <p>Test Configuration</p>
        <hr />

        <form onsubmit="submitForm(event)">
          <label for="browser">Browser:</label>
          <select id="browser" name="browser">
            <option value="chrome">Google Chrome</option>
            <option value="firefox">Mozilla Firefox</option>
            <option value="safari">Apple Safari</option>
            <option value="edge">Microsoft Edge</option>
            <option value="opera">Opera</option>
          </select>
          <br />
          <label for="version">Version:</label>
          <select id="version" name="version">
            <option value="latest">Latest</option>
            <option value="older">Older</option>
          </select>

          <div>
            <label>Step by Step Debugging:</label>
            <input type="checkbox" id="debugging" name="debugging" />
          </div>
          <div>
            <label>Enable Logs:</label>
            <input type="checkbox" id="logs" name="logs" />
          </div>
          <div>
            <label>Parallelism:</label>
            <input type="checkbox" id="Parallelism" name="Parallelism" />
          </div>
          <br />
          <input className={"disabled:bg-gray-400 ${isClickable1 ? 'cursor-pointer' : 'cursor-default'"} type="submit" value="Submit" 
           disabled={!isClickable1}
          />
        </form>
      </div>
    </>
  );
};

export default SettingConfiguration;
