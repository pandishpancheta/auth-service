package com.stockchain.stockchainauth.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OrderDTO {
    private UUID id;
    private UUID listingId;
    private UUID userId;
    private String token_uri;
    private String status;
}
