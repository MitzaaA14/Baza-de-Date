-- MySQL Workbench Synchronization
-- Generated: 2025-01-06 18:43
-- Model: New Model
-- Version: 1.0
-- Project: Name of the project
-- Author: Cornel Damian

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `robby` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;

CREATE TABLE IF NOT EXISTS `robby`.`album` (
  `id_album` INT(11) NOT NULL,
  `titlu` TINYTEXT NOT NULL,
  `gen` TINYTEXT NULL DEFAULT NULL,
  `data_lansare` DATE NULL DEFAULT NULL,
  `casa_de_discuri` TINYTEXT NULL DEFAULT NULL,
  `link_coperta_album` TINYTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id_album`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`bilete` (
  `id_bilet` INT(11) NOT NULL,
  `id_concert` INT(11) NOT NULL,
  `id_fan` INT(11) NULL DEFAULT NULL,
  `tip_bilet` TINYTEXT NULL DEFAULT NULL,
  `pret` DECIMAL(8,2) NOT NULL,
  PRIMARY KEY (`id_bilet`),
  INDEX `id_concert` (`id_concert` ASC) VISIBLE,
  INDEX `id_fan` (`id_fan` ASC) VISIBLE,
  CONSTRAINT `bilete_ibfk_1`
    FOREIGN KEY (`id_concert`)
    REFERENCES `robby`.`concerte` (`id_concert`)
    ON DELETE CASCADE,
  CONSTRAINT `bilete_ibfk_2`
    FOREIGN KEY (`id_fan`)
    REFERENCES `robby`.`comunitate` (`id_fan`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`comunitate` (
  `id_fan` INT(11) NOT NULL,
  `nume_fan` TINYTEXT NOT NULL,
  `prenume_fan` TINYTEXT NULL DEFAULT NULL,
  `data_nastere` DATE NULL DEFAULT NULL,
  `email` TINYTEXT NULL DEFAULT NULL,
  `data_cumparare` DATE NULL DEFAULT NULL,
  `telefon` TINYTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id_fan`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`concerte` (
  `id_concert` INT(11) NOT NULL,
  `id_turneu` INT(11) NOT NULL,
  `tara` TINYTEXT NULL DEFAULT NULL,
  `oras` TINYTEXT NULL DEFAULT NULL,
  `locatie` TINYTEXT NULL DEFAULT NULL,
  `data_concert` DATE NULL DEFAULT NULL,
  `capacitate` INT(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id_concert`),
  INDEX `concerte_ibfk_1_idx` (`id_turneu` ASC) VISIBLE,
  CONSTRAINT `concerte_ibfk_1`
    FOREIGN KEY (`id_turneu`)
    REFERENCES `robby`.`turnee` (`id_turneu`)
    ON DELETE RESTRICT)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`echipamente` (
  `id_echipament` INT(11) NOT NULL,
  `tip_echipament` TINYTEXT NULL DEFAULT NULL,
  `nume_echipament` TINYTEXT NULL DEFAULT NULL,
  `producator_echipament` TINYTEXT NULL DEFAULT NULL,
  `model_echipament` TINYTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id_echipament`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`foloseste` (
  `id_concert` INT(11) NOT NULL,
  `id_echipament` INT(11) NOT NULL,
  `cantitate` TINYINT(4) NOT NULL,
  PRIMARY KEY (`id_concert`, `id_echipament`),
  INDEX `id_echipament` (`id_echipament` ASC) VISIBLE,
  CONSTRAINT `foloseste_ibfk_1`
    FOREIGN KEY (`id_concert`)
    REFERENCES `robby`.`concerte` (`id_concert`)
    ON DELETE CASCADE,
  CONSTRAINT `foloseste_ibfk_2`
    FOREIGN KEY (`id_echipament`)
    REFERENCES `robby`.`echipamente` (`id_echipament`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`lista_piese_concert` (
  `id_concert` INT(11) NOT NULL,
  `numar_ordine_piesa` TINYINT(4) NOT NULL,
  `id_piesa` INT(11) NOT NULL,
  `durata_piesa_live` TINYTEXT NOT NULL,
  PRIMARY KEY (`id_piesa`, `id_concert`),
  INDEX `id_concert` (`id_concert` ASC) VISIBLE,
  CONSTRAINT `lista_piese_concert_ibfk_1`
    FOREIGN KEY (`id_concert`)
    REFERENCES `robby`.`concerte` (`id_concert`)
    ON DELETE CASCADE,
  CONSTRAINT `lista_piese_concert_ibfk_2`
    FOREIGN KEY (`id_piesa`)
    REFERENCES `robby`.`piese` (`id_piesa`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`lucreaza` (
  `id_membru` INT(11) NOT NULL,
  `id_concert` INT(11) NOT NULL,
  `castig_concert` DECIMAL(14,2) NOT NULL,
  PRIMARY KEY (`id_membru`, `id_concert`),
  INDEX `id_concert` (`id_concert` ASC) VISIBLE,
  CONSTRAINT `lucreaza_ibfk_1`
    FOREIGN KEY (`id_membru`)
    REFERENCES `robby`.`staff` (`id_membru`),
  CONSTRAINT `lucreaza_ibfk_2`
    FOREIGN KEY (`id_concert`)
    REFERENCES `robby`.`concerte` (`id_concert`)
    ON DELETE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`piese` (
  `id_piesa` INT(11) NOT NULL,
  `id_album` INT(11) NOT NULL,
  `titlu` TINYTEXT NOT NULL,
  `durata` TINYTEXT NOT NULL,
  `compozitor` TINYTEXT NULL DEFAULT NULL,
  `link_piesa` TINYTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id_piesa`),
  INDEX `id_album` (`id_album` ASC) VISIBLE,
  CONSTRAINT `piese_ibfk_1`
    FOREIGN KEY (`id_album`)
    REFERENCES `robby`.`album` (`id_album`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`premii` (
  `id_premiu` INT(11) NOT NULL,
  `id_piesa` INT(11) NULL DEFAULT NULL,
  `denumire_premiu` TINYTEXT NULL DEFAULT NULL,
  `data_acordare` DATE NULL DEFAULT NULL,
  PRIMARY KEY (`id_premiu`),
  INDEX `id_piesa` (`id_piesa` ASC) VISIBLE,
  CONSTRAINT `premii_ibfk_1`
    FOREIGN KEY (`id_piesa`)
    REFERENCES `robby`.`piese` (`id_piesa`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`staff` (
  `id_membru` INT(11) NOT NULL,
  `nume_membru` TINYTEXT NOT NULL,
  `prenume_membru` TINYTEXT NULL DEFAULT NULL,
  `rol` TINYTEXT NULL DEFAULT NULL,
  `email` TINYTEXT NULL DEFAULT NULL,
  `salariu` DECIMAL(8,2) NULL DEFAULT NULL,
  `telefon` TINYTEXT NULL DEFAULT NULL,
  PRIMARY KEY (`id_membru`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robby`.`turnee` (
  `id_turneu` INT(11) NOT NULL,
  `nume_turneu` TINYTEXT NULL DEFAULT NULL,
  `data_inceput` DATE NULL DEFAULT NULL,
  `data_sfarsit` DATE NULL DEFAULT NULL,
  PRIMARY KEY (`id_turneu`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
