# Discord Message Deleter

This tool allows you to bulk delete messages from a specific Discord channel using the Discord API. The program is written in Go and provides an easy way to delete messages based on the user's ID.

## How It Works

The tool fetches messages from a specified Discord channel and deletes those that match the provided user ID. It's a simple and effective way to clean up messages that you have sent in a channel.

### Key Features:

- **Bulk Deletion**: Automatically delete all messages in a channel sent by a specific user.
- **Secure**: The tool uses your Discord token to authenticate, ensuring that only authorized deletions occur.
- **Simple**: Input your token, channel ID, and user ID to start deleting messages.

## Getting Started

### Prerequisites

To use this tool, you'll need:

- Go installed on your system.
- Your Discord token.
- The channel ID where you want to delete messages.
- Your user ID.

### Installation

Clone the repository and navigate into it:

```bash
git clone https://github.com/mio-dokuhaki/discordmessagedeleter.git
cd discordmessagedeleter
go build -o discordmessagedeleter discordmessagedeleter.go
./discordmessagedeleter
```

You'll be prompted to enter your Discord token, the target channel ID, and your user ID like this:

```
Enter your token: YOUR_BOT_TOKEN
Enter target channelID: YOUR_CHANNEL_ID
Enter your userID: YOUR_USER_ID
```

The program will fetch the messages from the specified channel and delete those sent by the specified user.

Notes:

1. Ensure that you have the necessary permissions to delete messages in the target channel.
2. Be careful when using this tool, as message deletion is permanent.
