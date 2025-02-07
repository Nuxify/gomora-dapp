CREATE TABLE `greeter_contract_event_logs` (
  `tx_hash` varchar(255) NOT NULL DEFAULT '',
  `log_index` int NOT NULL DEFAULT 0,
  `contract_address` varchar(42) NOT NULL DEFAULT '',
  `event` enum('LogSetGreeting') NOT NULL,
  `metadata` json NOT NULL,
  `block_timestamp` timestamp NOT NULL,
  PRIMARY KEY (`tx_hash`, `log_index`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;