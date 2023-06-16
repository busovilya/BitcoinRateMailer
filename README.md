# How to run
```
docker build -t rate_mailer .
docker run -p 10000:10000 -e ENV_SENDER_EMAIL='<email>' -e ENV_SENDER_PASSWORD='<password>' rate_mailer
```

# REST API
## Get rate
Returns exchange rate of the coin to the given currency
* **Request**

    `GET rate/<coin>/<currency>`

* **URL params**

    **Required:**

    `coin=[string]`

    `currency=[string]`

* **Success response**
  * **Code**: 200<br />
  **Content**: `<int>`
* **Error response**
  * **Code**: 400<br />
  **Content** `{"error":<string>}`
### Example

**Request**
    curl http://localhost:10000/rate/bitcoin/uah

**Response**

    959270

## Get supported coins
Returns the list of supported coins
* **Request**

    `GET rate/coins`

* **Success response**
  * **Code**: 200<br />
  **Content**: `[<string>, <string>, ...]`
* **Error response**
  * **Code**: 400<br />
  **Content** ` `
### Example

**Request**
    curl http://localhost:10000/coins

**Response**

    ["adazoo","add-xyz-new","adex"]

## Get supported currencies
Returns the list of supported currencies
* **Request**

    `GET rate/currencies`

* **Success response**
  * **Code**: 200<br />
  **Content**: `[<string>, <string>, ...]`
* **Error response**
  * **Code**: 400<br />
  **Content** ` `
### Example

**Request**
    curl http://localhost:10000/currencies

**Response**

    ["usd", "eur", "uah"]

## Subscribe on mailing list
Add emails to the mailing list of coin/currency pair exchange rate 
* **Request**

    `POST rate/subscribe`

* **Data params**

    **Required**

    `email=[string]`

    `coin=[string]`
    
    `currency=[string]`

* **Success response**
  * **Code**: 200<br />
  **Content**: ` `
* **Error response**
  * **Code**: 400<br />
  **Content** ` `
* **Error response**
  * **Code**: 409<br />
  **Content** `{"error":<string>}`

## Send emamils to subscribers 
Send email to each subscriber in the mailing list with information they subscribed for
* **Request**

    `POST rate/sendEmails`

* **Success response**
  * **Code**: 200<br />
  **Content**: ` `
* **Error response**
  * **Code**: 400<br />
  **Content** ` `