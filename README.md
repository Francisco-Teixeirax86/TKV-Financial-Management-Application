# TKV-Financial-Management-Application
# Project Overview
The DADTKV Financial Management Application is a distributed, transactional key-value store system adapted into a financial transaction management app. This project demonstrates the implementation of a backend using a distributed system architecture, combined with an  front-end interface for managing financial transactions. This is a project developed as a form of me learning and aquiring new skills. Maybe it is ambitious but it's the way I gain motivation to learn.

# Key Features
- Distributed Key-Value Store: Implements a distributed system for storing and managing key-value pairs where each key represents an account, and the value represents the account balance.
- Transactional Consistency: Ensures that all financial transactions (like transfers between accounts) are executed with strict serializability, maintaining consistency across the system.
- Concurrency Control: Uses leases and consensus algorithms (e.g., Paxos) to manage concurrent transactions, ensuring no conflicting operations occur.
- Fault Tolerance: The system is designed to handle failures, ensuring that transactions are completed reliably, even in the presence of node failures.
- Front-End Interface: A user-friendly web application that allows users to:
    - View account balances.
    - Perform transactions (transfers, deposits, withdrawals).
    - View transaction history and account logs.
- Security: Incorporates basic security measures, such as authentication and secure communication, to protect sensitive financial data.

# Tech Stack
Backend
  - Programming Language: Go (Golang)
  - Networking: gRPC for communication between distributed components.
  - Concurrency: Go’s goroutines and channels for handling concurrent transactions.
  - Consensus Algorithm: Paxos (for lease management and transaction consistency).
  - Data Storage: In-memory key-value store for rapid access and updates.

Front-End
  - Framework: React.js / Angular (TBD) 
  - State Management: Redux 
  - UI Components: Material-UI or Bootstrap 

Security
  - Authentication: JWT (JSON Web Tokens) for secure user authentication.
  - Encryption: TLS/SSL for secure communication between the client and server.

Developed using Jet Brains IDE's and Linux Ubuntu.

License
This project is licensed under the MIT License - see the LICENSE file for details.
