import test from 'node:test';
import React, { useState } from 'react';

const GetScreenShot = ({ data }: any) => {
  var Username = "abdullahsaleem40404"
  var Projectname = "FYP-SAMPLE SELENIUM TESTS"
  var testfile = "python-1"
  // Check if data is null or empty
  if (!data || data.length === 0) {
    return <div>Loading...</div>; // Placeholder content while waiting for data
  }


  console.log(data);

  console.log("1----1");
  console.log(data);
  console.log("1----1");
  const pageSize = 2; // Number of items per page
  const [currentPage, setCurrentPage] = useState(1);


  const startIndex = (currentPage - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  const currentData = data.slice(startIndex, endIndex);

  const totalPages = Math.ceil(data.length / pageSize);

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };


  // const [id_c, setid_C] = useState(1);
  // const [testfile, settestfile] = useState(data[0].testfile);
  // const [duration, setduration] = useState(data[0].duration);
  // const [videourl, setvideourl] = useState(`https://drive.google.com/uc?export=download&id=${data[0].urlid}`);

  // /*Set Values*/

  // const handleNextPage = () => {

  //   if(id_c>=data.length-1){}
  //   else{
  //   setid_C(id_c+1);
  //   settestfile(data[id_c+1].testfile)
  //   setduration(data[id_c+1].duration)
  //   setvideourl(`https://drive.google.com/uc?export=download&id=${data[id_c+1].urlid}`)
  //   }
  // };

  // const handlePreviousPage = () => {

  //   if(id_c<=1){}
  //   else{
  //   setid_C(id_c-1);
  //   settestfile(data[id_c-1].testfile)
  //   setduration(data[id_c-1].duration)
  //   setvideourl(`https://drive.google.com/uc?export=download&id=${data[id_c-1].urlid}`)
  //   }
  // };


  // const videoId = '1qszo05uDKL4VXNhGKcStRQIyOEAnZUri';
  // const videoUrl = `https://drive.google.com/uc?export=download&id=${videoId}`;

  return (

    <>


      {currentData.map((project: any, idx: any) => {

        return (

          <ul>
            <div className="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md list-none">
              <li className="screenshot-text-item">
                <div className="screenshot">
                  <img src={`https://drive.google.com/uc?export=download&id=${project.urlid}`} alt="Screenshot 1"
                
                  />
                </div>
                <div className="text">
                  <div className="grid grid-cols-1 grid-rows-1 gap-0 shadow-md rounded">
                    <button className="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold py-2 px-4">
                      ScreenShot 1
                    </button>
                  </div>

                  <div className="grid grid-cols-2 gap-2 justify-center items-center mr-20 mt-10 mb-[-2rem]">
                    <p className="text-right font-bold">Username:</p>
                    <p><span className="text-sm">{project.username}</span></p>
                    <p className="text-right font-bold">Project:</p>
                    <p><span className="text-sm">{project.project}</span></p>
                    <p className="text-right font-bold">Test file:</p>
                    <p><span className="text-sm">{project.testfile}</span></p>
                  </div>


                </div>
              </li></div>



          </ul>

        );
      })}

      <div className="px-5 py-5 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between">
        <span className="text-xs xs:text-sm text-gray-900">
          Showing {startIndex + 1} to {Math.min(endIndex, data.length)} of {data.length} Entries
        </span>

        <div className="inline-flex mt-2 xs:mt-0">
          <button className="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-l"
            onClick={handlePreviousPage}
            disabled={currentPage === 1}
          >
            Prev
          </button>
          &nbsp; &nbsp;
          <button className="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-r"
            onClick={handleNextPage}
            disabled={currentPage === totalPages}
          >
            Next
          </button>
        </div>

      </div>


    </>

  );
};

export default GetScreenShot;
{/* <link rel="stylesheet" href="SettingsAndConfig.css"></link> */ }

{/* <div className="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md ">
        <div className="grid grid-cols-4 grid-rows-3 gap-0">
          <div className="text-lg font-medium">TEST RUN</div>
          <div className=""></div>
          <div className="p-1/2 mt-1 ml-4 font-semibold">Status:</div>
          <div className="p-1/2 mt-1 ml-4 font-semibold">Duration</div>
          <div className="font-semibold">Test id:</div>
          <div className="pt-1 ">{id_c}</div>
          <div className="pt-1 ml-7 text-green-500">

            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 512 512"
              className="flex-shrink-0 w-6 h-5 mt-1 text-green-500 transition duration-75 dark:text-green-400 group-hover:text-green-900 dark:group-hover:text-green fill-current"
            >
              <path
                d="M256 512A256 256 0 1 0 256 0a256 256 0 1 0 0 512zM369 
                209L241 337c-9.4 9.4-24.6 9.4-33.9 0l-64-64c-9.4-9.4-9.4-24.6 0-33.9s24.6-9.4 33.9 0l47 47L335 175c9.4-9.4 
                24.6-9.4 33.9 0s9.4 24.6 0 33.9z"
              />
            </svg>

          </div>
          <div className="pt-1 ml-4">{duration}</div>
          <div className="font-semibold">Test file</div>
          <div className="text-xs pt-1">{testfile}</div>
          <div className="ml-1 font-semibold">Completed</div>
          <div className=""></div>
        </div>

      </div>

      <div className="flex items-center flex-row">


        <div className="video-container">
          <video controls key={videourl}>
            <source src={videourl} type="video/mp4" />
          </video>
        </div>


      </div>

      <div className="flex justify-center">

        <div className="mb-3 inline-flex mt-2 xs:mt-0">
          <button className="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-l"
            onClick={handlePreviousPage}
          >
            Prev
          </button>
          &nbsp; &nbsp;
          <button className="text-sm text-indigo-50 transition duration-150 hover:bg-indigo-500 bg-indigo-600 font-semibold py-2 px-4 rounded-r"
            onClick={handleNextPage}
          >
            Next
          </button>
        </div>

      </div> */}


