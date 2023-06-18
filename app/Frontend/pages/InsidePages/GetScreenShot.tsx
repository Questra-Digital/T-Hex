import test from 'node:test';
import React, { useState } from 'react';

const GetScreenShot = ({ data }: any) => {
  var Username = "abdullahsaleem40404"
  var Projectname = "FYP-SAMPLE SELENIUM TESTS"
  var testfile = "python-1"

  if (!data || data.length === 0) {
    return <div>Loading...</div>; 
  }


  console.log(data);

  console.log("1----1");
  console.log(data);
  console.log("1----1");
  const pageSize = 2; 
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
