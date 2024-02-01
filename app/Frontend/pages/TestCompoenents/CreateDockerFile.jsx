import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
const CreateDockerFile = () => {
  const [InputDockerfile, setInputDockerfile] = useState("");
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
        <div className="flex flex-row border border-gray-300 rounded bg-gray-100 p-4 mt-10">
          <div className="w-1/2 flex flex-col justify-center">
            <span className="text-black-50 ml-3 text-xl font-bold">
              Dockerfile
            </span>
            <hr className="border-gray-300 my-2 w-4/5" />
            <div className="w-4/5 mt-5 bg-white rounded-lg">
              <textarea
                className="w-full h-48 p-4 text-sm text-black bg-white-100 rounded-lg resize-none"
                defaultValue={""}
                placeholder={`e.g
                FROM ubuntu:latest
                RUN apt-get update
                RUN apt-get install python3 -y
                RUN apt-get install python3-pip -y
                RUN pip install selenium
                WORKDIR /usr/app/src
                `}
                onChange={(e) => setInputDockerfile(e.target.value)}
                value={InputDockerfile}
              ></textarea>
            </div>
          </div>
          <div className="w-1/2 flex flex-col items-center justify-center">
            <div className="w-2/3 h-2/3 flex items-center justify-center">
              <img
                src="DockerFile.jpeg"
                alt="Description of the image"
                className="w-128 rounded-full"
              />
            </div>
            <div className="text-center mt-3">
              <p className="text-sm text-gray-500">Total Test Cases: 10</p>
              <p className="text-sm text-gray-500">Language: Go</p>
              <button
                className="mt-1 px-5 font-bold text-sm text-white bg-blue-500 rounded
                        hover:bg-blue-700 disabled:bg-gray-400 ${isClickable1 ? 'cursor-pointer' : 'cursor-default'"
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