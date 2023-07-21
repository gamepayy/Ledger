package commands

const usersTableCreation = "CREATE TABLE `users` ( `address` varchar(255) NOT NULL, `balance` varchar(255), PRIMARY KEY (`address`) )"
const tokensBalancesTableCreation = "CREATE TABLE `tokens_balances` ( `address` varchar(255) NOT NULL, `token_address` varchar(255) NOT NULL, `balance` varchar(255), PRIMARY KEY (`address`, `token_address`) )"
const tokensTableCreation = "CREATE TABLE `tokens` (`address` varchar(256) NOT NULL, `name` varchar(256),`symbol` varchar(256), `decimals` int, PRIMARY KEY (`address`))"
const pendingWithdrawalsTableCreation = "CREATE TABLE pending_withdrawals ( id INT AUTO_INCREMENT, address VARCHAR(256) NOT NULL, token_address VARCHAR(256) NOT NULL, amount VARCHAR(256), pending TINYINT(1), PRIMARY KEY(id) );"
const authTableCreation = "CREATE TABLE `auth` (`address` varchar(256) NOT NULL, `HASH` varchar(256), PRIMARY KEY (`address`))"
