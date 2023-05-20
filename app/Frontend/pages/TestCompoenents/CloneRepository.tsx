import React from 'react';
// import SideNavbar from '../Components/SideNavbar';
const CloneRepository = (props) => {

    const [showModal, setShowModal] = React.useState(false);
    const [isClickable, setIsClickable] = React.useState(true);


    const handleButtonClick=()=>{
        setShowModal(true);
        // <SideNavbar addItemToNavbar="a" />
        setIsClickable(false);
    }

    return (

        <>
            <div className="bg-white p-8 rounded-md w-full mt-5">
                <div className=" flex items-center justify-between pb-6">
                    <div>
                        <h1 className="text-gray-700 font-semibold text-xl">Enter URL of Repository</ h1>
                    </div>
                    <div className="flex items-center justify-between">
                        <svg
                            aria-hidden="true"
                            className="mr-2 flex-shrink-0 w-8 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            fill="currentColor"
                            viewBox="0 0 20 20"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                fill-rule="evenodd"
                                d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
                                clip-rule="evenodd"
                            ></path>
                        </svg>
                        <div className="wiflex bg-gray-50 items-center p-2 rounded-md">
                            <input
                                className="bg-gray-50 outline-none ml-1 block w-9/10 "
                                type="text"
                                name=""
                                id=""
                                placeholder="GitHub or GitLab URL..."
                            />

                        </div>
                        <div className="lg:ml-40 ml-10 space-x-8">
                            <button /*onClick={handleWorkspaceSetting}*/ 
                                onClick={handleButtonClick}
                                disabled={!isClickable}
                                className="bg-teal-600 hover:bg-teal-500 focus:bg-teal-700 active:bg-teal-800 disabled:bg-gray-400 px-4 py-2 rounded-md text-white font-semibold tracking-wide
                                ${isClickable ? 'cursor-pointer' : 'cursor-default'}
                                "
                            >
                                Add Repository
                            </button>
                        </div>
                        <div className="lg:ml-40 ml-10 space-x-8">
                            {/* <button className="bg-indigo-600 px-4 py-2 rounded-md text-white font-semibold tracking-wide cursor-pointer">
                Add a Project
              </button> */}
                            {/* <AddProject /> */}
                        </div>
                    </div>
                </div>
                {/* add project here */}
                {/* </div> */}
                {showModal === false ? (
                    <div>
                        <div className="mt-5 -mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">

                            <div className="text-center p-10 border border-gray-300 rounded bg-gray-100">
                                <h2 className="text-2xl text-gray-900">No projects found</h2>
                                <p className="text-gray-600">Please add one.</p>
                            </div>

                        </div>
                    </div>
                ) :

                    <div>
                        <div className="mt-5 -mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">

                            <div className="text-center p-10 border border-gray-300 rounded bg-gray-100">
                                <h2 className="text-2xl text-green-700">Project had been added successfully.
                                </h2>
                                <div className="mt-3 flex justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg"  width="24" height="24" fill="#2ecc71" viewBox="0 0 24 24"><path d="M12 0c-6.627 0-12 5.373-12 12s5.373 12 12 12 12-5.373 12-12-5.373-12-12-12zm-1.25 16.518l-4.5-4.319 1.396-1.435 3.078 2.937 6.105-6.218 1.421 1.409-7.5 7.626z"/></svg>
                                </div>
                            </div>

                        </div>
                    </div>

                }
                {/* add project here */}
            </div>

        </>
    );
}

export default CloneRepository;

