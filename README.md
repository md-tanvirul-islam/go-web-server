# Go HTTP Form Server

A simple HTTP server in Go that handles form submissions and serves static files.

## Features

- Serves static files from `./static` directory
- Handles GET requests to `/hello` with a simple response
- Displays a form at `/form` for GET requests
- Processes form submissions at `/form` for POST requests
- Bootstrap 5 styled form interface

## Prerequisites

- Go 1.16 or higher
- Web browser for testing

## Installation & Usage

1. Clone the repository or copy the code
2. Create a `static` directory and add your `form.html` file
3. Run the server:

```bash
go run main.go
```

4. Access the application at:
   - Form page: http://localhost:8080/form
   - Hello endpoint: http://localhost:8080/hello
   - Static files: http://localhost:8080/ (serves files from ./static)

## Project Structure

```
.
├── main.go            # Main server code
└── static/
    ├── form.html      # Bootstrap form template
    └── ...            # Other static files (CSS, JS, images)
```

## Endpoints

- `GET /` - Serves static files
- `GET /hello` - Returns "Hello!" response
- `GET /form` - Displays submission form
- `POST /form` - Processes form data

## Form Fields

The form expects these fields:
- `name` (text)
- `email` (email)
- `address` (text)

## Example Response

After form submission, you'll see:
```
POST request successful.
Name: John Doe
Email: john@example.com
Address: 123 Main St
```