import React from 'react';
const SettingConfiguration = () => {
    return (
        <>

            <div class="container">
                <p>Test Configuration</p>
                <hr />

                <form>
                    <label for="browser">Browser:</label>
                    <select id="browser" name="browser">
                        <option value="chrome">Google Chrome</option>
                        <option value="firefox">Mozilla Firefox</option>
                        <option value="safari">Apple Safari</option>
                        <option value="edge">Microsoft Edge</option>
                        <option value="opera">Opera</option>
                    </select>

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
                        <label >Enable Logs:</label>
                        <input type="checkbox" id="logs" name="logs" />
                    </div>
                    <div>
                        <label>Parallelism:</label>
                        <input type="checkbox" id="Parallelism" name="Parallelism" />
                    </div>
                    {/* <div id="Parallelism-details" style="display:none;">
                        <label for="parallel-tests">Number of parallel tests:</label>
                        <select id="parallel-tests" name="parallel-tests">
                            <option value="1">1</option>
                            <option value="2">2</option>
                            <option value="4">4</option>
                            <option value="8">8</option>
                        </select>
                    </div>
                    */}
                    <br />
                    <input type="submit" value="Submit" /> 
                </form>
            </div>
        </>

    );
}

export default SettingConfiguration;