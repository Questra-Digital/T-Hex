import React, { useEffect, useState } from "react";
import { useRouter } from 'next/router';
import { gql } from "@apollo/client";
import client from "../../appolo-client";
// import data from "./data/sampledata.json";
import Link from 'next/link';
import GetResults from "../InsidePages/GetResults";
const Results = () => {
  
  // useRouter().reload();
  console.log("resultpage");

  const [dataToPass, setDataToPass] = useState(null);
  const handleResultPageLoad = async () => {
    // console.log("path: " + inputCloneGitRep);
    // make request to set the settings of test environment in the database
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

    // console.log("0101---------------0101")
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

    // console.log(list);
    console.log(data);
    // console.log("0101---------------0101")
    setDataToPass(list);
    // datatopass=data;
    // console.log("------------------------")
    // const jsonData = JSON.stringify(data);
    // console.log(data["tests"][0]["project"]);
    // console.log("------------------------")

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
