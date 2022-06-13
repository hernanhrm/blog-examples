INSERT INTO drinks (name, description)
VALUES ('Soda', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit'),
       ('Water', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit'),
       ('Wine', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit'),
       ('Fruit juice', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit');

INSERT INTO meals (name, description, drink_id)
VALUES ('Fried Chicken', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
        (SELECT id FROM drinks WHERE drinks.name = 'Fruit juice')),
       ('Chicken Soup', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
        (SELECT id FROM drinks WHERE drinks.name = 'Water')),
       ('Bacon', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
        (SELECT id FROM drinks WHERE drinks.name = 'Soda')),
       ('Meat', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
        (SELECT id FROM drinks WHERE drinks.name = 'Wine'));
