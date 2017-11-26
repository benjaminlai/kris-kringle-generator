# Kris Kringle Name Generator - With Rules/Blacklist

A Golang app that will read a list of participants and randomly assign another participant as the receiver. 
It will not assign a receiver who is on the blacklist

## Getting Started

Rename ./app/resources/sample.json to participants.json
Sample ./app/resources/participants.json payload
```json[
    {
        "name":"Participant1",
        "phone":"1111111",
        "blacklist": [
            {
                "name": "Receiver1",
                "phone": "222222"
            },
            {
                "name": "Receiver2",
                "phone": "22222222"
            }
        ]
    },
    {
        "name":"Participant2",
        "phone":"2222222",
        "blacklist": [
            {
                "name": "Receiver3",
                "phone": "3333333"
            }
        ]
    }
]```

### Prerequisites

Must have: 
- Go
- ./app/resources/participants.json to be populated
- Twilio Account
- Add participants to Twilio Verified Numbers

Requires following ENV Parameters to be stored in .env file in root directory
```
TWILIO_SMS_BASEURL=https://api.twilio.com/2010-04-01/Accounts
TWILIO_ACCOUNT_SID=ACxxxxxxxxxx
TWILIO_AUTH_TOKEN=xxxxxxxxxxxxxxx
TWILIO_NUMBER_FROM=+61411234567
```

### Installing

- Go Build 
- ./kris-kringle-generator

## Authors

* **Benjamin Lai** - [benjaminlai](https://github.com/benjaminlai)

## Acknowledgments

* Twilio
