# uptime

# Goals

- Server uptime monitoring
  - Frequency (seconds/minutes/hours/days)
  - url
  - http/https
  - headers

Api's internal, but if required should be done easily.

Api's
## Auth
- Register
- Login
- Logout
- Forgot password
- Change password

## Uptime monitoring
- Add monitoring url. 
  - Url
  - Protocol
  - Frequency
  - Region [Later]
- Real time notification when monitoring site goes offline
  - Email
  - Webhook
- Stats page [Later]

## Db models

### MonitoringUrl

|protocol|url|frequency|

### Notification

|type|

type - Email/webhook

**Type - email**

|type|emailId|

**Type - webhook**

|type|webhookUrl|
