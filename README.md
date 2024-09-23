# Algorithm-Visualizer Project

## Overview

The Algorithm-Visualizer project is a web-based application that provides an interactive platform for visualizing and exploring various algorithms. The project consists of two main components:

* Backend API Service: Built using Go and the Fiber framework, this service provides a RESTful API for handling algorithm-related requests.
* Frontend Web App: Built using Next.js and React, this web application provides a user-friendly interface for interacting with the algorithm visualizer.

## Features

* Algorithm Visualization: The application allows users to visualize and explore various algorithms, including sorting algorithms.
* Interactive Interface: The web app provides an interactive interface for users to input data, select algorithms, and visualize the results.
* API-Driven: The backend API service handles requests from the frontend web app, providing a scalable and maintainable architecture.

## Getting Started

### Backend API Service

1. Clone the repository: `git clone https://github.com/PoowadolDev/Algorithm-Visualizer.git`
2. Navigate to the api-service directory: `cd api-service`
3. Run the API service: `go run main.go`

### Frontend Web App

1. Clone the repository: `git clone https://github.com/PoowadolDev/Algorithm-Visualizer.git`
2. Navigate to the web-app directory: `cd web-app`
3. Install dependencies: `npm install` or `yarn install`
4. Run the web app: `npm run dev` or `yarn dev`

## API Documentation

- ### GenerateData

```
  GET /generateData?size={size}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `size`      | `string` | **Required**. Size of generate array data |

- ### SortProblem

```
  POST /SortProblem
```

#### Request Body (Selection, Merge, Inserted, Bubble)
```json
{
    "sortType": "Selection",
    "dataList": [1,2,3,4]
}
```

## Contributing

Contributions are welcome! If you'd like to contribute to the project, please fork the repository and submit a pull request.

## License

The Algorithm-Visualizer project is licensed under the MIT License.

## Acknowledgments

* The project uses the Fiber framework for building the backend API service.
* The project uses Next.js and React for building the frontend web app.
* The project uses Chart.js for creating interactive charts.