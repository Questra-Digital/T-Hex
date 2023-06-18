// import React from "react";
// import datalog from './data/logdata.json'
// import GetLogs from "../InsidePages/GetLogs";
// const Logs = () => {  

//     return (
//         <>
//         <GetLogs data={datalog}/>
//         </>
//     );
// };

// export default Logs;


import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
import GetLogs from "../InsidePages/GetLogs";
const Logs = () => {

    const [dataToPass, setDataToPass] = useState(null);
    const handleResultPageLoad = async () => {

      const { data } = await client.query({
        query: gql`
          query Tests {
            tests {
              id
              username
              project
              type
              urlid
              testfile
              status
              duration
            }
          }
        `,
      });
  
      let list = [];
  
    //   let visited = [];
      for (let i = 0; i < data.tests.length; i++) {
        let value = data["tests"][i].type;
        if (value === 'log') {
            // visited.push(value);
            list.push(data["tests"][i])
        }
      }

    //   console.log("1-----------1");
    //   console.log(data);
    //   console.log("1-----------1");

    //   console.log("2-----------2");
    //   console.log(list);
    //   console.log("2-----------2");

      setDataToPass(list);
    };
    useEffect(() => {
      handleResultPageLoad();
    }, []); 
  



    return (
        <>
            <GetLogs data={dataToPass}/>
        </>
    );
};

export default Logs;
