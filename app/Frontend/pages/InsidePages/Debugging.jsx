import React from 'react';
const Debugging = () => {
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

            </div>

            <ul >
                <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md list-none">
                    <li class="screenshot-text-item">
                        <div class="screenshot">
                            <img src="screenshot1.png" alt="Screenshot 1" />
                        </div>
                        <div class="text">
                            <div class="grid grid-cols-4 grid-rows-1 gap-0 shadow-md rounded">
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold py-2 px-4">
                                    Test Logs
                                </button>
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold">
                                    Selenium Logs
                                </button>
                            </div>

                            <div class="text-xs pt-10">This is Test Logs...................</div>
                            <div class="text-xs">This is Selenium Logs...................</div>

                        </div>
                    </li></div>

                <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md list-none">
                    <li class="screenshot-text-item">
                        <div class="text">
                            <div class="grid grid-cols-4 grid-rows-1 gap-0 shadow-md rounded">
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold py-2 px-4">
                                    Test Logs
                                </button>
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold">
                                    Selenium Logs
                                </button>
                            </div>

                            <div class="text-xs pt-10">This is Test Logs...................</div>
                            <div class="text-xs">This is Selenium Logs...................</div>

                        </div>
                        <div class="screenshot">
                            <img src="screenshot2.png" alt="Screenshot 2" />
                        </div>
                    </li></div>

                <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md list-none">
                    <li class="screenshot-text-item">
                        <div class="screenshot">
                            <img src="screenshot3.png" alt="Screenshot 3" />
                        </div>
                        <div class="text">
                            <div class="grid grid-cols-4 grid-rows-1 gap-0 shadow-md rounded">
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold py-2 px-4">
                                    Test Logs
                                </button>
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold">
                                    Selenium Logs
                                </button>
                            </div>

                            <div class="text-xs pt-10">This is Test Logs...................</div>
                            <div class="text-xs">This is Selenium Logs...................</div>

                        </div>
                    </li></div>

                    <div class="mx-auto mt-20 w-2/3 p-4 shadow-lg dark:shadow-xl bg-white rounded-md list-none">
                    <li class="screenshot-text-item">
                        <div class="text">
                            <div class="grid grid-cols-4 grid-rows-1 gap-0 shadow-md rounded">
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold py-2 px-4">
                                    Test Logs
                                </button>
                                <button class="sm:text-sm truncate bg-white-500 hover:bg-gray-400 shadow-md text-black font-bold">
                                    Selenium Logs
                                </button>
                            </div>

                            <div class="text-xs pt-10">This is Test Logs...................</div>
                            <div class="text-xs">This is Selenium Logs...................</div>

                        </div>
                        <div class="screenshot">
                            <img src="screenshot2.png" alt="Screenshot 2" />
                        </div>
                    </li></div>
            </ul>

        </>
    );
}

export default Debugging;