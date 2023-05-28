# How to run
```
docker build -t bitcoin_rate_mailer .
docker run -p 10000:10000 -e ENV_SENDER_EMAIL='<email>' -e ENV_SENDER_PASSWORD='<password>' bitcoin_rate_mailer
```

For the testing purposes following credentials can be used
```
email: btcratemailer@gmail.com
password: tqzeyskmfhuvzfnc
```

Example requests
```
curl 127.0.0.1:10000/rate
{"btcuah":1004698}
```
```
curl -d "email=email@mail.com" 127.0.0.1:10000/subscribe
{"result":"email added"}

curl -d "email=email@mail.com" 127.0.0.1:10000/subscribe
{"result":"email exists"}
```
```
curl 127.0.0.1:10000/sendEmails
{"result":"emails were sent"}
```

# Implementation logic
`handlers.RateHandler` handles requests to /rate endpoint. It requests BTC/UAH rate from the external API


`handlers.SubscribeHandler` handles requests to /subscribe endpoint. It reads email parameter from POST request data, reads emails.data file and adds email to the list in the `emails.data` file if it doesn't exist yet. Before insertion to the file, passed `email` argument is validated and it is added to the file only if email has correct format. If `emails.data` doesn't exist, then it is created before email insertion. If there are some issues with file or `email` argument is abscent in the request, the issue is described in the response data in the `result` field.  


`handlers.SendEmailsHandler` handles requests to /sendEmails endpoint. It reads the list of emails from emails.data file, retrieve BTC/UAH rate from external API and sends mail to each of email from the list via Gmail SMTP server. If it failed to read data from file or retrieve BTC rate, it return respons with status 500 and data in JSON format with error description in `result` field