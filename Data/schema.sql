CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`everyday_nutrition` (
    `Name` varchar(20),
    `Analysis_item` varchar(20),
    `Unit` varchar(4),
    `Content_per_unit` float,
    `Weight_per_unit` varchar(20),
    `Content_per_unit_weight` float,
    PRIMARY KEY (`Name`, `Analysis_item`)
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;