import Header_page from './Components/Header';
import Footer_page from './Components/Footer';
import SettingConfiguration from './InsidePages/SettingConfiguration';
import StartTest from './InsidePages/StartTest';
import Debugging from './InsidePages/Debugging';
import HomePage from './InsidePages/HomePage';
import React from 'react';
import { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

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
        {/* <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css"></link> */}
      </>
    );
  }

  // Code that uses `document` goes here
  return (
    <Router>
      <Header_page />
      <Routes>
        <Route exact path="/Debugging" element={<Debugging />} />
        <Route exact path="/HomePage" element={<HomePage/>} />
        <Route exact path="/Settings" element={<SettingConfiguration />} />
        <Route exact path="/StartTest" element={<StartTest />} />
        
      </Routes>
      <Footer_page />
    </Router>
  );
}
