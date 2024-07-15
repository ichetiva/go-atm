# ATM API
## Version: 1.0


### /accounts

#### POST
##### Summary:

CreateAccount

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Success | [responder.Response](#responder.Response) |

### /accounts/{id}/balance

#### GET
##### Summary:

Get account balance

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | a202482a-bf68-4e17-a4c4-5a268b2e15d6 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Success | [responder.Response](#responder.Response) |
| 400 | Bad request | [responder.Response](#responder.Response) |

### /accounts/{id}/deposit

#### POST
##### Summary:

Deposit cash to account

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | a202482a-bf68-4e17-a4c4-5a268b2e15d6 | Yes | string |
| amount | body | Deposit amount | Yes | number |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Success | [responder.Response](#responder.Response) |
| 400 | Bad request | [responder.Response](#responder.Response) |
| 500 | Internal error | [responder.Response](#responder.Response) |

### /accounts/{id}/withdraw

#### POST
##### Summary:

Withdraw cash from account

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | a202482a-bf68-4e17-a4c4-5a268b2e15d6 | Yes | string |
| amount | body | Withdraw amount | Yes | number |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Success | [responder.Response](#responder.Response) |
| 400 | Bad request | [responder.Response](#responder.Response) |
| 500 | Internal error | [responder.Response](#responder.Response) |

### Models


#### responder.Response

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data |  |  | No |
| message | string |  | No |
| ok | boolean |  | No |