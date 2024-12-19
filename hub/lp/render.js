const React = require('react');
const ReactDOMServer = require('react-dom/server');
const fs = require('fs');
const path = require('path');

const THexLandingPage = require('./App.compiled.js').default;

// Render the component to static HTML
const html = ReactDOMServer.renderToStaticMarkup(
  React.createElement(THexLandingPage)
);

// Add basic HTML structure
const fullHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>T-Hex Landing Page</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css">
</head>
<body>
  ${html}
</body>
</html>
`;

// Write to a file
const outputPath = path.join(__dirname, 'index.html');
fs.writeFileSync(outputPath, fullHtml, 'utf-8');
console.log(`Static HTML generated at ${outputPath}`);
