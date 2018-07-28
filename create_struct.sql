/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 * Author:  dima
 * Created: 28.07.2018
 */

CREATE TABLE products(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    name  TEXT NOT NULL,
    product_id TEXT NOT NULL UNIQUE,
    category TEXT NOT NULL,
    quanto INTEGER NOT NULL,
    creationDate TEXT NOT NULL,
    expirationDate TEXT NOT NULL
);


