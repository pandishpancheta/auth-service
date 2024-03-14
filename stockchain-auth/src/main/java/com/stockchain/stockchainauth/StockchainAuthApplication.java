package com.stockchain.stockchainauth;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class StockchainAuthApplication {

	public static void main(String[] args) {
		SpringApplication.run(StockchainAuthApplication.class, args);
	}

}
