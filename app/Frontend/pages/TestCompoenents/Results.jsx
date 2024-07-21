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
    setDataToPass(data.tests); // Pass all test data
  };

  useEffect(() => {
    handleResultPageLoad();
  }, []);

  return (
    <>
      <GetResults data={dataToPass} />
      {/* Remove or update this Link component if not intended */}
      <Link href="/results">Show Result</Link>
    </>
  );
};

export default Results;
