import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
import VideoPlayer from "../InsidePages/VideoPlayer";
const Videos = () => {

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
        if (value === 'video') {
            list.push(data["tests"][i])
        }
      }

      console.log("1-----------1");
      console.log(list);
      console.log("1-----------1");

      setDataToPass(list);
    };
    useEffect(() => {
      handleResultPageLoad();
    }, []); 
  



    return (
        <>
            <VideoPlayer data={dataToPass}/>
        </>
    );
};

export default Videos;
