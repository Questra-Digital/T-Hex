# Scalable Web Testing Platform

In today's software development landscape, software testing has become essential. However, existing solutions for automatic script testing often include excessive and underutilized features, resulting in high costs. This project aims to address this issue by offering a streamlined platform that focuses on essential key features, affordability, and desirability. Our platform provides an easy start-up and scalable test execution, utilizing a cloud-based infrastructure. It supports cross-web testing on multiple platforms and enables parallel test execution. By prioritizing efficiency and cost-effectiveness, our project aims to meet the needs of developers and testers effectively. With a user-friendly interface and cost-effective solution, our platform aims to provide a desirable and accessible option for automated script testing. By optimizing resource utilization, our project ensures efficient testing processes while minimizing financial burdens for users.

# Goals & Objectives

The primary goals and objectives of the project are:
- 	To provide an easy start-up/set-up for scalable test execution.
- 	To provide a cloud-based/hosted platform.
- 	To provide an ability to run on multiple platforms (Cross-web testing).
- 	To provide the ability to run multiple tests in parallel.
- 	To make it affordable


## Getting Started

To get started with the Scalable Web Testing Platform, follow these steps:

1. **Prerequisites**: Ensure you have the following components installed:
   - Node.js
   - Next.js
   - Docker Desktop
   - GraphQL
   - GoLang
   - Psql Cli (for Accessing GraphQL DB in Cmd line)

2. **Clone the Repository**: Clone this GitLab repository to your local machine.

3. **Install Dependencies**: Navigate to the project directory and install the required dependencies by running:
   ```
   npm install
   ```

4. **Configure the Platform**: Open the configuration file `config.js` and customize the settings according to your requirements. Specify the browsers, devices, and test environment details in this file.

5. **Run FE**: Start the Front End by running the following command:
   ```
   cd app/frontend
   npm run dev  
   ```
    It will run on http://localhost:3000

6. **Run Api Gateway**: Start the Appolo Server by running the following command:
   ```
   cd app/Backend/api_gateway
   npm start
   ```
    It will run on http://localhost:4000

7. **Run BE Test Executor**: Start the Test Executor Microservice by running the following command:
   ```
   cd app/Backend/test_executor
   go run server.go
   ```
    It will run on http://localhost:8181

    NOTE: If you make any changes in GraphQl schema run the following Command: `go run github.com/99designs/gqlgen`

8. **Run BE Settings**: Start the Seetings Microservice by running the following command:
   ```
   cd app/Backend/Settings
   go run server.go
   ```
    It will run on http://localhost:9090
    
    NOTE: If you make any changes in GraphQl schema run the following Command: `go run github.com/99designs/gqlgen`

9. **Start Docker**: Start the Docker Desktop.

10. **Trigger Test Execution**: Execute the tests by following the procedure:
   ```
1- Goto http://localhost:3000
2- Click Start Test.
3- Open CloneRepository and Enter to be tested project details and press save.
4- Now enter your github project repository and press Add Repository.
5- Now open customize Settings from sidebar and Configure your own settings or use default.
6- Now open Write DockerFile and write docker script for your project test cases.
7- Click on Start Test button in StartTest page.
8- Results of the test cases will be displayed in the results page.
9- Videos, Logs and Screenshots will be displayed in their respected pages.
10- End the Session be Clicking on EndSession Button.
   ```

## Documentation

For detailed information on using the Scalable Web Testing Platform, refer to the [documentation](https://www.notion.so/T-Hex-Testing-Cloud-Platform-2eaf6146d77a41aca90c4a398fc95443)

## Working CODE Videos

Part-I

https://drive.google.com/file/d/1gD6YvqQyUgUx4AC982P6X0jakV4i-llk/view?usp=drive_link


Part-II

https://drive.google.com/file/d/1RrjonALipR0euRfy-CSmivFWEqeT7IOX/view?usp=drive_link

## Contributions

Contributions to the Scalable Web Testing Platform are welcome! If you find a bug or want to suggest an improvement, please create an issue in the GitLab repository. Feel free to submit pull requests with new features or bug fixes as well.

Please ensure that your contributions align with the project's coding standards and guidelines.


## Acknowledgements

We would like to express our gratitude to the open-source community for providing invaluable tools and libraries that have made this project possible.

## Contact

If you have any questions, suggestions, or feedback, please contact us at [[abdullahsaleem40404@gmail.com, mahmoodahmed2085@gmail.com].

Happy testing!

# Contributors list 
[<img src="https://github.com/F200413.png" width="60px;"/><br /><sub><a href="https://github.com/F200413"> Abdullah Nadeem </a></sub>](https://github.com/F200413/)

[<img src="https://github.com/huzaifaroman.png" width="60px;"/><br /><sub><a href="https://github.com/huzaifaroman"> Huzaifa Roman </a></sub>](https://github.com/huzaifaroman/)