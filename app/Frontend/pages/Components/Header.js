import { Link } from 'react-router-dom';

function HeaderComponent() {
    return (
        <nav className="p-5 bg-white shadow md:flex md:items-center md:justify-between bg-sky-900 relative">
            <div className="flex justify-between items-center">
                <span className="text-2xl font-[Poppins] text-black cursor-pointer">
                    <img className="h-14 inline" src="log1.png" alt="Web Test Hub Logo" />
                    WEB TEST HUB 
                </span>

                <span className="text-3xl text-black cursor-pointer mx-2 md:hidden block">
                    <ion-icon name="menu" />
                </span>
            </div>

            <ul className="bg-white md:flex md:items-center z-[-1] md:z-auto md:static absolute to-transparent left-0 md:w-auto md:py-0 py-4 md:pl-0 pl-7 md:opacity-100 opacity-0 top-[-400px] transition-all ease-in duration-500">
                <li className="mx-4 my-6 md:my-0">
                    <Link to="/" className="text-lg text-black hover:text-cyan-500 duration-500">Home</Link>
                </li>
                <li className="mx-4 my-6 md:my-0">
                    <Link to="/CloneRepository" className="text-lg text-black hover:text-cyan-500 duration-500">Start Test</Link>
                </li>

                <li className="mx-4 my-6 md:my-0">
                    <Link to="/ContactUs" className="text-lg text-black hover:text-cyan-500 duration-500">Contact Us</Link>
                </li>

                <li className="mx-4 my-6 md:my-0">
                     <Link to="/Login" className="text-lg text-black hover:text-cyan-500 duration-500">Sign in</Link>
                </li>
            </ul>
        </nav>
    );
}

export default HeaderComponent;
