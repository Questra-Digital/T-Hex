// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('Dockerfile');

Scenario('test something on the homepage', async ({ I }) => {
  // Navigate to the homepage
  I.amOnPage('http://localhost:3000');

  // Wait for a certain amount of time (e.g., 2 seconds)
  I.wait(2);

  // Assert that the "Web Test Hub" title is present on the page
  I.see('Web Test Hub');
  // Click the "Get Started" button
  I.click('Get Started');
  I.click('Write DockerFile');
  I.wait(2);
  // Wait for a certain amount of time (e.g., 3 seconds) after clicking the button
  I.click('Confirm');

});
