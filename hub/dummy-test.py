import sys
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# Get executor URL from command line
executor_url = sys.argv[1] if len(sys.argv) > 1 else "http://localhost:8888/"

chrome_options = Options()
chrome_options.add_argument("--start-maximized")

driver = webdriver.Remote(
	command_executor=executor_url,
	options=chrome_options
)

try:
	driver.get("https://www.example.com")
	WebDriverWait(driver, 10).until(
			EC.visibility_of(driver.find_element(By.CSS_SELECTOR, "h1")))
	assert "Example Domain" in driver.title, "Title does not match expected value."
	more_info = driver.find_element(By.CSS_SELECTOR, "a")
	print("Found link with text:", more_info.text)
	print("Test passed: Example.com loaded successfully and title is as expected.")

except AssertionError as e:
	print("Test failed:", e)

finally:
	driver.quit()
