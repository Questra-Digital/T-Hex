import React, { useEffect, useState } from 'react';
import SettingConfiguration from '../InsidePages/SettingConfiguration';
const CustomizeSettings = () => {
    const [showConfig, setShowConfig] = useState(false);
    const [showSettings, setShowSettings] = useState(false);

    const [isClickable1, setIsClickable1] = React.useState(true);
    const [isClickable2, setIsClickable2] = React.useState(true);

    const submitted =() =>{
        if(showConfig === true){
            setShowSettings(true);}


        setIsClickable1(false);     
        console.log("Clicking"); 
 
        // window.scrollTo(0, document.body.scrollHeight);
    }

    const [selectedOption, setSelectedOption] = useState('default');

    const handleRadioChange = (e) => {
      setSelectedOption(e.target.value);
      if(e.target.value === 'configure'){
        setShowConfig(true);
      }
    };
  

    useEffect(() => {
        window.scrollTo(0, document.body.scrollHeight);
      }, [showSettings]);
    

    return (
        <>
            <div className="bg-white p-8 rounded-md w-full mt-5">
                <div className="flex flex-col border border-gray-300 rounded bg-gray-100 p-4">

                    <label className="flex items-center" >
                        <input
                            type="radio"
                            name="config"
                            value="default"
                            checked={selectedOption === 'default'}
                            onChange={handleRadioChange}     
                        />
                        <span className="ml-4 whitespace-nowrap ">Run the test with default settings</span>
                    </label>
                    <label className="flex items-center" >
                        <input
                            type="radio"
                            name="config"
                            value="configure"
                            checked={selectedOption === 'configure'}
                            onChange={handleRadioChange}
                        />
                        <span className="ml-4">Configure your own settings</span>
                    </label>

                    <div className="flex">
                        <button className="h-10 mt-20 px-4 py-2 font-bold text-white bg-blue-500 rounded 
                        hover:bg-blue-700 disabled:bg-gray-400 ${isClickable1 ? 'cursor-pointer' : 'cursor-default'"
                        onClick={submitted}
                        disabled={!isClickable1}
                        >
                            Confirm
                        </button>

                        <img src="setting.png" alt="" className="ml-auto w-40" />
                    </div>
                </div>


                {showSettings && (
                    <SettingConfiguration />
                )}
            </div >
        </>
    );
}

export default CustomizeSettings;
