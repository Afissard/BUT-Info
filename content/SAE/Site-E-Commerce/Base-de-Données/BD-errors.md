---
title: BD-errors
draft: 
description: 
tags:
---
# Brand

| Error Code | Tested Value                                  | Error Message          |
| ---------- | --------------------------------------------- | ---------------------- |
| 45001      | count(*) from Brand                           | No brands found        |
| 45002      | count(*) from Brand where id_brand = brand_id | No brands found        |
| 45003      | row_count()                                   | Failed to add brand    |
| 45004      | row_count()                                   | Failed to update brand |
| 45005      | row_count()                                   | Failed to delete brand |
# Cart
| Error Code | Tested Value                               | Error Message                      |
| ---------- | ------------------------------------------ | ---------------------------------- |
| 46001      | count(*) from Cart where id_user = user_id | No products found in cart          |
| 46002      | count(*) from Cart where id_user = user_id | Cart is empty                      |
| 46003      | row_count()                                | Failed to add product to cart      |
| 46004      | user_founded                               | User not found                     |
| 46005      | product_founded                            | Product not found                  |
| 46006      | product_quantity_available < quantity      | Not enough quantity available      |
| 46007      | row_count()                                | Failed to remove product from cart |
| 46008      | cart_size                                  | Cart is empty                      |
| 46009      | row_count()                                | Failed to clear cart               |
# Image
| Error Code | Tested Value | Error Message               |
|------------|--------------|-----------------------------|
| 47001      | count(*) from Image where id_product = product_id | No image found for product    |
| 47002      | row_count()  | Failed to add image         |
| 47003      | row_count()  | Failed to update image      |
| 47004      | row_count()  | Failed to delete image      |
# Order
| Error Code | Tested Value                                 | Error Message            |
| ---------- | -------------------------------------------- | ------------------------ |
| 48001      | count(*) from Orders where id_user = user_id | No orders found for user |
| 48002      | row_count()                                  | Failed to add order      |
| 48003      | row_count()                                  | Failed to update order   |
# Product
| Error Code | Tested Value                                                                                                                                                                                                                                                                   | Error Message                           |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --------------------------------------- |
| 49001      | count(*) from `Product`                                                                                                                                                                                                                                                        | No products found                       |
| 49002      | count(*) from `Product` where id_product = product_id                                                                                                                                                                                                                          | Product not found                       |
| 49003      | count(*) from `Product` where name = product_name                                                                                                                                                                                                                              | Product not found                       |
| 49004      | count(*) from `Product` where id_subtype = subtype_id and type_id = (select id_type from `Subtype` where id_subtype = subtype_id) and family_id = (select id_family from `Type` where id_type = type_id) and id_brand = brand_id and price >= min_price and price <= max_price | No products found with the given filter |
| 49005      | count(*) from Review where id_product = product_id                                                                                                                                                                                                                             | No reviews found for product            |
| 49006      | row_count()                                                                                                                                                                                                                                                                    | Failed to add product                   |
| 49007      | row_count()                                                                                                                                                                                                                                                                    | Failed to update product                |
| 49008      | row_count()                                                                                                                                                                                                                                                                    | Product not found                       |
| 49009      | new_price < 0                                                                                                                                                                                                                                                                  | Price cannot be negative                |
| 49010      | row_count()                                                                                                                                                                                                                                                                    | Failed to delete product                |
# Review
| Error Code | Tested Value                                       | Error Message                  |
| ---------- | -------------------------------------------------- | ------------------------------ |
| 50001      | count(*) from Review where id_product = product_id | No reviews found for product   |
| 50002      | count(*) from Review where id_user = user_id       | No reviews found for user      |
| 50003      | row_count()                                        | Failed to add review           |
| 50004      | rating < 0 or rating > 5                           | Rating must be between 0 and 5 |
| 50005      | row_count()                                        | Failed to update review        |
| 50006      | rating_val < 0 or rating_val > 5                   | Rating must be between 0 and 5 |
| 50007      | row_count()                                        | Failed to delete review        |
# Shipping
| Error Code | Tested Value                                  | Error Message                   |
| ---------- | --------------------------------------------- | ------------------------------- |
| 52001      | count(*) from Located where id_user = user_id | No location found for user      |
| 52002      | row_count()                                   | Failed to add location          |
| 52003      | row_count()                                   | Failed to link location to user |
| 52004      | row_count()                                   | Failed to update location       |
| 52005      | row_count()                                   | Failed to delete location link  |
| 52006      | row_count()                                   | Failed to delete location       |
# User
| Error Code | Tested Value                                | Error Message             |
| ---------- | ------------------------------------------- | ------------------------- |
| 53001      | count(*) from Users where id_user = user_id | User not found            |
| 53002      | row_count()                                 | Failed to add user        |
| 53003      | row_count()                                 | Failed to update user     |
| 53004      | row_count()                                 | Failed to delete user     |
| 53005      | row_count()                                 | Failed to update password |