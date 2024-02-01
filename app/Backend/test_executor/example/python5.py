import logging
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By

logging.basicConfig(level=logging.INFO)

hub_url = "http://Selenium_Container:4444/wd/hub"

# Create a new instance of the Chrome driver
chrome_options = Options()
chrome_options.add_argument('--browserName=chrome')

driver = webdriver.Remote(command_executor=hub_url, options=chrome_options)

try:
    # Get the Google homepage
    driver.get("https://www.browserstack.com")

    # Capture a screenshot of the page
    screenshot_path = "T-HEX/Screenshot.png"
    driver.save_screenshot(screenshot_path)
    logging.info(f"Screenshot saved: {screenshot_path}")

finally:
    driver.quit()
