package commands

const usersTableCreation = "CREATE TABLE `users` ( `address` varchar(255) NOT NULL, `balance` varchar(255), PRIMARY KEY (`address`) )"
const tokensBalancesTableCreation = "CREATE TABLE `tokens_balances` ( `address` varchar(255) NOT NULL, `token_address` varchar(255) NOT NULL, `balance` varchar(255), PRIMARY KEY (`address`, `token_address`) )"
const tokensTableCreation = "CREATE TABLE `tokens` (`address` varchar(256) NOT NULL, `name` varchar(256),`symbol` varchar(256), `decimals` int, PRIMARY KEY (`address`))"
const pendingWithdrawalsTableCreation = "CREATE TABLE `pending_withdrawals` (`address` varchar(256) NOT NULL, `token_address` varchar(256) NOT NULL, `amount` varchar(256), `pending` BOOL, PRIMARY KEY (`address`, `token_address`))"
