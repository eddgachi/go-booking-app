# Conference Ticket Booking Application

## Overview
This is a simple command-line application written in Go that allows users to book tickets for a conference. It manages ticket availability, validates user input, and sends confirmation messages asynchronously using Goroutines.

## Features
- Allows users to book conference tickets.
- Validates user input (name, email, and ticket count).
- Displays the list of first names of attendees.
- Sends a confirmation message asynchronously.
- Prevents overbooking by tracking available tickets.

## Technologies Used
- Go (Golang)
- Goroutines (for asynchronous operations)
- Basic input handling via `fmt.Scan`.

## Installation
1. Install [Go](https://go.dev/) if you haven’t already.
2. Clone the repository:
   ```sh
   git clone <repository_url>
   ```
3. Navigate to the project directory:
   ```sh
   cd conference-booking
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## Usage
1. Run the program and follow the on-screen prompts to enter your details.
2. Enter your first name, last name, email, and the number of tickets you want to book.
3. If valid, the application will book your tickets and send a confirmation asynchronously.
4. If the entered details are incorrect, the application will display appropriate error messages.
5. The program will continue running until all tickets are booked.

## Example Output
```sh
Welcome to our Go conference booking application
We have a total of 50 tickets and 50 are still available
Get your tickets here to attend

Enter your first name: John
Enter your last name: Doe
Enter your email address: john.doe@example.com
Enter number of tickets: 2

Thank you John Doe for booking 2 tickets. You will receive a confirmation email at john.doe@example.com
48 tickets remaining for Go conference
```

## License
This project is licensed under the MIT License.


