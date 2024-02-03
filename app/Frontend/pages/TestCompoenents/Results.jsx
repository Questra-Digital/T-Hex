import React, { useEffect, useState } from "react";
import { useRouter } from 'next/router';
import { gql } from "@apollo/client";
import client from "../../appolo-client";

import Link from 'next/link';
import GetResults from "../InsidePages/GetResults";
const Results = () => {
  
  console.log("resultpage");

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

    console.log(data);
    let list = [];

    let visited = [];
    for (let i = 0; i < data.tests.length; i++) {
      let value = data["tests"][i].testfile;
      if (!visited.includes(value)) {
          visited.push(value);
          list.push(data["tests"][i])
      }
    }

    console.log(data);
    setDataToPass(list);
    
  };
  useEffect(() => {
    handleResultPageLoad();
  }, []); 

  return (
    <>
      <GetResults data={dataToPass} />
      <Link href="/results">Show Result</Link>
    </>
  );
};

export default Results;
