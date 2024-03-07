// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('CloneRepository');

Scenario('test CloneRepository component functionality', async ({ I }) => {
  // Navigate to the page containing the CloneRepository component
  I.amOnPage('http://localhost:3000');
  I.click('Get Started');
  // Wait for a certain amount of time (e.g., 2 seconds)
  I.wait(2);

  // Verify that the "Enter URL of Repository" title is present on the page
  I.see('Enter URL of Repository');

  // Perform additional assertions or interactions as needed

  // Example: Fill the input field with a repository URL
 I.fillField('repoinput', 'https://github.com/F200413/FYP-Sample-Python-Test.git');
  I.wait(2);

  // Click the "Add Repository" button
  I.click('Add Repository');

  // Wait for a certain amount of time (e.g., 3 seconds) after clicking the button
  I.wait(3);

  // Verify that the modal is displayed
  I.see('Project had been added successfully.');
I.wait(1);
  // Add more assertions or interactions as needed

  // Close the modal or perform other actions
});
