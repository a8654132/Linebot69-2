package main

type USER_MAC struct{
	LINDID	string
	USER 	[]URL
	CRAWLER []KEYWORD
}

type URL struct{
	NAME 	string
	CONTENT string
	KEYWORD string
	TIME	int
}

type KEYWORD struct{
	NAME	string
	CONTENT string
	TIME	int
}
