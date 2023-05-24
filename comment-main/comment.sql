CREATE DATABASE  IF NOT EXISTS `comment` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `comment`;
-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: localhost    Database: comment
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `reply_content`
--

DROP TABLE IF EXISTS `reply_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply_content` (
  `rpid` bigint NOT NULL AUTO_INCREMENT,
  `message` mediumtext NOT NULL,
  `ats` varchar(255) DEFAULT NULL,
  `ip` bigint NOT NULL,
  `plat` bit(1) NOT NULL,
  `device` varchar(255) NOT NULL,
  `version` varchar(255) NOT NULL,
  `ctime` time NOT NULL,
  `mtime` time NOT NULL,
  `topics` tinytext NOT NULL,
  `addr` varchar(255) NOT NULL,
  KEY `idx_content` (`rpid`),
  CONSTRAINT `reply_content_ibfk_1` FOREIGN KEY (`rpid`) REFERENCES `reply_index` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reply_index`
--

DROP TABLE IF EXISTS `reply_index`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply_index` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `oid` bigint NOT NULL,
  `type` bit(1) NOT NULL,
  `mid` bigint NOT NULL,
  `root` bigint NOT NULL,
  `parent` bigint NOT NULL,
  `floor` int NOT NULL,
  `count` int NOT NULL,
  `rcount` int NOT NULL,
  `likes` int NOT NULL,
  `hates` int NOT NULL,
  `report_count` int NOT NULL,
  `state` bit(1) NOT NULL,
  `ctime` time NOT NULL,
  `mtime` time NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_index` (`id`,`oid`,`type`,`mid`,`report_count`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reply_subject`
--

DROP TABLE IF EXISTS `reply_subject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply_subject` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `oid` bigint NOT NULL,
  `type` bit(1) NOT NULL,
  `mid` bigint NOT NULL,
  `count` int NOT NULL,
  `root_count` int NOT NULL,
  `acount` int NOT NULL,
  `state` bit(1) NOT NULL,
  `ctime` timestamp NOT NULL,
  `mtime` timestamp NOT NULL,
  `attrs` int NOT NULL,
  `meta` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_subject` (`id`,`oid`,`type`,`mid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-20 23:32:50
