import test from 'node:test';
import React, { useState } from 'react';

const VideoPlayer = ({ data }: any) => {


  // Check if data is null or empty
  if (!data || data.length === 0) {
    return <div>Loading...</div>; // Placeholder content while waiting for data
  }

  const [id_c, setid_C] = useState(1);
  const [testfile, settestfile] = useState(data[0].testfile);
  const [duration, setduration] = useState(data[0].duration);
  const [videourl, setvideourl] = useState(`https://drive.google.com/uc?export=download&id=${data[0].urlid}`);

  /*Set Values*/

  const handleNextPage = () => {
   
    if(id_c>=data.length-1){}
    else{
    setid_C(id_c+1);
    settestfile(data[id_c+1].testfile)
    setduration(data[id_c+1].duration)
    setvideourl(`https://drive.google.com/uc?export=download&id=${data[id_c+1].urlid}`)
    }
  };

  const handlePreviousPage = () => {
    
    if(id_c<=1){}
    else{
    setid_C(id_c-1);
    settestfile(data[id_c-1].testfile)
    setduration(data[id_c-1].duration)
    setvideourl(`https://drive.google.com/uc?export=download&id=${data[id_c-1].urlid}`)
    }
  };


  // const videoId = '1qszo05uDKL4VXNhGKcStRQIyOEAnZUri';
  // const videoUrl = `https://drive.google.com/uc?export=download&id=${videoId}`;

  return (

    <>
      {/* <link rel="stylesheet" href="SettingsAndConfig.css"></link> */}

      <div className="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md ">
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

      </div>

    </>

    // <div>
    //   <video controls>
    //     <source src={videoUrl} type="video/mp4" />
    //   </video>
    // </div>
  );
};

export default VideoPlayer;


// import React from 'react';

// const Video = () => {
//     return (
//         <>

//             <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md ">
//                 <div class="grid grid-cols-4 grid-rows-3 gap-0">
//                     <div class="text-lg font-medium">TEST RUN</div>
//                     <div class=""></div>
//                     <div class="p-1/2 mt-1 ml-4 font-semibold">Status:</div>
//                     <div class="p-1/2 mt-1 ml-4 font-semibold">Duration</div>
//                     <div class="font-semibold">Test by:</div>
//                     <div class="pt-1 ">Abdullah</div>
//                     <div class="pt-1 ml-7 text-green-500"><i class="fas fa-check-circle"></i></div>
//                     <div class="pt-1 ml-4">16 sec</div>
//                     <div class="font-semibold">Test id:</div>
//                     <div class="text-xs pt-1">242141232212</div>
//                     <div class="ml-1 font-semibold">Completed</div>
//                     <div class=""></div>
//                 </div>

//             </div><div class="video-container">

//                 <video controls>
//                     <source src="https://drive.google.com/file/d/1prbjCDbri1oYUyzHYqDF-7jYqGOAZw6y/view?usp=sharing" type="video/mp4" />
//                 </video>

//                 {/* <video src="https://drive.google.com/file/d/1prbjCDbri1oYUyzHYqDF-7jYqGOAZw6y/view?usp=sharing" controls></video> */}
//             </div><div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-2xl bg-white rounded-md ">
//                 <div class="grid grid-cols-2 grid-rows-3 gap-4">
//                     <div class="text_field">Test Logs</div>
//                     <div class="text_field">Selenium Logs</div>
//                     <div class="text-xs">This is Test Logs...................</div>
//                     <div class="text-xs">This is Selenium Logs...................</div>
//                     <div class="Download_field">Download</div>
//                     <div class="Download_field">Download</div>
//                 </div>

//             </div>
//         </>
//     );
// }

// export default Video;