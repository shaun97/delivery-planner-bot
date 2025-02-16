# Delivery Planner Bot

This is a Telegram bot for a route delivery planning service.

## Setup

1. Clone the repository.
2. Create a `.env` file with your Telegram bot token:
   ```
   TELEGRAM_TOKEN=your_telegram_token_here
   ```
3. Run the bot:
   ```
   go run main.go
   ```

## Project Structure

- `main.go`: Entry point of the application.
- `config/`: Directory to hold configuration-related files.
- `handlers/`: Directory to hold handler functions for different bot commands.
- `bot/`: Directory to hold bot-related initialization and utility functions.
- `services/`: Directory to hold business logic and services.
- `models/`: Directory to hold data models.
- `utils/`: Directory to hold utility functions.
- `.env`: Environment variables file to store sensitive information like the Telegram token.
- `README.md`: Documentation for the project.
- `.gitignore`: Git ignore file to exclude unnecessary files from version control.
