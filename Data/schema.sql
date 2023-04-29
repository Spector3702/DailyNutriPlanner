CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`foodstuff` (
    `Sno.` char(8),
    `Name` varchar(19),
    `Other_name` varchar(60),
    `English_name` varchar(30),
    `Describe` varchar(60),
    PRIMARY KEY (`Sno.`)
    UNIQUE (`Name`)
    
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`everyday_nutrition` (
    `Name` varchar(19),
    `Analysis_item` varchar(12),
    `Unit` varchar(4),
    `Content_per_unit` varchar(20),
    `Weight_per_unit` varchar(10),
    `Content_per_unit_weight` float(24),
    PRIMARY KEY (`Name`,`Analysis_item`),
    FOREIGN KEY (`Name`) REFERENCES `foodstuff`(`Name`) 
    ON DELETE SET NULL ON UPDATE CASCADE
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
