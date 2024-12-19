import React from 'react';
import { UserPlus, Download, Rocket, Cloud, Zap, Server, Github, Mail } from 'lucide-react';

const THexLandingPage = () => {
	const features = [
		{
			icon: <Cloud className="w-8 h-8 text-blue-500" />,
			title: "Cloud-Powered Selenium Grid",
			description: "Instantly scale your Selenium test infrastructure without managing complex setups."
		},
		{
			icon: <Zap className="w-8 h-8 text-yellow-500" />,
			title: "Parallel Test Execution",
			description: "Run multiple Selenium tests simultaneously, dramatically reducing your total test runtime."
		},
		{
			icon: <Server className="w-8 h-8 text-green-500" />,
			title: "CI/CD Integration",
			description: "Seamless integration with popular CI/CD platforms for streamlined testing workflows."
		}
	];

	const steps = [
		{
			icon: <UserPlus className="w-12 h-12 text-blue-600" />,
			title: "Register Account",
			description: "Create your T-Hex account in minutes. Get instant access to our cloud Selenium Grid."
		},
		{
			icon: <Download className="w-12 h-12 text-green-600" />,
			title: "Download T-Hex Runner",
			description: "Simple one-click download for Windows, Mac, or Linux. No complex installations or dependencies required."
		},
		{
			icon: <Rocket className="w-12 h-12 text-purple-600" />,
			title: "Run Tests in Parallel",
			description: "Execute your existing Selenium tests without any modifications. Our binary handles cloud distribution automatically."
		}
	];

	return (
		<div className="min-h-screen bg-white flex flex-col">
			{/* Navigation Header */}
			<nav className="bg-white shadow-sm fixed w-full z-10">
				<div className="container mx-auto px-4">
					<div className="flex justify-between items-center h-16">
						<a href="/" className="text-2xl font-bold text-blue-600">T-Hex</a>
						<div className="flex space-x-4 items-center">
							<a href="/contactus" className="text-gray-600 hover:text-blue-600 transition">
								Contact Us
							</a>
							<a
								href="/signup"
								className="bg-blue-600 text-white px-4 py-2 rounded-full hover:bg-blue-700 transition"
							>
								Sign Up
							</a>
						</div>
					</div>
				</div>
			</nav>

			<div className="pt-16">
				{/* Hero Section */}
				<header className="bg-gradient-to-br from-blue-600 to-purple-700 text-white">
					<div className="container mx-auto px-4 py-16 text-center">
						<h1 className="text-5xl font-bold mb-4">T-Hex: Supercharge Your Selenium Testing</h1>
						<p className="text-xl max-w-2xl mx-auto mb-8">
							Run parallel Selenium Grid tests effortlessly in the cloud. Scale your test infrastructure with a single click.
						</p>
						<div className="flex justify-center">
							<a
								href="/signup?next=getstarted"
								className="bg-white text-blue-600 px-6 py-3 rounded-full font-semibold hover:bg-blue-50 transition"
							>
								Get Started
							</a>
						</div>
					</div>
				</header>

				{/* Features Section */}
				<section className="container mx-auto px-4 py-16">
					<h2 className="text-3xl font-bold text-center mb-12">Why Choose T-Hex?</h2>
					<div className="grid md:grid-cols-3 gap-8">
						{features.map((feature, index) => (
							<div key={index} className="text-center p-6 bg-gray-100 rounded-lg shadow-md">
								<div className="flex justify-center mb-4">
									{feature.icon}
								</div>
								<h3 className="text-xl font-semibold mb-3">{feature.title}</h3>
								<p className="text-gray-600">{feature.description}</p>
							</div>
						))}
					</div>
				</section>

				{/* Steps Section */}
				<section className="container mx-auto px-4 py-16 bg-gray-50">
					<h2 className="text-3xl font-bold text-center mb-12">Get Started in 3 Simple Steps</h2>
					<div className="grid md:grid-cols-3 gap-8">
						{steps.map((step, index) => (
							<div key={index} className="text-center p-6 bg-white rounded-lg shadow-md border border-gray-200">
								<div className="flex justify-center mb-4">
									{step.icon}
								</div>
								<h3 className="text-xl font-semibold mb-3">{step.title}</h3>
								<p className="text-gray-600">{step.description}</p>
								<div className="mt-4 text-5xl font-bold text-gray-300">
									0{index + 1}
								</div>
							</div>
						))}
					</div>
				</section>

				{/* CTA Section */}
				<section className="bg-gradient-to-r from-blue-600 to-purple-700 text-white py-16">
					<div className="container mx-auto px-4 text-center">
						<h2 className="text-4xl font-bold mb-6">Ready to Accelerate Your Testing?</h2>
						<p className="text-xl max-w-2xl mx-auto mb-8">
							Transform your Selenium testing workflow with cloud-powered parallel execution. Start shipping faster today.
						</p>
						<div className="flex justify-center">
							<a
								href="/signup?next=getstarted"
								className="bg-white text-blue-600 px-8 py-4 rounded-full font-semibold hover:bg-blue-50 transition"
							>
								Get Started
							</a>
						</div>
					</div>
				</section>
			</div>

			{/* Footer */}
			<footer className="bg-gray-800 text-gray-300 py-12 mt-auto">
				<div className="container mx-auto px-4">
					<div className="grid md:grid-cols-2 gap-8">
						<div>
							<h3 className="text-xl font-semibold mb-4">About T-Hex</h3>
							<p className="mb-4">
								T-Hex is built on and works with Selenium WebDriver™, an open source project made possible
								by the Selenium community. Selenium™ and WebDriver™ are trademarks of the Software
								Freedom Conservancy.
							</p>
							<p>
								We are not affiliated with or endorsed by the Selenium project or Software Freedom
								Conservancy.
							</p>
						</div>
						<div className="md:text-right">
							<h3 className="text-xl font-semibold mb-4">Connect With Us</h3>
							<div className="flex flex-col items-end space-y-3">
								<a
									href="https://github.com/Questra-Digital/t-hex"
									target="_blank"
									rel="noopener noreferrer"
									className="inline-flex items-center space-x-2 text-gray-300 hover:text-white transition"
								>
									<Github className="w-6 h-6" />
									<span>View on GitHub</span>
								</a>
								<a
									href="/contactus"
									className="inline-flex items-center space-x-2 text-gray-300 hover:text-white transition"
								>
									<Mail className="w-6 h-6" />
									<span>Contact Us</span>
								</a>
							</div>
						</div>
					</div>
					<div className="border-t border-gray-700 mt-8 pt-8 text-center">
						<p>&copy; {new Date().getFullYear()} T-Hex. All rights reserved.</p>
					</div>
				</div>
			</footer>
		</div>
	);
};

export default THexLandingPage;
