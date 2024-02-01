import { Link } from 'react-router-dom';

function header_component() {

    return (

        <nav class="p-5 bg-blue shadow md:flex md:items-center md:justify-between bg-sky-900">
            <div class="flex justify-between items-center ">
                <span class="text-2xl font-[Poppins] text-slate-50 cursor-pointer">
                    <img class="h-10 inline" src="T_Hex.png" />
                    T-HEX
                </span>

                <span class="text-3xl text-slate-50  cursor-pointer mx-2 md:hidden block">
                    <ion-icon name="menu" />
                </span>
            </div>

            <ul
                class="md:flex md:items-center z-[-1] md:z-auto md:static absolute to-transparent left-0 md:w-auto md:py-0 py-4 md:pl-0 pl-7 md:opacity-100 opacity-0 top-[-400px] transition-all ease-in duration-500">
                <li class="mx-4 my-6 md:my-0 ">
                    <Link to="/" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Home</Link>
                </li>
                <li class="mx-4 my-6 md:my-0">
                    <Link to="/CloneRepository" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Start Test</Link>
                </li>

                <li class="mx-4 my-6 md:my-0">
                    <Link to="/ContactUs" class="text-l text-slate-50 hover:text-cyan-500 duration-500">Contact Us</Link>
                </li>

                <a href="http://localhost:4455/ui/login">
                    <button class="bg-cyan-400 text-white font-[Poppins] duration-500 px-6 py-2 mx-4 hover:bg-cyan-500 rounded  ">
                        Sign in
                    </button>
                </a>

            </ul>
        </nav>

    );
}

export default header_component;
