# Bank Sampoerna Application

This project consists of a backend API built with Go and a frontend application built with Next.js. The application allows users to manage bank accounts (Rekening) and transactions (Transaksi) through a user-friendly interface.

## Table of Contents

- [Bank Sampoerna Application](#bank-sampoerna-application)
  - [Table of Contents](#table-of-contents)
  - [Backend Setup](#backend-setup)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running the Application](#running-the-application)
    - [API Endpoints](#api-endpoints)
  - [Frontend Setup](#frontend-setup)
    - [Prerequisites](#prerequisites-1)
    - [Installation](#installation-1)
    - [Running the Application](#running-the-application-1)
  - [Project Structure](#project-structure)
  - [Usage](#usage)
  - [Documentation](#documentation)
  - [License](#license)

## Backend Setup

### Prerequisites

- Go 1.16+
- PostgreSQL
- PowerShell (if using Windows)

### Installation

1. Unzip Folder :
   test-bank-sampoerna
   cd bank-sampoerna/backend -> backend
   cd bank-sampoerna/backend -> backend

2. Install Dependencies
   go mod tidy
3. Create and configure the database:
    # Open PostgreSQL CLI psql -U yourusername -d postgres 
    # Create the database CREATE DATABASE bank_db; 
    # Exit PostgreSQL CLI \q
    # Run the application go run main.go
    # Open the browser http://localhost:8000 or any

API ENDPOINT
    Create Rekening: POST /rekening
    Read Rekening: GET /rekening/{id}
    Update Rekening: PUT /rekening/{id}
    Delete Rekening: DELETE /rekening/{id}

    Create Transaksi: POST /transaksi
    Read Transaksi: GET /transaksi/{id}
    Update Transaksi: PUT /transaksi/{id}
    Delete Transaksi: DELETE /transaksi/{id}


FRONTEND SETUP

Prerequisites
Node.js 18+
npm 10+

Installation
1. Clone the repository:
   git clone https://github.com/ramchild1998/bank-sampoerna.git
   cd bank-sampoerna/frontend
2. Install Dependencies
   npm install
3. Run the application
   npm run dev
   The application will be available at http://localhost:3000 or any

PROJECT STRUCTURE

bank-sampoerna/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── models/
│   ├── handlers/
│   ├── database/
│   ├── router/
│   ├── migrations/
│   ├── tests/
│   └── docs/
│       └── swagger/
├── frontend/
│   ├── components/
│   ├── pages/
│   ├── public/
│   ├── services/
│   ├── styles/
│   ├── .gitignore
│   ├── package.json
│   └── README.md

Usage
Open the frontend application at http://localhost:3000.

Use the interface to:

Add, view, update, and delete Rekening.

Add, view, update, and delete Transaksi for each Rekening.

Documentation
API documentation is available via Swagger at http://localhost:8000/swagger/index.html.

License
This project is licensed under the MIT License.

