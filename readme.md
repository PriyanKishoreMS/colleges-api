# Indian Colleges API

Click here to go to the API: [https://colleges-api.onrender.com](https://colleges-api.onrender.com/)

This API gives you access to a database containing 43,000 colleges in India and allows you to search and list them by state and district, giving you a complete list of all the colleges in a certain state or district.

This API uses data sourced from data.gov.in/catalog/institutions-aishe-survey, which has been made publicly available by the government of India.

## Endpoints

### Search for colleges

This endpoint allows you to search for colleges by name. You can also specify a state or district to narrow down your search.

```
GET /colleges?search=[college name]
GET /colleges/[state]?search=[college name]
GET /colleges/[state]/[district]?search=[college name]
```

#### Optional query parameters:

- `page`: The page number of the search results (default is 1).
- `limit`: The number of results per page (default is 10).

### Get all states

This endpoint returns a list of all states in India.

```
GET /colleges/states
```

### Get districts by state

This endpoint returns a list of all districts in a given state.

```
GET /colleges/[state]/districts
```

### Get all colleges in state

This endpoint returns a list of all colleges in a given state.

```
GET /colleges/[state]
```

#### Optional query parameters:

- `page`: The page number
- `limit`: The number of results per page (default is 10).

### Get all colleges in district

This endpoint returns a list of all colleges in a given district.

```
GET /colleges/[state]/[district]
```

#### Optional query parameters:

- `page`: The page number
- `limit`: The number of results per page (default is 10).

## Example usage

```bash
curl -X GET https://colleges-api.onrender.com/colleges?limit=2&page=69
```

```json
{
	"colleges": [
		{
			"Name": "21ST CENTURY INTERNATIONAL SCHOOL TRUST KANJIRANGAL",
			"State": "Tamil Nadu",
			"City": "Sivagangai",
			"Address_line1": "Rani Velunachiyar Nagar,",
			"Address_line2": "Kanchirangal"
		},
		{
			"Name": "220028-JSS'S ARTS,SCIENCE & COMMERCE COLLEGE, NANDURBAR.",
			"State": "Maharashtra",
			"City": "Nandurbar",
			"Address_line1": "Shikshak Colony,",
			"Address_line2": "Waghoda Road"
		}
	],
	"count": 42938,
	"currentPage": 69,
	"pages": 21470
}
```

```bash
curl -X GET https://colleges-api.onrender.com/colleges/tamil%20nadu/Kancheepuram?search=hindustan%20institute%20of%20tech
```

```json
{
	"colleges": [
		{
			"Name": "HINDUSTAN INSTITUTE OF TECHNOLOGY AND SCIENCE",
			"State": "Tamil Nadu",
			"City": "Kancheepuram",
			"Address_line1": "Hindustan Nagar,",
			"Address_line2": "Vandalur-Kelambakkam Road"
		}
	],
	"count": 1,
	"currentPage": 1,
	"pages": 1
}
```
