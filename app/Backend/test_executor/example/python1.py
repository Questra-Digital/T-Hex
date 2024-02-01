import logging
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
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

    # Capture a screenshot
    screenshot_path1 = "T-HEX/Screenshot1.png"
    driver.save_screenshot(screenshot_path1)
    logging.info(f"Screenshot saved: {screenshot_path1}")

    # Find the search field and type "calculator"
    search_field = driver.find_element(By.NAME, "q")
    search_field.send_keys("calculator")

    # Capture a screenshot
    screenshot_path2 = "T-HEX/Screenshot2.png"
    driver.save_screenshot(screenshot_path2)
    logging.info(f"Screenshot saved: {screenshot_path2}")

    # Press Enter to perform the search
    search_field.send_keys(Keys.RETURN)

    # Wait for the search results to load
    WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.ID, "result-stats")))

    # Get the title of the search results page
    title = driver.title
    logging.info(f"Python-1 Project says: Search results page title: {title}")

    # Capture a screenshot
    screenshot_path3 = "T-HEX/Screenshot3.png"
    driver.save_screenshot(screenshot_path3)
    logging.info(f"Screenshot saved: {screenshot_path3}")

finally:
    driver.quit()
