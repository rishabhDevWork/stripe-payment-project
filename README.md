# Stripe Payment Integration

This project is a full-stack Stripe integration with a Go backend and React frontend.

## Structure
- backend/: Go server that handles the Stripe API
- frontend/: React app that renders the checkout UI

## Prerequisites
- Node.js and npm (for frontend)
- Go (for backend)
- Stripe account with API keys

## Installation

### Frontend
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```

### Backend
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

### Frontend
1. From the frontend directory, start the React development server:
   ```bash
   npm start
   ```
   The frontend will run on `http://localhost:3000` by default.

### Backend
1. From the backend directory, run the Go server:
   ```bash
   go run .
   ```
   The backend will run on `http://localhost:8080` by default (ensure port matches your configuration).

## Notes
- Ensure Stripe API keys are configured in the backend (e.g., in a `.env` file or environment variables).
- The frontend and backend must be running simultaneously for the integration to work.
