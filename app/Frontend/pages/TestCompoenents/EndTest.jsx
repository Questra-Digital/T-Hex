import React, { useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";

const EndTest = (props) => {
  const [isClickable, setIsClickable] = useState(true);

  const handleButtonClick = () => {
    setIsClickable(false);
    handleEndTest();
  };

  const handleEndTest = async () => {
    console.log("handleEndTest");

    const { data } = await client.mutate({
      mutation: gql`
        mutation DestroyAll {
          destroyAll
        }
      `,
    });
    console.log("Destroyed All Tests: \n");
  };

  return (
    <>
      <div className="bg-white p-8 rounded-md w-full mt-5">
       
        <div>
          <div className="mt-20 -mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
            <div className="text-center p-10 border border-gray-300 rounded bg-gray-50">
              <h2 className="text-xl text-green-800 font-semibold text-gray-900">
                Now click on the following button to end the testing.
              </h2>

              <button
                onClick={handleButtonClick}
                disabled={!isClickable}
                className="mt-4 bg-green-600 hover:bg-green-500 focus:bg-green-700 active:bg-green-800 disabled:bg-gray-400 px-4 py-2 rounded-md text-white font-semibold tracking-wide
                                ${isClickable ? 'cursor-pointer' : 'cursor-default'}
                                "
              >
                End Testing
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default EndTest;
