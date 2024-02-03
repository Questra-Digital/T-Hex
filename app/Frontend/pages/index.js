import Header_page from './Components/Header';
import Footer_page1 from './Components/Footer1';
import Footer_page from './Components/Footer';
import SettingConfiguration from './InsidePages/SettingConfiguration';
import Debugging from './InsidePages/Debugging';
import HomePage from './InsidePages/HomePage';
import RunTest from './TestCompoenents/RunTest';
import React from 'react';
import { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Routes, useLocation } from 'react-router-dom'
import SideNavbar from './Components/SideNavbar';
import CloneRepository from './TestCompoenents/CloneRepository';
import CustomizeSettings from './TestCompoenents/CustomizeSettings';
import CreateDockerFile from './TestCompoenents/CreateDockerFile';
import Results from './TestCompoenents/Results'

import ContactUs from './InsidePages/ContactUs'
import Logs from './TestCompoenents/Logs'
import Screenshots from './TestCompoenents/Screenshots'
import Videos from './TestCompoenents/Videos'
import EndTest from './TestCompoenents/EndTest'

export default function Home() {

  const [isBrowser, setIsBrowser] = useState(false);

  useEffect(() => {
    setIsBrowser(true);
  }, []);

  if (!isBrowser) {
    return (
      <>
        <script src="https://cdn.tailwindcss.com"></script>
        <script type="module" src="https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js"></script>
      </>
    );
  }

  return (
    <Router>

      <div className='fixed w-full'>
        <Header_page />
      </div>

      <Routes>
        <Route exact path="/" element={<HomePage />} />
      </Routes>
      <div className="flex flex-row">
        <div className="mt-20 fixed basis-1/5">
          <SideNavbar />
        </div>

        <div className="mt-20 ml-auto basis-4/5">
          <Routes>
            <Route path="/CustomizeSettings" element={<CustomizeSettings />} />
            <Route path="/CloneRepository" element={<CloneRepository />} />
            <Route path="/RunTest" element={<RunTest />} />
            <Route path="/CreateDockerFile" element={<CreateDockerFile />} />
            <Route path="/Results" element={<Results />} />
            <Route path="/ContactUs" element={<ContactUs />} />

            <Route path="/Logs" element={<Logs />} />
            <Route path="/Screenshots" element={<Screenshots />} />
            <Route path="/Videos" element={<Videos />} />
            <Route path="/EndTest" element={<EndTest />} />
          </Routes>
        </div>
      </div>
      
    </Router >
  );
}
