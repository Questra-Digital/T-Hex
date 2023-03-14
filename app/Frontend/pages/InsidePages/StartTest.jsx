import React from 'react';
const StartTest = () => {
    return (
        <>

            <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md ">
                <div class="grid grid-cols-4 grid-rows-3 gap-0">
                    <div class="text-lg font-medium">TEST RUN</div>
                    <div class=""></div>
                    <div class="p-1/2 mt-1 ml-4 font-semibold">Status:</div>
                    <div class="p-1/2 mt-1 ml-4 font-semibold">Duration</div>
                    <div class="font-semibold">Test by:</div>
                    <div class="pt-1 ">Abdullah</div>
                    <div class="pt-1 ml-7 text-green-500"><i class="fas fa-check-circle"></i></div>
                    <div class="pt-1 ml-4">16 sec</div>
                    <div class="font-semibold">Test id:</div>
                    <div class="text-xs pt-1">242141232212</div>
                    <div class="ml-1 font-semibold">Completed</div>
                    <div class=""></div>
                </div>

            </div><div class="video-container">
                <video src="v1.mp4" controls></video>
            </div><div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-2xl bg-white rounded-md ">
                <div class="grid grid-cols-2 grid-rows-3 gap-4">
                    <div class="text_field">Test Logs</div>
                    <div class="text_field">Selenium Logs</div>
                    <div class="text-xs">This is Test Logs...................</div>
                    <div class="text-xs">This is Selenium Logs...................</div>
                    <div class="Download_field">Download</div>
                    <div class="Download_field">Download</div>
                </div>

            </div>
        </>
    );
}

export default StartTest;