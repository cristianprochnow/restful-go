# RESTful
🐰 REST Api sample made with Golang.

# Routes

## `/cars` GET

### Response

```json
[
	{
		"id": 1,
		"model": "Civic",
		"brand": "Honda",
		"price": 164000.5
	},
	{
		"id": 2,
		"model": "Corola",
		"brand": "Toyota",
		"price": 145000.8
	},
	{
		"id": 3,
		"model": "Cerato",
		"brand": "Kia",
		"price": 123500.5
	},
	{
		"id": 4,
		"model": "Elantra",
		"brand": "Hyundai",
		"price": 118560.5
	},
	{
		"id": 5,
		"model": "HB20",
		"brand": "Hyundai",
		"price": 100000.67
	},
	{
		"id": 6,
		"model": "HB20",
		"brand": "Hyundai",
		"price": 100000.67
	}
]
```

## `/cars` POST

### Body

```json
{
	"price": 100000.67,
	"model": "Cerato",
	"brand": "Kia"
}
```

### Response

```json
{
	"id": 6,
	"model": "HB20",
	"brand": "Hyundai",
	"price": 100000.67
}
```

## `/cars` PUT

### Body

```json
{
	"price": 100000.67,
	"model": "Corola",
	"brand": "Toyota"
}
```

### Response

```json
{
	"id": 2,
	"model": "Corola",
	"brand": "Toyota",
	"price": 100000.67
}
```