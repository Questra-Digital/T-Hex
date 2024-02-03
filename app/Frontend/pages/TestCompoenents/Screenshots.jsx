import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
import GetScreenShot from "../InsidePages/GetScreenShot";
const Screenshots = () => {

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
        if (value === 'screenshot') {
            list.push(data["tests"][i])
        }
      }

      console.log("S-----------S");
      console.log(data);
      console.log("S-----------S");

      setDataToPass(list);
    };
    useEffect(() => {
      handleResultPageLoad();
    }, []); 
  



    return (
        <>
            <GetScreenShot data={dataToPass}/>
        </>
    );
};

export default Screenshots;
