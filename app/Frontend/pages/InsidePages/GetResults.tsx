import React, { useState } from 'react';
import ListHead from '../ShowData/ListHead';
import ListBody from '../ShowData/ListBody';
const GetResults = ({ data }: any) => {


  if (!data || data.length === 0) {
    return <div>Loading...</div>; 
  }


  const pageSize = 3; 
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

      <div className="m-0 inline-block min-w-full shadow rounded-lg overflow-hidden">
        <table className="min-w-full leading-normal">

          <ListHead />
          <ListBody data={currentData} />

        </table>
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
      </div>

    </>
  );
}

export default GetResults;