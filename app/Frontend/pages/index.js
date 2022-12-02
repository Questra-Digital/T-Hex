
import Image from './cars';
export default function Home() {


  
  function Menu(e) {
  
  }

  return (

    <><script src="https://cdn.tailwindcss.com"></script>
      <script type="module" src="https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js"></script>
    <nav class="p-5 bg-white shadow md:flex md:items-center md:justify-between bg-sky-900">
      <div class="flex justify-between items-center ">
        <span class="text-2xl font-[Poppins] text-slate-50 cursor-pointer">
          <img class="h-10 inline" src="T_Hex.png" />
          T-HEX
        </span>

        <span class="text-3xl text-slate-50  cursor-pointer mx-2 md:hidden block">
          <ion-icon name="menu" onclick={Menu(this)}></ion-icon>
        </span>
      </div>

      <ul
        class="md:flex md:items-center z-[-1] md:z-auto md:static absolute to-transparent left-0 md:w-auto md:py-0 py-4 md:pl-0 pl-7 md:opacity-100 opacity-0 top-[-400px] transition-all ease-in duration-500">
        <li class="mx-4 my-6 md:my-0 ">
          <a href="#" class="text-l text-slate-50  hover:text-cyan-500 duration-500">Products</a>
        </li>
        <li class="mx-4 my-6 md:my-0">
          <a href="#" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Developers</a>
        </li>
        <li class="mx-4 my-6 md:my-0">
          <a href="#" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Live for Teams</a>
        </li>
        <li class="mx-4 my-6 md:my-0">
          <a href="#" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Pricing</a>
        </li>


        <button class="bg-cyan-400 text-white font-[Poppins] duration-500 px-6 py-2 mx-4 hover:bg-cyan-500 rounded  ">
          Sign in
        </button>


      </ul>
    </nav>
    
    <footer class="mt-80">
    <div class="p-10 bg-gray-800 text-gray-200">

      <div>
        <div class="grid grid-cols-1 lg:grid-cols-5 ">
          <div>
            <ul class="text-gray-200">
              <h4 class="text-2xl pb-4">PRODUCTS</h4>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Live</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Automate</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Percy</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>App Live</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>App Automate</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Screenshots</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Responsive</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Enterprise</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>SpeedLab</li>
            </ul>

          </div>

          <div>
            <ul class="text-gray-200">
              <h4 class="text-2xl pb-4">PLATFROM</h4>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Browsers and Devices</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Data Centers</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Mobile Features</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Security</li>
            </ul>
          </div>


          <div>
            <ul class="text-gray-200">
              <h4 class="text-2xl pb-4">SOLUTIONS</h4>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test on iPhone</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test on iPad</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test on Galaxy</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test on IE</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Android Testing</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>iOS Testing</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Cross Browser Testing</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Emulators and Simulators</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Selenium</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Cypress</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Android Emulators</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Visual Testing</li>
            </ul>
          </div>

          <div>
            <ul class="text-gray-200">
              <h4 class="text-2xl pb-4">RESOURCES</h4>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test on Right Devices</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Support</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Status</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Release Notes</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Case Studies</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Blog</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Events</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Test University</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Champions</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Mobile Emulators</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Guide</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Responsive Design</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Nightwatch</li>
            </ul>

          </div>
          <div>
            <ul class="text-gray-200">
              <h4 class="text-2xl pb-4">COMPANY</h4>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>About Us</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Customers</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Careers</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Open Source</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Partners</li>
              <li class="pb-4"><i class="fa fa-chevron-right text-yellow-500"></i><a href="#"
                  class="hover:text-yellow-500"/>Press</li>

            </ul>



          </div>

          

        </div>

      </div>

    </div>
  </footer>
 
    
    </>

  )
}
