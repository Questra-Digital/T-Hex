import React, { useState, useEffect } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";

const RunTest = (props) => {
  const [isClickable, setIsClickable] = useState(true);
  const [progress, setProgress] = useState(0);

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

  useEffect(() => {
    const interval = setInterval(() => {
      if (progress < 100) {
        setProgress(progress + 1);
      }
    }, 10000); // 500 ms interval for faster animation

    // Clean up the interval on component unmount
    return () => clearInterval(interval);
  }, [progress]);

  const runVideo = "/darl.mp4";

  return (
    <>
      {/* Video Background */}
      <video
        autoPlay
        loop
        muted
        style={{
          position: "absolute",
          width: "80%",
          height: "88%",
          objectFit: "cover",
          zIndex: -1,
        }}
      >
        <source src={runVideo} type="video/mp4" />
        Your browser does not support the video tag.
      </video>

      <div className="relative h-screen flex flex-col items-center justify-center text-white" >
        
        {/* Text */}
        <h1 className="text-2xl font-semibold mb-8 p-3 bg-black rounded-lg text-white "  
       

        >
          Ready to unleash the power of your tests?
        </h1>

        {/* Button */}
        <button
          onClick={handleButtonClick}
          disabled={!isClickable}
          className={`bg-blue-600 hover:bg-purple-500 focus:bg-purple-700 active:bg-blue-800 disabled:bg-gray-400 px-8 py-4 rounded-md text-white font-semibold tracking-wide ${
            isClickable ? "cursor-pointer" : "cursor-default"
          }`}
          style={{ backgroundColor: isClickable ? 'purple' : "#cccccc" }}
        >
          Start Test
        </button>

        {/* Progress Bar */}
        <div className="mt-8 w-64">
          <div className="relative h-8 bg-gray-300 rounded-full overflow-hidden">
            <div
              className="absolute top-0 left-0 bg-purple-500 h-full transition-all duration-500"
              style={{ width: `${progress}%` }}
            ></div>
            <span className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 font-semibold">
              {progress}%
            </span>
          </div>
        </div>
      </div>
    </>
  );
};

export default RunTest;
