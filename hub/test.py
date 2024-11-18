from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# Set Chrome options
chrome_options = Options()
chrome_options.add_argument("--start-maximized")  # Maximize the window

# Configure Selenium Grid
driver = webdriver.Remote(
	command_executor="http://localhost:8081/wd/hub",
	options=chrome_options
)

try:
	driver.get("http://172.17.0.1:8080")
	print(driver.page_source)
	driver.find_element(By.NAME, "username").send_keys("space")
	driver.find_element(By.NAME, "password").send_keys("12245589")
	driver.find_element(By.XPATH, '//input[@value="Login"]').click()
	WebDriverWait(driver, 10).until(
			EC.invisibility_of_element((By.XPATH, "//input[@value='Login']")))
	assert "Dashboard" in driver.page_source

	# Logout
	assert "Logout" in driver.page_source

	driver.find_element(By.LINK_TEXT, 'Logout').click()
	assert "Login" in driver.page_source

finally:
	# Close the browser session
	driver.quit()
