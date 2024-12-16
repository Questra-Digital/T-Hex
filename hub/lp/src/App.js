import React from 'react';
import { Check, Cloud, Play, Server, Zap, UserPlus, Download, Rocket } from 'lucide-react';

const THexLandingPage = () => {
	const features = [
		{
			icon: <Cloud className="w-8 h-8 text-blue-500" />,
			title: "Cloud-Powered Selenium Grid",
			description: "Instantly scale your Selenium WebDriver tests without managing complex setups."
		},
		{
			icon: <Zap className="w-8 h-8 text-yellow-500" />,
			title: "Parallel Test Execution",
			description: "Run multiple Selenium WebDriver tests simultaneously, dramatically reducing your total test runtime."
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
			title: "Install Test Runner",
			description: "Simple standalone executable binary, available across Windows, Mac, and Linux. Compatible with all major testing frameworks."
		},
		{
			icon: <Rocket className="w-12 h-12 text-purple-600" />,
			title: "Run Tests in Parallel",
			description: "Execute your existing test suites without any code modifications. Instantly scale across our global cloud infrastructure."
		}
	];

	return (
		<div className="min-h-screen bg-white">
		{/* Hero Section */}
		<header className="bg-gradient-to-br from-blue-600 to-purple-700 text-white">
		<div className="container mx-auto px-4 py-16 text-center">
		<h1 className="text-5xl font-bold mb-4">T-Hex: Supercharge Your Selenium Testing</h1>
		<p className="text-xl max-w-2xl mx-auto mb-8">
		Run parallel Selenium Grid tests effortlessly in the cloud. Scale your test infrastructure with a single click.
		</p>
		<div className="flex justify-center space-x-4">
		<button className="bg-white text-blue-600 px-6 py-3 rounded-full font-semibold hover:bg-blue-50 transition">
		Start Free Trial
		</button>
		<button className="border border-white text-white px-6 py-3 rounded-full hover:bg-white hover:text-blue-600 transition">
		View Documentation
		</button>
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
		Join hundreds of development teams using T-Hex to make their Selenium testing faster, more reliable, and effortless.
		</p>
		<div className="flex justify-center space-x-4">
		<button className="bg-white text-blue-600 px-8 py-4 rounded-full font-semibold hover:bg-blue-50 transition">
		Start Free Trial
		</button>
		<button className="border border-white text-white px-8 py-4 rounded-full hover:bg-white hover:text-blue-600 transition">
		Schedule Demo
		</button>
		</div>
		</div>
		</section>
		</div>
	);
};

export default THexLandingPage;
