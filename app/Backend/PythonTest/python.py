import logging
from selenium import webdriver

logging.basicConfig(level=logging.INFO)

hub_url = "http://Selenium_Container:4444/wd/hub"

# Create a new instance of the Chrome driver
caps = {'browserName': 'chrome'}
driver = webdriver.Remote(desired_capabilities=caps, command_executor=hub_url)

try:
    # Get the Google homepage
    driver.get("https://www.google.com")

    # Find the search field and enter "Selenium"
    search_field = driver.find_element_by_name("q")
    search_field.send_keys("Selenium")

    # Submit the search form
    search_field.submit()

    # Wait for the search results to load
    # time.sleep(5)

    # Get the title of the search results page
    title = driver.title
    logging.info(f"Search results page title: {title}")
finally:
    driver.quit()
