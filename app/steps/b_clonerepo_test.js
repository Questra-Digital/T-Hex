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
  I.fillField('email', 'Jawadkhalid@cfd.nu.edu.pk');
  I.wait(2);
  I.fillField('test', 'FYP-Sample-Python-Tesstts');
  I.wait(2);
  I.fillField('number', '3');
  I.wait(2);
  I.fillField('fname', 'python');
  I.wait(2);
  I.fillField('projectlang', 'python');
  I.wait(2);
  I.click('Save Details');
  
});