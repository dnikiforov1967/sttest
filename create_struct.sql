/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 * Author:  dima
 * Created: 28.07.2018
 */

DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS products;

CREATE TABLE products(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    name  TEXT NOT NULL,
    product_id TEXT NOT NULL UNIQUE,
    category TEXT NOT NULL,
    quanto NUMERIC NOT NULL,
    creationDate TEXT NOT NULL,
    expirationDate TEXT NOT NULL
);

CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    parent_id INTEGER NOT NULL,
    eventType TEXT NOT NULL,
	terminal NUMERIC NOT NULL,
	kind TEXT NOT NULL,
	origin TEXT NOT NULL,
	execType TEXT NOT NULL,
	path TEXT NOT NULL,
	cashType TEXT NOT NULL,
	paymentType TEXT NOT NULL,
	method TEXT NOT NULL,
	algorithmId TEXT NOT NULL,
    FOREIGN KEY(parent_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX events_idx0 ON events(parent_id);

