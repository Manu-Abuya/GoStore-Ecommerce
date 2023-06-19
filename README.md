# GoStore-Ecommerce

Welcome to GoStore-Ecommerce! This is an e-commerce project built with Go that allows users to create an online store and sell products. 

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- User registration and authentication
- Store creation and management
- Product listing and searching
- Shopping cart functionality
- Secure payment processing
- Order management and tracking
- User reviews and ratings
- Admin dashboard for site management

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/GoStore-Ecommerce.git
   
2. Install dependancies:
    ```bash
    go mod download
    ```
    
3. Set up the database:

  - Create a new PostgreSQL database.
  - Update the database configuration in config/config.go with your database credentials.
  
4. Build and run the application:

    ```bash
    go build
    ./shopgo
    ```
    
5. Access the application:

      Open your web browser and navigate to http://localhost:8080.
      
      
## Usage

   - Register a new user account or log in with existing credentials.
   - Create a new store or browse existing stores.
   - Add products to your store or explore products from other stores.
   - Manage your shopping cart and proceed to checkout.
   - Track your orders and update their status.
   - Leave reviews and ratings for products you have purchased.
   - Access the admin dashboard (if authorized) to manage site settings and user accounts.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please submit an issue or create a pull request. Follow these steps to contribute:

   - Fork the repository.
   - Create a new branch for your feature or bug fix.
   - Make your changes and test thoroughly.
   - Commit your changes with descriptive commit messages.
   - Push your branch to your forked repository.
   - Open a pull request to the main repository's main branch.

