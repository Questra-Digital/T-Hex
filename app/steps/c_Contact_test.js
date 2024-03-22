const { I } = require('codeceptjs');

Feature('Dockerfile');

Scenario('Test Dockerfile', async ({ I }) => {
  I.amOnPage('http://localhost:3000');  
  // Click on the "Contact" link
  I.click('Contact');
  I.wait(2);
  I.executeScript(() => {
    window.scrollBy(0, 400); // Adjust the value as needed
  });
  I.wait(2);
  I.executeScript(() => {
    window.scrollBy(800, 0); 
  });
  I.wait(2);

  I.fillField('username', 'Mr jawad Khalid Cheema');
  I.wait(1);
  I.fillField('Email', 'JawadKhalid@cfd.nu.edu.pk');
  I.wait(1);
  I.fillField('message', 'Hello Abdullah Ramsha And Huzaifa You Guys Are doing Good now Well Done Thankyou <3');
  I.wait(2);
 
  I.click('SEND');
});
