## Cron Job in SMTP
 - What is the SMTP

 `
     SMTP (Simple Mail Transfer Protocol) is an internet standard protocol used for sending emails between servers. It defines how emails are sent, relayed, and delivered over the internet.
 `

 - what is the Cron Job

 `
    A cron job is a scheduled task that runs automatically at a specified time or interval on a Unix/Linux system. It is used for automating repetitive tasks such as backups, log rotation, or running scripts.
 `

## A cron job is written as:

```
* * * * * command-to-execute
| | | | |
| | | | └── Day of the week (0 - 6) [Sunday = 0]
| | | └──── Month (1 - 12)
| | └────── Day of the month (1 - 31)
| └──────── Hour (0 - 23)
└────────── Minute (0 - 59)
```

- In this Code Sender is Send the Email Every `5 Second` and After the `1 minute` it is stop Automatic. Every `5 Second` the receiver is received the Email.

## run The Code

```bash
    Go run main.go
```

