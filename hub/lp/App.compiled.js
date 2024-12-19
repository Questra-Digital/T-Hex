"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;
var _react = _interopRequireDefault(require("react"));
var _lucideReact = require("lucide-react");
function _interopRequireDefault(e) { return e && e.__esModule ? e : { default: e }; }
const THexLandingPage = () => {
  const features = [{
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.Cloud, {
      className: "w-8 h-8 text-blue-500"
    }),
    title: "Cloud-Powered Selenium Grid",
    description: "Instantly scale your Selenium test infrastructure without managing complex setups."
  }, {
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.Zap, {
      className: "w-8 h-8 text-yellow-500"
    }),
    title: "Parallel Test Execution",
    description: "Run multiple Selenium tests simultaneously, dramatically reducing your total test runtime."
  }, {
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.Server, {
      className: "w-8 h-8 text-green-500"
    }),
    title: "CI/CD Integration",
    description: "Seamless integration with popular CI/CD platforms for streamlined testing workflows."
  }];
  const steps = [{
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.UserPlus, {
      className: "w-12 h-12 text-blue-600"
    }),
    title: "Register Account",
    description: "Create your T-Hex account in minutes. Get instant access to our cloud Selenium Grid."
  }, {
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.Download, {
      className: "w-12 h-12 text-green-600"
    }),
    title: "Download T-Hex Runner",
    description: "Simple one-click download for Windows, Mac, or Linux. No complex installations or dependencies required."
  }, {
    icon: /*#__PURE__*/_react.default.createElement(_lucideReact.Rocket, {
      className: "w-12 h-12 text-purple-600"
    }),
    title: "Run Tests in Parallel",
    description: "Execute your existing Selenium tests without any modifications. Our binary handles cloud distribution automatically."
  }];
  return /*#__PURE__*/_react.default.createElement("div", {
    className: "min-h-screen bg-white flex flex-col"
  }, /*#__PURE__*/_react.default.createElement("nav", {
    className: "bg-white shadow-sm fixed w-full z-10"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "container mx-auto px-4"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "flex justify-between items-center h-16"
  }, /*#__PURE__*/_react.default.createElement("a", {
    href: "/",
    className: "text-2xl font-bold text-blue-600"
  }, "T-Hex"), /*#__PURE__*/_react.default.createElement("div", {
    className: "flex space-x-4 items-center"
  }, /*#__PURE__*/_react.default.createElement("a", {
    href: "/contactus",
    className: "text-gray-600 hover:text-blue-600 transition"
  }, "Contact Us"), /*#__PURE__*/_react.default.createElement("a", {
    href: "/signup",
    className: "bg-blue-600 text-white px-4 py-2 rounded-full hover:bg-blue-700 transition"
  }, "Sign Up"))))), /*#__PURE__*/_react.default.createElement("div", {
    className: "pt-16"
  }, /*#__PURE__*/_react.default.createElement("header", {
    className: "bg-gradient-to-br from-blue-600 to-purple-700 text-white"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "container mx-auto px-4 py-16 text-center"
  }, /*#__PURE__*/_react.default.createElement("h1", {
    className: "text-5xl font-bold mb-4"
  }, "T-Hex: Supercharge Your Selenium Testing"), /*#__PURE__*/_react.default.createElement("p", {
    className: "text-xl max-w-2xl mx-auto mb-8"
  }, "Run parallel Selenium Grid tests effortlessly in the cloud. Scale your test infrastructure with a single click."), /*#__PURE__*/_react.default.createElement("div", {
    className: "flex justify-center"
  }, /*#__PURE__*/_react.default.createElement("a", {
    href: "/signup?next=getstarted",
    className: "bg-white text-blue-600 px-6 py-3 rounded-full font-semibold hover:bg-blue-50 transition"
  }, "Get Started")))), /*#__PURE__*/_react.default.createElement("section", {
    className: "container mx-auto px-4 py-16"
  }, /*#__PURE__*/_react.default.createElement("h2", {
    className: "text-3xl font-bold text-center mb-12"
  }, "Why Choose T-Hex?"), /*#__PURE__*/_react.default.createElement("div", {
    className: "grid md:grid-cols-3 gap-8"
  }, features.map((feature, index) => /*#__PURE__*/_react.default.createElement("div", {
    key: index,
    className: "text-center p-6 bg-gray-100 rounded-lg shadow-md"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "flex justify-center mb-4"
  }, feature.icon), /*#__PURE__*/_react.default.createElement("h3", {
    className: "text-xl font-semibold mb-3"
  }, feature.title), /*#__PURE__*/_react.default.createElement("p", {
    className: "text-gray-600"
  }, feature.description))))), /*#__PURE__*/_react.default.createElement("section", {
    className: "container mx-auto px-4 py-16 bg-gray-50"
  }, /*#__PURE__*/_react.default.createElement("h2", {
    className: "text-3xl font-bold text-center mb-12"
  }, "Get Started in 3 Simple Steps"), /*#__PURE__*/_react.default.createElement("div", {
    className: "grid md:grid-cols-3 gap-8"
  }, steps.map((step, index) => /*#__PURE__*/_react.default.createElement("div", {
    key: index,
    className: "text-center p-6 bg-white rounded-lg shadow-md border border-gray-200"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "flex justify-center mb-4"
  }, step.icon), /*#__PURE__*/_react.default.createElement("h3", {
    className: "text-xl font-semibold mb-3"
  }, step.title), /*#__PURE__*/_react.default.createElement("p", {
    className: "text-gray-600"
  }, step.description), /*#__PURE__*/_react.default.createElement("div", {
    className: "mt-4 text-5xl font-bold text-gray-300"
  }, "0", index + 1))))), /*#__PURE__*/_react.default.createElement("section", {
    className: "bg-gradient-to-r from-blue-600 to-purple-700 text-white py-16"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "container mx-auto px-4 text-center"
  }, /*#__PURE__*/_react.default.createElement("h2", {
    className: "text-4xl font-bold mb-6"
  }, "Ready to Accelerate Your Testing?"), /*#__PURE__*/_react.default.createElement("p", {
    className: "text-xl max-w-2xl mx-auto mb-8"
  }, "Transform your Selenium testing workflow with cloud-powered parallel execution. Start shipping faster today."), /*#__PURE__*/_react.default.createElement("div", {
    className: "flex justify-center"
  }, /*#__PURE__*/_react.default.createElement("a", {
    href: "/signup?next=getstarted",
    className: "bg-white text-blue-600 px-8 py-4 rounded-full font-semibold hover:bg-blue-50 transition"
  }, "Get Started"))))), /*#__PURE__*/_react.default.createElement("footer", {
    className: "bg-gray-800 text-gray-300 py-12 mt-auto"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "container mx-auto px-4"
  }, /*#__PURE__*/_react.default.createElement("div", {
    className: "grid md:grid-cols-2 gap-8"
  }, /*#__PURE__*/_react.default.createElement("div", null, /*#__PURE__*/_react.default.createElement("h3", {
    className: "text-xl font-semibold mb-4"
  }, "About T-Hex"), /*#__PURE__*/_react.default.createElement("p", {
    className: "mb-4"
  }, "T-Hex is built on and works with Selenium WebDriver\u2122, an open source project made possible by the Selenium community. Selenium\u2122 and WebDriver\u2122 are trademarks of the Software Freedom Conservancy."), /*#__PURE__*/_react.default.createElement("p", null, "We are not affiliated with or endorsed by the Selenium project or Software Freedom Conservancy.")), /*#__PURE__*/_react.default.createElement("div", {
    className: "md:text-right"
  }, /*#__PURE__*/_react.default.createElement("h3", {
    className: "text-xl font-semibold mb-4"
  }, "Connect With Us"), /*#__PURE__*/_react.default.createElement("div", {
    className: "flex flex-col items-end space-y-3"
  }, /*#__PURE__*/_react.default.createElement("a", {
    href: "https://github.com/Questra-Digital/t-hex",
    target: "_blank",
    rel: "noopener noreferrer",
    className: "inline-flex items-center space-x-2 text-gray-300 hover:text-white transition"
  }, /*#__PURE__*/_react.default.createElement(_lucideReact.Github, {
    className: "w-6 h-6"
  }), /*#__PURE__*/_react.default.createElement("span", null, "View on GitHub")), /*#__PURE__*/_react.default.createElement("a", {
    href: "/contactus",
    className: "inline-flex items-center space-x-2 text-gray-300 hover:text-white transition"
  }, /*#__PURE__*/_react.default.createElement(_lucideReact.Mail, {
    className: "w-6 h-6"
  }), /*#__PURE__*/_react.default.createElement("span", null, "Contact Us"))))), /*#__PURE__*/_react.default.createElement("div", {
    className: "border-t border-gray-700 mt-8 pt-8 text-center"
  }, /*#__PURE__*/_react.default.createElement("p", null, "\xA9 ", new Date().getFullYear(), " T-Hex. All rights reserved.")))));
};
var _default = exports.default = THexLandingPage;
