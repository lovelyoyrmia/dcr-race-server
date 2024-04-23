CREATE TABLE `users_location` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `uid` varchar(255) UNIQUE NOT NULL,
  `category` varchar(255) NOT NULL,
  `fullname` varchar(255) NOT NULL,
  `latitude` float,
  `longitude` float,
  `altitude` float,
  `timestamp` timestamp DEFAULT CURRENT_TIMESTAMP
);
