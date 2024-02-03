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

      for (let i = 0; i < data.tests.length; i++) {
        let value = data["tests"][i].type;
        if (value === 'log') {
            list.push(data["tests"][i])
        }
      }

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
