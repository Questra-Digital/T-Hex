import logging
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

logging.basicConfig(level=logging.INFO)

hub_url = "http://Selenium_Container:4444/wd/hub"

# Create a new instance of the Chrome driver
chrome_options = Options()
chrome_options.add_argument('--browserName=chrome')

driver = webdriver.Remote(command_executor=hub_url, options=chrome_options)

try:
    # Get the Google homepage
    driver.get("https://www.google.com")

    title = driver.title
    logging.info(f"Python Project says: Search results page title: {title}")
finally:
    driver.quit()
