CREATE TABLE articles 
(
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    content text NOT NULL,
    date_created datetime NOT NULL,
    tags json default NULL,
    PRIMARY KEY (id)
);
