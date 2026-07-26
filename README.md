<div align="center">
    <img src=".public/logo.svg">
</div>

<h1 align="center">Hydron</h1>

<div align="center">

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

</div>

---  

A lightweight, high-throughput traffic generation tool built in Go. Designed to test target servers, custom firewalls, and reverse proxies using rate-limiting, concurrency pools, and header-based client IP simulation.

**Note**: Actively under development. Features and large chunks of code are subject to change.  

## Features

- **Direct IP and Port Targeting**, bypassing DNS lookup  
- Worker Pool Architecture supporting **concurrent execution** with **goroutines** 
- **Rate Limiting** to control requests per second    
- Simulated Source IPs **(X-Forwarded-For)** to simulate realistic traffic patterns for proxy/firewall evaluation.

## Prerequisites
- Go 1.20+

## Usage

Clone the repository

```
git clone https://github.com/edmuri/Hydron.git
cd Hydron  
```

```
go run .
```