import React from 'react'
import * as SoIcons from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import styles from './Footer1.module.css'
import { Link } from 'react-router-dom'


function Footer() {
    return (
        <>
            <div className={styles.fBody}>
                <footer className={styles.footer}>
                    <div className={styles.fContainer}>
                        <div className={styles.fRow}>
                            <div className={styles.fCol}>
                                <h4>Features</h4>
                                <ul>

                                <li><Link to="/CloneRepository" >Clone Git Repository</Link></li>
                                <li><Link to="/CustomizeSettings" >Configure Settings</Link></li>
                                <li><Link to="/CreateDockerFile" >Write Docker File</Link></li>
                                <li><Link to="/RunTest" >start test</Link></li>
                               </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>Results</h4>
                                <ul>
                                <li><Link to="/Results" >Test Result</Link></li>
                                <li><Link to="/Videos" >Test Videos</Link></li>
                                <li><Link to="/Logs" >Test Logs</Link></li>
                                <li><Link to="/Screenshots" >Test Screenshots</Link></li>
                                </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>Get help</h4>
                                <ul>
                                <li><Link to="/ContactUs" >Contact Us</Link></li>
                                <li><Link to="/ContactUs" >About Us</Link></li>                                
                                </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>follow us</h4>
                                <div className={styles.fSocialLinks}>
                                    <a href="https://www.facebook.com/"><FontAwesomeIcon icon={SoIcons.faFacebook} /></a>
                                    <a href="https://twitter.com/"><FontAwesomeIcon icon={SoIcons.faTwitter} /></a>
                                    <a href="https://www.instagram.com/"><FontAwesomeIcon icon={SoIcons.faInstagram} /></a>
                                    <a href="https://pk.linkedin.com/"><FontAwesomeIcon icon={SoIcons.faLinkedinIn} /></a>
                                </div>
                            </div>
                        </div>
                    </div>
                </footer>
            </div>
        </>
    )
}

export default Footer

















// function footer_component() {
    
//     return (

//         <footer class="mt-80">
//             <div class="p-10 bg-gray-800 text-gray-200">

//                 <div>
//                     <div class="grid grid-cols-1 lg:grid-cols-5 ">
//                         <div>
//                             <ul class="text-gray-200">
//                                 <h4 class="text-2xl pb-4">PRODUCTS</h4>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Live</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Automate</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Percy</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />App Live</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />App Automate</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Screenshots</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Responsive</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Enterprise</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />SpeedLab</li>
//                             </ul>

//                         </div>

//                         <div>
//                             <ul class="text-gray-200">
//                                 <h4 class="text-2xl pb-4">PLATFROM</h4>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Browsers and Devices</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Data Centers</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Mobile Features</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Security</li>
//                             </ul>
//                         </div>


//                         <div>
//                             <ul class="text-gray-200">
//                                 <h4 class="text-2xl pb-4">SOLUTIONS</h4>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test on iPhone</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test on iPad</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test on Galaxy</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test on IE</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Android Testing</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />iOS Testing</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Cross Browser Testing</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Emulators and Simulators</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Selenium</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Cypress</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Android Emulators</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Visual Testing</li>
//                             </ul>
//                         </div>

//                         <div>
//                             <ul class="text-gray-200">
//                                 <h4 class="text-2xl pb-4">RESOURCES</h4>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test on Right Devices</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Support</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Status</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Release Notes</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Case Studies</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Blog</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Events</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Test University</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Champions</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Mobile Emulators</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Guide</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Responsive Design</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Nightwatch</li>
//                             </ul>

//                         </div>
//                         <div>
//                             <ul class="text-gray-200">
//                                 <h4 class="text-2xl pb-4">COMPANY</h4>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />About Us</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Customers</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Careers</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Open Source</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Partners</li>
//                                 <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
//                                     class="hover:text-yellow-500" />Press</li>

//                             </ul>



//                         </div>



//                     </div>

//                 </div>

//             </div>
//         </footer>

//     );
// }

// export default footer_component;
