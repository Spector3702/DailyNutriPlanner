CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`foodstuff` (
    `Sno` char(8),
    `Name` varchar(60), -- select ? 
    `Other_name` varchar(60),
    `English_name` varchar(30),
    `Describe` varchar(60), 
    PRIMARY KEY (`Sno`),
    UNIQUE (`Name`)
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`everyday_nutrition` (
    `Name` varchar(60), -- select ? 
    `Analysis_item` varchar(12),
    `Unit` varchar(4),
    `Content_per_unit` varchar(20),
    `Weight_per_unit` varchar(10),
    `Content_per_unit_weight` float,
    PRIMARY KEY (`Name`,`Analysis_item`),
    FOREIGN KEY (`Name`) REFERENCES `foodstuff`(`Name`) 
    ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`recommended_nutrition` (
    `Gender` ENUM('male','female'),
    `Age` ENUM('<1','1-6','7-12','13-15','16-18','19-44','45-64','65-74','75+'),
    `Numbers_of_people` int,
    `Calorie` float,
    `Protein` float,
    `Fat` float,
    `Carbohydrate` float,
    `VitaminB1` float,
    `VitaminB2` float,
    `VitaminC` float,
    `Nicotine` float,
    `VitaminB6` float,
    `VitaminA` float,
    `VitaminE` float,
    `Ca` float,
    `P` float,
    `Fe` float,
    `Mg` float,
    `Zn` float,
    `Na` float,
    `K` float,
    PRIMARY KEY (`Gender`, `Age`)
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`personal_info` (
    `Email` varchar(60),
    `Gender` ENUM('male','female'),
    `Age` ENUM('<1','1-6','7-12','13-15','16-18','19-44','45-64','65-74','75+'),
    `Weight` float,
    `Height` float,
    `Work_load` ENUM('low','medium','high'),
    PRIMARY KEY (`Email`),
    FOREIGN KEY (`Gender`, `Age`) REFERENCES `recommended_nutrition`(`Gender`, `Age`)
    ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
