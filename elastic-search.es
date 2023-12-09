// POST PRODUCT
POST /products/_doc
{
    "name": "Awesome Glasses",
    "description": "This is an awesome glasses for casual wear.",
    "price": 199.99,
    "category": "Clothing",
    "brand": "Example Brand"
}

// GET ALL PRODUCT
GET /products/_search
{
    "query": {"match": {
        "name": "Awes Glass"
    }},
    "sort": [
        {"price": "asc"}
    ],
    "from": 10,
    "size": 10
}

// GET DETAIL PRODUCT
GET /products/_doc/KpdcTIwB7ZPJJ2wh4pQd

// DELETE DETAIL PRODUCT
DELETE /products/_doc/KpdcTIwB7ZPJJ2wh4pQd