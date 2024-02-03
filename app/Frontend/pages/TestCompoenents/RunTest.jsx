import React, { useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";

const RunTest = (props) => {
  const [isClickable, setIsClickable] = useState(true);

  const handleButtonClick = () => {
    setIsClickable(false);
    handleInitiateTest();
  };

  const handleInitiateTest = async () => {
    console.log("handleInitiateTest");
    const { data } = await client.mutate({
      mutation: gql`
        mutation InitiateTest {
          initiateTest
        }
      `,
    });
    console.log("InititateTest Result: \n", data);
  };

  return (
    <>
      <div className="bg-white p-8 rounded-md w-full mt-5">
        <div className=" flex pb-6">
          <div>
            <h1 className="text-gray-700 font-semibold text-2xl">
              Everything is done
            </h1>
          </div>
          <div className="flex">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
              className="flex-shrink-0 w-8 h-6 mt-2 ml-5  text-green-500 transition duration-75 dark:text-green-400 group-hover:text-green-900 dark:group-hover:text-green fill-current"
            >
              <path
                d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM369 
                        209L241 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L335 175c9.4-9.4 
                        24.6-9.4 33.9 0s9.4 24.6 0 33.9z"
              />
            </svg>
          </div>
        </div>

        <div>
          <div className="mt-5 -mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
            <div className="text-center p-10 border border-gray-300 rounded bg-gray-50">
              <h2 className="text-xl text-green-800 font-semibold text-gray-900">
                Now click on the following button to start the test.
              </h2>

              <button
                onClick={handleButtonClick}
                disabled={!isClickable}
                className="mt-4 bg-green-600 hover:bg-green-500 focus:bg-green-700 active:bg-green-800 disabled:bg-gray-400 px-4 py-2 rounded-md text-white font-semibold tracking-wide
                                ${isClickable ? 'cursor-pointer' : 'cursor-default'}
                                "
              >
                Start Test
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default RunTest;
