import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
const CreateDockerFile = () => {
  const [InputDockerfile, setInputDockerfile] = useState(`
  FROM ubuntu:latest
  RUN apt-get update
  RUN apt-get install python3 -y
  RUN apt-get install python3-pip -y
  RUN pip install selenium
  WORKDIR /usr/app/src
  `);
  const handlebuttonClick = () => {
    console.log("Button Click ");
    handleDockerInput();
  };
  const handleDockerInput = async () => {
    console.log("path: " + InputDockerfile);
    const { data } = await client.mutate({
      mutation: gql`
        mutation CreateDockerFile($input: createDockerFileInput!) {
          createDockerFile(input: $input)
        }
      `,
      variables: {
        input: {
          dockerCmnds: InputDockerfile,
        },
      },
    });
  
    setInputDockerfile("");
  };

  return (
    <>
      <div className="bg-white p-8 rounded-md w-full">
        <div className="flex flex-row border border-purple-900 rounded bg-white p-4 mt-10"  style={{ borderRadius: '50px' }}>
          <div className="w-1/2 flex flex-col justify-center   " >
            <span className="text-black-50 ml-3 text-xl font-bold text-black">
              Dockerfile
            </span>
            <hr className="border-black my-2 w-4/5 " />
            <div className="w-4/5 mt-5 bg-white rounded-lg" >
              <textarea
                className="w-full h-48 p-4 text-sm text-black bg-black-100 rounded-lg resize-none"
                defaultValue={""}           
                value={InputDockerfile}
                readOnly={true}
                onChange={(e) => setInputDockerfile(e.target.value)}  
              ></textarea>
            </div>
          </div>
          <div className="w-1/2 flex flex-col items-center justify-center  rounded-lg">
            <div className="w-2/3 h-2/3 flex items-center justify-center ">
              <img
                src="DockerFile.jpeg"
                alt="Description of the image"
                className="w-128 rounded-full"
              />
            </div>
            <div className="text-center mt-3">
              <p className=" text-black">Click Confirm to Create Docker Files</p>
              <p className=" text-black">Language: python</p>
              <button
                className="mt-1 px-5 font-bold text-sm text-white bg-purple-500 rounded
                        hover:bg-Purple-700 disabled:bg-gray-400 ${isClickable1 ? 'cursor-pointer' : 'cursor-default'"
                onClick={handlebuttonClick}
              >
                Confirm
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
export default CreateDockerFile;