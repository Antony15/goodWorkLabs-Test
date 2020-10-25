# goodworklabs-test

Built with Golang for microservices and used Redis for cache memory.

Used the Here Maps Places Search API’s to find Parking spots, Charging Stations and Restaurants near the user provided location.

Built a microservice, that will expose one end point for capturing the request and return the response as per the business rules below.

Business rules:
 Output must contain 3 closest POI’s of each type in the response.
 The calls to all the three services must be done in parallel.
 To avoid multiple calls to the provider (Here Maps) the results must ideally be cached in memory.

Input : 
{
    "latitude":"37.7942",
    "longitude":"-122.4070"
}

Output:
{
    "results": {
        "items": [
            {
                "type": "Parking spots",
                "closest_POIs": [
                    {
                        "distance": 803,
                        "title": "Parking",
                        "averageRating": 0,
                        "id": "840jx7ps-28aedefdc9bd011739cd0d8eddef8eca"
                    },
                    {
                        "distance": 1294,
                        "title": "Lombardi Sports",
                        "averageRating": 0,
                        "id": "8409q8zn-4cff04fcd0da4f3080a9b7ea19a9fd14"
                    }
                ]
            },
            {
                "type": "Charging Stations",
                "closest_POIs": [
                    {
                        "distance": 179,
                        "title": "Ccsf Sfmta / Portsmouth Sq 2",
                        "averageRating": 0,
                        "id": "840if4l9-a373ffa524ab076756e9c983ca6ffd5d"
                    },
                    {
                        "distance": 183,
                        "title": "CCSF SFMTA / Portsmouth Squr",
                        "averageRating": 0,
                        "id": "8409q8zn-54aaf78542804c12b628f41eb8fd77a7"
                    },
                    {
                        "distance": 246,
                        "title": "Columbia Reit - 650 California, Llc",
                        "averageRating": 0,
                        "id": "8409q8yy-66f342437d334834b66185a0014a9892"
                    }
                ]
            },
            {
                "type": "Restaurants",
                "closest_POIs": [
                    {
                        "distance": 219,
                        "title": "Z & Y Restaurant",
                        "averageRating": 0,
                        "id": "8409q8zn-9d9117b1642b4ff28135ee00620646ec"
                    },
                    {
                        "distance": 230,
                        "title": "Great Eastern Restaurant",
                        "averageRating": 0,
                        "id": "8409q8zn-22aaf504dab24fa88334633518a3fa86"
                    },
                    {
                        "distance": 254,
                        "title": "Schilling & Co. Cafe",
                        "averageRating": 0,
                        "id": "840aabd1-f4c8e3f2bf420851aa0f7ada80e358b6"
                    }
                ]
            }
        ]
    }
}

Package and delivered the final outcome as a Docker container => https://hub.docker.com/r/antony15/goodworklabs-test/tags
