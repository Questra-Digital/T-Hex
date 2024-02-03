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
    driver.get("https://www.google.com")

    # Find the search input field and enter the search query
    search_input = driver.find_element(By.NAME, "Tobefailed")
    search_input.send_keys("browserstack")
    search_input.submit()

    # Wait for the search results to load
    driver.implicitly_wait(10)

    # Find the first search result link and click on it
    first_link = driver.find_element(By.CSS_SELECTOR, 'a[href*="browserstack"]')
    first_link.click()

    # Capture a screenshot of the page
    screenshot_path = "T-HEX/Screenshot.png"
    driver.save_screenshot(screenshot_path)
    logging.info(f"Screenshot saved: {screenshot_path}")

finally:
    driver.quit()
