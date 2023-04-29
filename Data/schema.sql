CREATE TABLE IF NOT EXISTS `DailyNutriPlanner`.`everyday_nutrition` (
    `Name` varchar(19),
    `Analysis_item` varchar(12),
    `Content_per_unit` varchar(20),
    `Unit` varchar(4),
    PRIMARY KEY (`Name`, `Analysis_item`)
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;